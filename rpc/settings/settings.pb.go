// Code generated by protoc-gen-go. DO NOT EDIT.
// source: settings/settings.proto

package settings

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RawData struct {
	// The settings, in JSON format.
	JsonData             string   `protobuf:"bytes,1,opt,name=jsonData,proto3" json:"jsonData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawData) Reset()         { *m = RawData{} }
func (m *RawData) String() string { return proto.CompactTextString(m) }
func (*RawData) ProtoMessage()    {}
func (*RawData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bfd59e429426d0, []int{0}
}

func (m *RawData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawData.Unmarshal(m, b)
}
func (m *RawData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawData.Marshal(b, m, deterministic)
}
func (m *RawData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawData.Merge(m, src)
}
func (m *RawData) XXX_Size() int {
	return xxx_messageInfo_RawData.Size(m)
}
func (m *RawData) XXX_DiscardUnknown() {
	xxx_messageInfo_RawData.DiscardUnknown(m)
}

var xxx_messageInfo_RawData proto.InternalMessageInfo

func (m *RawData) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

type Value struct {
	// The key of the setting.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The setting, in JSON format.
	JsonData             string   `protobuf:"bytes,2,opt,name=jsonData,proto3" json:"jsonData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bfd59e429426d0, []int{1}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Value) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

type GetAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllRequest) Reset()         { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()    {}
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bfd59e429426d0, []int{2}
}

func (m *GetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRequest.Unmarshal(m, b)
}
func (m *GetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRequest.Marshal(b, m, deterministic)
}
func (m *GetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRequest.Merge(m, src)
}
func (m *GetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRequest.Size(m)
}
func (m *GetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRequest proto.InternalMessageInfo

type GetValueRequest struct {
	// The key of the setting.
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetValueRequest) Reset()         { *m = GetValueRequest{} }
func (m *GetValueRequest) String() string { return proto.CompactTextString(m) }
func (*GetValueRequest) ProtoMessage()    {}
func (*GetValueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bfd59e429426d0, []int{3}
}

func (m *GetValueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetValueRequest.Unmarshal(m, b)
}
func (m *GetValueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetValueRequest.Marshal(b, m, deterministic)
}
func (m *GetValueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetValueRequest.Merge(m, src)
}
func (m *GetValueRequest) XXX_Size() int {
	return xxx_messageInfo_GetValueRequest.Size(m)
}
func (m *GetValueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetValueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetValueRequest proto.InternalMessageInfo

func (m *GetValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type MergeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MergeResponse) Reset()         { *m = MergeResponse{} }
func (m *MergeResponse) String() string { return proto.CompactTextString(m) }
func (*MergeResponse) ProtoMessage()    {}
func (*MergeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bfd59e429426d0, []int{4}
}

func (m *MergeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MergeResponse.Unmarshal(m, b)
}
func (m *MergeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MergeResponse.Marshal(b, m, deterministic)
}
func (m *MergeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MergeResponse.Merge(m, src)
}
func (m *MergeResponse) XXX_Size() int {
	return xxx_messageInfo_MergeResponse.Size(m)
}
func (m *MergeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MergeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MergeResponse proto.InternalMessageInfo

type SetValueResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetValueResponse) Reset()         { *m = SetValueResponse{} }
func (m *SetValueResponse) String() string { return proto.CompactTextString(m) }
func (*SetValueResponse) ProtoMessage()    {}
func (*SetValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bfd59e429426d0, []int{5}
}

func (m *SetValueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetValueResponse.Unmarshal(m, b)
}
func (m *SetValueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetValueResponse.Marshal(b, m, deterministic)
}
func (m *SetValueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetValueResponse.Merge(m, src)
}
func (m *SetValueResponse) XXX_Size() int {
	return xxx_messageInfo_SetValueResponse.Size(m)
}
func (m *SetValueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetValueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetValueResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RawData)(nil), "cc.arduino.cli.settings.RawData")
	proto.RegisterType((*Value)(nil), "cc.arduino.cli.settings.Value")
	proto.RegisterType((*GetAllRequest)(nil), "cc.arduino.cli.settings.GetAllRequest")
	proto.RegisterType((*GetValueRequest)(nil), "cc.arduino.cli.settings.GetValueRequest")
	proto.RegisterType((*MergeResponse)(nil), "cc.arduino.cli.settings.MergeResponse")
	proto.RegisterType((*SetValueResponse)(nil), "cc.arduino.cli.settings.SetValueResponse")
}

func init() { proto.RegisterFile("settings/settings.proto", fileDescriptor_a4bfd59e429426d0) }

var fileDescriptor_a4bfd59e429426d0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0x2d, 0x29,
	0xc9, 0xcc, 0x4b, 0x2f, 0xd6, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0x93,
	0x93, 0xf5, 0x12, 0x8b, 0x52, 0x4a, 0x33, 0xf3, 0xf2, 0xf5, 0x92, 0x73, 0x32, 0xf5, 0x60, 0xd2,
	0x4a, 0xaa, 0x5c, 0xec, 0x41, 0x89, 0xe5, 0x2e, 0x89, 0x25, 0x89, 0x42, 0x52, 0x5c, 0x1c, 0x59,
	0xc5, 0xf9, 0x79, 0x20, 0xb6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0xaf, 0x64, 0xca,
	0xc5, 0x1a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x09, 0x95, 0x07,
	0x31, 0x51, 0xb4, 0x31, 0xa1, 0x69, 0xe3, 0xe7, 0xe2, 0x75, 0x4f, 0x2d, 0x71, 0xcc, 0xc9, 0x09,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x52, 0xe6, 0xe2, 0x77, 0x4f, 0x2d, 0x01, 0x1b, 0x05,
	0x15, 0xc2, 0x34, 0x11, 0xa4, 0xcb, 0x37, 0xb5, 0x28, 0x3d, 0x35, 0x28, 0xb5, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x55, 0x49, 0x88, 0x4b, 0x20, 0x18, 0xae, 0x0b, 0x22, 0x66, 0x74, 0x8f, 0x89, 0x8b,
	0x23, 0x18, 0xea, 0x0b, 0xa1, 0x20, 0x2e, 0x36, 0x88, 0x3d, 0x42, 0x6a, 0x7a, 0x38, 0x7c, 0xaa,
	0x87, 0xe2, 0x10, 0x29, 0x05, 0x9c, 0xea, 0x60, 0xc1, 0x11, 0xc8, 0xc5, 0x0a, 0x76, 0x85, 0x10,
	0x41, 0xa5, 0x52, 0xb8, 0x2d, 0x45, 0xf1, 0x87, 0x50, 0x08, 0x17, 0x07, 0xcc, 0xf7, 0x42, 0x1a,
	0xf8, 0x1c, 0x8a, 0x1c, 0x40, 0x52, 0x72, 0x38, 0x55, 0x42, 0x4c, 0x0a, 0x05, 0x07, 0x04, 0x84,
	0x4d, 0x40, 0xad, 0x94, 0x26, 0x4e, 0x79, 0xf4, 0x00, 0x76, 0xd2, 0x8d, 0xd2, 0x4e, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0xea, 0x81, 0xd1, 0xba, 0xc9, 0x39, 0x99,
	0xfa, 0x45, 0x05, 0xc9, 0xf0, 0x74, 0x96, 0xc4, 0x06, 0x4e, 0x68, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x77, 0x53, 0x26, 0x69, 0x83, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SettingsClient is the client API for Settings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SettingsClient interface {
	// List all the settings.
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*RawData, error)
	// Set multiple settings values at once.
	Merge(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*MergeResponse, error)
	// Get the value of a specific setting.
	GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*Value, error)
	// Set the value of a specific setting.
	SetValue(ctx context.Context, in *Value, opts ...grpc.CallOption) (*SetValueResponse, error)
}

type settingsClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsClient(cc grpc.ClientConnInterface) SettingsClient {
	return &settingsClient{cc}
}

func (c *settingsClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*RawData, error) {
	out := new(RawData)
	err := c.cc.Invoke(ctx, "/cc.arduino.cli.settings.Settings/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) Merge(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*MergeResponse, error) {
	out := new(MergeResponse)
	err := c.cc.Invoke(ctx, "/cc.arduino.cli.settings.Settings/Merge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/cc.arduino.cli.settings.Settings/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) SetValue(ctx context.Context, in *Value, opts ...grpc.CallOption) (*SetValueResponse, error) {
	out := new(SetValueResponse)
	err := c.cc.Invoke(ctx, "/cc.arduino.cli.settings.Settings/SetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServer is the server API for Settings service.
type SettingsServer interface {
	// List all the settings.
	GetAll(context.Context, *GetAllRequest) (*RawData, error)
	// Set multiple settings values at once.
	Merge(context.Context, *RawData) (*MergeResponse, error)
	// Get the value of a specific setting.
	GetValue(context.Context, *GetValueRequest) (*Value, error)
	// Set the value of a specific setting.
	SetValue(context.Context, *Value) (*SetValueResponse, error)
}

// UnimplementedSettingsServer can be embedded to have forward compatible implementations.
type UnimplementedSettingsServer struct {
}

func (*UnimplementedSettingsServer) GetAll(ctx context.Context, req *GetAllRequest) (*RawData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedSettingsServer) Merge(ctx context.Context, req *RawData) (*MergeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Merge not implemented")
}
func (*UnimplementedSettingsServer) GetValue(ctx context.Context, req *GetValueRequest) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (*UnimplementedSettingsServer) SetValue(ctx context.Context, req *Value) (*SetValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}

func RegisterSettingsServer(s *grpc.Server, srv SettingsServer) {
	s.RegisterService(&_Settings_serviceDesc, srv)
}

func _Settings_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.arduino.cli.settings.Settings/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_Merge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).Merge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.arduino.cli.settings.Settings/Merge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).Merge(ctx, req.(*RawData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.arduino.cli.settings.Settings/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetValue(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.arduino.cli.settings.Settings/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).SetValue(ctx, req.(*Value))
	}
	return interceptor(ctx, in, info, handler)
}

var _Settings_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cc.arduino.cli.settings.Settings",
	HandlerType: (*SettingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _Settings_GetAll_Handler,
		},
		{
			MethodName: "Merge",
			Handler:    _Settings_Merge_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _Settings_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _Settings_SetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "settings/settings.proto",
}
