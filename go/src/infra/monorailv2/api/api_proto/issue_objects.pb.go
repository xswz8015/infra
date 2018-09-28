// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api_proto/issue_objects.proto

package monorail

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Next available tag: 8
type ApprovalStatus int32

const (
	ApprovalStatus_NOT_SET          ApprovalStatus = 0
	ApprovalStatus_NEEDS_REVIEW     ApprovalStatus = 1
	ApprovalStatus_NA               ApprovalStatus = 2
	ApprovalStatus_REVIEW_REQUESTED ApprovalStatus = 3
	ApprovalStatus_REVIEW_STARTED   ApprovalStatus = 4
	ApprovalStatus_NEED_INFO        ApprovalStatus = 5
	ApprovalStatus_APPROVED         ApprovalStatus = 6
	ApprovalStatus_NOT_APPROVED     ApprovalStatus = 7
)

var ApprovalStatus_name = map[int32]string{
	0: "NOT_SET",
	1: "NEEDS_REVIEW",
	2: "NA",
	3: "REVIEW_REQUESTED",
	4: "REVIEW_STARTED",
	5: "NEED_INFO",
	6: "APPROVED",
	7: "NOT_APPROVED",
}

var ApprovalStatus_value = map[string]int32{
	"NOT_SET":          0,
	"NEEDS_REVIEW":     1,
	"NA":               2,
	"REVIEW_REQUESTED": 3,
	"REVIEW_STARTED":   4,
	"NEED_INFO":        5,
	"APPROVED":         6,
	"NOT_APPROVED":     7,
}

func (x ApprovalStatus) String() string {
	return proto.EnumName(ApprovalStatus_name, int32(x))
}

func (ApprovalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{0}
}

// Next available tag: 7
type CannedQuery int32

const (
	CannedQuery_ALL       CannedQuery = 0
	CannedQuery_NEW       CannedQuery = 1
	CannedQuery_OPEN      CannedQuery = 2
	CannedQuery_OWNED     CannedQuery = 3
	CannedQuery_REPORTED  CannedQuery = 4
	CannedQuery_STARRED   CannedQuery = 5
	CannedQuery_TO_VERIFY CannedQuery = 6
)

var CannedQuery_name = map[int32]string{
	0: "ALL",
	1: "NEW",
	2: "OPEN",
	3: "OWNED",
	4: "REPORTED",
	5: "STARRED",
	6: "TO_VERIFY",
}

var CannedQuery_value = map[string]int32{
	"ALL":       0,
	"NEW":       1,
	"OPEN":      2,
	"OWNED":     3,
	"REPORTED":  4,
	"STARRED":   5,
	"TO_VERIFY": 6,
}

func (x CannedQuery) String() string {
	return proto.EnumName(CannedQuery_name, int32(x))
}

func (CannedQuery) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{1}
}

// Next available tag: 8
type Approval struct {
	FieldRef             *FieldRef      `protobuf:"bytes,1,opt,name=field_ref,json=fieldRef,proto3" json:"field_ref,omitempty"`
	ApproverRefs         []*UserRef     `protobuf:"bytes,2,rep,name=approver_refs,json=approverRefs,proto3" json:"approver_refs,omitempty"`
	Status               ApprovalStatus `protobuf:"varint,3,opt,name=status,proto3,enum=monorail.ApprovalStatus" json:"status,omitempty"`
	SetOn                uint32         `protobuf:"fixed32,4,opt,name=set_on,json=setOn,proto3" json:"set_on,omitempty"`
	SetterRef            *UserRef       `protobuf:"bytes,5,opt,name=setter_ref,json=setterRef,proto3" json:"setter_ref,omitempty"`
	PhaseRef             *PhaseRef      `protobuf:"bytes,7,opt,name=phase_ref,json=phaseRef,proto3" json:"phase_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Approval) Reset()         { *m = Approval{} }
func (m *Approval) String() string { return proto.CompactTextString(m) }
func (*Approval) ProtoMessage()    {}
func (*Approval) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{0}
}
func (m *Approval) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Approval.Unmarshal(m, b)
}
func (m *Approval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Approval.Marshal(b, m, deterministic)
}
func (m *Approval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Approval.Merge(m, src)
}
func (m *Approval) XXX_Size() int {
	return xxx_messageInfo_Approval.Size(m)
}
func (m *Approval) XXX_DiscardUnknown() {
	xxx_messageInfo_Approval.DiscardUnknown(m)
}

var xxx_messageInfo_Approval proto.InternalMessageInfo

func (m *Approval) GetFieldRef() *FieldRef {
	if m != nil {
		return m.FieldRef
	}
	return nil
}

func (m *Approval) GetApproverRefs() []*UserRef {
	if m != nil {
		return m.ApproverRefs
	}
	return nil
}

func (m *Approval) GetStatus() ApprovalStatus {
	if m != nil {
		return m.Status
	}
	return ApprovalStatus_NOT_SET
}

func (m *Approval) GetSetOn() uint32 {
	if m != nil {
		return m.SetOn
	}
	return 0
}

func (m *Approval) GetSetterRef() *UserRef {
	if m != nil {
		return m.SetterRef
	}
	return nil
}

func (m *Approval) GetPhaseRef() *PhaseRef {
	if m != nil {
		return m.PhaseRef
	}
	return nil
}

// This message is only suitable for displaying the amendment to users.
// We don't currently offer structured amendments that client code can
// reason about, field names can be ambiguous, and we don't have
// old_value for most changes.
// Next available tag: 4
type Amendment struct {
	// This may be the name of a built-in or custom field, or relative to
	// an approval field name.
	FieldName string `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	// This may be a new value that overwrote the old value, e.g., "Assigned",
	// or it may be a space-separated list of changes, e.g., "Size-L -Size-S".
	NewOrDeltaValue string `protobuf:"bytes,2,opt,name=new_or_delta_value,json=newOrDeltaValue,proto3" json:"new_or_delta_value,omitempty"`
	// old_value is only used when the user changes the summary.
	OldValue             string   `protobuf:"bytes,3,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Amendment) Reset()         { *m = Amendment{} }
func (m *Amendment) String() string { return proto.CompactTextString(m) }
func (*Amendment) ProtoMessage()    {}
func (*Amendment) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{1}
}
func (m *Amendment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Amendment.Unmarshal(m, b)
}
func (m *Amendment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Amendment.Marshal(b, m, deterministic)
}
func (m *Amendment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Amendment.Merge(m, src)
}
func (m *Amendment) XXX_Size() int {
	return xxx_messageInfo_Amendment.Size(m)
}
func (m *Amendment) XXX_DiscardUnknown() {
	xxx_messageInfo_Amendment.DiscardUnknown(m)
}

var xxx_messageInfo_Amendment proto.InternalMessageInfo

func (m *Amendment) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *Amendment) GetNewOrDeltaValue() string {
	if m != nil {
		return m.NewOrDeltaValue
	}
	return ""
}

func (m *Amendment) GetOldValue() string {
	if m != nil {
		return m.OldValue
	}
	return ""
}

// Next available tag: 9
type Attachment struct {
	AttachmentId         uint64   `protobuf:"varint,1,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	Filename             string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Size                 uint64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	ContentType          string   `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	IsDeleted            bool     `protobuf:"varint,5,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	ThumbnailUrl         string   `protobuf:"bytes,6,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	ViewUrl              string   `protobuf:"bytes,7,opt,name=view_url,json=viewUrl,proto3" json:"view_url,omitempty"`
	DownloadUrl          string   `protobuf:"bytes,8,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attachment) Reset()         { *m = Attachment{} }
func (m *Attachment) String() string { return proto.CompactTextString(m) }
func (*Attachment) ProtoMessage()    {}
func (*Attachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{2}
}
func (m *Attachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attachment.Unmarshal(m, b)
}
func (m *Attachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attachment.Marshal(b, m, deterministic)
}
func (m *Attachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attachment.Merge(m, src)
}
func (m *Attachment) XXX_Size() int {
	return xxx_messageInfo_Attachment.Size(m)
}
func (m *Attachment) XXX_DiscardUnknown() {
	xxx_messageInfo_Attachment.DiscardUnknown(m)
}

var xxx_messageInfo_Attachment proto.InternalMessageInfo

func (m *Attachment) GetAttachmentId() uint64 {
	if m != nil {
		return m.AttachmentId
	}
	return 0
}

func (m *Attachment) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Attachment) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Attachment) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Attachment) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *Attachment) GetThumbnailUrl() string {
	if m != nil {
		return m.ThumbnailUrl
	}
	return ""
}

func (m *Attachment) GetViewUrl() string {
	if m != nil {
		return m.ViewUrl
	}
	return ""
}

func (m *Attachment) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

// Next available tag: 14
type Comment struct {
	ProjectName    string        `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	LocalId        uint32        `protobuf:"varint,2,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	SequenceNum    uint32        `protobuf:"varint,3,opt,name=sequence_num,json=sequenceNum,proto3" json:"sequence_num,omitempty"`
	IsDeleted      bool          `protobuf:"varint,4,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	Commenter      *UserRef      `protobuf:"bytes,5,opt,name=commenter,proto3" json:"commenter,omitempty"`
	Timestamp      uint32        `protobuf:"fixed32,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Content        string        `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	InboundMessage string        `protobuf:"bytes,8,opt,name=inbound_message,json=inboundMessage,proto3" json:"inbound_message,omitempty"`
	Amendments     []*Amendment  `protobuf:"bytes,9,rep,name=amendments,proto3" json:"amendments,omitempty"`
	Attachments    []*Attachment `protobuf:"bytes,10,rep,name=attachments,proto3" json:"attachments,omitempty"`
	ApprovalRef    *FieldRef     `protobuf:"bytes,11,opt,name=approval_ref,json=approvalRef,proto3" json:"approval_ref,omitempty"`
	// If set, this comment is an issue description.
	DescriptionNum       uint32   `protobuf:"varint,12,opt,name=description_num,json=descriptionNum,proto3" json:"description_num,omitempty"`
	IsSpam               bool     `protobuf:"varint,13,opt,name=is_spam,json=isSpam,proto3" json:"is_spam,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{3}
}
func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *Comment) GetLocalId() uint32 {
	if m != nil {
		return m.LocalId
	}
	return 0
}

func (m *Comment) GetSequenceNum() uint32 {
	if m != nil {
		return m.SequenceNum
	}
	return 0
}

func (m *Comment) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *Comment) GetCommenter() *UserRef {
	if m != nil {
		return m.Commenter
	}
	return nil
}

func (m *Comment) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Comment) GetInboundMessage() string {
	if m != nil {
		return m.InboundMessage
	}
	return ""
}

func (m *Comment) GetAmendments() []*Amendment {
	if m != nil {
		return m.Amendments
	}
	return nil
}

func (m *Comment) GetAttachments() []*Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Comment) GetApprovalRef() *FieldRef {
	if m != nil {
		return m.ApprovalRef
	}
	return nil
}

func (m *Comment) GetDescriptionNum() uint32 {
	if m != nil {
		return m.DescriptionNum
	}
	return 0
}

func (m *Comment) GetIsSpam() bool {
	if m != nil {
		return m.IsSpam
	}
	return false
}

// Next available tag: 5
type FieldValue struct {
	FieldRef             *FieldRef `protobuf:"bytes,1,opt,name=field_ref,json=fieldRef,proto3" json:"field_ref,omitempty"`
	Value                string    `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	IsDerived            bool      `protobuf:"varint,3,opt,name=is_derived,json=isDerived,proto3" json:"is_derived,omitempty"`
	PhaseRef             *PhaseRef `protobuf:"bytes,4,opt,name=phase_ref,json=phaseRef,proto3" json:"phase_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FieldValue) Reset()         { *m = FieldValue{} }
func (m *FieldValue) String() string { return proto.CompactTextString(m) }
func (*FieldValue) ProtoMessage()    {}
func (*FieldValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{4}
}
func (m *FieldValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldValue.Unmarshal(m, b)
}
func (m *FieldValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldValue.Marshal(b, m, deterministic)
}
func (m *FieldValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValue.Merge(m, src)
}
func (m *FieldValue) XXX_Size() int {
	return xxx_messageInfo_FieldValue.Size(m)
}
func (m *FieldValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValue.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValue proto.InternalMessageInfo

func (m *FieldValue) GetFieldRef() *FieldRef {
	if m != nil {
		return m.FieldRef
	}
	return nil
}

func (m *FieldValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *FieldValue) GetIsDerived() bool {
	if m != nil {
		return m.IsDerived
	}
	return false
}

func (m *FieldValue) GetPhaseRef() *PhaseRef {
	if m != nil {
		return m.PhaseRef
	}
	return nil
}

// Next available tag: 24
type Issue struct {
	ProjectName           string          `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	LocalId               uint32          `protobuf:"varint,2,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	Summary               string          `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	StatusRef             *StatusRef      `protobuf:"bytes,4,opt,name=status_ref,json=statusRef,proto3" json:"status_ref,omitempty"`
	OwnerRef              *UserRef        `protobuf:"bytes,5,opt,name=owner_ref,json=ownerRef,proto3" json:"owner_ref,omitempty"`
	CcRefs                []*UserRef      `protobuf:"bytes,6,rep,name=cc_refs,json=ccRefs,proto3" json:"cc_refs,omitempty"`
	LabelRefs             []*LabelRef     `protobuf:"bytes,7,rep,name=label_refs,json=labelRefs,proto3" json:"label_refs,omitempty"`
	ComponentRefs         []*ComponentRef `protobuf:"bytes,8,rep,name=component_refs,json=componentRefs,proto3" json:"component_refs,omitempty"`
	BlockedOnIssueRefs    []*IssueRef     `protobuf:"bytes,9,rep,name=blocked_on_issue_refs,json=blockedOnIssueRefs,proto3" json:"blocked_on_issue_refs,omitempty"`
	BlockingIssueRefs     []*IssueRef     `protobuf:"bytes,10,rep,name=blocking_issue_refs,json=blockingIssueRefs,proto3" json:"blocking_issue_refs,omitempty"`
	DanglingBlockedOnRefs []*IssueRef     `protobuf:"bytes,23,rep,name=dangling_blocked_on_refs,json=danglingBlockedOnRefs,proto3" json:"dangling_blocked_on_refs,omitempty"`
	MergedIntoIssueRef    *IssueRef       `protobuf:"bytes,11,opt,name=merged_into_issue_ref,json=mergedIntoIssueRef,proto3" json:"merged_into_issue_ref,omitempty"`
	FieldValues           []*FieldValue   `protobuf:"bytes,12,rep,name=field_values,json=fieldValues,proto3" json:"field_values,omitempty"`
	IsDeleted             bool            `protobuf:"varint,13,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	ReporterRef           *UserRef        `protobuf:"bytes,14,opt,name=reporter_ref,json=reporterRef,proto3" json:"reporter_ref,omitempty"`
	OpenedTimestamp       uint32          `protobuf:"fixed32,15,opt,name=opened_timestamp,json=openedTimestamp,proto3" json:"opened_timestamp,omitempty"`
	ClosedTimestamp       uint32          `protobuf:"fixed32,16,opt,name=closed_timestamp,json=closedTimestamp,proto3" json:"closed_timestamp,omitempty"`
	ModifiedTimestamp     uint32          `protobuf:"fixed32,17,opt,name=modified_timestamp,json=modifiedTimestamp,proto3" json:"modified_timestamp,omitempty"`
	StarCount             uint32          `protobuf:"varint,18,opt,name=star_count,json=starCount,proto3" json:"star_count,omitempty"`
	IsSpam                bool            `protobuf:"varint,19,opt,name=is_spam,json=isSpam,proto3" json:"is_spam,omitempty"`
	AttachmentCount       uint32          `protobuf:"varint,20,opt,name=attachment_count,json=attachmentCount,proto3" json:"attachment_count,omitempty"`
	ApprovalValues        []*Approval     `protobuf:"bytes,21,rep,name=approval_values,json=approvalValues,proto3" json:"approval_values,omitempty"`
	Phases                []*PhaseDef     `protobuf:"bytes,22,rep,name=phases,proto3" json:"phases,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}        `json:"-"`
	XXX_unrecognized      []byte          `json:"-"`
	XXX_sizecache         int32           `json:"-"`
}

func (m *Issue) Reset()         { *m = Issue{} }
func (m *Issue) String() string { return proto.CompactTextString(m) }
func (*Issue) ProtoMessage()    {}
func (*Issue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{5}
}
func (m *Issue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Issue.Unmarshal(m, b)
}
func (m *Issue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Issue.Marshal(b, m, deterministic)
}
func (m *Issue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Issue.Merge(m, src)
}
func (m *Issue) XXX_Size() int {
	return xxx_messageInfo_Issue.Size(m)
}
func (m *Issue) XXX_DiscardUnknown() {
	xxx_messageInfo_Issue.DiscardUnknown(m)
}

var xxx_messageInfo_Issue proto.InternalMessageInfo

func (m *Issue) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *Issue) GetLocalId() uint32 {
	if m != nil {
		return m.LocalId
	}
	return 0
}

func (m *Issue) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Issue) GetStatusRef() *StatusRef {
	if m != nil {
		return m.StatusRef
	}
	return nil
}

func (m *Issue) GetOwnerRef() *UserRef {
	if m != nil {
		return m.OwnerRef
	}
	return nil
}

func (m *Issue) GetCcRefs() []*UserRef {
	if m != nil {
		return m.CcRefs
	}
	return nil
}

func (m *Issue) GetLabelRefs() []*LabelRef {
	if m != nil {
		return m.LabelRefs
	}
	return nil
}

func (m *Issue) GetComponentRefs() []*ComponentRef {
	if m != nil {
		return m.ComponentRefs
	}
	return nil
}

func (m *Issue) GetBlockedOnIssueRefs() []*IssueRef {
	if m != nil {
		return m.BlockedOnIssueRefs
	}
	return nil
}

func (m *Issue) GetBlockingIssueRefs() []*IssueRef {
	if m != nil {
		return m.BlockingIssueRefs
	}
	return nil
}

func (m *Issue) GetDanglingBlockedOnRefs() []*IssueRef {
	if m != nil {
		return m.DanglingBlockedOnRefs
	}
	return nil
}

func (m *Issue) GetMergedIntoIssueRef() *IssueRef {
	if m != nil {
		return m.MergedIntoIssueRef
	}
	return nil
}

func (m *Issue) GetFieldValues() []*FieldValue {
	if m != nil {
		return m.FieldValues
	}
	return nil
}

func (m *Issue) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *Issue) GetReporterRef() *UserRef {
	if m != nil {
		return m.ReporterRef
	}
	return nil
}

func (m *Issue) GetOpenedTimestamp() uint32 {
	if m != nil {
		return m.OpenedTimestamp
	}
	return 0
}

func (m *Issue) GetClosedTimestamp() uint32 {
	if m != nil {
		return m.ClosedTimestamp
	}
	return 0
}

func (m *Issue) GetModifiedTimestamp() uint32 {
	if m != nil {
		return m.ModifiedTimestamp
	}
	return 0
}

func (m *Issue) GetStarCount() uint32 {
	if m != nil {
		return m.StarCount
	}
	return 0
}

func (m *Issue) GetIsSpam() bool {
	if m != nil {
		return m.IsSpam
	}
	return false
}

func (m *Issue) GetAttachmentCount() uint32 {
	if m != nil {
		return m.AttachmentCount
	}
	return 0
}

func (m *Issue) GetApprovalValues() []*Approval {
	if m != nil {
		return m.ApprovalValues
	}
	return nil
}

func (m *Issue) GetPhases() []*PhaseDef {
	if m != nil {
		return m.Phases
	}
	return nil
}

// Next available tag: 4
type IssueSummary struct {
	ProjectName          string   `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	LocalId              uint32   `protobuf:"varint,2,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	Summary              string   `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueSummary) Reset()         { *m = IssueSummary{} }
func (m *IssueSummary) String() string { return proto.CompactTextString(m) }
func (*IssueSummary) ProtoMessage()    {}
func (*IssueSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{6}
}
func (m *IssueSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueSummary.Unmarshal(m, b)
}
func (m *IssueSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueSummary.Marshal(b, m, deterministic)
}
func (m *IssueSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueSummary.Merge(m, src)
}
func (m *IssueSummary) XXX_Size() int {
	return xxx_messageInfo_IssueSummary.Size(m)
}
func (m *IssueSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueSummary.DiscardUnknown(m)
}

var xxx_messageInfo_IssueSummary proto.InternalMessageInfo

func (m *IssueSummary) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *IssueSummary) GetLocalId() uint32 {
	if m != nil {
		return m.LocalId
	}
	return 0
}

func (m *IssueSummary) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

// Next available tag: 3
type PhaseDef struct {
	PhaseRef             *PhaseRef `protobuf:"bytes,1,opt,name=phase_ref,json=phaseRef,proto3" json:"phase_ref,omitempty"`
	Rank                 uint32    `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PhaseDef) Reset()         { *m = PhaseDef{} }
func (m *PhaseDef) String() string { return proto.CompactTextString(m) }
func (*PhaseDef) ProtoMessage()    {}
func (*PhaseDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{7}
}
func (m *PhaseDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhaseDef.Unmarshal(m, b)
}
func (m *PhaseDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhaseDef.Marshal(b, m, deterministic)
}
func (m *PhaseDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhaseDef.Merge(m, src)
}
func (m *PhaseDef) XXX_Size() int {
	return xxx_messageInfo_PhaseDef.Size(m)
}
func (m *PhaseDef) XXX_DiscardUnknown() {
	xxx_messageInfo_PhaseDef.DiscardUnknown(m)
}

var xxx_messageInfo_PhaseDef proto.InternalMessageInfo

func (m *PhaseDef) GetPhaseRef() *PhaseRef {
	if m != nil {
		return m.PhaseRef
	}
	return nil
}

func (m *PhaseDef) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

// Next available tag: 2
type PhaseRef struct {
	PhaseName            string   `protobuf:"bytes,1,opt,name=phase_name,json=phaseName,proto3" json:"phase_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhaseRef) Reset()         { *m = PhaseRef{} }
func (m *PhaseRef) String() string { return proto.CompactTextString(m) }
func (*PhaseRef) ProtoMessage()    {}
func (*PhaseRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab61f0576fd6c44, []int{8}
}
func (m *PhaseRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhaseRef.Unmarshal(m, b)
}
func (m *PhaseRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhaseRef.Marshal(b, m, deterministic)
}
func (m *PhaseRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhaseRef.Merge(m, src)
}
func (m *PhaseRef) XXX_Size() int {
	return xxx_messageInfo_PhaseRef.Size(m)
}
func (m *PhaseRef) XXX_DiscardUnknown() {
	xxx_messageInfo_PhaseRef.DiscardUnknown(m)
}

var xxx_messageInfo_PhaseRef proto.InternalMessageInfo

func (m *PhaseRef) GetPhaseName() string {
	if m != nil {
		return m.PhaseName
	}
	return ""
}

func init() {
	proto.RegisterType((*Approval)(nil), "monorail.Approval")
	proto.RegisterType((*Amendment)(nil), "monorail.Amendment")
	proto.RegisterType((*Attachment)(nil), "monorail.Attachment")
	proto.RegisterType((*Comment)(nil), "monorail.Comment")
	proto.RegisterType((*FieldValue)(nil), "monorail.FieldValue")
	proto.RegisterType((*Issue)(nil), "monorail.Issue")
	proto.RegisterType((*IssueSummary)(nil), "monorail.IssueSummary")
	proto.RegisterType((*PhaseDef)(nil), "monorail.PhaseDef")
	proto.RegisterType((*PhaseRef)(nil), "monorail.PhaseRef")
	proto.RegisterEnum("monorail.ApprovalStatus", ApprovalStatus_name, ApprovalStatus_value)
	proto.RegisterEnum("monorail.CannedQuery", CannedQuery_name, CannedQuery_value)
}

func init() { proto.RegisterFile("api/api_proto/issue_objects.proto", fileDescriptor_8ab61f0576fd6c44) }

var fileDescriptor_8ab61f0576fd6c44 = []byte{
	// 1290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4f, 0x73, 0xdb, 0x44,
	0x14, 0xaf, 0x13, 0xc7, 0x92, 0x9e, 0x65, 0x5b, 0xd9, 0x26, 0xad, 0x5a, 0xe8, 0x4c, 0x62, 0x0e,
	0x4d, 0xc3, 0x90, 0x94, 0x14, 0xca, 0x81, 0xe1, 0xe0, 0xc6, 0xea, 0x8c, 0x87, 0x62, 0xa7, 0xb2,
	0x93, 0x0e, 0x17, 0x34, 0x8a, 0xb4, 0x4e, 0x45, 0xa5, 0x5d, 0xa1, 0x95, 0x93, 0x09, 0x37, 0x2e,
	0x9c, 0xf9, 0x06, 0x1c, 0xf8, 0x60, 0x7c, 0x15, 0x66, 0xdf, 0x4a, 0x96, 0x62, 0x08, 0x85, 0x61,
	0xb8, 0xed, 0xfe, 0xde, 0xef, 0xfd, 0xd9, 0xb7, 0xfb, 0x7e, 0x12, 0xec, 0xfa, 0x69, 0x74, 0xe8,
	0xa7, 0x91, 0x97, 0x66, 0x3c, 0xe7, 0x87, 0x91, 0x10, 0x0b, 0xea, 0xf1, 0xf3, 0xef, 0x69, 0x90,
	0x8b, 0x03, 0xc4, 0x88, 0x9e, 0x70, 0xc6, 0x33, 0x3f, 0x8a, 0x1f, 0x3e, 0xbc, 0x49, 0x0e, 0x78,
	0x92, 0x70, 0xa6, 0x58, 0xfd, 0x5f, 0xd7, 0x40, 0x1f, 0xa4, 0x69, 0xc6, 0x2f, 0xfd, 0x98, 0x1c,
	0x82, 0x31, 0x8f, 0x68, 0x1c, 0x7a, 0x19, 0x9d, 0xdb, 0x8d, 0x9d, 0xc6, 0x5e, 0xfb, 0x88, 0x1c,
	0x94, 0x61, 0x0e, 0x5e, 0x4a, 0x93, 0x4b, 0xe7, 0xae, 0x3e, 0x2f, 0x56, 0xe4, 0x39, 0x74, 0x7c,
	0x74, 0xa6, 0x99, 0xf4, 0x11, 0xf6, 0xda, 0xce, 0xfa, 0x5e, 0xfb, 0x68, 0xb3, 0x72, 0x3a, 0x15,
	0x34, 0x93, 0x3e, 0x66, 0xc9, 0x73, 0xe9, 0x5c, 0x90, 0xa7, 0xd0, 0x12, 0xb9, 0x9f, 0x2f, 0x84,
	0xbd, 0xbe, 0xd3, 0xd8, 0xeb, 0x1e, 0xd9, 0x95, 0x43, 0x59, 0xcc, 0x14, 0xed, 0x6e, 0xc1, 0x23,
	0xdb, 0xd0, 0x12, 0x34, 0xf7, 0x38, 0xb3, 0x9b, 0x3b, 0x8d, 0x3d, 0xcd, 0xdd, 0x10, 0x34, 0x9f,
	0x30, 0xf2, 0x14, 0x40, 0xd0, 0x3c, 0x57, 0xe9, 0xed, 0x0d, 0x2c, 0xf9, 0x2f, 0xb2, 0x1b, 0x8a,
	0x24, 0x4b, 0x3e, 0x04, 0x23, 0x7d, 0xeb, 0x0b, 0x8a, 0x0e, 0xda, 0xea, 0x19, 0x4f, 0xa4, 0x09,
	0xcf, 0x98, 0x16, 0xab, 0x7e, 0x0e, 0xc6, 0x20, 0xa1, 0x2c, 0x4c, 0x28, 0xcb, 0xc9, 0x23, 0x00,
	0xd5, 0x21, 0xe6, 0x27, 0x14, 0x5b, 0x64, 0xb8, 0xaa, 0x67, 0x63, 0x3f, 0xa1, 0xe4, 0x63, 0x20,
	0x8c, 0x5e, 0x79, 0x3c, 0xf3, 0x42, 0x1a, 0xe7, 0xbe, 0x77, 0xe9, 0xc7, 0x0b, 0x6a, 0xaf, 0x21,
	0xad, 0xc7, 0xe8, 0xd5, 0x24, 0x1b, 0x4a, 0xfc, 0x4c, 0xc2, 0xe4, 0x03, 0x30, 0x78, 0x1c, 0x16,
	0x9c, 0x75, 0xe4, 0xe8, 0x3c, 0x0e, 0xd1, 0xd8, 0xff, 0x79, 0x0d, 0x60, 0x90, 0xe7, 0x7e, 0xf0,
	0x16, 0xf3, 0x7e, 0x04, 0x1d, 0x7f, 0xb9, 0xf3, 0xa2, 0x10, 0x53, 0x37, 0x5d, 0xb3, 0x02, 0x47,
	0x21, 0x79, 0x08, 0xfa, 0x3c, 0x8a, 0x29, 0x96, 0xa6, 0x72, 0x2e, 0xf7, 0x84, 0x40, 0x53, 0x44,
	0x3f, 0xaa, 0x3c, 0x4d, 0x17, 0xd7, 0x64, 0x17, 0xcc, 0x80, 0xb3, 0x5c, 0x46, 0xcc, 0xaf, 0x53,
	0x8a, 0x9d, 0x35, 0xdc, 0x76, 0x81, 0xcd, 0xae, 0x53, 0x2a, 0xcf, 0x1b, 0x09, 0x79, 0x18, 0x9a,
	0xd3, 0x10, 0xfb, 0xab, 0xbb, 0x46, 0x24, 0x86, 0x0a, 0x90, 0x65, 0xe5, 0x6f, 0x17, 0xc9, 0x39,
	0xf3, 0xa3, 0xd8, 0x5b, 0x64, 0xb1, 0xdd, 0xc2, 0x10, 0xe6, 0x12, 0x3c, 0xcd, 0x62, 0xf2, 0x00,
	0xf4, 0xcb, 0x88, 0x5e, 0xa1, 0x5d, 0x43, 0xbb, 0x26, 0xf7, 0xd2, 0xb4, 0x0b, 0x66, 0xc8, 0xaf,
	0x58, 0xcc, 0xfd, 0x10, 0xcd, 0xba, 0xaa, 0xa0, 0xc4, 0x4e, 0xb3, 0xb8, 0xff, 0x53, 0x13, 0xb4,
	0x63, 0x9e, 0x60, 0x17, 0x76, 0xc1, 0x4c, 0x33, 0x2e, 0x1f, 0x79, 0xbd, 0xff, 0xed, 0x02, 0xc3,
	0x1b, 0x78, 0x00, 0x7a, 0xcc, 0x03, 0x3f, 0x96, 0x3d, 0x92, 0x3d, 0xe8, 0xb8, 0x1a, 0xee, 0x47,
	0xa1, 0xf4, 0x16, 0xf4, 0x87, 0x05, 0x65, 0x01, 0xf5, 0xd8, 0x22, 0xc1, 0x56, 0x74, 0xdc, 0x76,
	0x89, 0x8d, 0x17, 0xc9, 0xca, 0x71, 0x9b, 0xab, 0xc7, 0x3d, 0x04, 0x23, 0x50, 0xa5, 0xd0, 0xec,
	0x6f, 0x1e, 0xdb, 0x92, 0x43, 0x3e, 0x04, 0x23, 0x8f, 0x12, 0x2a, 0x72, 0x3f, 0x49, 0xb1, 0x37,
	0x9a, 0x5b, 0x01, 0xc4, 0x06, 0xad, 0xe8, 0x75, 0xd9, 0x97, 0x62, 0x4b, 0x1e, 0x43, 0x2f, 0x62,
	0xe7, 0x7c, 0xc1, 0x42, 0x2f, 0xa1, 0x42, 0xf8, 0x17, 0xb4, 0x68, 0x4d, 0xb7, 0x80, 0xbf, 0x51,
	0x28, 0x79, 0x06, 0xe0, 0x97, 0x8f, 0x53, 0xd8, 0x06, 0x4e, 0xdf, 0xdd, 0xda, 0x30, 0x95, 0x36,
	0xb7, 0x46, 0x23, 0xcf, 0xa1, 0x5d, 0xbd, 0x1b, 0x61, 0x03, 0x7a, 0x6d, 0xd5, 0xbc, 0x96, 0x46,
	0xb7, 0x4e, 0x24, 0x9f, 0x43, 0x31, 0xc5, 0x7e, 0x8c, 0xd3, 0xd3, 0xbe, 0x55, 0x21, 0xda, 0x25,
	0x4f, 0x4e, 0xdc, 0x63, 0xe8, 0x85, 0x54, 0x04, 0x59, 0x94, 0xe6, 0x11, 0x67, 0xd8, 0x7a, 0x13,
	0x5b, 0xdf, 0xad, 0xc1, 0xb2, 0xfb, 0xf7, 0x41, 0x8b, 0x84, 0x27, 0x52, 0x3f, 0xb1, 0x3b, 0xd8,
	0xfa, 0x56, 0x24, 0xa6, 0xa9, 0x9f, 0xf4, 0x7f, 0x6b, 0x00, 0x60, 0x6c, 0x35, 0x38, 0xff, 0x5a,
	0xa6, 0xb6, 0x60, 0xa3, 0x3e, 0x89, 0x6a, 0xb3, 0xbc, 0xec, 0x2c, 0xba, 0xa4, 0x21, 0xbe, 0x86,
	0xe2, 0xb2, 0x11, 0xb8, 0x29, 0x14, 0xcd, 0x7f, 0x20, 0x14, 0xbf, 0xeb, 0xb0, 0x31, 0x92, 0x42,
	0xfc, 0x1f, 0xdf, 0xa9, 0x0d, 0x9a, 0x58, 0x24, 0x89, 0x9f, 0x5d, 0x17, 0xaa, 0x50, 0x6e, 0xc9,
	0x11, 0x80, 0x92, 0xc3, 0x5a, 0x4d, 0xb5, 0xdb, 0x2e, 0x24, 0x13, 0xf5, 0xae, 0x5c, 0x92, 0x03,
	0x30, 0xf8, 0x15, 0x7b, 0x9f, 0x40, 0xea, 0xc8, 0x91, 0xfc, 0x7d, 0xd0, 0x82, 0x40, 0x89, 0x79,
	0xeb, 0x36, 0x31, 0x6f, 0x05, 0x01, 0xca, 0xf8, 0xa7, 0x00, 0xb1, 0x7f, 0x4e, 0x63, 0x45, 0xd7,
	0x90, 0x5e, 0xeb, 0xd1, 0x2b, 0x69, 0xc3, 0x72, 0xe2, 0x62, 0x25, 0xc8, 0x57, 0xd0, 0x0d, 0x78,
	0x92, 0x72, 0x26, 0x55, 0x07, 0xdd, 0x74, 0x74, 0xbb, 0x57, 0xb9, 0x1d, 0x97, 0x76, 0xe9, 0xda,
	0x09, 0x6a, 0x3b, 0x41, 0x1c, 0xd8, 0x3e, 0x8f, 0x79, 0xf0, 0x8e, 0x86, 0x1e, 0x67, 0x9e, 0xfa,
	0xec, 0x61, 0x14, 0x63, 0x35, 0x39, 0xde, 0x84, 0x8c, 0x40, 0x0a, 0x87, 0x09, 0x2b, 0x21, 0x41,
	0x5e, 0xc0, 0x5d, 0x44, 0x23, 0x76, 0x51, 0x0f, 0x02, 0xb7, 0x06, 0xd9, 0x2c, 0xe9, 0x55, 0x8c,
	0xaf, 0xc1, 0x0e, 0x7d, 0x76, 0x11, 0xcb, 0x18, 0xb5, 0x9a, 0x30, 0xd0, 0xfd, 0x5b, 0x03, 0x6d,
	0x97, 0x3e, 0x2f, 0xca, 0xaa, 0xca, 0x73, 0x25, 0x34, 0xbb, 0xa0, 0xa1, 0x17, 0xb1, 0x9c, 0x57,
	0x35, 0xfd, 0x79, 0xc6, 0xaa, 0x73, 0x29, 0x87, 0x11, 0xcb, 0x79, 0x89, 0x91, 0x2f, 0xc0, 0x54,
	0x93, 0x81, 0x2f, 0x5c, 0xd8, 0xe6, 0xea, 0x68, 0x57, 0x53, 0xe4, 0xb6, 0xe7, 0xcb, 0xb5, 0x58,
	0x11, 0xbe, 0xce, 0xaa, 0xf0, 0x7d, 0x06, 0x66, 0x46, 0x53, 0x9e, 0x95, 0x1f, 0xda, 0xee, 0x6d,
	0xef, 0xa8, 0x5d, 0xd2, 0x64, 0x35, 0x4f, 0xc0, 0xe2, 0x29, 0x65, 0x34, 0xf4, 0x2a, 0x11, 0xec,
	0xa1, 0x08, 0xf6, 0x14, 0x3e, 0x5b, 0x4a, 0xe1, 0x13, 0xb0, 0x82, 0x98, 0x8b, 0x1b, 0x54, 0x4b,
	0x51, 0x15, 0x5e, 0x51, 0x3f, 0x01, 0x92, 0xf0, 0x30, 0x9a, 0x47, 0x37, 0xc8, 0x9b, 0x48, 0xde,
	0x2c, 0x2d, 0x15, 0xfd, 0x11, 0xce, 0x4c, 0xe6, 0x05, 0x7c, 0xc1, 0x72, 0x9b, 0xe0, 0xa8, 0xc9,
	0xf1, 0xc8, 0x8e, 0x25, 0x50, 0xd7, 0x9c, 0xbb, 0x75, 0xcd, 0x91, 0x15, 0xd5, 0xbe, 0xb8, 0xca,
	0x7b, 0x0b, 0xbd, 0x7b, 0x15, 0xae, 0x62, 0x7c, 0x09, 0xbd, 0xa5, 0x2e, 0x16, 0x8d, 0xdf, 0x5e,
	0x7d, 0x00, 0xe5, 0x6f, 0x8d, 0xdb, 0x2d, 0xa9, 0x45, 0xe7, 0xf7, 0xa1, 0x85, 0x0a, 0x22, 0xec,
	0x7b, 0xab, 0x3e, 0xa8, 0x31, 0x43, 0x39, 0x6f, 0x8a, 0xd1, 0x9f, 0x83, 0x89, 0x57, 0x3d, 0x2d,
	0xf4, 0xe0, 0x7f, 0xd2, 0x99, 0xfe, 0x04, 0xf4, 0x32, 0xf7, 0x4d, 0x19, 0x6c, 0xbc, 0x5f, 0x06,
	0xe5, 0x9f, 0x46, 0xe6, 0xb3, 0x77, 0x45, 0x36, 0x5c, 0xf7, 0x9f, 0x14, 0x01, 0xa5, 0xfd, 0x11,
	0x80, 0x0a, 0x58, 0xff, 0x85, 0x42, 0x44, 0x16, 0xbc, 0xff, 0x4b, 0x03, 0xba, 0x37, 0xff, 0x01,
	0x49, 0x1b, 0xb4, 0xf1, 0x64, 0xe6, 0x4d, 0x9d, 0x99, 0x75, 0x87, 0x58, 0x60, 0x8e, 0x1d, 0x67,
	0x38, 0xf5, 0x5c, 0xe7, 0x6c, 0xe4, 0xbc, 0xb1, 0x1a, 0xa4, 0x05, 0x6b, 0xe3, 0x81, 0xb5, 0x46,
	0xb6, 0xc0, 0x52, 0x98, 0xe7, 0x3a, 0xaf, 0x4f, 0x9d, 0xe9, 0xcc, 0x19, 0x5a, 0xeb, 0x84, 0x40,
	0xb7, 0x40, 0xa7, 0xb3, 0x81, 0x2b, 0xb1, 0x26, 0xe9, 0x80, 0x21, 0x63, 0x78, 0xa3, 0xf1, 0xcb,
	0x89, 0xb5, 0x41, 0x4c, 0xd0, 0x07, 0x27, 0x27, 0xee, 0xe4, 0xcc, 0x19, 0x5a, 0x2d, 0x4c, 0x30,
	0x99, 0x79, 0x4b, 0x44, 0xdb, 0xff, 0x0e, 0xda, 0xc7, 0x3e, 0x63, 0x34, 0x7c, 0xbd, 0xa0, 0xd9,
	0x35, 0xd1, 0x60, 0x7d, 0xf0, 0xea, 0x95, 0x75, 0x47, 0x2e, 0xc6, 0x58, 0x81, 0x0e, 0xcd, 0xc9,
	0x89, 0x33, 0xb6, 0xd6, 0x88, 0x01, 0x1b, 0x93, 0x37, 0x63, 0x4c, 0x6c, 0x82, 0xee, 0x3a, 0x27,
	0x93, 0x22, 0x65, 0x1b, 0x34, 0x99, 0xdf, 0x75, 0x86, 0xd6, 0x86, 0xcc, 0x3f, 0x9b, 0x78, 0x67,
	0x8e, 0x3b, 0x7a, 0xf9, 0xad, 0xd5, 0x3a, 0x6f, 0xe1, 0xaf, 0xf8, 0xb3, 0x3f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x10, 0x98, 0xa5, 0x10, 0xd5, 0x0b, 0x00, 0x00,
}
