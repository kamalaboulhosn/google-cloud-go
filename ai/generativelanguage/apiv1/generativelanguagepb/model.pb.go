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
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/ai/generativelanguage/v1/model.proto

package generativelanguagepb

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

// Information about a Generative Language Model.
type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the `Model`.
	//
	// Format: `models/{model}` with a `{model}` naming convention of:
	//
	// * "{base_model_id}-{version}"
	//
	// Examples:
	//
	// * `models/chat-bison-001`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the base model, pass this to the generation request.
	//
	// Examples:
	//
	// * `chat-bison`
	BaseModelId string `protobuf:"bytes,2,opt,name=base_model_id,json=baseModelId,proto3" json:"base_model_id,omitempty"`
	// Required. The version number of the model.
	//
	// This represents the major version
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// The human-readable name of the model. E.g. "Chat Bison".
	//
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A short description of the model.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Maximum number of input tokens allowed for this model.
	InputTokenLimit int32 `protobuf:"varint,6,opt,name=input_token_limit,json=inputTokenLimit,proto3" json:"input_token_limit,omitempty"`
	// Maximum number of output tokens available for this model.
	OutputTokenLimit int32 `protobuf:"varint,7,opt,name=output_token_limit,json=outputTokenLimit,proto3" json:"output_token_limit,omitempty"`
	// The model's supported generation methods.
	//
	// The method names are defined as Pascal case
	// strings, such as `generateMessage` which correspond to API methods.
	SupportedGenerationMethods []string `protobuf:"bytes,8,rep,name=supported_generation_methods,json=supportedGenerationMethods,proto3" json:"supported_generation_methods,omitempty"`
	// Controls the randomness of the output.
	//
	// Values can range over `[0.0,1.0]`, inclusive. A value closer to `1.0` will
	// produce responses that are more varied, while a value closer to `0.0` will
	// typically result in less surprising responses from the model.
	// This value specifies default to be used by the backend while making the
	// call to the model.
	Temperature *float32 `protobuf:"fixed32,9,opt,name=temperature,proto3,oneof" json:"temperature,omitempty"`
	// For Nucleus sampling.
	//
	// Nucleus sampling considers the smallest set of tokens whose probability
	// sum is at least `top_p`.
	// This value specifies default to be used by the backend while making the
	// call to the model.
	TopP *float32 `protobuf:"fixed32,10,opt,name=top_p,json=topP,proto3,oneof" json:"top_p,omitempty"`
	// For Top-k sampling.
	//
	// Top-k sampling considers the set of `top_k` most probable tokens.
	// This value specifies default to be used by the backend while making the
	// call to the model.
	// If empty, indicates the model doesn't use top-k sampling, and `top_k` isn't
	// allowed as a generation parameter.
	TopK *int32 `protobuf:"varint,11,opt,name=top_k,json=topK,proto3,oneof" json:"top_k,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ai_generativelanguage_v1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_google_ai_generativelanguage_v1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_google_ai_generativelanguage_v1_model_proto_rawDescGZIP(), []int{0}
}

func (x *Model) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Model) GetBaseModelId() string {
	if x != nil {
		return x.BaseModelId
	}
	return ""
}

func (x *Model) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Model) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Model) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Model) GetInputTokenLimit() int32 {
	if x != nil {
		return x.InputTokenLimit
	}
	return 0
}

func (x *Model) GetOutputTokenLimit() int32 {
	if x != nil {
		return x.OutputTokenLimit
	}
	return 0
}

func (x *Model) GetSupportedGenerationMethods() []string {
	if x != nil {
		return x.SupportedGenerationMethods
	}
	return nil
}

func (x *Model) GetTemperature() float32 {
	if x != nil && x.Temperature != nil {
		return *x.Temperature
	}
	return 0
}

func (x *Model) GetTopP() float32 {
	if x != nil && x.TopP != nil {
		return *x.TopP
	}
	return 0
}

func (x *Model) GetTopK() int32 {
	if x != nil && x.TopK != nil {
		return *x.TopK
	}
	return 0
}

var File_google_ai_generativelanguage_v1_model_proto protoreflect.FileDescriptor

var file_google_ai_generativelanguage_v1_model_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x04, 0x0a, 0x05, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x40, 0x0a, 0x1c, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1a, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x04,
	0x74, 0x6f, 0x70, 0x50, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x6b,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x4b, 0x88, 0x01,
	0x01, 0x3a, 0x3c, 0xea, 0x41, 0x39, 0x0a, 0x27, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x0e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x7d, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f,
	0x70, 0x5f, 0x6b, 0x42, 0x8e, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x70, 0x62, 0x3b,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ai_generativelanguage_v1_model_proto_rawDescOnce sync.Once
	file_google_ai_generativelanguage_v1_model_proto_rawDescData = file_google_ai_generativelanguage_v1_model_proto_rawDesc
)

func file_google_ai_generativelanguage_v1_model_proto_rawDescGZIP() []byte {
	file_google_ai_generativelanguage_v1_model_proto_rawDescOnce.Do(func() {
		file_google_ai_generativelanguage_v1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ai_generativelanguage_v1_model_proto_rawDescData)
	})
	return file_google_ai_generativelanguage_v1_model_proto_rawDescData
}

var file_google_ai_generativelanguage_v1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ai_generativelanguage_v1_model_proto_goTypes = []any{
	(*Model)(nil), // 0: google.ai.generativelanguage.v1.Model
}
var file_google_ai_generativelanguage_v1_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ai_generativelanguage_v1_model_proto_init() }
func file_google_ai_generativelanguage_v1_model_proto_init() {
	if File_google_ai_generativelanguage_v1_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ai_generativelanguage_v1_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Model); i {
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
	file_google_ai_generativelanguage_v1_model_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ai_generativelanguage_v1_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ai_generativelanguage_v1_model_proto_goTypes,
		DependencyIndexes: file_google_ai_generativelanguage_v1_model_proto_depIdxs,
		MessageInfos:      file_google_ai_generativelanguage_v1_model_proto_msgTypes,
	}.Build()
	File_google_ai_generativelanguage_v1_model_proto = out.File
	file_google_ai_generativelanguage_v1_model_proto_rawDesc = nil
	file_google_ai_generativelanguage_v1_model_proto_goTypes = nil
	file_google_ai_generativelanguage_v1_model_proto_depIdxs = nil
}
