// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/sync_rpc_service.proto

package protos

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

type GatewayRequest struct {
	GwId string `protobuf:"bytes,1,opt,name=gwId,proto3" json:"gwId,omitempty"`
	// sync_rpc_service leverages the fact that grpc is built on top of http/2.
	// As pseudo header :method will always be POST, and :scheme will always be http,
	// they are excluded from the message definition. grpc uses pseudo header :authority
	// and :path to route the rpc calls to the corresponding services
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	Path      string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// non-pseudo headers
	Headers              map[string]string `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload              []byte            `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GatewayRequest) Reset()         { *m = GatewayRequest{} }
func (m *GatewayRequest) String() string { return proto.CompactTextString(m) }
func (*GatewayRequest) ProtoMessage()    {}
func (*GatewayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22887391e8a5ac6c, []int{0}
}

func (m *GatewayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayRequest.Unmarshal(m, b)
}
func (m *GatewayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayRequest.Marshal(b, m, deterministic)
}
func (m *GatewayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayRequest.Merge(m, src)
}
func (m *GatewayRequest) XXX_Size() int {
	return xxx_messageInfo_GatewayRequest.Size(m)
}
func (m *GatewayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayRequest proto.InternalMessageInfo

func (m *GatewayRequest) GetGwId() string {
	if m != nil {
		return m.GwId
	}
	return ""
}

func (m *GatewayRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *GatewayRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GatewayRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *GatewayRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type GatewayResponse struct {
	// pseudo header :status
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// non-pseudo headers
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// response body
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// err message to return to the caller, if any
	Err string `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	// keepConnActive is used to tell the client that the connection is still active, even if no response has
	// been received for a while
	KeepConnActive       bool     `protobuf:"varint,5,opt,name=keepConnActive,proto3" json:"keepConnActive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayResponse) Reset()         { *m = GatewayResponse{} }
func (m *GatewayResponse) String() string { return proto.CompactTextString(m) }
func (*GatewayResponse) ProtoMessage()    {}
func (*GatewayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22887391e8a5ac6c, []int{1}
}

func (m *GatewayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayResponse.Unmarshal(m, b)
}
func (m *GatewayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayResponse.Marshal(b, m, deterministic)
}
func (m *GatewayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayResponse.Merge(m, src)
}
func (m *GatewayResponse) XXX_Size() int {
	return xxx_messageInfo_GatewayResponse.Size(m)
}
func (m *GatewayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayResponse proto.InternalMessageInfo

func (m *GatewayResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GatewayResponse) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *GatewayResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *GatewayResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *GatewayResponse) GetKeepConnActive() bool {
	if m != nil {
		return m.KeepConnActive
	}
	return false
}

// SyncRPCRequest is sent down to gateway from cloud
type SyncRPCRequest struct {
	// unique request Id passed in from cloud
	ReqId   uint32          `protobuf:"varint,1,opt,name=reqId,proto3" json:"reqId,omitempty"`
	ReqBody *GatewayRequest `protobuf:"bytes,2,opt,name=reqBody,proto3" json:"reqBody,omitempty"`
	// cloud will send a heartBeat every minute, if no other requests are sent
	// down to the gateway
	HeartBeat bool `protobuf:"varint,3,opt,name=heartBeat,proto3" json:"heartBeat,omitempty"`
	// connClosed is set to true when the client closes the connection
	ConnClosed           bool     `protobuf:"varint,4,opt,name=connClosed,proto3" json:"connClosed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRPCRequest) Reset()         { *m = SyncRPCRequest{} }
func (m *SyncRPCRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRPCRequest) ProtoMessage()    {}
func (*SyncRPCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22887391e8a5ac6c, []int{2}
}

func (m *SyncRPCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRPCRequest.Unmarshal(m, b)
}
func (m *SyncRPCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRPCRequest.Marshal(b, m, deterministic)
}
func (m *SyncRPCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRPCRequest.Merge(m, src)
}
func (m *SyncRPCRequest) XXX_Size() int {
	return xxx_messageInfo_SyncRPCRequest.Size(m)
}
func (m *SyncRPCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRPCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRPCRequest proto.InternalMessageInfo

func (m *SyncRPCRequest) GetReqId() uint32 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *SyncRPCRequest) GetReqBody() *GatewayRequest {
	if m != nil {
		return m.ReqBody
	}
	return nil
}

func (m *SyncRPCRequest) GetHeartBeat() bool {
	if m != nil {
		return m.HeartBeat
	}
	return false
}

func (m *SyncRPCRequest) GetConnClosed() bool {
	if m != nil {
		return m.ConnClosed
	}
	return false
}

// SyncRPCResponse is sent from gateway to cloud
type SyncRPCResponse struct {
	ReqId    uint32           `protobuf:"varint,1,opt,name=reqId,proto3" json:"reqId,omitempty"`
	RespBody *GatewayResponse `protobuf:"bytes,2,opt,name=respBody,proto3" json:"respBody,omitempty"`
	// gateway will send a heartBeat if it hasn't received SyncRPCRequests from cloud for a while.
	// If it's a heartbeat, reqId and respBody will be ignored.
	HeartBeat            bool     `protobuf:"varint,3,opt,name=heartBeat,proto3" json:"heartBeat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRPCResponse) Reset()         { *m = SyncRPCResponse{} }
func (m *SyncRPCResponse) String() string { return proto.CompactTextString(m) }
func (*SyncRPCResponse) ProtoMessage()    {}
func (*SyncRPCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22887391e8a5ac6c, []int{3}
}

func (m *SyncRPCResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRPCResponse.Unmarshal(m, b)
}
func (m *SyncRPCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRPCResponse.Marshal(b, m, deterministic)
}
func (m *SyncRPCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRPCResponse.Merge(m, src)
}
func (m *SyncRPCResponse) XXX_Size() int {
	return xxx_messageInfo_SyncRPCResponse.Size(m)
}
func (m *SyncRPCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRPCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRPCResponse proto.InternalMessageInfo

func (m *SyncRPCResponse) GetReqId() uint32 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *SyncRPCResponse) GetRespBody() *GatewayResponse {
	if m != nil {
		return m.RespBody
	}
	return nil
}

func (m *SyncRPCResponse) GetHeartBeat() bool {
	if m != nil {
		return m.HeartBeat
	}
	return false
}

func init() {
	proto.RegisterType((*GatewayRequest)(nil), "magma.orc8r.GatewayRequest")
	proto.RegisterMapType((map[string]string)(nil), "magma.orc8r.GatewayRequest.HeadersEntry")
	proto.RegisterType((*GatewayResponse)(nil), "magma.orc8r.GatewayResponse")
	proto.RegisterMapType((map[string]string)(nil), "magma.orc8r.GatewayResponse.HeadersEntry")
	proto.RegisterType((*SyncRPCRequest)(nil), "magma.orc8r.SyncRPCRequest")
	proto.RegisterType((*SyncRPCResponse)(nil), "magma.orc8r.SyncRPCResponse")
}

func init() {
	proto.RegisterFile("orc8r/protos/sync_rpc_service.proto", fileDescriptor_22887391e8a5ac6c)
}

var fileDescriptor_22887391e8a5ac6c = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcb, 0x6e, 0xda, 0x40,
	0x14, 0xed, 0x40, 0x78, 0xe4, 0x92, 0x92, 0x6a, 0x14, 0x45, 0x6e, 0x88, 0x2a, 0x44, 0xa5, 0xca,
	0xdd, 0x98, 0x8a, 0xaa, 0x12, 0xca, 0xae, 0xa0, 0xa8, 0x8f, 0x55, 0x35, 0x59, 0xb5, 0x9b, 0x68,
	0xb0, 0xaf, 0xb0, 0x15, 0xe3, 0x31, 0x33, 0x03, 0x91, 0x57, 0xfd, 0x90, 0x7e, 0x48, 0x7f, 0xab,
	0x5f, 0x50, 0x55, 0x9e, 0xb1, 0xc1, 0x44, 0x88, 0x2e, 0x58, 0x31, 0xf7, 0xcc, 0x7d, 0x9c, 0x73,
	0x0f, 0x1e, 0x78, 0x2d, 0xa4, 0x3f, 0x96, 0xc3, 0x54, 0x0a, 0x2d, 0xd4, 0x50, 0x65, 0x89, 0x7f,
	0x2f, 0x53, 0xff, 0x5e, 0xa1, 0x5c, 0x47, 0x3e, 0x7a, 0x06, 0xa7, 0x9d, 0x05, 0x9f, 0x2f, 0xb8,
	0x67, 0x52, 0x07, 0x7f, 0x08, 0x74, 0x3f, 0x71, 0x8d, 0x8f, 0x3c, 0x63, 0xb8, 0x5c, 0xa1, 0xd2,
	0x94, 0xc2, 0xc9, 0xfc, 0xf1, 0x4b, 0xe0, 0x90, 0x3e, 0x71, 0x4f, 0x99, 0x39, 0xd3, 0x6b, 0x38,
	0xe5, 0x2b, 0x1d, 0x0a, 0x19, 0xe9, 0xcc, 0xa9, 0x99, 0x8b, 0x2d, 0x90, 0x57, 0xa4, 0x5c, 0x87,
	0x4e, 0xdd, 0x56, 0xe4, 0x67, 0x3a, 0x81, 0x56, 0x88, 0x3c, 0x40, 0xa9, 0x9c, 0x46, 0xbf, 0xee,
	0x76, 0x46, 0xae, 0x57, 0x99, 0xeb, 0xed, 0xce, 0xf4, 0x3e, 0xdb, 0xd4, 0xdb, 0x44, 0xcb, 0x8c,
	0x95, 0x85, 0xd4, 0x81, 0x56, 0xca, 0xb3, 0x58, 0xf0, 0xc0, 0x69, 0xf6, 0x89, 0x7b, 0xc6, 0xca,
	0xf0, 0xea, 0x06, 0xce, 0xaa, 0x25, 0xf4, 0x05, 0xd4, 0x1f, 0x30, 0x2b, 0x28, 0xe7, 0x47, 0x7a,
	0x01, 0x8d, 0x35, 0x8f, 0x57, 0x58, 0xb0, 0xb5, 0xc1, 0x4d, 0x6d, 0x4c, 0x06, 0x7f, 0x09, 0x9c,
	0x6f, 0xc6, 0xab, 0x54, 0x24, 0x0a, 0xe9, 0x25, 0x34, 0x95, 0xe6, 0x7a, 0xa5, 0x8a, 0x16, 0x45,
	0x44, 0xa7, 0x5b, 0x15, 0x35, 0xa3, 0xe2, 0xed, 0x7e, 0x15, 0xb6, 0xcd, 0xff, 0x65, 0xd4, 0x77,
	0x64, 0xe4, 0xb4, 0x51, 0x4a, 0xe7, 0xc4, 0xd2, 0x46, 0x29, 0xe9, 0x1b, 0xe8, 0x3e, 0x20, 0xa6,
	0x53, 0x91, 0x24, 0x1f, 0x7d, 0x1d, 0xad, 0xd1, 0x69, 0xf4, 0x89, 0xdb, 0x66, 0x4f, 0xd0, 0xa3,
	0x16, 0xf0, 0x8b, 0x40, 0xf7, 0x2e, 0x4b, 0x7c, 0xf6, 0x6d, 0x5a, 0x7a, 0x7e, 0x01, 0x0d, 0x89,
	0xcb, 0xc2, 0xf4, 0xe7, 0xcc, 0x06, 0xf4, 0x03, 0xb4, 0x24, 0x2e, 0x27, 0x22, 0xb0, 0x9e, 0x77,
	0x46, 0xbd, 0x03, 0x1e, 0xb2, 0x32, 0x37, 0xff, 0xb3, 0x84, 0xc8, 0xa5, 0x9e, 0x20, 0xd7, 0x46,
	0x71, 0x9b, 0x6d, 0x01, 0xfa, 0x0a, 0xc0, 0x17, 0x49, 0x32, 0x8d, 0x85, 0xc2, 0xc0, 0x48, 0x6f,
	0xb3, 0x0a, 0x32, 0xf8, 0x09, 0xe7, 0x1b, 0x72, 0x85, 0x3b, 0xfb, 0xd9, 0x8d, 0xa1, 0x2d, 0x51,
	0xa5, 0x15, 0x7a, 0xd7, 0x87, 0xcc, 0x61, 0x9b, 0xec, 0xc3, 0x04, 0x47, 0xbf, 0xb7, 0xeb, 0xb9,
	0xb3, 0x1f, 0x0e, 0xfd, 0x0e, 0x97, 0xb7, 0x4a, 0xf3, 0x59, 0x1c, 0xa9, 0xb0, 0xbc, 0xd2, 0x12,
	0xf9, 0x82, 0xee, 0x8e, 0x7c, 0x42, 0xfc, 0xaa, 0xb7, 0xff, 0xd6, 0xec, 0x6b, 0xf0, 0xcc, 0x25,
	0xef, 0x08, 0xfd, 0x0a, 0xad, 0x02, 0x3f, 0xba, 0xd7, 0xa4, 0xf7, 0xe3, 0xa5, 0xc9, 0x19, 0xda,
	0x67, 0x20, 0x8e, 0x66, 0xc3, 0xb9, 0x28, 0x5e, 0x83, 0x59, 0xd3, 0xfc, 0xbe, 0xff, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x16, 0x73, 0xb2, 0x1e, 0x24, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SyncRPCServiceClient is the client API for SyncRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncRPCServiceClient interface {
	// creates a bidirectional stream from gateway to cloud
	// so cloud can send in SyncRPCRequest, and wait for SyncRPCResponse.
	// This will be the underlying service for Synchronous RPC from the cloud.
	EstablishSyncRPCStream(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_EstablishSyncRPCStreamClient, error)
	// same method as EstablishSyncRPCStream, but named differently for backward compatibility
	SyncRPC(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_SyncRPCClient, error)
}

type syncRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncRPCServiceClient(cc grpc.ClientConnInterface) SyncRPCServiceClient {
	return &syncRPCServiceClient{cc}
}

func (c *syncRPCServiceClient) EstablishSyncRPCStream(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_EstablishSyncRPCStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SyncRPCService_serviceDesc.Streams[0], "/magma.orc8r.SyncRPCService/EstablishSyncRPCStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncRPCServiceEstablishSyncRPCStreamClient{stream}
	return x, nil
}

type SyncRPCService_EstablishSyncRPCStreamClient interface {
	Send(*SyncRPCResponse) error
	Recv() (*SyncRPCRequest, error)
	grpc.ClientStream
}

type syncRPCServiceEstablishSyncRPCStreamClient struct {
	grpc.ClientStream
}

func (x *syncRPCServiceEstablishSyncRPCStreamClient) Send(m *SyncRPCResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *syncRPCServiceEstablishSyncRPCStreamClient) Recv() (*SyncRPCRequest, error) {
	m := new(SyncRPCRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *syncRPCServiceClient) SyncRPC(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_SyncRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SyncRPCService_serviceDesc.Streams[1], "/magma.orc8r.SyncRPCService/SyncRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncRPCServiceSyncRPCClient{stream}
	return x, nil
}

type SyncRPCService_SyncRPCClient interface {
	Send(*SyncRPCResponse) error
	Recv() (*SyncRPCRequest, error)
	grpc.ClientStream
}

type syncRPCServiceSyncRPCClient struct {
	grpc.ClientStream
}

func (x *syncRPCServiceSyncRPCClient) Send(m *SyncRPCResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *syncRPCServiceSyncRPCClient) Recv() (*SyncRPCRequest, error) {
	m := new(SyncRPCRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SyncRPCServiceServer is the server API for SyncRPCService service.
type SyncRPCServiceServer interface {
	// creates a bidirectional stream from gateway to cloud
	// so cloud can send in SyncRPCRequest, and wait for SyncRPCResponse.
	// This will be the underlying service for Synchronous RPC from the cloud.
	EstablishSyncRPCStream(SyncRPCService_EstablishSyncRPCStreamServer) error
	// same method as EstablishSyncRPCStream, but named differently for backward compatibility
	SyncRPC(SyncRPCService_SyncRPCServer) error
}

// UnimplementedSyncRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSyncRPCServiceServer struct {
}

func (*UnimplementedSyncRPCServiceServer) EstablishSyncRPCStream(srv SyncRPCService_EstablishSyncRPCStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EstablishSyncRPCStream not implemented")
}
func (*UnimplementedSyncRPCServiceServer) SyncRPC(srv SyncRPCService_SyncRPCServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncRPC not implemented")
}

func RegisterSyncRPCServiceServer(s *grpc.Server, srv SyncRPCServiceServer) {
	s.RegisterService(&_SyncRPCService_serviceDesc, srv)
}

func _SyncRPCService_EstablishSyncRPCStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SyncRPCServiceServer).EstablishSyncRPCStream(&syncRPCServiceEstablishSyncRPCStreamServer{stream})
}

type SyncRPCService_EstablishSyncRPCStreamServer interface {
	Send(*SyncRPCRequest) error
	Recv() (*SyncRPCResponse, error)
	grpc.ServerStream
}

type syncRPCServiceEstablishSyncRPCStreamServer struct {
	grpc.ServerStream
}

func (x *syncRPCServiceEstablishSyncRPCStreamServer) Send(m *SyncRPCRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *syncRPCServiceEstablishSyncRPCStreamServer) Recv() (*SyncRPCResponse, error) {
	m := new(SyncRPCResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SyncRPCService_SyncRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SyncRPCServiceServer).SyncRPC(&syncRPCServiceSyncRPCServer{stream})
}

type SyncRPCService_SyncRPCServer interface {
	Send(*SyncRPCRequest) error
	Recv() (*SyncRPCResponse, error)
	grpc.ServerStream
}

type syncRPCServiceSyncRPCServer struct {
	grpc.ServerStream
}

func (x *syncRPCServiceSyncRPCServer) Send(m *SyncRPCRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *syncRPCServiceSyncRPCServer) Recv() (*SyncRPCResponse, error) {
	m := new(SyncRPCResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SyncRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.SyncRPCService",
	HandlerType: (*SyncRPCServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EstablishSyncRPCStream",
			Handler:       _SyncRPCService_EstablishSyncRPCStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SyncRPC",
			Handler:       _SyncRPCService_SyncRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orc8r/protos/sync_rpc_service.proto",
}