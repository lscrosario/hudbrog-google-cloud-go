// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: google/iam/admin/v1/audit_data.proto

package adminpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Audit log information specific to Cloud IAM admin APIs. This message is
// serialized as an `Any` type in the `ServiceData` message of an
// `AuditLog` message.
type AuditData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The permission_delta when when creating or updating a Role.
	PermissionDelta *AuditData_PermissionDelta `protobuf:"bytes,1,opt,name=permission_delta,json=permissionDelta,proto3" json:"permission_delta,omitempty"`
}

func (x *AuditData) Reset() {
	*x = AuditData{}
	mi := &file_google_iam_admin_v1_audit_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditData) ProtoMessage() {}

func (x *AuditData) ProtoReflect() protoreflect.Message {
	mi := &file_google_iam_admin_v1_audit_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditData.ProtoReflect.Descriptor instead.
func (*AuditData) Descriptor() ([]byte, []int) {
	return file_google_iam_admin_v1_audit_data_proto_rawDescGZIP(), []int{0}
}

func (x *AuditData) GetPermissionDelta() *AuditData_PermissionDelta {
	if x != nil {
		return x.PermissionDelta
	}
	return nil
}

// A PermissionDelta message to record the added_permissions and
// removed_permissions inside a role.
type AuditData_PermissionDelta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Added permissions.
	AddedPermissions []string `protobuf:"bytes,1,rep,name=added_permissions,json=addedPermissions,proto3" json:"added_permissions,omitempty"`
	// Removed permissions.
	RemovedPermissions []string `protobuf:"bytes,2,rep,name=removed_permissions,json=removedPermissions,proto3" json:"removed_permissions,omitempty"`
}

func (x *AuditData_PermissionDelta) Reset() {
	*x = AuditData_PermissionDelta{}
	mi := &file_google_iam_admin_v1_audit_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditData_PermissionDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditData_PermissionDelta) ProtoMessage() {}

func (x *AuditData_PermissionDelta) ProtoReflect() protoreflect.Message {
	mi := &file_google_iam_admin_v1_audit_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditData_PermissionDelta.ProtoReflect.Descriptor instead.
func (*AuditData_PermissionDelta) Descriptor() ([]byte, []int) {
	return file_google_iam_admin_v1_audit_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AuditData_PermissionDelta) GetAddedPermissions() []string {
	if x != nil {
		return x.AddedPermissions
	}
	return nil
}

func (x *AuditData_PermissionDelta) GetRemovedPermissions() []string {
	if x != nil {
		return x.RemovedPermissions
	}
	return nil
}

var File_google_iam_admin_v1_audit_data_proto protoreflect.FileDescriptor

var file_google_iam_admin_v1_audit_data_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xd7, 0x01, 0x0a, 0x09,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x59, 0x0a, 0x10, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x6c, 0x74, 0x61, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x6c, 0x74, 0x61, 0x1a, 0x6f, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x10, 0x61, 0x64, 0x64, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x98, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x42, 0x0e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x33, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62,
	0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x49, 0x61, 0x6d, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_iam_admin_v1_audit_data_proto_rawDescOnce sync.Once
	file_google_iam_admin_v1_audit_data_proto_rawDescData = file_google_iam_admin_v1_audit_data_proto_rawDesc
)

func file_google_iam_admin_v1_audit_data_proto_rawDescGZIP() []byte {
	file_google_iam_admin_v1_audit_data_proto_rawDescOnce.Do(func() {
		file_google_iam_admin_v1_audit_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_iam_admin_v1_audit_data_proto_rawDescData)
	})
	return file_google_iam_admin_v1_audit_data_proto_rawDescData
}

var file_google_iam_admin_v1_audit_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_iam_admin_v1_audit_data_proto_goTypes = []any{
	(*AuditData)(nil),                 // 0: google.iam.admin.v1.AuditData
	(*AuditData_PermissionDelta)(nil), // 1: google.iam.admin.v1.AuditData.PermissionDelta
}
var file_google_iam_admin_v1_audit_data_proto_depIdxs = []int32{
	1, // 0: google.iam.admin.v1.AuditData.permission_delta:type_name -> google.iam.admin.v1.AuditData.PermissionDelta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_iam_admin_v1_audit_data_proto_init() }
func file_google_iam_admin_v1_audit_data_proto_init() {
	if File_google_iam_admin_v1_audit_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_iam_admin_v1_audit_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_iam_admin_v1_audit_data_proto_goTypes,
		DependencyIndexes: file_google_iam_admin_v1_audit_data_proto_depIdxs,
		MessageInfos:      file_google_iam_admin_v1_audit_data_proto_msgTypes,
	}.Build()
	File_google_iam_admin_v1_audit_data_proto = out.File
	file_google_iam_admin_v1_audit_data_proto_rawDesc = nil
	file_google_iam_admin_v1_audit_data_proto_goTypes = nil
	file_google_iam_admin_v1_audit_data_proto_depIdxs = nil
}
