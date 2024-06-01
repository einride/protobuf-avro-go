// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: einride/avro/example/v1/example_timeofday.proto

package examplev1

import (
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
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

type ExampleTimeOfDay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeOfDay *timeofday.TimeOfDay `protobuf:"bytes,1,opt,name=time_of_day,json=timeOfDay,proto3" json:"time_of_day,omitempty"`
}

func (x *ExampleTimeOfDay) Reset() {
	*x = ExampleTimeOfDay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_avro_example_v1_example_timeofday_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleTimeOfDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleTimeOfDay) ProtoMessage() {}

func (x *ExampleTimeOfDay) ProtoReflect() protoreflect.Message {
	mi := &file_einride_avro_example_v1_example_timeofday_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleTimeOfDay.ProtoReflect.Descriptor instead.
func (*ExampleTimeOfDay) Descriptor() ([]byte, []int) {
	return file_einride_avro_example_v1_example_timeofday_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleTimeOfDay) GetTimeOfDay() *timeofday.TimeOfDay {
	if x != nil {
		return x.TimeOfDay
	}
	return nil
}

var File_einride_avro_example_v1_example_timeofday_proto protoreflect.FileDescriptor

var file_einride_avro_example_v1_example_timeofday_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x61, 0x76, 0x72, 0x6f, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x66, 0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x72, 0x6f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x66, 0x64, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x10, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x12, 0x36, 0x0a, 0x0b, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x66,
	0x44, 0x61, 0x79, 0x42, 0x5d, 0x5a, 0x5b, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d,
	0x61, 0x76, 0x72, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x61, 0x76, 0x72, 0x6f, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_avro_example_v1_example_timeofday_proto_rawDescOnce sync.Once
	file_einride_avro_example_v1_example_timeofday_proto_rawDescData = file_einride_avro_example_v1_example_timeofday_proto_rawDesc
)

func file_einride_avro_example_v1_example_timeofday_proto_rawDescGZIP() []byte {
	file_einride_avro_example_v1_example_timeofday_proto_rawDescOnce.Do(func() {
		file_einride_avro_example_v1_example_timeofday_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_avro_example_v1_example_timeofday_proto_rawDescData)
	})
	return file_einride_avro_example_v1_example_timeofday_proto_rawDescData
}

var file_einride_avro_example_v1_example_timeofday_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_avro_example_v1_example_timeofday_proto_goTypes = []interface{}{
	(*ExampleTimeOfDay)(nil),    // 0: einride.avro.example.v1.ExampleTimeOfDay
	(*timeofday.TimeOfDay)(nil), // 1: google.type.TimeOfDay
}
var file_einride_avro_example_v1_example_timeofday_proto_depIdxs = []int32{
	1, // 0: einride.avro.example.v1.ExampleTimeOfDay.time_of_day:type_name -> google.type.TimeOfDay
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_einride_avro_example_v1_example_timeofday_proto_init() }
func file_einride_avro_example_v1_example_timeofday_proto_init() {
	if File_einride_avro_example_v1_example_timeofday_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_avro_example_v1_example_timeofday_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleTimeOfDay); i {
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
			RawDescriptor: file_einride_avro_example_v1_example_timeofday_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_avro_example_v1_example_timeofday_proto_goTypes,
		DependencyIndexes: file_einride_avro_example_v1_example_timeofday_proto_depIdxs,
		MessageInfos:      file_einride_avro_example_v1_example_timeofday_proto_msgTypes,
	}.Build()
	File_einride_avro_example_v1_example_timeofday_proto = out.File
	file_einride_avro_example_v1_example_timeofday_proto_rawDesc = nil
	file_einride_avro_example_v1_example_timeofday_proto_goTypes = nil
	file_einride_avro_example_v1_example_timeofday_proto_depIdxs = nil
}
