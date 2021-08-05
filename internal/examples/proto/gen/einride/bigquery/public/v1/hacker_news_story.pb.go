// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.2
// source: einride/bigquery/public/v1/hacker_news_story.proto

package publicv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
//  bigquery-public-data.hacker_news.stories
type HackerNewsStory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                      // INTEGER NULLABLE
	By          string                 `protobuf:"bytes,2,opt,name=by,proto3" json:"by,omitempty"`                       // STRING NULLABLE
	Score       int32                  `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`                // INTEGER NULLABLE
	Time        int64                  `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`                  // INTEGER NULLABLE
	TimeTs      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time_ts,json=timeTs,proto3" json:"time_ts,omitempty"` // TIMESTAMP NULLABLE
	Title       string                 `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`                 // STRING NULLABLE
	Url         string                 `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`                     // STRING NULLABLE
	Text        string                 `protobuf:"bytes,8,opt,name=text,proto3" json:"text,omitempty"`                   // STRING NULLABLE
	Deleted     bool                   `protobuf:"varint,9,opt,name=deleted,proto3" json:"deleted,omitempty"`            // BOOLEAN NULLABLE
	Dead        bool                   `protobuf:"varint,10,opt,name=dead,proto3" json:"dead,omitempty"`                 // BOOLEAN NULLABLE
	Descendants int32                  `protobuf:"varint,11,opt,name=descendants,proto3" json:"descendants,omitempty"`   // INTEGER NULLABLE
	Author      string                 `protobuf:"bytes,12,opt,name=author,proto3" json:"author,omitempty"`              // STRING NULLABLE
}

func (x *HackerNewsStory) Reset() {
	*x = HackerNewsStory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_public_v1_hacker_news_story_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HackerNewsStory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HackerNewsStory) ProtoMessage() {}

func (x *HackerNewsStory) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_hacker_news_story_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HackerNewsStory.ProtoReflect.Descriptor instead.
func (*HackerNewsStory) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_hacker_news_story_proto_rawDescGZIP(), []int{0}
}

func (x *HackerNewsStory) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HackerNewsStory) GetBy() string {
	if x != nil {
		return x.By
	}
	return ""
}

func (x *HackerNewsStory) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *HackerNewsStory) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *HackerNewsStory) GetTimeTs() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeTs
	}
	return nil
}

func (x *HackerNewsStory) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *HackerNewsStory) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HackerNewsStory) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *HackerNewsStory) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *HackerNewsStory) GetDead() bool {
	if x != nil {
		return x.Dead
	}
	return false
}

func (x *HackerNewsStory) GetDescendants() int32 {
	if x != nil {
		return x.Descendants
	}
	return 0
}

func (x *HackerNewsStory) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

var File_einride_bigquery_public_v1_hacker_news_story_proto protoreflect.FileDescriptor

var file_einride_bigquery_public_v1_hacker_news_story_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x0f, 0x48, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x62, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x74, 0x69,
	0x6d, 0x65, 0x54, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65,
	0x61, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x61, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x65, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x61, 0x76, 0x72, 0x6f, 0x2d, 0x67, 0x6f, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_public_v1_hacker_news_story_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_hacker_news_story_proto_rawDescData = file_einride_bigquery_public_v1_hacker_news_story_proto_rawDesc
)

func file_einride_bigquery_public_v1_hacker_news_story_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_hacker_news_story_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_hacker_news_story_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_public_v1_hacker_news_story_proto_rawDescData)
	})
	return file_einride_bigquery_public_v1_hacker_news_story_proto_rawDescData
}

var file_einride_bigquery_public_v1_hacker_news_story_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_public_v1_hacker_news_story_proto_goTypes = []interface{}{
	(*HackerNewsStory)(nil),       // 0: einride.bigquery.public.v1.HackerNewsStory
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_einride_bigquery_public_v1_hacker_news_story_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.public.v1.HackerNewsStory.time_ts:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_hacker_news_story_proto_init() }
func file_einride_bigquery_public_v1_hacker_news_story_proto_init() {
	if File_einride_bigquery_public_v1_hacker_news_story_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_bigquery_public_v1_hacker_news_story_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HackerNewsStory); i {
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
			RawDescriptor: file_einride_bigquery_public_v1_hacker_news_story_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_hacker_news_story_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_hacker_news_story_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_hacker_news_story_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_hacker_news_story_proto = out.File
	file_einride_bigquery_public_v1_hacker_news_story_proto_rawDesc = nil
	file_einride_bigquery_public_v1_hacker_news_story_proto_goTypes = nil
	file_einride_bigquery_public_v1_hacker_news_story_proto_depIdxs = nil
}
