// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subscribe.proto

package proto

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

type SubscribeRequest struct {
	// TODO: Add the variable(s) in the subscribe message
	ClientAddr           string   `protobuf:"bytes,1,opt,name=ClientAddr,proto3" json:"ClientAddr,omitempty"`
	RefreshRate          int64    `protobuf:"varint,2,opt,name=RefreshRate,proto3" json:"RefreshRate,omitempty"`
	Function             string   `protobuf:"bytes,3,opt,name=Function,proto3" json:"Function,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d2980c9543da44, []int{0}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

func (m *SubscribeRequest) GetRefreshRate() int64 {
	if m != nil {
		return m.RefreshRate
	}
	return 0
}

func (m *SubscribeRequest) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

type Notification struct {
	// TODO: Add the variable(s) in the notification message
	Channel              []string `protobuf:"bytes,1,rep,name=Channel,proto3" json:"Channel,omitempty"`
	Viewers              []int64  `protobuf:"varint,2,rep,packed,name=Viewers,proto3" json:"Viewers,omitempty"`
	Duration             []string `protobuf:"bytes,3,rep,name=Duration,proto3" json:"Duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d2980c9543da44, []int{1}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetChannel() []string {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *Notification) GetViewers() []int64 {
	if m != nil {
		return m.Viewers
	}
	return nil
}

func (m *Notification) GetDuration() []string {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscribeRequest)(nil), "proto.SubscribeRequest")
	proto.RegisterType((*Notification)(nil), "proto.Notification")
}

func init() { proto.RegisterFile("subscribe.proto", fileDescriptor_38d2980c9543da44) }

var fileDescriptor_38d2980c9543da44 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4b, 0x4b, 0xc4, 0x30,
	0x14, 0x85, 0xcd, 0x04, 0x1f, 0xbd, 0x0e, 0x28, 0x71, 0x61, 0x98, 0x85, 0x84, 0xae, 0xb2, 0x1a,
	0x44, 0xd7, 0x2e, 0x64, 0xc4, 0x9d, 0x2e, 0x22, 0xb8, 0x14, 0xd2, 0xce, 0x2d, 0x0d, 0x94, 0xa4,
	0xe6, 0x81, 0x7f, 0x5f, 0x12, 0xdb, 0x52, 0x5c, 0x85, 0xef, 0x1c, 0xee, 0x21, 0x1f, 0x5c, 0x85,
	0xd4, 0x84, 0xd6, 0x9b, 0x06, 0xf7, 0xa3, 0x77, 0xd1, 0xb1, 0xd3, 0xf2, 0xd4, 0x23, 0x5c, 0x7f,
	0xcc, 0x8d, 0xc2, 0xef, 0x84, 0x21, 0xb2, 0x3b, 0x80, 0xc3, 0x60, 0xd0, 0xc6, 0xe7, 0xe3, 0xd1,
	0x73, 0x22, 0x88, 0xac, 0xd4, 0x2a, 0x61, 0x02, 0x2e, 0x15, 0x76, 0x1e, 0x43, 0xaf, 0x74, 0x44,
	0xbe, 0x11, 0x44, 0x52, 0xb5, 0x8e, 0xd8, 0x0e, 0x2e, 0x5e, 0x93, 0x6d, 0xa3, 0x71, 0x96, 0xd3,
	0x72, 0xbf, 0x70, 0xfd, 0x05, 0xdb, 0x77, 0x17, 0x4d, 0x67, 0x5a, 0x9d, 0x99, 0x71, 0x38, 0x3f,
	0xf4, 0xda, 0x5a, 0x1c, 0x38, 0x11, 0x54, 0x56, 0x6a, 0xc6, 0xdc, 0x7c, 0x1a, 0xfc, 0x41, 0x1f,
	0xf8, 0x46, 0x50, 0x49, 0xd5, 0x8c, 0x79, 0xff, 0x25, 0x79, 0x3d, 0xed, 0xe7, 0xa3, 0x85, 0x1f,
	0xde, 0x60, 0x3b, 0x19, 0x8d, 0x65, 0xff, 0x09, 0xaa, 0xc5, 0x90, 0xdd, 0xfe, 0xd9, 0xef, 0xff,
	0x3b, 0xef, 0x6e, 0xa6, 0x62, 0xfd, 0xb5, 0xfa, 0xe4, 0x9e, 0x34, 0x67, 0x25, 0x7f, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x87, 0x51, 0xd9, 0x2a, 0x41, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SubscriptionClient is the client API for Subscription service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscriptionClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Subscription_SubscribeClient, error)
}

type subscriptionClient struct {
	cc *grpc.ClientConn
}

func NewSubscriptionClient(cc *grpc.ClientConn) SubscriptionClient {
	return &subscriptionClient{cc}
}

func (c *subscriptionClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Subscription_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Subscription_serviceDesc.Streams[0], "/proto.Subscription/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Subscription_SubscribeClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type subscriptionSubscribeClient struct {
	grpc.ClientStream
}

func (x *subscriptionSubscribeClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SubscriptionServer is the server API for Subscription service.
type SubscriptionServer interface {
	Subscribe(*SubscribeRequest, Subscription_SubscribeServer) error
}

// UnimplementedSubscriptionServer can be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServer struct {
}

func (*UnimplementedSubscriptionServer) Subscribe(req *SubscribeRequest, srv Subscription_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterSubscriptionServer(s *grpc.Server, srv SubscriptionServer) {
	s.RegisterService(&_Subscription_serviceDesc, srv)
}

func _Subscription_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionServer).Subscribe(m, &subscriptionSubscribeServer{stream})
}

type Subscription_SubscribeServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type subscriptionSubscribeServer struct {
	grpc.ServerStream
}

func (x *subscriptionSubscribeServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

var _Subscription_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Subscription",
	HandlerType: (*SubscriptionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Subscription_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "subscribe.proto",
}