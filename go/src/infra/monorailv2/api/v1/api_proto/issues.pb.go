// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/api_proto/issues.proto

package monorail_v1

import prpc "go.chromium.org/luci/grpc/prpc"

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

// Next available tag: 2
type GetIssueRequest struct {
	// The name of the issue to request, see Issue.name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIssueRequest) Reset()         { *m = GetIssueRequest{} }
func (m *GetIssueRequest) String() string { return proto.CompactTextString(m) }
func (*GetIssueRequest) ProtoMessage()    {}
func (*GetIssueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da7675e39ac6c753, []int{0}
}

func (m *GetIssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIssueRequest.Unmarshal(m, b)
}
func (m *GetIssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIssueRequest.Marshal(b, m, deterministic)
}
func (m *GetIssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIssueRequest.Merge(m, src)
}
func (m *GetIssueRequest) XXX_Size() int {
	return xxx_messageInfo_GetIssueRequest.Size(m)
}
func (m *GetIssueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIssueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIssueRequest proto.InternalMessageInfo

func (m *GetIssueRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetIssueRequest)(nil), "monorail.v1.GetIssueRequest")
}

func init() { proto.RegisterFile("api/v1/api_proto/issues.proto", fileDescriptor_da7675e39ac6c753) }

var fileDescriptor_da7675e39ac6c753 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0xc8, 0xd4,
	0x2f, 0x33, 0xd4, 0x4f, 0x2c, 0xc8, 0x8c, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x2c, 0x2e,
	0x2e, 0x4d, 0x2d, 0xd6, 0x03, 0x73, 0x84, 0xb8, 0x73, 0xf3, 0xf3, 0xf2, 0x8b, 0x12, 0x33, 0x73,
	0xf4, 0xca, 0x0c, 0xa5, 0x54, 0xb0, 0xab, 0x8d, 0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0x81, 0x6a,
	0x51, 0x52, 0xe5, 0xe2, 0x77, 0x4f, 0x2d, 0xf1, 0x04, 0xc9, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x81, 0xd9, 0x46, 0x1e, 0x5c, 0x6c, 0x60, 0x35, 0xc5, 0x42, 0x76, 0x5c, 0x1c, 0x30, 0x0d, 0x42,
	0x32, 0x7a, 0x48, 0x16, 0xea, 0xa1, 0x99, 0x23, 0x25, 0x84, 0x22, 0x0b, 0x96, 0x52, 0x62, 0x48,
	0x62, 0x03, 0xdb, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x02, 0x5c, 0xe5, 0xda, 0xcb, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IssuesClient is the client API for Issues service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IssuesClient interface {
	GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*Issue, error)
}
type issuesPRPCClient struct {
	client *prpc.Client
}

func NewIssuesPRPCClient(client *prpc.Client) IssuesClient {
	return &issuesPRPCClient{client}
}

func (c *issuesPRPCClient) GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*Issue, error) {
	out := new(Issue)
	err := c.client.Call(ctx, "monorail.v1.Issues", "GetIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type issuesClient struct {
	cc *grpc.ClientConn
}

func NewIssuesClient(cc *grpc.ClientConn) IssuesClient {
	return &issuesClient{cc}
}

func (c *issuesClient) GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*Issue, error) {
	out := new(Issue)
	err := c.cc.Invoke(ctx, "/monorail.v1.Issues/GetIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssuesServer is the server API for Issues service.
type IssuesServer interface {
	GetIssue(context.Context, *GetIssueRequest) (*Issue, error)
}

// UnimplementedIssuesServer can be embedded to have forward compatible implementations.
type UnimplementedIssuesServer struct {
}

func (*UnimplementedIssuesServer) GetIssue(ctx context.Context, req *GetIssueRequest) (*Issue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIssue not implemented")
}

func RegisterIssuesServer(s prpc.Registrar, srv IssuesServer) {
	s.RegisterService(&_Issues_serviceDesc, srv)
}

func _Issues_GetIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuesServer).GetIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v1.Issues/GetIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuesServer).GetIssue(ctx, req.(*GetIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Issues_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monorail.v1.Issues",
	HandlerType: (*IssuesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIssue",
			Handler:    _Issues_GetIssue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api_proto/issues.proto",
}
