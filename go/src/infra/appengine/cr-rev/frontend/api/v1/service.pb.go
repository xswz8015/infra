// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This proto file defines cr-rev API v1. It is based on previous version of
// REST cr-rev API and it has the goal to keep clients fully compatible.
// json fields are used to simplify conversion of legacy naming scheme.
// API v2 needs be introduced to fix flaws in API v1 design.

// All Request and Response parameters are required.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: infra/appengine/cr-rev/frontend/api/v1/service.proto

package api

import prpc "go.chromium.org/luci/grpc/prpc"

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

type RedirectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// query is equal to URL path of crrev main redirect logic. For example, it
	// can be "/3" (redirect to chromium/src commit with position 3) or
	// /{some_commit_hash} (redirect to some repository based on
	// some_commit_hash), etc.
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *RedirectRequest) Reset() {
	*x = RedirectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectRequest) ProtoMessage() {}

func (x *RedirectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectRequest.ProtoReflect.Descriptor instead.
func (*RedirectRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *RedirectRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type RedirectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// git_hash is a full git commit hash of matched commit that is used for
	// redirect.
	GitHash string `protobuf:"bytes,1,opt,name=git_hash,json=git_sha,proto3" json:"git_hash,omitempty"`
	// host is googlesource host (e.g. chromium).
	Host string `protobuf:"bytes,2,opt,name=host,json=project,proto3" json:"host,omitempty"`
	// repository is Git repository (e.g. chromium/src)
	Repository string `protobuf:"bytes,3,opt,name=repository,json=repo,proto3" json:"repository,omitempty"`
	// redirect_url is Gitiles URL of the commit, the same URL that user is
	// redirected to when using crrev's main redirect logic.
	RedirectUrl string `protobuf:"bytes,4,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *RedirectResponse) Reset() {
	*x = RedirectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectResponse) ProtoMessage() {}

func (x *RedirectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectResponse.ProtoReflect.Descriptor instead.
func (*RedirectResponse) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *RedirectResponse) GetGitHash() string {
	if x != nil {
		return x.GitHash
	}
	return ""
}

func (x *RedirectResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RedirectResponse) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *RedirectResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

type NumberingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// host is googlesource host (e.g. chromium).
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// repository is Git repository (e.g. chromium/src)
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	// position_ref is name of position defined in value of git-footer git-svn-id
	// or Cr-Commit-Position (e.g. refs/heads/master,
	// svn://svn.chromium.org/chrome/trunk/src)
	PositionRef string `protobuf:"bytes,3,opt,name=position_ref,json=positionRef,proto3" json:"position_ref,omitempty"`
	// position_number is sequential identifier of commit in given branch
	// (position_ref).
	PositionNumber int64 `protobuf:"varint,4,opt,name=position_number,json=positionNumber,proto3" json:"position_number,omitempty"`
}

func (x *NumberingRequest) Reset() {
	*x = NumberingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberingRequest) ProtoMessage() {}

func (x *NumberingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberingRequest.ProtoReflect.Descriptor instead.
func (*NumberingRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *NumberingRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *NumberingRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *NumberingRequest) GetPositionRef() string {
	if x != nil {
		return x.PositionRef
	}
	return ""
}

func (x *NumberingRequest) GetPositionNumber() int64 {
	if x != nil {
		return x.PositionNumber
	}
	return 0
}

type NumberingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// git_hash is a full git commit hash of matched commit that is used for
	// redirect.
	GitHash string `protobuf:"bytes,1,opt,name=git_hash,json=git_sha,proto3" json:"git_hash,omitempty"`
	// position_number matches NumberingRequest position_number. Probably useless
	// to clients, but kept for backward compatibility.
	PositionNumber int64 `protobuf:"varint,2,opt,name=position_number,json=number,proto3" json:"position_number,omitempty"`
	// host matches NumberingRequest host. Probably useless to clients, but kept
	// for backward compatibility.
	Host string `protobuf:"bytes,3,opt,name=host,json=project,proto3" json:"host,omitempty"`
	// repository matches NumberingRequest repository. Probably useless to
	// clients, but kept for backward compatibility.
	Repository string `protobuf:"bytes,4,opt,name=repository,json=repo,proto3" json:"repository,omitempty"`
}

func (x *NumberingResponse) Reset() {
	*x = NumberingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberingResponse) ProtoMessage() {}

func (x *NumberingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberingResponse.ProtoReflect.Descriptor instead.
func (*NumberingResponse) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *NumberingResponse) GetGitHash() string {
	if x != nil {
		return x.GitHash
	}
	return ""
}

func (x *NumberingResponse) GetPositionNumber() int64 {
	if x != nil {
		return x.PositionNumber
	}
	return 0
}

func (x *NumberingResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *NumberingResponse) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

type CommitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// git_hash is a full git commit hash of desired commit.
	GitHash string `protobuf:"bytes,1,opt,name=git_hash,json=gitHash,proto3" json:"git_hash,omitempty"`
}

func (x *CommitRequest) Reset() {
	*x = CommitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitRequest) ProtoMessage() {}

func (x *CommitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitRequest.ProtoReflect.Descriptor instead.
func (*CommitRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *CommitRequest) GetGitHash() string {
	if x != nil {
		return x.GitHash
	}
	return ""
}

type CommitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// git_hash is a full git commit hash of matched commit that is used for
	// redirect.
	GitHash string `protobuf:"bytes,1,opt,name=git_hash,json=git_sha,proto3" json:"git_hash,omitempty"`
	// host is googlesource host (e.g. chromium).
	Host string `protobuf:"bytes,2,opt,name=host,json=project,proto3" json:"host,omitempty"`
	// repository is Git repository (e.g. chromium/src)
	Repository string `protobuf:"bytes,3,opt,name=repository,json=repo,proto3" json:"repository,omitempty"`
	// position_number is sequential identifier of commit in given branch
	// (position_ref).
	PositionNumber int64 `protobuf:"varint,4,opt,name=position_number,json=number,proto3" json:"position_number,omitempty"`
	// redirect_url is Gitiles URL of the commit, the same URL that user is
	// redirected to when using crrev's main redirect logic.
	RedirectUrl string `protobuf:"bytes,5,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *CommitResponse) Reset() {
	*x = CommitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitResponse) ProtoMessage() {}

func (x *CommitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitResponse.ProtoReflect.Descriptor instead.
func (*CommitResponse) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *CommitResponse) GetGitHash() string {
	if x != nil {
		return x.GitHash
	}
	return ""
}

func (x *CommitResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *CommitResponse) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *CommitResponse) GetPositionNumber() int64 {
	if x != nil {
		return x.PositionNumber
	}
	return 0
}

func (x *CommitResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

var File_infra_appengine_cr_rev_frontend_api_v1_service_proto protoreflect.FileDescriptor

var file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x63, 0x72, 0x2d, 0x72, 0x65, 0x76, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x72, 0x72, 0x65, 0x76, 0x22, 0x27, 0x0a,
	0x0f, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x81, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x69, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67,
	0x69, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x12, 0x15, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x92, 0x01, 0x0a, 0x10, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x80, 0x01, 0x0a, 0x11, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x69, 0x74, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x69, 0x74, 0x5f, 0x73, 0x68, 0x61,
	0x12, 0x1f, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x15, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x22, 0x2a, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x69, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0xa0,
	0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x69, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x69, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x12, 0x15, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x1f, 0x0a,
	0x0f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72,
	0x6c, 0x32, 0xbb, 0x01, 0x0a, 0x05, 0x43, 0x72, 0x72, 0x65, 0x76, 0x12, 0x3b, 0x0a, 0x08, 0x52,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x72, 0x72, 0x65, 0x76, 0x2e,
	0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x63, 0x72, 0x72, 0x65, 0x76, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x63, 0x72, 0x72, 0x65, 0x76, 0x2e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x63, 0x72, 0x72, 0x65, 0x76, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x72, 0x72, 0x65, 0x76, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x72, 0x72, 0x65, 0x76,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2c, 0x5a, 0x2a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x63, 0x72, 0x2d, 0x72, 0x65, 0x76, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescOnce sync.Once
	file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescData = file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDesc
)

func file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescGZIP() []byte {
	file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescOnce.Do(func() {
		file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescData)
	})
	return file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDescData
}

var file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_infra_appengine_cr_rev_frontend_api_v1_service_proto_goTypes = []interface{}{
	(*RedirectRequest)(nil),   // 0: crrev.RedirectRequest
	(*RedirectResponse)(nil),  // 1: crrev.RedirectResponse
	(*NumberingRequest)(nil),  // 2: crrev.NumberingRequest
	(*NumberingResponse)(nil), // 3: crrev.NumberingResponse
	(*CommitRequest)(nil),     // 4: crrev.CommitRequest
	(*CommitResponse)(nil),    // 5: crrev.CommitResponse
}
var file_infra_appengine_cr_rev_frontend_api_v1_service_proto_depIdxs = []int32{
	0, // 0: crrev.Crrev.Redirect:input_type -> crrev.RedirectRequest
	2, // 1: crrev.Crrev.Numbering:input_type -> crrev.NumberingRequest
	4, // 2: crrev.Crrev.Commit:input_type -> crrev.CommitRequest
	1, // 3: crrev.Crrev.Redirect:output_type -> crrev.RedirectResponse
	3, // 4: crrev.Crrev.Numbering:output_type -> crrev.NumberingResponse
	5, // 5: crrev.Crrev.Commit:output_type -> crrev.CommitResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infra_appengine_cr_rev_frontend_api_v1_service_proto_init() }
func file_infra_appengine_cr_rev_frontend_api_v1_service_proto_init() {
	if File_infra_appengine_cr_rev_frontend_api_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectRequest); i {
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
		file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectResponse); i {
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
		file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberingRequest); i {
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
		file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberingResponse); i {
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
		file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitRequest); i {
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
		file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitResponse); i {
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
			RawDescriptor: file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_appengine_cr_rev_frontend_api_v1_service_proto_goTypes,
		DependencyIndexes: file_infra_appengine_cr_rev_frontend_api_v1_service_proto_depIdxs,
		MessageInfos:      file_infra_appengine_cr_rev_frontend_api_v1_service_proto_msgTypes,
	}.Build()
	File_infra_appengine_cr_rev_frontend_api_v1_service_proto = out.File
	file_infra_appengine_cr_rev_frontend_api_v1_service_proto_rawDesc = nil
	file_infra_appengine_cr_rev_frontend_api_v1_service_proto_goTypes = nil
	file_infra_appengine_cr_rev_frontend_api_v1_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CrrevClient is the client API for Crrev service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrrevClient interface {
	// Redirect implements the same logic as the main crrev redirect, but returns
	// redirect and commit information in body instead of HTTP redirect.
	Redirect(ctx context.Context, in *RedirectRequest, opts ...grpc.CallOption) (*RedirectResponse, error)
	// Numbering returns commit that matches desired position of commit, based on
	// NumberingRequest parameters. Commit position is based on git-footer
	// git-svn-id or Cr-Commit-Position.
	Numbering(ctx context.Context, in *NumberingRequest, opts ...grpc.CallOption) (*NumberingResponse, error)
	// Commit returns commit with desired commit hash. If there are multiple
	// commits with the same commit hash (which happens with forks and mirrors),
	// it checks priorities based on config. It is possible that priorities are
	// the same. In such case, there is no guarantee which one will be returned.
	Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error)
}
type crrevPRPCClient struct {
	client *prpc.Client
}

func NewCrrevPRPCClient(client *prpc.Client) CrrevClient {
	return &crrevPRPCClient{client}
}

func (c *crrevPRPCClient) Redirect(ctx context.Context, in *RedirectRequest, opts ...grpc.CallOption) (*RedirectResponse, error) {
	out := new(RedirectResponse)
	err := c.client.Call(ctx, "crrev.Crrev", "Redirect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crrevPRPCClient) Numbering(ctx context.Context, in *NumberingRequest, opts ...grpc.CallOption) (*NumberingResponse, error) {
	out := new(NumberingResponse)
	err := c.client.Call(ctx, "crrev.Crrev", "Numbering", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crrevPRPCClient) Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error) {
	out := new(CommitResponse)
	err := c.client.Call(ctx, "crrev.Crrev", "Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type crrevClient struct {
	cc grpc.ClientConnInterface
}

func NewCrrevClient(cc grpc.ClientConnInterface) CrrevClient {
	return &crrevClient{cc}
}

func (c *crrevClient) Redirect(ctx context.Context, in *RedirectRequest, opts ...grpc.CallOption) (*RedirectResponse, error) {
	out := new(RedirectResponse)
	err := c.cc.Invoke(ctx, "/crrev.Crrev/Redirect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crrevClient) Numbering(ctx context.Context, in *NumberingRequest, opts ...grpc.CallOption) (*NumberingResponse, error) {
	out := new(NumberingResponse)
	err := c.cc.Invoke(ctx, "/crrev.Crrev/Numbering", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crrevClient) Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error) {
	out := new(CommitResponse)
	err := c.cc.Invoke(ctx, "/crrev.Crrev/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrrevServer is the server API for Crrev service.
type CrrevServer interface {
	// Redirect implements the same logic as the main crrev redirect, but returns
	// redirect and commit information in body instead of HTTP redirect.
	Redirect(context.Context, *RedirectRequest) (*RedirectResponse, error)
	// Numbering returns commit that matches desired position of commit, based on
	// NumberingRequest parameters. Commit position is based on git-footer
	// git-svn-id or Cr-Commit-Position.
	Numbering(context.Context, *NumberingRequest) (*NumberingResponse, error)
	// Commit returns commit with desired commit hash. If there are multiple
	// commits with the same commit hash (which happens with forks and mirrors),
	// it checks priorities based on config. It is possible that priorities are
	// the same. In such case, there is no guarantee which one will be returned.
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)
}

// UnimplementedCrrevServer can be embedded to have forward compatible implementations.
type UnimplementedCrrevServer struct {
}

func (*UnimplementedCrrevServer) Redirect(context.Context, *RedirectRequest) (*RedirectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redirect not implemented")
}
func (*UnimplementedCrrevServer) Numbering(context.Context, *NumberingRequest) (*NumberingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Numbering not implemented")
}
func (*UnimplementedCrrevServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}

func RegisterCrrevServer(s prpc.Registrar, srv CrrevServer) {
	s.RegisterService(&_Crrev_serviceDesc, srv)
}

func _Crrev_Redirect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedirectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrrevServer).Redirect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crrev.Crrev/Redirect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrrevServer).Redirect(ctx, req.(*RedirectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crrev_Numbering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrrevServer).Numbering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crrev.Crrev/Numbering",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrrevServer).Numbering(ctx, req.(*NumberingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crrev_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrrevServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crrev.Crrev/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrrevServer).Commit(ctx, req.(*CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crrev_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crrev.Crrev",
	HandlerType: (*CrrevServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Redirect",
			Handler:    _Crrev_Redirect_Handler,
		},
		{
			MethodName: "Numbering",
			Handler:    _Crrev_Numbering_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Crrev_Commit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/appengine/cr-rev/frontend/api/v1/service.proto",
}
