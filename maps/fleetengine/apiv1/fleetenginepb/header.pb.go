// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/maps/fleetengine/v1/header.proto

package fleetenginepb

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

// Possible types of SDK.
type RequestHeader_SdkType int32

const (
	// The default value. This value is used if the `sdk_type` is omitted.
	RequestHeader_SDK_TYPE_UNSPECIFIED RequestHeader_SdkType = 0
	// The calling SDK is Consumer.
	RequestHeader_CONSUMER RequestHeader_SdkType = 1
	// The calling SDK is Driver.
	RequestHeader_DRIVER RequestHeader_SdkType = 2
	// The calling SDK is JavaScript.
	RequestHeader_JAVASCRIPT RequestHeader_SdkType = 3
)

// Enum value maps for RequestHeader_SdkType.
var (
	RequestHeader_SdkType_name = map[int32]string{
		0: "SDK_TYPE_UNSPECIFIED",
		1: "CONSUMER",
		2: "DRIVER",
		3: "JAVASCRIPT",
	}
	RequestHeader_SdkType_value = map[string]int32{
		"SDK_TYPE_UNSPECIFIED": 0,
		"CONSUMER":             1,
		"DRIVER":               2,
		"JAVASCRIPT":           3,
	}
)

func (x RequestHeader_SdkType) Enum() *RequestHeader_SdkType {
	p := new(RequestHeader_SdkType)
	*p = x
	return p
}

func (x RequestHeader_SdkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestHeader_SdkType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_fleetengine_v1_header_proto_enumTypes[0].Descriptor()
}

func (RequestHeader_SdkType) Type() protoreflect.EnumType {
	return &file_google_maps_fleetengine_v1_header_proto_enumTypes[0]
}

func (x RequestHeader_SdkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestHeader_SdkType.Descriptor instead.
func (RequestHeader_SdkType) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_fleetengine_v1_header_proto_rawDescGZIP(), []int{0, 0}
}

// The platform of the calling SDK.
type RequestHeader_Platform int32

const (
	// The default value. This value is used if the platform is omitted.
	RequestHeader_PLATFORM_UNSPECIFIED RequestHeader_Platform = 0
	// The request is coming from Android.
	RequestHeader_ANDROID RequestHeader_Platform = 1
	// The request is coming from iOS.
	RequestHeader_IOS RequestHeader_Platform = 2
	// The request is coming from the web.
	RequestHeader_WEB RequestHeader_Platform = 3
)

// Enum value maps for RequestHeader_Platform.
var (
	RequestHeader_Platform_name = map[int32]string{
		0: "PLATFORM_UNSPECIFIED",
		1: "ANDROID",
		2: "IOS",
		3: "WEB",
	}
	RequestHeader_Platform_value = map[string]int32{
		"PLATFORM_UNSPECIFIED": 0,
		"ANDROID":              1,
		"IOS":                  2,
		"WEB":                  3,
	}
)

func (x RequestHeader_Platform) Enum() *RequestHeader_Platform {
	p := new(RequestHeader_Platform)
	*p = x
	return p
}

func (x RequestHeader_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestHeader_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_fleetengine_v1_header_proto_enumTypes[1].Descriptor()
}

func (RequestHeader_Platform) Type() protoreflect.EnumType {
	return &file_google_maps_fleetengine_v1_header_proto_enumTypes[1]
}

func (x RequestHeader_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestHeader_Platform.Descriptor instead.
func (RequestHeader_Platform) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_fleetengine_v1_header_proto_rawDescGZIP(), []int{0, 1}
}

// A RequestHeader contains fields common to all Fleet Engine RPC requests.
type RequestHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The BCP-47 language code, such as en-US or sr-Latn. For more information,
	// see http://www.unicode.org/reports/tr35/#Unicode_locale_identifier. If none
	// is specified, the response may be in any language, with a preference for
	// English if such a name exists. Field value example: `en-US`.
	LanguageCode string `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Required. CLDR region code of the region where the request originates.
	// Field value example: `US`.
	RegionCode string `protobuf:"bytes,2,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// Version of the calling SDK, if applicable.
	// The version format is "major.minor.patch", example: `1.1.2`.
	SdkVersion string `protobuf:"bytes,3,opt,name=sdk_version,json=sdkVersion,proto3" json:"sdk_version,omitempty"`
	// Version of the operating system on which the calling SDK is running.
	// Field value examples: `4.4.1`, `12.1`.
	OsVersion string `protobuf:"bytes,4,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	// Model of the device on which the calling SDK is running.
	// Field value examples: `iPhone12,1`, `SM-G920F`.
	DeviceModel string `protobuf:"bytes,5,opt,name=device_model,json=deviceModel,proto3" json:"device_model,omitempty"`
	// The type of SDK sending the request.
	SdkType RequestHeader_SdkType `protobuf:"varint,6,opt,name=sdk_type,json=sdkType,proto3,enum=maps.fleetengine.v1.RequestHeader_SdkType" json:"sdk_type,omitempty"`
	// Version of the MapSDK which the calling SDK depends on, if applicable.
	// The version format is "major.minor.patch", example: `5.2.1`.
	MapsSdkVersion string `protobuf:"bytes,7,opt,name=maps_sdk_version,json=mapsSdkVersion,proto3" json:"maps_sdk_version,omitempty"`
	// Version of the NavSDK which the calling SDK depends on, if applicable.
	// The version format is "major.minor.patch", example: `2.1.0`.
	NavSdkVersion string `protobuf:"bytes,8,opt,name=nav_sdk_version,json=navSdkVersion,proto3" json:"nav_sdk_version,omitempty"`
	// Platform of the calling SDK.
	Platform RequestHeader_Platform `protobuf:"varint,9,opt,name=platform,proto3,enum=maps.fleetengine.v1.RequestHeader_Platform" json:"platform,omitempty"`
	// Manufacturer of the Android device from the calling SDK, only applicable
	// for the Android SDKs.
	// Field value example: `Samsung`.
	Manufacturer string `protobuf:"bytes,10,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	// Android API level of the calling SDK, only applicable for the Android SDKs.
	// Field value example: `23`.
	AndroidApiLevel int32 `protobuf:"varint,11,opt,name=android_api_level,json=androidApiLevel,proto3" json:"android_api_level,omitempty"`
	// Optional ID that can be provided for logging purposes in order to identify
	// the request.
	TraceId string `protobuf:"bytes,12,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *RequestHeader) Reset() {
	*x = RequestHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_fleetengine_v1_header_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeader) ProtoMessage() {}

func (x *RequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_fleetengine_v1_header_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeader.ProtoReflect.Descriptor instead.
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return file_google_maps_fleetengine_v1_header_proto_rawDescGZIP(), []int{0}
}

func (x *RequestHeader) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *RequestHeader) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *RequestHeader) GetSdkVersion() string {
	if x != nil {
		return x.SdkVersion
	}
	return ""
}

func (x *RequestHeader) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *RequestHeader) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *RequestHeader) GetSdkType() RequestHeader_SdkType {
	if x != nil {
		return x.SdkType
	}
	return RequestHeader_SDK_TYPE_UNSPECIFIED
}

func (x *RequestHeader) GetMapsSdkVersion() string {
	if x != nil {
		return x.MapsSdkVersion
	}
	return ""
}

func (x *RequestHeader) GetNavSdkVersion() string {
	if x != nil {
		return x.NavSdkVersion
	}
	return ""
}

func (x *RequestHeader) GetPlatform() RequestHeader_Platform {
	if x != nil {
		return x.Platform
	}
	return RequestHeader_PLATFORM_UNSPECIFIED
}

func (x *RequestHeader) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *RequestHeader) GetAndroidApiLevel() int32 {
	if x != nil {
		return x.AndroidApiLevel
	}
	return 0
}

func (x *RequestHeader) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

var File_google_maps_fleetengine_v1_header_proto protoreflect.FileDescriptor

var file_google_maps_fleetengine_v1_header_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x05, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x64, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x45, 0x0a, 0x08, 0x73, 0x64, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2a, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x64, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x73,
	0x64, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x70, 0x73, 0x5f, 0x73,
	0x64, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6d, 0x61, 0x70, 0x73, 0x53, 0x64, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x61, 0x76, 0x5f, 0x73, 0x64, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x61, 0x76, 0x53, 0x64,
	0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x41, 0x70, 0x69, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x07,
	0x53, 0x64, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x44, 0x4b, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x52, 0x49, 0x56, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4a,
	0x41, 0x56, 0x41, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x10, 0x03, 0x22, 0x43, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x54, 0x46,
	0x4f, 0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x03,
	0x42, 0x92, 0x01, 0x0a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73,
	0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x46, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x6d, 0x61, 0x70, 0x73, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x70, 0x62, 0x3b, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x70, 0x62, 0xa2, 0x02, 0x03, 0x43, 0x46, 0x45, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_fleetengine_v1_header_proto_rawDescOnce sync.Once
	file_google_maps_fleetengine_v1_header_proto_rawDescData = file_google_maps_fleetengine_v1_header_proto_rawDesc
)

func file_google_maps_fleetengine_v1_header_proto_rawDescGZIP() []byte {
	file_google_maps_fleetengine_v1_header_proto_rawDescOnce.Do(func() {
		file_google_maps_fleetengine_v1_header_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_fleetengine_v1_header_proto_rawDescData)
	})
	return file_google_maps_fleetengine_v1_header_proto_rawDescData
}

var file_google_maps_fleetengine_v1_header_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_maps_fleetengine_v1_header_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_fleetengine_v1_header_proto_goTypes = []interface{}{
	(RequestHeader_SdkType)(0),  // 0: maps.fleetengine.v1.RequestHeader.SdkType
	(RequestHeader_Platform)(0), // 1: maps.fleetengine.v1.RequestHeader.Platform
	(*RequestHeader)(nil),       // 2: maps.fleetengine.v1.RequestHeader
}
var file_google_maps_fleetengine_v1_header_proto_depIdxs = []int32{
	0, // 0: maps.fleetengine.v1.RequestHeader.sdk_type:type_name -> maps.fleetengine.v1.RequestHeader.SdkType
	1, // 1: maps.fleetengine.v1.RequestHeader.platform:type_name -> maps.fleetengine.v1.RequestHeader.Platform
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_maps_fleetengine_v1_header_proto_init() }
func file_google_maps_fleetengine_v1_header_proto_init() {
	if File_google_maps_fleetengine_v1_header_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_fleetengine_v1_header_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeader); i {
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
			RawDescriptor: file_google_maps_fleetengine_v1_header_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_fleetengine_v1_header_proto_goTypes,
		DependencyIndexes: file_google_maps_fleetengine_v1_header_proto_depIdxs,
		EnumInfos:         file_google_maps_fleetengine_v1_header_proto_enumTypes,
		MessageInfos:      file_google_maps_fleetengine_v1_header_proto_msgTypes,
	}.Build()
	File_google_maps_fleetengine_v1_header_proto = out.File
	file_google_maps_fleetengine_v1_header_proto_rawDesc = nil
	file_google_maps_fleetengine_v1_header_proto_goTypes = nil
	file_google_maps_fleetengine_v1_header_proto_depIdxs = nil
}
