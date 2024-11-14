// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: einride/avro/example/v1/example_map.proto

package examplev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExampleMap_Enum int32

const (
	ExampleMap_ENUM_UNSPECIFIED ExampleMap_Enum = 0
	ExampleMap_ENUM_VALUE1      ExampleMap_Enum = 1
	ExampleMap_ENUM_VALUE2      ExampleMap_Enum = 2
)

// Enum value maps for ExampleMap_Enum.
var (
	ExampleMap_Enum_name = map[int32]string{
		0: "ENUM_UNSPECIFIED",
		1: "ENUM_VALUE1",
		2: "ENUM_VALUE2",
	}
	ExampleMap_Enum_value = map[string]int32{
		"ENUM_UNSPECIFIED": 0,
		"ENUM_VALUE1":      1,
		"ENUM_VALUE2":      2,
	}
)

func (x ExampleMap_Enum) Enum() *ExampleMap_Enum {
	p := new(ExampleMap_Enum)
	*p = x
	return p
}

func (x ExampleMap_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExampleMap_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_avro_example_v1_example_map_proto_enumTypes[0].Descriptor()
}

func (ExampleMap_Enum) Type() protoreflect.EnumType {
	return &file_einride_avro_example_v1_example_map_proto_enumTypes[0]
}

func (x ExampleMap_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExampleMap_Enum.Descriptor instead.
func (ExampleMap_Enum) EnumDescriptor() ([]byte, []int) {
	return file_einride_avro_example_v1_example_map_proto_rawDescGZIP(), []int{0, 0}
}

type ExampleMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringToString     map[string]string                 `protobuf:"bytes,1,rep,name=string_to_string,json=stringToString,proto3" json:"string_to_string,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StringToNested     map[string]*ExampleMap_Nested     `protobuf:"bytes,2,rep,name=string_to_nested,json=stringToNested,proto3" json:"string_to_nested,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StringToEnum       map[string]ExampleMap_Enum        `protobuf:"bytes,3,rep,name=string_to_enum,json=stringToEnum,proto3" json:"string_to_enum,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=einride.avro.example.v1.ExampleMap_Enum"`
	Int32ToString      map[int32]string                  `protobuf:"bytes,4,rep,name=int32_to_string,json=int32ToString,proto3" json:"int32_to_string,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Int64ToString      map[int64]string                  `protobuf:"bytes,5,rep,name=int64_to_string,json=int64ToString,proto3" json:"int64_to_string,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Uint32ToString     map[uint32]string                 `protobuf:"bytes,6,rep,name=uint32_to_string,json=uint32ToString,proto3" json:"uint32_to_string,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BoolToString       map[bool]string                   `protobuf:"bytes,7,rep,name=bool_to_string,json=boolToString,proto3" json:"bool_to_string,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StringToFloatValue map[string]*wrapperspb.FloatValue `protobuf:"bytes,8,rep,name=string_to_float_value,json=stringToFloatValue,proto3" json:"string_to_float_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExampleMap) Reset() {
	*x = ExampleMap{}
	mi := &file_einride_avro_example_v1_example_map_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleMap) ProtoMessage() {}

func (x *ExampleMap) ProtoReflect() protoreflect.Message {
	mi := &file_einride_avro_example_v1_example_map_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleMap.ProtoReflect.Descriptor instead.
func (*ExampleMap) Descriptor() ([]byte, []int) {
	return file_einride_avro_example_v1_example_map_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleMap) GetStringToString() map[string]string {
	if x != nil {
		return x.StringToString
	}
	return nil
}

func (x *ExampleMap) GetStringToNested() map[string]*ExampleMap_Nested {
	if x != nil {
		return x.StringToNested
	}
	return nil
}

func (x *ExampleMap) GetStringToEnum() map[string]ExampleMap_Enum {
	if x != nil {
		return x.StringToEnum
	}
	return nil
}

func (x *ExampleMap) GetInt32ToString() map[int32]string {
	if x != nil {
		return x.Int32ToString
	}
	return nil
}

func (x *ExampleMap) GetInt64ToString() map[int64]string {
	if x != nil {
		return x.Int64ToString
	}
	return nil
}

func (x *ExampleMap) GetUint32ToString() map[uint32]string {
	if x != nil {
		return x.Uint32ToString
	}
	return nil
}

func (x *ExampleMap) GetBoolToString() map[bool]string {
	if x != nil {
		return x.BoolToString
	}
	return nil
}

func (x *ExampleMap) GetStringToFloatValue() map[string]*wrapperspb.FloatValue {
	if x != nil {
		return x.StringToFloatValue
	}
	return nil
}

type ExampleMap_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringToString map[string]string `protobuf:"bytes,1,rep,name=string_to_string,json=stringToString,proto3" json:"string_to_string,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExampleMap_Nested) Reset() {
	*x = ExampleMap_Nested{}
	mi := &file_einride_avro_example_v1_example_map_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleMap_Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleMap_Nested) ProtoMessage() {}

func (x *ExampleMap_Nested) ProtoReflect() protoreflect.Message {
	mi := &file_einride_avro_example_v1_example_map_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleMap_Nested.ProtoReflect.Descriptor instead.
func (*ExampleMap_Nested) Descriptor() ([]byte, []int) {
	return file_einride_avro_example_v1_example_map_proto_rawDescGZIP(), []int{0, 8}
}

func (x *ExampleMap_Nested) GetStringToString() map[string]string {
	if x != nil {
		return x.StringToString
	}
	return nil
}

var File_einride_avro_example_v1_example_map_proto protoreflect.FileDescriptor

var file_einride_avro_example_v1_example_map_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x61, 0x76, 0x72, 0x6f, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x0d, 0x0a, 0x0a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x4d, 0x61, 0x70, 0x12, 0x61, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d,
	0x61, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x61, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x6f, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x54, 0x6f, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x5b, 0x0a, 0x0e, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x45,
	0x6e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x54, 0x6f, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x5e, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f,
	0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x4d, 0x61, 0x70, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x6f,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x5e, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f,
	0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x4d, 0x61, 0x70, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x61, 0x0a, 0x10, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x5b, 0x0a, 0x0e, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x54, 0x6f, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x54, 0x6f,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x6e, 0x0a, 0x15, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x6f, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x61, 0x76, 0x72, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x54, 0x6f, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x12, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x41, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x6d, 0x0a, 0x13, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x54, 0x6f, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x40, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x69, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x54, 0x6f, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x4d, 0x61, 0x70, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x13, 0x55, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x42, 0x6f,
	0x6f, 0x6c, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x62, 0x0a, 0x17, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0xb5, 0x01, 0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x68, 0x0a, 0x10, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61,
	0x76, 0x72, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x41, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x31, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x32, 0x10, 0x02, 0x42, 0x5d, 0x5a, 0x5b, 0x67, 0x6f, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2d, 0x61, 0x76, 0x72, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x61, 0x76, 0x72,
	0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_avro_example_v1_example_map_proto_rawDescOnce sync.Once
	file_einride_avro_example_v1_example_map_proto_rawDescData = file_einride_avro_example_v1_example_map_proto_rawDesc
)

func file_einride_avro_example_v1_example_map_proto_rawDescGZIP() []byte {
	file_einride_avro_example_v1_example_map_proto_rawDescOnce.Do(func() {
		file_einride_avro_example_v1_example_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_avro_example_v1_example_map_proto_rawDescData)
	})
	return file_einride_avro_example_v1_example_map_proto_rawDescData
}

var file_einride_avro_example_v1_example_map_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_einride_avro_example_v1_example_map_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_einride_avro_example_v1_example_map_proto_goTypes = []any{
	(ExampleMap_Enum)(0),          // 0: einride.avro.example.v1.ExampleMap.Enum
	(*ExampleMap)(nil),            // 1: einride.avro.example.v1.ExampleMap
	nil,                           // 2: einride.avro.example.v1.ExampleMap.StringToStringEntry
	nil,                           // 3: einride.avro.example.v1.ExampleMap.StringToNestedEntry
	nil,                           // 4: einride.avro.example.v1.ExampleMap.StringToEnumEntry
	nil,                           // 5: einride.avro.example.v1.ExampleMap.Int32ToStringEntry
	nil,                           // 6: einride.avro.example.v1.ExampleMap.Int64ToStringEntry
	nil,                           // 7: einride.avro.example.v1.ExampleMap.Uint32ToStringEntry
	nil,                           // 8: einride.avro.example.v1.ExampleMap.BoolToStringEntry
	nil,                           // 9: einride.avro.example.v1.ExampleMap.StringToFloatValueEntry
	(*ExampleMap_Nested)(nil),     // 10: einride.avro.example.v1.ExampleMap.Nested
	nil,                           // 11: einride.avro.example.v1.ExampleMap.Nested.StringToStringEntry
	(*wrapperspb.FloatValue)(nil), // 12: google.protobuf.FloatValue
}
var file_einride_avro_example_v1_example_map_proto_depIdxs = []int32{
	2,  // 0: einride.avro.example.v1.ExampleMap.string_to_string:type_name -> einride.avro.example.v1.ExampleMap.StringToStringEntry
	3,  // 1: einride.avro.example.v1.ExampleMap.string_to_nested:type_name -> einride.avro.example.v1.ExampleMap.StringToNestedEntry
	4,  // 2: einride.avro.example.v1.ExampleMap.string_to_enum:type_name -> einride.avro.example.v1.ExampleMap.StringToEnumEntry
	5,  // 3: einride.avro.example.v1.ExampleMap.int32_to_string:type_name -> einride.avro.example.v1.ExampleMap.Int32ToStringEntry
	6,  // 4: einride.avro.example.v1.ExampleMap.int64_to_string:type_name -> einride.avro.example.v1.ExampleMap.Int64ToStringEntry
	7,  // 5: einride.avro.example.v1.ExampleMap.uint32_to_string:type_name -> einride.avro.example.v1.ExampleMap.Uint32ToStringEntry
	8,  // 6: einride.avro.example.v1.ExampleMap.bool_to_string:type_name -> einride.avro.example.v1.ExampleMap.BoolToStringEntry
	9,  // 7: einride.avro.example.v1.ExampleMap.string_to_float_value:type_name -> einride.avro.example.v1.ExampleMap.StringToFloatValueEntry
	10, // 8: einride.avro.example.v1.ExampleMap.StringToNestedEntry.value:type_name -> einride.avro.example.v1.ExampleMap.Nested
	0,  // 9: einride.avro.example.v1.ExampleMap.StringToEnumEntry.value:type_name -> einride.avro.example.v1.ExampleMap.Enum
	12, // 10: einride.avro.example.v1.ExampleMap.StringToFloatValueEntry.value:type_name -> google.protobuf.FloatValue
	11, // 11: einride.avro.example.v1.ExampleMap.Nested.string_to_string:type_name -> einride.avro.example.v1.ExampleMap.Nested.StringToStringEntry
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_einride_avro_example_v1_example_map_proto_init() }
func file_einride_avro_example_v1_example_map_proto_init() {
	if File_einride_avro_example_v1_example_map_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_einride_avro_example_v1_example_map_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_avro_example_v1_example_map_proto_goTypes,
		DependencyIndexes: file_einride_avro_example_v1_example_map_proto_depIdxs,
		EnumInfos:         file_einride_avro_example_v1_example_map_proto_enumTypes,
		MessageInfos:      file_einride_avro_example_v1_example_map_proto_msgTypes,
	}.Build()
	File_einride_avro_example_v1_example_map_proto = out.File
	file_einride_avro_example_v1_example_map_proto_rawDesc = nil
	file_einride_avro_example_v1_example_map_proto_goTypes = nil
	file_einride_avro_example_v1_example_map_proto_depIdxs = nil
}
