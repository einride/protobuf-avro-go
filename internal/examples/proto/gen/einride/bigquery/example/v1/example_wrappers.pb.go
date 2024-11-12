// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: einride/bigquery/example/v1/example_wrappers.proto

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

type ExampleWrappers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FloatValue  *wrapperspb.FloatValue  `protobuf:"bytes,1,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	DoubleValue *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	StringValue *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	BytesValue  *wrapperspb.BytesValue  `protobuf:"bytes,4,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
	Int32Value  *wrapperspb.Int32Value  `protobuf:"bytes,5,opt,name=int32_value,json=int32Value,proto3" json:"int32_value,omitempty"`
	Int64Value  *wrapperspb.Int64Value  `protobuf:"bytes,6,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	Uint32Value *wrapperspb.UInt32Value `protobuf:"bytes,7,opt,name=uint32_value,json=uint32Value,proto3" json:"uint32_value,omitempty"`
	Uint64Value *wrapperspb.UInt64Value `protobuf:"bytes,8,opt,name=uint64_value,json=uint64Value,proto3" json:"uint64_value,omitempty"`
	BoolValue   *wrapperspb.BoolValue   `protobuf:"bytes,9,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
}

func (x *ExampleWrappers) Reset() {
	*x = ExampleWrappers{}
	mi := &file_einride_bigquery_example_v1_example_wrappers_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleWrappers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleWrappers) ProtoMessage() {}

func (x *ExampleWrappers) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_wrappers_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleWrappers.ProtoReflect.Descriptor instead.
func (*ExampleWrappers) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_wrappers_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleWrappers) GetFloatValue() *wrapperspb.FloatValue {
	if x != nil {
		return x.FloatValue
	}
	return nil
}

func (x *ExampleWrappers) GetDoubleValue() *wrapperspb.DoubleValue {
	if x != nil {
		return x.DoubleValue
	}
	return nil
}

func (x *ExampleWrappers) GetStringValue() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue
	}
	return nil
}

func (x *ExampleWrappers) GetBytesValue() *wrapperspb.BytesValue {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

func (x *ExampleWrappers) GetInt32Value() *wrapperspb.Int32Value {
	if x != nil {
		return x.Int32Value
	}
	return nil
}

func (x *ExampleWrappers) GetInt64Value() *wrapperspb.Int64Value {
	if x != nil {
		return x.Int64Value
	}
	return nil
}

func (x *ExampleWrappers) GetUint32Value() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Uint32Value
	}
	return nil
}

func (x *ExampleWrappers) GetUint64Value() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Uint64Value
	}
	return nil
}

func (x *ExampleWrappers) GetBoolValue() *wrapperspb.BoolValue {
	if x != nil {
		return x.BoolValue
	}
	return nil
}

var File_einride_bigquery_example_v1_example_wrappers_proto protoreflect.FileDescriptor

var file_einride_bigquery_example_v1_example_wrappers_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc8, 0x04, 0x0a, 0x0f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x57, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x3f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x61, 0x5a, 0x5f,
	0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x61, 0x76, 0x72, 0x6f, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_example_v1_example_wrappers_proto_rawDescOnce sync.Once
	file_einride_bigquery_example_v1_example_wrappers_proto_rawDescData = file_einride_bigquery_example_v1_example_wrappers_proto_rawDesc
)

func file_einride_bigquery_example_v1_example_wrappers_proto_rawDescGZIP() []byte {
	file_einride_bigquery_example_v1_example_wrappers_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_example_v1_example_wrappers_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_example_v1_example_wrappers_proto_rawDescData)
	})
	return file_einride_bigquery_example_v1_example_wrappers_proto_rawDescData
}

var file_einride_bigquery_example_v1_example_wrappers_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_example_v1_example_wrappers_proto_goTypes = []any{
	(*ExampleWrappers)(nil),        // 0: einride.bigquery.example.v1.ExampleWrappers
	(*wrapperspb.FloatValue)(nil),  // 1: google.protobuf.FloatValue
	(*wrapperspb.DoubleValue)(nil), // 2: google.protobuf.DoubleValue
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*wrapperspb.BytesValue)(nil),  // 4: google.protobuf.BytesValue
	(*wrapperspb.Int32Value)(nil),  // 5: google.protobuf.Int32Value
	(*wrapperspb.Int64Value)(nil),  // 6: google.protobuf.Int64Value
	(*wrapperspb.UInt32Value)(nil), // 7: google.protobuf.UInt32Value
	(*wrapperspb.UInt64Value)(nil), // 8: google.protobuf.UInt64Value
	(*wrapperspb.BoolValue)(nil),   // 9: google.protobuf.BoolValue
}
var file_einride_bigquery_example_v1_example_wrappers_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.example.v1.ExampleWrappers.float_value:type_name -> google.protobuf.FloatValue
	2, // 1: einride.bigquery.example.v1.ExampleWrappers.double_value:type_name -> google.protobuf.DoubleValue
	3, // 2: einride.bigquery.example.v1.ExampleWrappers.string_value:type_name -> google.protobuf.StringValue
	4, // 3: einride.bigquery.example.v1.ExampleWrappers.bytes_value:type_name -> google.protobuf.BytesValue
	5, // 4: einride.bigquery.example.v1.ExampleWrappers.int32_value:type_name -> google.protobuf.Int32Value
	6, // 5: einride.bigquery.example.v1.ExampleWrappers.int64_value:type_name -> google.protobuf.Int64Value
	7, // 6: einride.bigquery.example.v1.ExampleWrappers.uint32_value:type_name -> google.protobuf.UInt32Value
	8, // 7: einride.bigquery.example.v1.ExampleWrappers.uint64_value:type_name -> google.protobuf.UInt64Value
	9, // 8: einride.bigquery.example.v1.ExampleWrappers.bool_value:type_name -> google.protobuf.BoolValue
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_einride_bigquery_example_v1_example_wrappers_proto_init() }
func file_einride_bigquery_example_v1_example_wrappers_proto_init() {
	if File_einride_bigquery_example_v1_example_wrappers_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_einride_bigquery_example_v1_example_wrappers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_example_v1_example_wrappers_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_example_v1_example_wrappers_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_example_v1_example_wrappers_proto_msgTypes,
	}.Build()
	File_einride_bigquery_example_v1_example_wrappers_proto = out.File
	file_einride_bigquery_example_v1_example_wrappers_proto_rawDesc = nil
	file_einride_bigquery_example_v1_example_wrappers_proto_goTypes = nil
	file_einride_bigquery_example_v1_example_wrappers_proto_depIdxs = nil
}
