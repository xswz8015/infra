// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/trove/api/testresults/api.proto

package testresults

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

type CollectTestResultsRequest struct {
	// Isolate identifies the isolated output to collect results from.
	Isolate *CollectTestResultsRequest_Isolate `protobuf:"bytes,1,opt,name=isolate,proto3" json:"isolate,omitempty"`
	// TODO: handle local workstation, FindIt, possibly others.
	//
	// Types that are valid to be assigned to Build:
	//	*CollectTestResultsRequest_Buildbucket_
	//	*CollectTestResultsRequest_Buildbot_
	Build isCollectTestResultsRequest_Build `protobuf_oneof:"build"`
	// Step is the step name to collect results for.
	Step                 string   `protobuf:"bytes,4,opt,name=step,proto3" json:"step,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectTestResultsRequest) Reset()         { *m = CollectTestResultsRequest{} }
func (m *CollectTestResultsRequest) String() string { return proto.CompactTextString(m) }
func (*CollectTestResultsRequest) ProtoMessage()    {}
func (*CollectTestResultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37344750934b981, []int{0}
}

func (m *CollectTestResultsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectTestResultsRequest.Unmarshal(m, b)
}
func (m *CollectTestResultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectTestResultsRequest.Marshal(b, m, deterministic)
}
func (m *CollectTestResultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectTestResultsRequest.Merge(m, src)
}
func (m *CollectTestResultsRequest) XXX_Size() int {
	return xxx_messageInfo_CollectTestResultsRequest.Size(m)
}
func (m *CollectTestResultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectTestResultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectTestResultsRequest proto.InternalMessageInfo

func (m *CollectTestResultsRequest) GetIsolate() *CollectTestResultsRequest_Isolate {
	if m != nil {
		return m.Isolate
	}
	return nil
}

type isCollectTestResultsRequest_Build interface {
	isCollectTestResultsRequest_Build()
}

type CollectTestResultsRequest_Buildbucket_ struct {
	Buildbucket *CollectTestResultsRequest_Buildbucket `protobuf:"bytes,2,opt,name=buildbucket,proto3,oneof"`
}

type CollectTestResultsRequest_Buildbot_ struct {
	Buildbot *CollectTestResultsRequest_Buildbot `protobuf:"bytes,3,opt,name=buildbot,proto3,oneof"`
}

func (*CollectTestResultsRequest_Buildbucket_) isCollectTestResultsRequest_Build() {}

func (*CollectTestResultsRequest_Buildbot_) isCollectTestResultsRequest_Build() {}

func (m *CollectTestResultsRequest) GetBuild() isCollectTestResultsRequest_Build {
	if m != nil {
		return m.Build
	}
	return nil
}

func (m *CollectTestResultsRequest) GetBuildbucket() *CollectTestResultsRequest_Buildbucket {
	if x, ok := m.GetBuild().(*CollectTestResultsRequest_Buildbucket_); ok {
		return x.Buildbucket
	}
	return nil
}

func (m *CollectTestResultsRequest) GetBuildbot() *CollectTestResultsRequest_Buildbot {
	if x, ok := m.GetBuild().(*CollectTestResultsRequest_Buildbot_); ok {
		return x.Buildbot
	}
	return nil
}

func (m *CollectTestResultsRequest) GetStep() string {
	if m != nil {
		return m.Step
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CollectTestResultsRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CollectTestResultsRequest_Buildbucket_)(nil),
		(*CollectTestResultsRequest_Buildbot_)(nil),
	}
}

type CollectTestResultsRequest_Isolate struct {
	Host      string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// One hash for each shard.
	Hash                 []string `protobuf:"bytes,3,rep,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectTestResultsRequest_Isolate) Reset()         { *m = CollectTestResultsRequest_Isolate{} }
func (m *CollectTestResultsRequest_Isolate) String() string { return proto.CompactTextString(m) }
func (*CollectTestResultsRequest_Isolate) ProtoMessage()    {}
func (*CollectTestResultsRequest_Isolate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37344750934b981, []int{0, 0}
}

func (m *CollectTestResultsRequest_Isolate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectTestResultsRequest_Isolate.Unmarshal(m, b)
}
func (m *CollectTestResultsRequest_Isolate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectTestResultsRequest_Isolate.Marshal(b, m, deterministic)
}
func (m *CollectTestResultsRequest_Isolate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectTestResultsRequest_Isolate.Merge(m, src)
}
func (m *CollectTestResultsRequest_Isolate) XXX_Size() int {
	return xxx_messageInfo_CollectTestResultsRequest_Isolate.Size(m)
}
func (m *CollectTestResultsRequest_Isolate) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectTestResultsRequest_Isolate.DiscardUnknown(m)
}

var xxx_messageInfo_CollectTestResultsRequest_Isolate proto.InternalMessageInfo

func (m *CollectTestResultsRequest_Isolate) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *CollectTestResultsRequest_Isolate) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CollectTestResultsRequest_Isolate) GetHash() []string {
	if m != nil {
		return m.Hash
	}
	return nil
}

// Buildbot specifies a buildbot build.
type CollectTestResultsRequest_Buildbot struct {
	Master               string   `protobuf:"bytes,1,opt,name=master,proto3" json:"master,omitempty"`
	Builder              string   `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
	BuildNumber          int64    `protobuf:"varint,3,opt,name=build_number,json=buildNumber,proto3" json:"build_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectTestResultsRequest_Buildbot) Reset()         { *m = CollectTestResultsRequest_Buildbot{} }
func (m *CollectTestResultsRequest_Buildbot) String() string { return proto.CompactTextString(m) }
func (*CollectTestResultsRequest_Buildbot) ProtoMessage()    {}
func (*CollectTestResultsRequest_Buildbot) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37344750934b981, []int{0, 1}
}

func (m *CollectTestResultsRequest_Buildbot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectTestResultsRequest_Buildbot.Unmarshal(m, b)
}
func (m *CollectTestResultsRequest_Buildbot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectTestResultsRequest_Buildbot.Marshal(b, m, deterministic)
}
func (m *CollectTestResultsRequest_Buildbot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectTestResultsRequest_Buildbot.Merge(m, src)
}
func (m *CollectTestResultsRequest_Buildbot) XXX_Size() int {
	return xxx_messageInfo_CollectTestResultsRequest_Buildbot.Size(m)
}
func (m *CollectTestResultsRequest_Buildbot) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectTestResultsRequest_Buildbot.DiscardUnknown(m)
}

var xxx_messageInfo_CollectTestResultsRequest_Buildbot proto.InternalMessageInfo

func (m *CollectTestResultsRequest_Buildbot) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *CollectTestResultsRequest_Buildbot) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func (m *CollectTestResultsRequest_Buildbot) GetBuildNumber() int64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

// Buildbucket specifies a buildbucket build.
type CollectTestResultsRequest_Buildbucket struct {
	BuildId              int64    `protobuf:"varint,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectTestResultsRequest_Buildbucket) Reset()         { *m = CollectTestResultsRequest_Buildbucket{} }
func (m *CollectTestResultsRequest_Buildbucket) String() string { return proto.CompactTextString(m) }
func (*CollectTestResultsRequest_Buildbucket) ProtoMessage()    {}
func (*CollectTestResultsRequest_Buildbucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37344750934b981, []int{0, 2}
}

func (m *CollectTestResultsRequest_Buildbucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectTestResultsRequest_Buildbucket.Unmarshal(m, b)
}
func (m *CollectTestResultsRequest_Buildbucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectTestResultsRequest_Buildbucket.Marshal(b, m, deterministic)
}
func (m *CollectTestResultsRequest_Buildbucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectTestResultsRequest_Buildbucket.Merge(m, src)
}
func (m *CollectTestResultsRequest_Buildbucket) XXX_Size() int {
	return xxx_messageInfo_CollectTestResultsRequest_Buildbucket.Size(m)
}
func (m *CollectTestResultsRequest_Buildbucket) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectTestResultsRequest_Buildbucket.DiscardUnknown(m)
}

var xxx_messageInfo_CollectTestResultsRequest_Buildbucket proto.InternalMessageInfo

func (m *CollectTestResultsRequest_Buildbucket) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

type CollectTestResultsResponse struct {
	NumResultsCollected  int64    `protobuf:"varint,1,opt,name=num_results_collected,json=numResultsCollected,proto3" json:"num_results_collected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectTestResultsResponse) Reset()         { *m = CollectTestResultsResponse{} }
func (m *CollectTestResultsResponse) String() string { return proto.CompactTextString(m) }
func (*CollectTestResultsResponse) ProtoMessage()    {}
func (*CollectTestResultsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37344750934b981, []int{1}
}

func (m *CollectTestResultsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectTestResultsResponse.Unmarshal(m, b)
}
func (m *CollectTestResultsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectTestResultsResponse.Marshal(b, m, deterministic)
}
func (m *CollectTestResultsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectTestResultsResponse.Merge(m, src)
}
func (m *CollectTestResultsResponse) XXX_Size() int {
	return xxx_messageInfo_CollectTestResultsResponse.Size(m)
}
func (m *CollectTestResultsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectTestResultsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CollectTestResultsResponse proto.InternalMessageInfo

func (m *CollectTestResultsResponse) GetNumResultsCollected() int64 {
	if m != nil {
		return m.NumResultsCollected
	}
	return 0
}

func init() {
	proto.RegisterType((*CollectTestResultsRequest)(nil), "testresults.CollectTestResultsRequest")
	proto.RegisterType((*CollectTestResultsRequest_Isolate)(nil), "testresults.CollectTestResultsRequest.Isolate")
	proto.RegisterType((*CollectTestResultsRequest_Buildbot)(nil), "testresults.CollectTestResultsRequest.Buildbot")
	proto.RegisterType((*CollectTestResultsRequest_Buildbucket)(nil), "testresults.CollectTestResultsRequest.Buildbucket")
	proto.RegisterType((*CollectTestResultsResponse)(nil), "testresults.CollectTestResultsResponse")
}

func init() {
	proto.RegisterFile("infra/appengine/trove/api/testresults/api.proto", fileDescriptor_d37344750934b981)
}

var fileDescriptor_d37344750934b981 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xa3, 0x30,
	0x10, 0xc6, 0x93, 0x25, 0x1b, 0xc2, 0xb0, 0x27, 0xaf, 0x76, 0x45, 0x50, 0x0f, 0x69, 0x0e, 0x6d,
	0x4e, 0x20, 0xd1, 0x37, 0x48, 0x2e, 0xc9, 0xa1, 0x7f, 0x84, 0xaa, 0x9e, 0x2a, 0x21, 0x43, 0xa6,
	0x0d, 0x0a, 0xd8, 0x14, 0x9b, 0xbe, 0x41, 0xdf, 0xbb, 0xca, 0x60, 0x1a, 0x0e, 0xad, 0x9a, 0xde,
	0x66, 0xbe, 0xf1, 0xf7, 0xb3, 0xe7, 0x03, 0x08, 0x73, 0xf1, 0x54, 0xf3, 0x90, 0x57, 0x15, 0x8a,
	0xe7, 0x5c, 0x60, 0xa8, 0x6b, 0xf9, 0x8a, 0x21, 0xaf, 0xf2, 0x50, 0xa3, 0xd2, 0x35, 0xaa, 0xa6,
	0xd0, 0xea, 0xd0, 0x07, 0x55, 0x2d, 0xb5, 0x64, 0x6e, 0x4f, 0x9e, 0xbf, 0x8d, 0x60, 0xba, 0x92,
	0x45, 0x81, 0x99, 0xbe, 0x47, 0xa5, 0xe3, 0x56, 0x8e, 0xf1, 0xa5, 0x41, 0xa5, 0xd9, 0x1a, 0xec,
	0x5c, 0xc9, 0x82, 0x6b, 0xf4, 0x86, 0xb3, 0xe1, 0xc2, 0x8d, 0x82, 0xa0, 0x67, 0x0e, 0xbe, 0x34,
	0x06, 0x9b, 0xd6, 0x15, 0x77, 0x76, 0xf6, 0x00, 0x6e, 0xda, 0xe4, 0xc5, 0x36, 0x6d, 0xb2, 0x3d,
	0x6a, 0xef, 0x17, 0xd1, 0xa2, 0x13, 0x69, 0xcb, 0xa3, 0x73, 0x3d, 0x88, 0xfb, 0x20, 0x76, 0x0d,
	0x93, 0xb6, 0x95, 0xda, 0xb3, 0x08, 0x1a, 0xfe, 0x08, 0x2a, 0x0f, 0xc4, 0x0f, 0x04, 0x63, 0x30,
	0x52, 0x1a, 0x2b, 0x6f, 0x34, 0x1b, 0x2e, 0x9c, 0x98, 0x6a, 0xff, 0x16, 0x6c, 0xb3, 0xce, 0x61,
	0xbc, 0x93, 0x4a, 0x53, 0x18, 0x4e, 0x4c, 0x35, 0x3b, 0x03, 0x47, 0xf0, 0x12, 0x55, 0xc5, 0x33,
	0xa4, 0xbd, 0x9c, 0xf8, 0x28, 0x90, 0x83, 0xab, 0x9d, 0x67, 0xcd, 0x2c, 0x72, 0x70, 0xb5, 0xf3,
	0x13, 0x98, 0x74, 0x97, 0xb3, 0xff, 0x30, 0x2e, 0xb9, 0xd2, 0x58, 0x1b, 0xa6, 0xe9, 0x98, 0x07,
	0x36, 0x3d, 0x0a, 0x6b, 0xc3, 0xec, 0x5a, 0x76, 0x0e, 0x7f, 0xa8, 0x4c, 0x44, 0x53, 0xa6, 0x58,
	0xd3, 0xd6, 0x96, 0x09, 0xe5, 0x86, 0x24, 0x7f, 0x01, 0x6e, 0x2f, 0x32, 0x36, 0x35, 0x19, 0x25,
	0xf9, 0x96, 0x6e, 0xb1, 0x0c, 0x6c, 0xb3, 0x5d, 0xda, 0xf0, 0x9b, 0xca, 0xf9, 0x1d, 0xf8, 0x9f,
	0x45, 0xa5, 0x2a, 0x29, 0x14, 0xb2, 0x08, 0xfe, 0x89, 0xa6, 0x4c, 0x4c, 0xa8, 0x49, 0xd6, 0x9e,
	0xc4, 0x0e, 0xf7, 0x57, 0x34, 0xa5, 0xb1, 0xac, 0xba, 0x51, 0xb4, 0x07, 0xb7, 0x87, 0x62, 0x8f,
	0x60, 0x9b, 0x19, 0xbb, 0x38, 0xed, 0x0b, 0xf9, 0x97, 0xdf, 0x9e, 0x6b, 0x9f, 0x37, 0x1f, 0xa4,
	0x63, 0xfa, 0xb5, 0xaf, 0xde, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xc7, 0x9d, 0x9f, 0x0d, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestResultsClient is the client API for TestResults service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestResultsClient interface {
	Collect(ctx context.Context, in *CollectTestResultsRequest, opts ...grpc.CallOption) (*CollectTestResultsResponse, error)
}
type testResultsPRPCClient struct {
	client *prpc.Client
}

func NewTestResultsPRPCClient(client *prpc.Client) TestResultsClient {
	return &testResultsPRPCClient{client}
}

func (c *testResultsPRPCClient) Collect(ctx context.Context, in *CollectTestResultsRequest, opts ...grpc.CallOption) (*CollectTestResultsResponse, error) {
	out := new(CollectTestResultsResponse)
	err := c.client.Call(ctx, "testresults.TestResults", "Collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type testResultsClient struct {
	cc grpc.ClientConnInterface
}

func NewTestResultsClient(cc grpc.ClientConnInterface) TestResultsClient {
	return &testResultsClient{cc}
}

func (c *testResultsClient) Collect(ctx context.Context, in *CollectTestResultsRequest, opts ...grpc.CallOption) (*CollectTestResultsResponse, error) {
	out := new(CollectTestResultsResponse)
	err := c.cc.Invoke(ctx, "/testresults.TestResults/Collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestResultsServer is the server API for TestResults service.
type TestResultsServer interface {
	Collect(context.Context, *CollectTestResultsRequest) (*CollectTestResultsResponse, error)
}

// UnimplementedTestResultsServer can be embedded to have forward compatible implementations.
type UnimplementedTestResultsServer struct {
}

func (*UnimplementedTestResultsServer) Collect(ctx context.Context, req *CollectTestResultsRequest) (*CollectTestResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterTestResultsServer(s prpc.Registrar, srv TestResultsServer) {
	s.RegisterService(&_TestResults_serviceDesc, srv)
}

func _TestResults_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectTestResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestResultsServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testresults.TestResults/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestResultsServer).Collect(ctx, req.(*CollectTestResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestResults_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testresults.TestResults",
	HandlerType: (*TestResultsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Collect",
			Handler:    _TestResults_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/appengine/trove/api/testresults/api.proto",
}
