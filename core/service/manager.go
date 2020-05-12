/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package service

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/mysteriumnetwork/node/core/policy"
	"github.com/mysteriumnetwork/node/core/service/servicestate"
	"github.com/mysteriumnetwork/node/identity"
	"github.com/mysteriumnetwork/node/market"
	"github.com/mysteriumnetwork/node/p2p"
	"github.com/mysteriumnetwork/node/session"
	"github.com/mysteriumnetwork/node/session/connectivity"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var (
	// ErrorLocation error indicates that action (i.e. disconnect)
	ErrorLocation = errors.New("failed to detect service location")
	// ErrUnsupportedServiceType indicates that manager tried to create an unsupported service type
	ErrUnsupportedServiceType = errors.New("unsupported service type")
	// ErrUnsupportedAccessPolicy indicates that manager tried to create service with unsupported access policy
	ErrUnsupportedAccessPolicy = errors.New("unsupported access policy")
)

// Service interface represents pluggable Mysterium service
type Service interface {
	Serve(instance *Instance) error
	Stop() error
	session.ConfigProvider
}

// DiscoveryFactory initiates instance which is able announce service discoverability
type DiscoveryFactory func() Discovery

// Discovery registers the service to the discovery api periodically
type Discovery interface {
	Start(ownIdentity identity.Identity, proposal market.ServiceProposal)
	Stop()
	Wait()
}

// WaitForNATHole blocks until NAT hole is punched towards consumer through local NAT or until hole punching failed
type WaitForNATHole func() error

// NewManager creates new instance of pluggable instances manager
func NewManager(
	serviceRegistry *Registry,
	discoveryFactory DiscoveryFactory,
	eventPublisher Publisher,
	policyOracle *policy.Oracle,
	p2pListener p2p.Listener,
	sessionManager func(proposal market.ServiceProposal, serviceID string, channel p2p.Channel) *session.Manager,
	statusStorage connectivity.StatusStorage,
) *Manager {
	return &Manager{
		serviceRegistry:  serviceRegistry,
		servicePool:      NewPool(eventPublisher),
		discoveryFactory: discoveryFactory,
		eventPublisher:   eventPublisher,
		policyOracle:     policyOracle,
		p2pListener:      p2pListener,
		sessionManager:   sessionManager,
		statusStorage:    statusStorage,
	}
}

// Manager entrypoint which knows how to start pluggable Mysterium instances
type Manager struct {
	serviceRegistry *Registry
	servicePool     *Pool

	discoveryFactory DiscoveryFactory
	eventPublisher   Publisher
	policyOracle     *policy.Oracle

	p2pListener    p2p.Listener
	sessionManager func(proposal market.ServiceProposal, serviceID string, channel p2p.Channel) *session.Manager
	statusStorage  connectivity.StatusStorage
}

// Start starts an instance of the given service type if knows one in service registry.
// It passes the options to the start method of the service.
// If an error occurs in the underlying service, the error is then returned.
func (manager *Manager) Start(providerID identity.Identity, serviceType string, policyIDs []string, options Options, pm market.PaymentMethod) (id ID, err error) {
	service, proposal, err := manager.serviceRegistry.Create(serviceType, options)
	if err != nil {
		return id, err
	}

	proposal.SetPaymentMethod(pm)
	proposal.SetAccessPolicies(nil)
	policyRules := policy.NewRepository()
	if len(policyIDs) > 0 {
		policies := manager.policyOracle.Policies(policyIDs)
		if err = manager.policyOracle.SubscribePolicies(policies, policyRules); err != nil {
			log.Warn().Err(err).Msg("Can't find given access policies")
			return id, ErrUnsupportedAccessPolicy
		}
		proposal.SetAccessPolicies(&policies)
	}

	proposal.SetProviderContacts(providerID, market.ContactList{manager.p2pListener.GetContact()})

	id, err = generateID()
	if err != nil {
		return id, err
	}

	discovery := manager.discoveryFactory()
	discovery.Start(providerID, proposal)

	instance := &Instance{
		id:             id,
		state:          servicestate.Starting,
		options:        options,
		service:        service,
		proposal:       proposal,
		policies:       policyRules,
		discovery:      discovery,
		eventPublisher: manager.eventPublisher,
	}

	channelHandlers := func(ch p2p.Channel) {
		instance.addP2PChannel(ch)
		mng := manager.sessionManager(proposal, string(id), ch)
		subscribeSessionCreate(mng, ch, service)
		subscribeSessionStatus(mng, ch, manager.statusStorage)
		subscribeSessionAcknowledge(mng, ch)
		subscribeSessionDestroy(mng, ch, func() {
			// Give some time for channel to finish sending last message.
			time.Sleep(10 * time.Second)
			instance.closeP2PChannel(ch)
		})
	}
	if err := manager.p2pListener.Listen(providerID, serviceType, channelHandlers); err != nil {
		return id, fmt.Errorf("could not subscribe to p2p channels: %w", err)
	}

	manager.servicePool.Add(instance)

	go func() {
		instance.setState(servicestate.Running)

		serveErr := service.Serve(instance)
		if serveErr != nil {
			log.Error().Err(serveErr).Msg("Service serve failed")
		}

		// TODO: fix https://github.com/mysteriumnetwork/node/issues/855
		stopErr := manager.servicePool.Stop(id)
		if stopErr != nil {
			log.Error().Err(stopErr).Msg("Service stop failed")
		}

		discovery.Wait()
	}()

	return id, nil
}

func generateID() (ID, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return ID(""), err
	}
	return ID(uid.String()), nil
}

// List returns array of running service instances.
func (manager *Manager) List() map[ID]*Instance {
	return manager.servicePool.List()
}

// Kill stops all services.
func (manager *Manager) Kill() error {
	return manager.servicePool.StopAll()
}

// Stop stops the service.
func (manager *Manager) Stop(id ID) error {
	err := manager.servicePool.Stop(id)
	if err != nil {
		return err
	}

	return nil
}

// Service returns a service instance by requested id.
func (manager *Manager) Service(id ID) *Instance {
	return manager.servicePool.Instance(id)
}
