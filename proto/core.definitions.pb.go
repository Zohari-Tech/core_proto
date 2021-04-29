// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: core.definitions.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SessionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phonenumber     string `protobuf:"bytes,1,opt,name=phonenumber,proto3" json:"phonenumber,omitempty"`
	ClientSessionid string `protobuf:"bytes,2,opt,name=clientSessionid,proto3" json:"clientSessionid,omitempty"`
	MessageId       string `protobuf:"bytes,3,opt,name=messageId,proto3" json:"messageId,omitempty"`
	SessionId       string `protobuf:"bytes,4,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Service         string `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *SessionData) Reset() {
	*x = SessionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_definitions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionData) ProtoMessage() {}

func (x *SessionData) ProtoReflect() protoreflect.Message {
	mi := &file_core_definitions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionData.ProtoReflect.Descriptor instead.
func (*SessionData) Descriptor() ([]byte, []int) {
	return file_core_definitions_proto_rawDescGZIP(), []int{0}
}

func (x *SessionData) GetPhonenumber() string {
	if x != nil {
		return x.Phonenumber
	}
	return ""
}

func (x *SessionData) GetClientSessionid() string {
	if x != nil {
		return x.ClientSessionid
	}
	return ""
}

func (x *SessionData) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *SessionData) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *SessionData) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type SessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input     string       `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Data      *SessionData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp int64        `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SessionRequest) Reset() {
	*x = SessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_definitions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionRequest) ProtoMessage() {}

func (x *SessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_definitions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionRequest.ProtoReflect.Descriptor instead.
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return file_core_definitions_proto_rawDescGZIP(), []int{1}
}

func (x *SessionRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *SessionRequest) GetData() *SessionData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SessionRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type SessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response     string       `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Data         *SessionData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	MessageTitle string       `protobuf:"bytes,3,opt,name=messageTitle,proto3" json:"messageTitle,omitempty"`
	Timestamp    int64        `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SessionResponse) Reset() {
	*x = SessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_definitions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionResponse) ProtoMessage() {}

func (x *SessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_definitions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionResponse.ProtoReflect.Descriptor instead.
func (*SessionResponse) Descriptor() ([]byte, []int) {
	return file_core_definitions_proto_rawDescGZIP(), []int{2}
}

func (x *SessionResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *SessionResponse) GetData() *SessionData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SessionResponse) GetMessageTitle() string {
	if x != nil {
		return x.MessageTitle
	}
	return ""
}

func (x *SessionResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_core_definitions_proto protoreflect.FileDescriptor

var file_core_definitions_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaf, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x6c, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x97, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x50, 0x0a, 0x11, 0x4e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x08, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_definitions_proto_rawDescOnce sync.Once
	file_core_definitions_proto_rawDescData = file_core_definitions_proto_rawDesc
)

func file_core_definitions_proto_rawDescGZIP() []byte {
	file_core_definitions_proto_rawDescOnce.Do(func() {
		file_core_definitions_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_definitions_proto_rawDescData)
	})
	return file_core_definitions_proto_rawDescData
}

var file_core_definitions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_definitions_proto_goTypes = []interface{}{
	(*SessionData)(nil),     // 0: proto.SessionData
	(*SessionRequest)(nil),  // 1: proto.SessionRequest
	(*SessionResponse)(nil), // 2: proto.SessionResponse
}
var file_core_definitions_proto_depIdxs = []int32{
	0, // 0: proto.SessionRequest.data:type_name -> proto.SessionData
	0, // 1: proto.SessionResponse.data:type_name -> proto.SessionData
	1, // 2: proto.NavigationService.Navigate:input_type -> proto.SessionRequest
	2, // 3: proto.NavigationService.Navigate:output_type -> proto.SessionResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_definitions_proto_init() }
func file_core_definitions_proto_init() {
	if File_core_definitions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_definitions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionData); i {
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
		file_core_definitions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionRequest); i {
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
		file_core_definitions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionResponse); i {
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
			RawDescriptor: file_core_definitions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_definitions_proto_goTypes,
		DependencyIndexes: file_core_definitions_proto_depIdxs,
		MessageInfos:      file_core_definitions_proto_msgTypes,
	}.Build()
	File_core_definitions_proto = out.File
	file_core_definitions_proto_rawDesc = nil
	file_core_definitions_proto_goTypes = nil
	file_core_definitions_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NavigationServiceClient is the client API for NavigationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NavigationServiceClient interface {
	Navigate(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionResponse, error)
}

type navigationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNavigationServiceClient(cc grpc.ClientConnInterface) NavigationServiceClient {
	return &navigationServiceClient{cc}
}

func (c *navigationServiceClient) Navigate(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, "/proto.NavigationService/Navigate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NavigationServiceServer is the server API for NavigationService service.
type NavigationServiceServer interface {
	Navigate(context.Context, *SessionRequest) (*SessionResponse, error)
}

// UnimplementedNavigationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNavigationServiceServer struct {
}

func (*UnimplementedNavigationServiceServer) Navigate(context.Context, *SessionRequest) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Navigate not implemented")
}

func RegisterNavigationServiceServer(s *grpc.Server, srv NavigationServiceServer) {
	s.RegisterService(&_NavigationService_serviceDesc, srv)
}

func _NavigationService_Navigate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).Navigate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NavigationService/Navigate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).Navigate(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NavigationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NavigationService",
	HandlerType: (*NavigationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Navigate",
			Handler:    _NavigationService_Navigate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.definitions.proto",
}
