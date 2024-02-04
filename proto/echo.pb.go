// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/echo.proto

package proto

import (
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

type PingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelaySeconds float64 `protobuf:"fixed64,1,opt,name=delaySeconds,proto3" json:"delaySeconds,omitempty"`
	Payload      string  `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PingMessage) Reset() {
	*x = PingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingMessage) ProtoMessage() {}

func (x *PingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingMessage.ProtoReflect.Descriptor instead.
func (*PingMessage) Descriptor() ([]byte, []int) {
	return file_proto_echo_proto_rawDescGZIP(), []int{0}
}

func (x *PingMessage) GetDelaySeconds() float64 {
	if x != nil {
		return x.DelaySeconds
	}
	return 0
}

func (x *PingMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type PongMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp  string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Hostname   string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	RemoteIp   string `protobuf:"bytes,4,opt,name=remoteIp,proto3" json:"remoteIp,omitempty"`
	RemotePort int32  `protobuf:"varint,5,opt,name=remotePort,proto3" json:"remotePort,omitempty"`
	// reserve a field for real client ip
	// string clientIp = 6;
	RequestId     string  `protobuf:"bytes,7,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Authority     string  `protobuf:"bytes,8,opt,name=authority,proto3" json:"authority,omitempty"`
	RequestMethod string  `protobuf:"bytes,9,opt,name=requestMethod,proto3" json:"requestMethod,omitempty"`
	RequestTime   float64 `protobuf:"fixed64,10,opt,name=requestTime,proto3" json:"requestTime,omitempty"`
	UserAgent     string  `protobuf:"bytes,11,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	DelaySeconds  float64 `protobuf:"fixed64,21,opt,name=delaySeconds,proto3" json:"delaySeconds,omitempty"`
	Payload       string  `protobuf:"bytes,22,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PongMessage) Reset() {
	*x = PongMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongMessage) ProtoMessage() {}

func (x *PongMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongMessage.ProtoReflect.Descriptor instead.
func (*PongMessage) Descriptor() ([]byte, []int) {
	return file_proto_echo_proto_rawDescGZIP(), []int{1}
}

func (x *PongMessage) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *PongMessage) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *PongMessage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PongMessage) GetRemoteIp() string {
	if x != nil {
		return x.RemoteIp
	}
	return ""
}

func (x *PongMessage) GetRemotePort() int32 {
	if x != nil {
		return x.RemotePort
	}
	return 0
}

func (x *PongMessage) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *PongMessage) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *PongMessage) GetRequestMethod() string {
	if x != nil {
		return x.RequestMethod
	}
	return ""
}

func (x *PongMessage) GetRequestTime() float64 {
	if x != nil {
		return x.RequestTime
	}
	return 0
}

func (x *PongMessage) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *PongMessage) GetDelaySeconds() float64 {
	if x != nil {
		return x.DelaySeconds
	}
	return 0
}

func (x *PongMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_proto_echo_proto protoreflect.FileDescriptor

var file_proto_echo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x22, 0x4b, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xfd, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x6e, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x49, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x49, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x36, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x2e, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e,
	0x50, 0x6f, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x70, 0x77,
	0x69, 0x6e, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_echo_proto_rawDescOnce sync.Once
	file_proto_echo_proto_rawDescData = file_proto_echo_proto_rawDesc
)

func file_proto_echo_proto_rawDescGZIP() []byte {
	file_proto_echo_proto_rawDescOnce.Do(func() {
		file_proto_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_echo_proto_rawDescData)
	})
	return file_proto_echo_proto_rawDescData
}

var file_proto_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_echo_proto_goTypes = []interface{}{
	(*PingMessage)(nil), // 0: echo.PingMessage
	(*PongMessage)(nil), // 1: echo.PongMessage
}
var file_proto_echo_proto_depIdxs = []int32{
	0, // 0: echo.Echo.Ping:input_type -> echo.PingMessage
	1, // 1: echo.Echo.Ping:output_type -> echo.PongMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_echo_proto_init() }
func file_proto_echo_proto_init() {
	if File_proto_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingMessage); i {
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
		file_proto_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongMessage); i {
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
			RawDescriptor: file_proto_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_echo_proto_goTypes,
		DependencyIndexes: file_proto_echo_proto_depIdxs,
		MessageInfos:      file_proto_echo_proto_msgTypes,
	}.Build()
	File_proto_echo_proto = out.File
	file_proto_echo_proto_rawDesc = nil
	file_proto_echo_proto_goTypes = nil
	file_proto_echo_proto_depIdxs = nil
}
