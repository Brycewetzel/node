// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: pb/p2p.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type P2PSignedMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`           // Holds data of P2PConfigExchange.
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"` // Signature of data field.
}

func (x *P2PSignedMsg) Reset() {
	*x = P2PSignedMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_p2p_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P2PSignedMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P2PSignedMsg) ProtoMessage() {}

func (x *P2PSignedMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pb_p2p_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P2PSignedMsg.ProtoReflect.Descriptor instead.
func (*P2PSignedMsg) Descriptor() ([]byte, []int) {
	return file_pb_p2p_proto_rawDescGZIP(), []int{0}
}

func (x *P2PSignedMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *P2PSignedMsg) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type P2PConfigExchangeMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey        string `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`               // Public key field which is send from both provider and consumer.
	ConfigCiphertext []byte `protobuf:"bytes,2,opt,name=configCiphertext,proto3" json:"configCiphertext,omitempty"` // Encrypted P2PConnectConfig data.
}

func (x *P2PConfigExchangeMsg) Reset() {
	*x = P2PConfigExchangeMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_p2p_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P2PConfigExchangeMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P2PConfigExchangeMsg) ProtoMessage() {}

func (x *P2PConfigExchangeMsg) ProtoReflect() protoreflect.Message {
	mi := &file_pb_p2p_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P2PConfigExchangeMsg.ProtoReflect.Descriptor instead.
func (*P2PConfigExchangeMsg) Descriptor() ([]byte, []int) {
	return file_pb_p2p_proto_rawDescGZIP(), []int{1}
}

func (x *P2PConfigExchangeMsg) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *P2PConfigExchangeMsg) GetConfigCiphertext() []byte {
	if x != nil {
		return x.ConfigCiphertext
	}
	return nil
}

type P2PConnectConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicIP      string  `protobuf:"bytes,1,opt,name=publicIP,proto3" json:"publicIP,omitempty"`
	Ports         []int32 `protobuf:"varint,2,rep,packed,name=ports,proto3" json:"ports,omitempty"`
	Compatibility int32   `protobuf:"varint,3,opt,name=compatibility,proto3" json:"compatibility,omitempty"`
}

func (x *P2PConnectConfig) Reset() {
	*x = P2PConnectConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_p2p_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P2PConnectConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P2PConnectConfig) ProtoMessage() {}

func (x *P2PConnectConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pb_p2p_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P2PConnectConfig.ProtoReflect.Descriptor instead.
func (*P2PConnectConfig) Descriptor() ([]byte, []int) {
	return file_pb_p2p_proto_rawDescGZIP(), []int{2}
}

func (x *P2PConnectConfig) GetPublicIP() string {
	if x != nil {
		return x.PublicIP
	}
	return ""
}

func (x *P2PConnectConfig) GetPorts() []int32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *P2PConnectConfig) GetCompatibility() int32 {
	if x != nil {
		return x.Compatibility
	}
	return 0
}

type P2PKeepAlivePing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
}

func (x *P2PKeepAlivePing) Reset() {
	*x = P2PKeepAlivePing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_p2p_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P2PKeepAlivePing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P2PKeepAlivePing) ProtoMessage() {}

func (x *P2PKeepAlivePing) ProtoReflect() protoreflect.Message {
	mi := &file_pb_p2p_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P2PKeepAlivePing.ProtoReflect.Descriptor instead.
func (*P2PKeepAlivePing) Descriptor() ([]byte, []int) {
	return file_pb_p2p_proto_rawDescGZIP(), []int{3}
}

func (x *P2PKeepAlivePing) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

type P2PChannelHandlersReady struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *P2PChannelHandlersReady) Reset() {
	*x = P2PChannelHandlersReady{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_p2p_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P2PChannelHandlersReady) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P2PChannelHandlersReady) ProtoMessage() {}

func (x *P2PChannelHandlersReady) ProtoReflect() protoreflect.Message {
	mi := &file_pb_p2p_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P2PChannelHandlersReady.ProtoReflect.Descriptor instead.
func (*P2PChannelHandlersReady) Descriptor() ([]byte, []int) {
	return file_pb_p2p_proto_rawDescGZIP(), []int{4}
}

func (x *P2PChannelHandlersReady) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type P2PChannelEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	StatusCode uint64 `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Topic      string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Msg        string `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	Data       []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *P2PChannelEnvelope) Reset() {
	*x = P2PChannelEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_p2p_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P2PChannelEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P2PChannelEnvelope) ProtoMessage() {}

func (x *P2PChannelEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_pb_p2p_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P2PChannelEnvelope.ProtoReflect.Descriptor instead.
func (*P2PChannelEnvelope) Descriptor() ([]byte, []int) {
	return file_pb_p2p_proto_rawDescGZIP(), []int{5}
}

func (x *P2PChannelEnvelope) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *P2PChannelEnvelope) GetStatusCode() uint64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *P2PChannelEnvelope) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *P2PChannelEnvelope) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *P2PChannelEnvelope) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pb_p2p_proto protoreflect.FileDescriptor

var file_pb_p2p_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x62, 0x2f, 0x70, 0x32, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x40, 0x0a, 0x0c, 0x50, 0x32, 0x50, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d,
	0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x60, 0x0a, 0x14, 0x50, 0x32, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x22, 0x6a, 0x0a, 0x10, 0x50, 0x32, 0x50, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x22, 0x30, 0x0a, 0x10, 0x50, 0x32, 0x50, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69,
	0x76, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x22, 0x2f, 0x0a, 0x17, 0x50, 0x32, 0x50, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x50, 0x32, 0x50, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_p2p_proto_rawDescOnce sync.Once
	file_pb_p2p_proto_rawDescData = file_pb_p2p_proto_rawDesc
)

func file_pb_p2p_proto_rawDescGZIP() []byte {
	file_pb_p2p_proto_rawDescOnce.Do(func() {
		file_pb_p2p_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_p2p_proto_rawDescData)
	})
	return file_pb_p2p_proto_rawDescData
}

var file_pb_p2p_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_p2p_proto_goTypes = []interface{}{
	(*P2PSignedMsg)(nil),            // 0: pb.P2PSignedMsg
	(*P2PConfigExchangeMsg)(nil),    // 1: pb.P2PConfigExchangeMsg
	(*P2PConnectConfig)(nil),        // 2: pb.P2PConnectConfig
	(*P2PKeepAlivePing)(nil),        // 3: pb.P2PKeepAlivePing
	(*P2PChannelHandlersReady)(nil), // 4: pb.P2PChannelHandlersReady
	(*P2PChannelEnvelope)(nil),      // 5: pb.P2PChannelEnvelope
}
var file_pb_p2p_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_p2p_proto_init() }
func file_pb_p2p_proto_init() {
	if File_pb_p2p_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_p2p_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P2PSignedMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_p2p_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P2PConfigExchangeMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_p2p_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P2PConnectConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_p2p_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P2PKeepAlivePing); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_p2p_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P2PChannelHandlersReady); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_p2p_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P2PChannelEnvelope); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_p2p_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_p2p_proto_goTypes,
		DependencyIndexes: file_pb_p2p_proto_depIdxs,
		MessageInfos:      file_pb_p2p_proto_msgTypes,
	}.Build()
	File_pb_p2p_proto = out.File
	file_pb_p2p_proto_rawDesc = nil
	file_pb_p2p_proto_goTypes = nil
	file_pb_p2p_proto_depIdxs = nil
}
