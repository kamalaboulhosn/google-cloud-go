// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: google/cloud/dialogflow/cx/v3beta1/advanced_settings.proto

package cxpb

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

// Hierarchical advanced settings for agent/flow/page/fulfillment/parameter.
// Settings exposed at lower level overrides the settings exposed at higher
// level. Overriding occurs at the sub-setting level. For example, the
// playback_interruption_settings at fulfillment level only overrides the
// playback_interruption_settings at the agent level, leaving other settings
// at the agent level unchanged.
//
// DTMF settings does not override each other. DTMF settings set at different
// levels define DTMF detections running in parallel.
//
// Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
type AdvancedSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Settings for logging.
	// Settings for Dialogflow History, Contact Center messages, StackDriver logs,
	// and speech logging.
	// Exposed at the following levels:
	// - Agent level.
	LoggingSettings *AdvancedSettings_LoggingSettings `protobuf:"bytes,6,opt,name=logging_settings,json=loggingSettings,proto3" json:"logging_settings,omitempty"`
}

func (x *AdvancedSettings) Reset() {
	*x = AdvancedSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvancedSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedSettings) ProtoMessage() {}

func (x *AdvancedSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedSettings.ProtoReflect.Descriptor instead.
func (*AdvancedSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescGZIP(), []int{0}
}

func (x *AdvancedSettings) GetLoggingSettings() *AdvancedSettings_LoggingSettings {
	if x != nil {
		return x.LoggingSettings
	}
	return nil
}

// Define behaviors on logging.
type AdvancedSettings_LoggingSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, StackDriver logging is currently enabled.
	EnableStackdriverLogging bool `protobuf:"varint,2,opt,name=enable_stackdriver_logging,json=enableStackdriverLogging,proto3" json:"enable_stackdriver_logging,omitempty"`
	// If true, DF Interaction logging is currently enabled.
	EnableInteractionLogging bool `protobuf:"varint,3,opt,name=enable_interaction_logging,json=enableInteractionLogging,proto3" json:"enable_interaction_logging,omitempty"`
}

func (x *AdvancedSettings_LoggingSettings) Reset() {
	*x = AdvancedSettings_LoggingSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvancedSettings_LoggingSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedSettings_LoggingSettings) ProtoMessage() {}

func (x *AdvancedSettings_LoggingSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedSettings_LoggingSettings.ProtoReflect.Descriptor instead.
func (*AdvancedSettings_LoggingSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AdvancedSettings_LoggingSettings) GetEnableStackdriverLogging() bool {
	if x != nil {
		return x.EnableStackdriverLogging
	}
	return false
}

func (x *AdvancedSettings_LoggingSettings) GetEnableInteractionLogging() bool {
	if x != nil {
		return x.EnableInteractionLogging
	}
	return false
}

var File_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x93, 0x02, 0x0a, 0x10, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x6f, 0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x8d, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x0a, 0x1a, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x42, 0xdd, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x15, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x63,
	0x78, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x78, 0x2e, 0x56, 0x33, 0x42, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02,
	0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x43, 0x58, 0x3a, 0x3a,
	0x56, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescData = file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDesc
)

func file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDescData
}

var file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_goTypes = []interface{}{
	(*AdvancedSettings)(nil),                 // 0: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings
	(*AdvancedSettings_LoggingSettings)(nil), // 1: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.LoggingSettings
}
var file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_depIdxs = []int32{
	1, // 0: google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.logging_settings:type_name -> google.cloud.dialogflow.cx.v3beta1.AdvancedSettings.LoggingSettings
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_init() }
func file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_init() {
	if File_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvancedSettings); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvancedSettings_LoggingSettings); i {
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
			RawDescriptor: file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_depIdxs,
		MessageInfos:      file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto = out.File
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_rawDesc = nil
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_goTypes = nil
	file_google_cloud_dialogflow_cx_v3beta1_advanced_settings_proto_depIdxs = nil
}
