package command_run

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	identity_handler "github.com/mysterium/node/cmd/mysterium_server/command_run/identity"
	"github.com/mysterium/node/communication"
	nats_dialog "github.com/mysterium/node/communication/nats/dialog"
	nats_discovery "github.com/mysterium/node/communication/nats/discovery"
	"github.com/mysterium/node/identity"
	"github.com/mysterium/node/ip"
	"github.com/mysterium/node/location"
	"github.com/mysterium/node/nat"
	"github.com/mysterium/node/openvpn"
	"github.com/mysterium/node/openvpn/middlewares/server/auth"
	"github.com/mysterium/node/openvpn/primitives"
	openvpn_session "github.com/mysterium/node/openvpn/session"
	"github.com/mysterium/node/server"
	"github.com/mysterium/node/session"
	"path/filepath"
)

// NewCommand function creates new server command by given options
func NewCommand(options CommandOptions) *CommandRun {
	return NewCommandWith(
		options,
		server.NewClient(),
		ip.NewResolver(),
		nat.NewService(),
	)
}

// NewCommandWith function creates new client command by given options + injects given dependencies
func NewCommandWith(
	options CommandOptions,
	mysteriumClient server.Client,
	ipResolver ip.Resolver,
	natService nat.NATService,
) *CommandRun {

	keystoreInstance := keystore.NewKeyStore(options.DirectoryKeystore, keystore.StandardScryptN, keystore.StandardScryptP)
	cache := identity.NewIdentityCache(options.DirectoryKeystore, "remember.json")
	identityManager := identity.NewIdentityManager(keystoreInstance)
	createSigner := func(id identity.Identity) identity.Signer {
		return identity.NewSigner(keystoreInstance, id)
	}
	identityHandler := identity_handler.NewHandler(
		identityManager,
		mysteriumClient,
		cache,
		createSigner,
	)

	var locationDetector location.Detector
	if options.LocationCountry != "" {
		locationDetector = location.NewDetectorFake(options.LocationCountry)
	} else if options.LocationDatabase != "" {
		locationDetector = location.NewDetector(filepath.Join(options.DirectoryConfig, options.LocationDatabase))
	} else {
		locationDetector = location.NewDetectorFake("")
	}

	// (Re)generate required security primitives before openvpn start
	openVPNPrimitives := primitives.NewOpenVPNSecPrimitives()
	openVPNPrimitives.Generate()

	return &CommandRun{
		identityLoader: func() (identity.Identity, error) {
			return identity_handler.LoadIdentity(identityHandler, options.NodeKey, options.Passphrase)
		},
		createSigner:     createSigner,
		locationDetector: locationDetector,
		ipResolver:       ipResolver,
		mysteriumClient:  mysteriumClient,
		natService:       natService,
		dialogWaiterFactory: func(myID identity.Identity) communication.DialogWaiter {
			return nats_dialog.NewDialogWaiter(
				nats_discovery.NewAddressGenerate(myID),
				identity.NewSigner(keystoreInstance, myID),
			)
		},

		sessionManagerFactory: func(vpnServerIP string) session.Manager {
			return openvpn_session.NewManager(
				openvpn.NewClientConfig(
					vpnServerIP,
					openVPNPrimitives,
				),
				&session.UUIDGenerator{},
			)
		},
		vpnServerFactory: func(manager session.Manager) *openvpn.Server {
			vpnServerConfig := openvpn.NewServerConfig(
				"10.8.0.0", "255.255.255.0",
				openVPNPrimitives,
			)
			sessionValidator := openvpn_session.NewSessionValidator(
				manager.FindSession,
				identity.NewExtractor(),
			)
			vpnMiddlewares := []openvpn.ManagementMiddleware{
				auth.NewMiddleware(sessionValidator),
			}
			return openvpn.NewServer(vpnServerConfig, options.DirectoryRuntime, vpnMiddlewares...)
		},
	}
}
