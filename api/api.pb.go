// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type StatsRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsRequest) Reset()         { *m = StatsRequest{} }
func (m *StatsRequest) String() string { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()    {}
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *StatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsRequest.Unmarshal(m, b)
}
func (m *StatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsRequest.Marshal(b, m, deterministic)
}
func (m *StatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsRequest.Merge(m, src)
}
func (m *StatsRequest) XXX_Size() int {
	return xxx_messageInfo_StatsRequest.Size(m)
}
func (m *StatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsRequest proto.InternalMessageInfo

func (m *StatsRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type StatsReply struct {
	Upload               uint64   `protobuf:"varint,1,opt,name=upload,proto3" json:"upload,omitempty"`
	Download             uint64   `protobuf:"varint,2,opt,name=download,proto3" json:"download,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsReply) Reset()         { *m = StatsReply{} }
func (m *StatsReply) String() string { return proto.CompactTextString(m) }
func (*StatsReply) ProtoMessage()    {}
func (*StatsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *StatsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsReply.Unmarshal(m, b)
}
func (m *StatsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsReply.Marshal(b, m, deterministic)
}
func (m *StatsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsReply.Merge(m, src)
}
func (m *StatsReply) XXX_Size() int {
	return xxx_messageInfo_StatsReply.Size(m)
}
func (m *StatsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatsReply proto.InternalMessageInfo

func (m *StatsReply) GetUpload() uint64 {
	if m != nil {
		return m.Upload
	}
	return 0
}

func (m *StatsReply) GetDownload() uint64 {
	if m != nil {
		return m.Download
	}
	return 0
}

func init() {
	proto.RegisterType((*StatsRequest)(nil), "api.StatsRequest")
	proto.RegisterType((*StatsReply)(nil), "api.StatsReply")
}

func init() {
	proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c)
}

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x06, 0x32, 0x95, 0xb4, 0xb8, 0x78, 0x82, 0x4b, 0x12,
	0x4b, 0x8a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38, 0x0a, 0x12, 0x8b,
	0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x07,
	0x2e, 0x2e, 0xa8, 0xda, 0x82, 0x9c, 0x4a, 0x21, 0x31, 0x2e, 0xb6, 0xd2, 0x82, 0x9c, 0xfc, 0x44,
	0x88, 0x3a, 0x96, 0x20, 0x28, 0x0f, 0x64, 0x42, 0x4a, 0x7e, 0x79, 0x1e, 0x58, 0x86, 0x09, 0x2c,
	0x03, 0xe7, 0x1b, 0x39, 0x73, 0xf1, 0x86, 0x14, 0xe5, 0x67, 0x25, 0xe6, 0x05, 0xa7, 0x16, 0x95,
	0x65, 0x26, 0xa7, 0x0a, 0x19, 0x71, 0x71, 0x05, 0x96, 0xa6, 0x16, 0x55, 0x82, 0xcd, 0x15, 0x12,
	0xd4, 0x03, 0xb9, 0x0e, 0xd9, 0x3d, 0x52, 0xfc, 0xc8, 0x42, 0x40, 0x6b, 0x95, 0x18, 0x92, 0xd8,
	0xc0, 0xce, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xea, 0x6a, 0xb8, 0xe4, 0xcb, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TrojanServiceClient is the client API for TrojanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrojanServiceClient interface {
	QueryStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error)
}

type trojanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrojanServiceClient(cc grpc.ClientConnInterface) TrojanServiceClient {
	return &trojanServiceClient{cc}
}

func (c *trojanServiceClient) QueryStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error) {
	out := new(StatsReply)
	err := c.cc.Invoke(ctx, "/api.TrojanService/QueryStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrojanServiceServer is the server API for TrojanService service.
type TrojanServiceServer interface {
	QueryStats(context.Context, *StatsRequest) (*StatsReply, error)
}

// UnimplementedTrojanServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTrojanServiceServer struct {
}

func (*UnimplementedTrojanServiceServer) QueryStats(ctx context.Context, req *StatsRequest) (*StatsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStats not implemented")
}

func RegisterTrojanServiceServer(s *grpc.Server, srv TrojanServiceServer) {
	s.RegisterService(&_TrojanService_serviceDesc, srv)
}

func _TrojanService_QueryStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrojanServiceServer).QueryStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TrojanService/QueryStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrojanServiceServer).QueryStats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrojanService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TrojanService",
	HandlerType: (*TrojanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryStats",
			Handler:    _TrojanService_QueryStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
