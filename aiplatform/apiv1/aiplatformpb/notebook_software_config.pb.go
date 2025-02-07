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
// source: google/cloud/aiplatform/v1/notebook_software_config.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a notebook runtime post startup script behavior.
type PostStartupScriptConfig_PostStartupScriptBehavior int32

const (
	// Unspecified post startup script behavior.
	PostStartupScriptConfig_POST_STARTUP_SCRIPT_BEHAVIOR_UNSPECIFIED PostStartupScriptConfig_PostStartupScriptBehavior = 0
	// Run post startup script after runtime is started.
	PostStartupScriptConfig_RUN_ONCE PostStartupScriptConfig_PostStartupScriptBehavior = 1
	// Run post startup script after runtime is stopped.
	PostStartupScriptConfig_RUN_EVERY_START PostStartupScriptConfig_PostStartupScriptBehavior = 2
	// Download and run post startup script every time runtime is started.
	PostStartupScriptConfig_DOWNLOAD_AND_RUN_EVERY_START PostStartupScriptConfig_PostStartupScriptBehavior = 3
)

// Enum value maps for PostStartupScriptConfig_PostStartupScriptBehavior.
var (
	PostStartupScriptConfig_PostStartupScriptBehavior_name = map[int32]string{
		0: "POST_STARTUP_SCRIPT_BEHAVIOR_UNSPECIFIED",
		1: "RUN_ONCE",
		2: "RUN_EVERY_START",
		3: "DOWNLOAD_AND_RUN_EVERY_START",
	}
	PostStartupScriptConfig_PostStartupScriptBehavior_value = map[string]int32{
		"POST_STARTUP_SCRIPT_BEHAVIOR_UNSPECIFIED": 0,
		"RUN_ONCE":                     1,
		"RUN_EVERY_START":              2,
		"DOWNLOAD_AND_RUN_EVERY_START": 3,
	}
)

func (x PostStartupScriptConfig_PostStartupScriptBehavior) Enum() *PostStartupScriptConfig_PostStartupScriptBehavior {
	p := new(PostStartupScriptConfig_PostStartupScriptBehavior)
	*p = x
	return p
}

func (x PostStartupScriptConfig_PostStartupScriptBehavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PostStartupScriptConfig_PostStartupScriptBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_notebook_software_config_proto_enumTypes[0].Descriptor()
}

func (PostStartupScriptConfig_PostStartupScriptBehavior) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_notebook_software_config_proto_enumTypes[0]
}

func (x PostStartupScriptConfig_PostStartupScriptBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PostStartupScriptConfig_PostStartupScriptBehavior.Descriptor instead.
func (PostStartupScriptConfig_PostStartupScriptBehavior) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDescGZIP(), []int{0, 0}
}

// Post startup script config.
type PostStartupScriptConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Post startup script to run after runtime is started.
	PostStartupScript string `protobuf:"bytes,1,opt,name=post_startup_script,json=postStartupScript,proto3" json:"post_startup_script,omitempty"`
	// Optional. Post startup script url to download. Example:
	// https://bucket/script.sh
	PostStartupScriptUrl string `protobuf:"bytes,2,opt,name=post_startup_script_url,json=postStartupScriptUrl,proto3" json:"post_startup_script_url,omitempty"`
	// Optional. Post startup script behavior that defines download and execution
	// behavior.
	PostStartupScriptBehavior PostStartupScriptConfig_PostStartupScriptBehavior `protobuf:"varint,3,opt,name=post_startup_script_behavior,json=postStartupScriptBehavior,proto3,enum=google.cloud.aiplatform.v1.PostStartupScriptConfig_PostStartupScriptBehavior" json:"post_startup_script_behavior,omitempty"`
}

func (x *PostStartupScriptConfig) Reset() {
	*x = PostStartupScriptConfig{}
	mi := &file_google_cloud_aiplatform_v1_notebook_software_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostStartupScriptConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostStartupScriptConfig) ProtoMessage() {}

func (x *PostStartupScriptConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_notebook_software_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostStartupScriptConfig.ProtoReflect.Descriptor instead.
func (*PostStartupScriptConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDescGZIP(), []int{0}
}

func (x *PostStartupScriptConfig) GetPostStartupScript() string {
	if x != nil {
		return x.PostStartupScript
	}
	return ""
}

func (x *PostStartupScriptConfig) GetPostStartupScriptUrl() string {
	if x != nil {
		return x.PostStartupScriptUrl
	}
	return ""
}

func (x *PostStartupScriptConfig) GetPostStartupScriptBehavior() PostStartupScriptConfig_PostStartupScriptBehavior {
	if x != nil {
		return x.PostStartupScriptBehavior
	}
	return PostStartupScriptConfig_POST_STARTUP_SCRIPT_BEHAVIOR_UNSPECIFIED
}

// Notebook Software Config.
type NotebookSoftwareConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Environment variables to be passed to the container.
	// Maximum limit is 100.
	Env []*EnvVar `protobuf:"bytes,1,rep,name=env,proto3" json:"env,omitempty"`
	// Optional. Post startup script config.
	PostStartupScriptConfig *PostStartupScriptConfig `protobuf:"bytes,2,opt,name=post_startup_script_config,json=postStartupScriptConfig,proto3" json:"post_startup_script_config,omitempty"`
}

func (x *NotebookSoftwareConfig) Reset() {
	*x = NotebookSoftwareConfig{}
	mi := &file_google_cloud_aiplatform_v1_notebook_software_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotebookSoftwareConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotebookSoftwareConfig) ProtoMessage() {}

func (x *NotebookSoftwareConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_notebook_software_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotebookSoftwareConfig.ProtoReflect.Descriptor instead.
func (*NotebookSoftwareConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDescGZIP(), []int{1}
}

func (x *NotebookSoftwareConfig) GetEnv() []*EnvVar {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *NotebookSoftwareConfig) GetPostStartupScriptConfig() *PostStartupScriptConfig {
	if x != nil {
		return x.PostStartupScriptConfig
	}
	return nil
}

var File_google_cloud_aiplatform_v1_notebook_software_config_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb1, 0x03, 0x0a, 0x17, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x75, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x33,
	0x0a, 0x13, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x5f, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x11, 0x70, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x12, 0x3a, 0x0a, 0x17, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x75, 0x70, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x14, 0x70, 0x6f, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x55, 0x72, 0x6c, 0x12,
	0x93, 0x01, 0x0a, 0x1c, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70,
	0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x19, 0x70, 0x6f, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x22, 0x8e, 0x01, 0x0a, 0x19, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x75, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x42, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x28, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x55, 0x50, 0x5f, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x5f, 0x42, 0x45, 0x48, 0x41, 0x56,
	0x49, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x55, 0x4e, 0x5f, 0x4f, 0x4e, 0x43, 0x45, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x52, 0x55, 0x4e, 0x5f, 0x45, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44,
	0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x55, 0x4e, 0x5f, 0x45, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x10, 0x03, 0x22, 0xca, 0x01, 0x0a, 0x16, 0x4e, 0x6f, 0x74, 0x65, 0x62,
	0x6f, 0x6f, 0x6b, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x39, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x56,
	0x61, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x75, 0x0a, 0x1a,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x5f, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x17, 0x70, 0x6f, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0xd9, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x1b, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b,
	0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDescData = file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_notebook_software_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1_notebook_software_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_aiplatform_v1_notebook_software_config_proto_goTypes = []any{
	(PostStartupScriptConfig_PostStartupScriptBehavior)(0), // 0: google.cloud.aiplatform.v1.PostStartupScriptConfig.PostStartupScriptBehavior
	(*PostStartupScriptConfig)(nil),                        // 1: google.cloud.aiplatform.v1.PostStartupScriptConfig
	(*NotebookSoftwareConfig)(nil),                         // 2: google.cloud.aiplatform.v1.NotebookSoftwareConfig
	(*EnvVar)(nil),                                         // 3: google.cloud.aiplatform.v1.EnvVar
}
var file_google_cloud_aiplatform_v1_notebook_software_config_proto_depIdxs = []int32{
	0, // 0: google.cloud.aiplatform.v1.PostStartupScriptConfig.post_startup_script_behavior:type_name -> google.cloud.aiplatform.v1.PostStartupScriptConfig.PostStartupScriptBehavior
	3, // 1: google.cloud.aiplatform.v1.NotebookSoftwareConfig.env:type_name -> google.cloud.aiplatform.v1.EnvVar
	1, // 2: google.cloud.aiplatform.v1.NotebookSoftwareConfig.post_startup_script_config:type_name -> google.cloud.aiplatform.v1.PostStartupScriptConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_notebook_software_config_proto_init() }
func file_google_cloud_aiplatform_v1_notebook_software_config_proto_init() {
	if File_google_cloud_aiplatform_v1_notebook_software_config_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1_env_var_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_notebook_software_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_notebook_software_config_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1_notebook_software_config_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1_notebook_software_config_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_notebook_software_config_proto = out.File
	file_google_cloud_aiplatform_v1_notebook_software_config_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_notebook_software_config_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_notebook_software_config_proto_depIdxs = nil
}
