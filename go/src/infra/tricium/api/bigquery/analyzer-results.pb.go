// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/tricium/api/bigquery/analyzer-results.proto

package analyzer_results

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	v1 "infra/tricium/api/v1"
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

type AnalysisRun struct {
	// The revision information for the Gerrit change being analyzed by this
	// analysis run.
	GerritRevision *v1.GerritRevision `protobuf:"bytes,1,opt,name=gerrit_revision,json=gerritRevision,proto3" json:"gerrit_revision,omitempty"`
	// The revision number. In Gerrit this is the change revision and is
	// displayed as the patchset number in PolyGerrit.
	RevisionNumber int32 `protobuf:"varint,2,opt,name=revision_number,json=revisionNumber,proto3" json:"revision_number,omitempty"`
	// All files in the change revision analyzed by the run.
	Files []*v1.Data_File `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
	// Time when the request was received.
	RequestedTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=requested_time,json=requestedTime,proto3" json:"requested_time,omitempty"`
	// Platform for which the result applies.
	ResultPlatform v1.Platform_Name `protobuf:"varint,5,opt,name=result_platform,json=resultPlatform,proto3,enum=tricium.Platform_Name" json:"result_platform,omitempty"`
	// Overall state for the run result. As results are only sent after
	// completion PENDING and RUNNING would never be used.
	ResultState v1.State `protobuf:"varint,6,opt,name=result_state,json=resultState,proto3,enum=tricium.State" json:"result_state,omitempty"`
	// Comments added to Gerrit during this analysis run.
	Comments             []*AnalysisRun_GerritComment `protobuf:"bytes,7,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *AnalysisRun) Reset()         { *m = AnalysisRun{} }
func (m *AnalysisRun) String() string { return proto.CompactTextString(m) }
func (*AnalysisRun) ProtoMessage()    {}
func (*AnalysisRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed1cac75aec4820e, []int{0}
}

func (m *AnalysisRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisRun.Unmarshal(m, b)
}
func (m *AnalysisRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisRun.Marshal(b, m, deterministic)
}
func (m *AnalysisRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisRun.Merge(m, src)
}
func (m *AnalysisRun) XXX_Size() int {
	return xxx_messageInfo_AnalysisRun.Size(m)
}
func (m *AnalysisRun) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisRun.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisRun proto.InternalMessageInfo

func (m *AnalysisRun) GetGerritRevision() *v1.GerritRevision {
	if m != nil {
		return m.GerritRevision
	}
	return nil
}

func (m *AnalysisRun) GetRevisionNumber() int32 {
	if m != nil {
		return m.RevisionNumber
	}
	return 0
}

func (m *AnalysisRun) GetFiles() []*v1.Data_File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *AnalysisRun) GetRequestedTime() *timestamp.Timestamp {
	if m != nil {
		return m.RequestedTime
	}
	return nil
}

func (m *AnalysisRun) GetResultPlatform() v1.Platform_Name {
	if m != nil {
		return m.ResultPlatform
	}
	return v1.Platform_ANY
}

func (m *AnalysisRun) GetResultState() v1.State {
	if m != nil {
		return m.ResultState
	}
	return v1.State_PENDING
}

func (m *AnalysisRun) GetComments() []*AnalysisRun_GerritComment {
	if m != nil {
		return m.Comments
	}
	return nil
}

type AnalysisRun_GerritComment struct {
	// The comment generated by the analysis run.
	Comment *v1.Data_Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	// Comment creation time.
	//
	// Comment creation time in terms of when it is tracked in the service not
	// when it is created by the analyzer.
	CreatedTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// Analyzer function name.
	//
	// e.g., "ClangTidy".
	Analyzer string `protobuf:"bytes,3,opt,name=analyzer,proto3" json:"analyzer,omitempty"`
	// Platforms this comment applies to.
	Platforms []v1.Platform_Name `protobuf:"varint,4,rep,packed,name=platforms,proto3,enum=tricium.Platform_Name" json:"platforms,omitempty"`
	// Has this comment been selected to be displayed on the review?
	Selected             bool     `protobuf:"varint,5,opt,name=selected,proto3" json:"selected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnalysisRun_GerritComment) Reset()         { *m = AnalysisRun_GerritComment{} }
func (m *AnalysisRun_GerritComment) String() string { return proto.CompactTextString(m) }
func (*AnalysisRun_GerritComment) ProtoMessage()    {}
func (*AnalysisRun_GerritComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed1cac75aec4820e, []int{0, 0}
}

func (m *AnalysisRun_GerritComment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisRun_GerritComment.Unmarshal(m, b)
}
func (m *AnalysisRun_GerritComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisRun_GerritComment.Marshal(b, m, deterministic)
}
func (m *AnalysisRun_GerritComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisRun_GerritComment.Merge(m, src)
}
func (m *AnalysisRun_GerritComment) XXX_Size() int {
	return xxx_messageInfo_AnalysisRun_GerritComment.Size(m)
}
func (m *AnalysisRun_GerritComment) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisRun_GerritComment.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisRun_GerritComment proto.InternalMessageInfo

func (m *AnalysisRun_GerritComment) GetComment() *v1.Data_Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *AnalysisRun_GerritComment) GetCreatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *AnalysisRun_GerritComment) GetAnalyzer() string {
	if m != nil {
		return m.Analyzer
	}
	return ""
}

func (m *AnalysisRun_GerritComment) GetPlatforms() []v1.Platform_Name {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func (m *AnalysisRun_GerritComment) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

func init() {
	proto.RegisterType((*AnalysisRun)(nil), "analyzer.results.AnalysisRun")
	proto.RegisterType((*AnalysisRun_GerritComment)(nil), "analyzer.results.AnalysisRun.GerritComment")
}

func init() {
	proto.RegisterFile("infra/tricium/api/bigquery/analyzer-results.proto", fileDescriptor_ed1cac75aec4820e)
}

var fileDescriptor_ed1cac75aec4820e = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0x75, 0xdd, 0x3a, 0x67, 0xcb, 0x90, 0x25, 0xc0, 0xca, 0xcd, 0xa2, 0x71, 0x41,
	0x24, 0x84, 0xad, 0x16, 0x6e, 0x11, 0x4c, 0x20, 0x76, 0x37, 0x21, 0xc3, 0x7d, 0xe5, 0xb6, 0xa7,
	0x91, 0xa5, 0x38, 0xe9, 0x6c, 0x67, 0xd2, 0x78, 0x0c, 0xde, 0x93, 0x77, 0x40, 0x76, 0x6c, 0x43,
	0x61, 0x62, 0x97, 0xe7, 0x3f, 0xff, 0x7f, 0x72, 0xce, 0xe7, 0xa0, 0xb9, 0xec, 0xb6, 0x5a, 0x30,
	0xab, 0xe5, 0x5a, 0x0e, 0x8a, 0x89, 0x9d, 0x64, 0x2b, 0xd9, 0xdc, 0x0e, 0xa0, 0xef, 0x99, 0xe8,
	0x44, 0x7b, 0xff, 0x1d, 0xf4, 0x6b, 0x0d, 0x66, 0x68, 0xad, 0xa1, 0x3b, 0xdd, 0xdb, 0x1e, 0x3f,
	0x89, 0x3a, 0x0d, 0x7a, 0x79, 0xf1, 0xef, 0x90, 0xbb, 0x39, 0xdb, 0x08, 0x2b, 0xc6, 0x48, 0xf9,
	0xe2, 0x41, 0xc3, 0xae, 0x15, 0x76, 0xdb, 0x6b, 0x15, 0x4c, 0x97, 0x0f, 0x9a, 0x42, 0x19, 0x3c,
	0x17, 0x4d, 0xdf, 0x37, 0x2d, 0x30, 0x5f, 0xad, 0x86, 0x2d, 0xb3, 0x52, 0x81, 0xb1, 0x42, 0xed,
	0x46, 0xc3, 0xe5, 0x8f, 0x29, 0xca, 0xaf, 0xdc, 0x7e, 0x46, 0x1a, 0x3e, 0x74, 0xf8, 0x03, 0x3a,
	0x6f, 0x40, 0x6b, 0x69, 0x97, 0x1a, 0xee, 0xa4, 0x91, 0x7d, 0x47, 0xb2, 0x2a, 0xab, 0xf3, 0xc5,
	0x73, 0x1a, 0x27, 0x5f, 0xfb, 0x3e, 0x0f, 0x6d, 0x5e, 0x34, 0x7b, 0x35, 0x7e, 0x89, 0xce, 0x63,
	0x74, 0xd9, 0x0d, 0x6a, 0x05, 0x9a, 0x1c, 0x54, 0x59, 0x3d, 0xe5, 0x45, 0x94, 0x6f, 0xbc, 0x8a,
	0x6b, 0x34, 0xdd, 0xca, 0x16, 0x0c, 0x99, 0x54, 0x93, 0x3a, 0x5f, 0xe0, 0xf4, 0x81, 0x4f, 0x0e,
	0xc4, 0x67, 0xd9, 0x02, 0x1f, 0x0d, 0xf8, 0x0a, 0x15, 0x1a, 0x6e, 0x07, 0x30, 0x16, 0x36, 0x4b,
	0x77, 0x01, 0x39, 0xf4, 0x3b, 0x95, 0x74, 0x3c, 0x8f, 0xc6, 0xf3, 0xe8, 0xb7, 0x78, 0x1e, 0x3f,
	0x4b, 0x09, 0xa7, 0xe1, 0xf7, 0x6e, 0x2b, 0x47, 0x7f, 0x19, 0x29, 0x92, 0x69, 0x95, 0xd5, 0xc5,
	0xe2, 0x59, 0xfa, 0xec, 0x97, 0x88, 0xf7, 0x46, 0x28, 0x70, 0xdb, 0x3a, 0x7b, 0x14, 0xf1, 0x1c,
	0x9d, 0x86, 0x01, 0xc6, 0x0a, 0x0b, 0xe4, 0xc8, 0xa7, 0x8b, 0x94, 0xfe, 0xea, 0x54, 0x9e, 0x8f,
	0x1e, 0x5f, 0xe0, 0x6b, 0x34, 0x5b, 0xf7, 0x4a, 0x41, 0x67, 0x0d, 0x39, 0xf6, 0x37, 0xbe, 0xa2,
	0x7f, 0xff, 0x0b, 0xf4, 0x0f, 0xf8, 0x81, 0xec, 0xc7, 0x31, 0xc3, 0x53, 0xb8, 0xfc, 0x99, 0xa1,
	0xb3, 0xbd, 0x1e, 0x66, 0xe8, 0x38, 0x74, 0xc3, 0xf3, 0x3c, 0xdd, 0xa7, 0x17, 0x67, 0x44, 0x17,
	0x7e, 0x87, 0x4e, 0xd7, 0x1a, 0x44, 0x02, 0x78, 0xf0, 0x28, 0xc0, 0x3c, 0xf8, 0x3d, 0xbe, 0x12,
	0xcd, 0xe2, 0xe6, 0x64, 0x52, 0x65, 0xf5, 0x09, 0x4f, 0x35, 0x7e, 0x8b, 0x4e, 0x22, 0x53, 0x43,
	0x0e, 0xab, 0xc9, 0x7f, 0xa0, 0xfe, 0x36, 0xba, 0x89, 0x06, 0x5a, 0x58, 0x5b, 0xd8, 0xf8, 0x97,
	0x98, 0xf1, 0x54, 0xaf, 0x8e, 0xfc, 0x3a, 0x6f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xff,
	0xc5, 0x29, 0x6d, 0x03, 0x00, 0x00,
}
