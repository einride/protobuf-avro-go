// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: einride/bigquery/public/v1/whos_on_first_geojson.proto

package publicv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Protobuf schema for the BigQuery public table:
//
//	bigquery-public-data.geo_whos_on_first.geojson
type WhosOnFirstGeoJson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Geoid                 string                 `protobuf:"bytes,1,opt,name=geoid,proto3" json:"geoid,omitempty"`                                                                // STRING NULLABLE
	Id                    int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`                                                                     // INTEGER NULLABLE
	Body                  *structpb.Struct       `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`                                                                  // STRING NULLABLE
	GeometryType          string                 `protobuf:"bytes,4,opt,name=geometry_type,json=geometryType,proto3" json:"geometry_type,omitempty"`                              // STRING NULLABLE
	BoundingBox           string                 `protobuf:"bytes,5,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`                                 // GEOGRAPHY NULLABLE
	Geom                  string                 `protobuf:"bytes,6,opt,name=geom,proto3" json:"geom,omitempty"`                                                                  // GEOGRAPHY NULLABLE
	LastModified          int64                  `protobuf:"varint,7,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`                             // INTEGER NULLABLE
	LastModifiedTimestamp *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=last_modified_timestamp,json=lastModifiedTimestamp,proto3" json:"last_modified_timestamp,omitempty"` // TIMESTAMP NULLABLE
}

func (x *WhosOnFirstGeoJson) Reset() {
	*x = WhosOnFirstGeoJson{}
	mi := &file_einride_bigquery_public_v1_whos_on_first_geojson_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WhosOnFirstGeoJson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhosOnFirstGeoJson) ProtoMessage() {}

func (x *WhosOnFirstGeoJson) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_whos_on_first_geojson_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhosOnFirstGeoJson.ProtoReflect.Descriptor instead.
func (*WhosOnFirstGeoJson) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDescGZIP(), []int{0}
}

func (x *WhosOnFirstGeoJson) GetGeoid() string {
	if x != nil {
		return x.Geoid
	}
	return ""
}

func (x *WhosOnFirstGeoJson) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WhosOnFirstGeoJson) GetBody() *structpb.Struct {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *WhosOnFirstGeoJson) GetGeometryType() string {
	if x != nil {
		return x.GeometryType
	}
	return ""
}

func (x *WhosOnFirstGeoJson) GetBoundingBox() string {
	if x != nil {
		return x.BoundingBox
	}
	return ""
}

func (x *WhosOnFirstGeoJson) GetGeom() string {
	if x != nil {
		return x.Geom
	}
	return ""
}

func (x *WhosOnFirstGeoJson) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

func (x *WhosOnFirstGeoJson) GetLastModifiedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModifiedTimestamp
	}
	return nil
}

var File_einride_bigquery_public_v1_whos_on_first_geojson_proto protoreflect.FileDescriptor

var file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDesc = []byte{
	0x0a, 0x36, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x68, 0x6f,
	0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x67, 0x65, 0x6f, 0x6a, 0x73,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x12, 0x57, 0x68, 0x6f, 0x73, 0x4f, 0x6e, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x47, 0x65, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65,
	0x6f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6f, 0x69, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2b, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62,
	0x6f, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x42, 0x6f, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x65, 0x6f, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x65, 0x6f, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x52,
	0x0a, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x15, 0x6c, 0x61, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x61,
	0x76, 0x72, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDescData = file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDesc
)

func file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDescData)
	})
	return file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDescData
}

var file_einride_bigquery_public_v1_whos_on_first_geojson_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_public_v1_whos_on_first_geojson_proto_goTypes = []any{
	(*WhosOnFirstGeoJson)(nil),    // 0: einride.bigquery.public.v1.WhosOnFirstGeoJson
	(*structpb.Struct)(nil),       // 1: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_einride_bigquery_public_v1_whos_on_first_geojson_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.public.v1.WhosOnFirstGeoJson.body:type_name -> google.protobuf.Struct
	2, // 1: einride.bigquery.public.v1.WhosOnFirstGeoJson.last_modified_timestamp:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_whos_on_first_geojson_proto_init() }
func file_einride_bigquery_public_v1_whos_on_first_geojson_proto_init() {
	if File_einride_bigquery_public_v1_whos_on_first_geojson_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_whos_on_first_geojson_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_whos_on_first_geojson_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_whos_on_first_geojson_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_whos_on_first_geojson_proto = out.File
	file_einride_bigquery_public_v1_whos_on_first_geojson_proto_rawDesc = nil
	file_einride_bigquery_public_v1_whos_on_first_geojson_proto_goTypes = nil
	file_einride_bigquery_public_v1_whos_on_first_geojson_proto_depIdxs = nil
}
