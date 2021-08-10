package protoavro

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/linkedin/goavro/v2"
	examplev1 "go.einride.tech/protobuf-avro/internal/examples/proto/gen/einride/avro/example/v1"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/genproto/googleapis/type/timeofday"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"
)

func Test_JSON(t *testing.T) {
	for _, tt := range []struct {
		name     string
		msg      proto.Message
		expected map[string]interface{}
	}{
		{
			name: "library.Book",
			msg: &library.Book{
				Name:   "books/1",
				Author: "J. K. Rowling",
				Title:  "Harry Potter",
				Read:   true,
			},
			expected: map[string]interface{}{
				"google.example.library.v1.Book": map[string]interface{}{
					"name":   map[string]interface{}{"string": "books/1"},
					"author": map[string]interface{}{"string": "J. K. Rowling"},
					"title":  map[string]interface{}{"string": "Harry Potter"},
					"read":   map[string]interface{}{"boolean": true},
				},
			},
		},
		{
			name: "library.UpdateBookRequest",
			msg: &library.UpdateBookRequest{
				Book: &library.Book{
					Name:   "books/1",
					Author: "J. K. Rowling",
					Title:  "Harry Potter",
					Read:   false,
				},
			},
			expected: map[string]interface{}{
				"google.example.library.v1.UpdateBookRequest": map[string]interface{}{
					"book": map[string]interface{}{
						"google.example.library.v1.Book": map[string]interface{}{
							"name":   map[string]interface{}{"string": "books/1"},
							"author": map[string]interface{}{"string": "J. K. Rowling"},
							"title":  map[string]interface{}{"string": "Harry Potter"},
							"read":   map[string]interface{}{"boolean": false},
						},
					},
					"update_mask": nil,
				},
			},
		},
		{
			name: "examplev1.ExampleAny",
			msg: &examplev1.ExampleAny{
				Any: mustAny(t, &examplev1.ExampleEnum{
					EnumValue: examplev1.ExampleEnum_ENUM_VALUE1,
				}),
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleAny": map[string]interface{}{
					"any": map[string]interface{}{
						"string": stableAny(t, `{
	"@type": "type.googleapis.com/einride.avro.example.v1.ExampleEnum",
	"enumValue": "ENUM_VALUE1"
}`),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleBytes",
			msg: &examplev1.ExampleBytes{
				Bytes: []byte{1, 2, 3},
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleBytes": map[string]interface{}{
					"bytes": map[string]interface{}{"bytes": []byte{1, 2, 3}},
				},
			},
		},
		{
			name: "examplev1.ExampleDate",
			msg: &examplev1.ExampleDate{
				Date: &date.Date{Year: 2021, Month: 6, Day: 27},
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleDate": map[string]interface{}{
					"date": map[string]interface{}{
						"int.date": int32(18805),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleDateTime",
			msg: &examplev1.ExampleDateTime{
				DateTime: &datetime.DateTime{
					Year:    2021,
					Month:   6,
					Day:     27,
					Hours:   18,
					Minutes: 6,
					Seconds: 30,
					Nanos:   2,
					TimeOffset: &datetime.DateTime_UtcOffset{
						UtcOffset: durationpb.New(time.Hour),
					},
				},
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleDateTime": map[string]interface{}{
					"date_time": map[string]interface{}{
						"google.type.DateTime": map[string]interface{}{
							"year":      map[string]interface{}{"int": int32(2021)},
							"month":     map[string]interface{}{"int": int32(6)},
							"day":       map[string]interface{}{"int": int32(27)},
							"hours":     map[string]interface{}{"int": int32(18)},
							"minutes":   map[string]interface{}{"int": int32(6)},
							"seconds":   map[string]interface{}{"int": int32(30)},
							"nanos":     map[string]interface{}{"int": int32(2)},
							"time_zone": nil,
							"utc_offset": map[string]interface{}{
								"float": float64(3600),
							},
						},
					},
				},
			},
		},
		{
			name: "examplev1.ExampleDuration",
			msg: &examplev1.ExampleDuration{
				Duration: durationpb.New(time.Hour * 3),
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleDuration": map[string]interface{}{
					"duration": map[string]interface{}{
						"float": float64(10800),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleEnum",
			msg: &examplev1.ExampleEnum{
				EnumValue: examplev1.ExampleEnum_ENUM_VALUE2,
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleEnum": map[string]interface{}{
					"enum_value": map[string]interface{}{"einride.avro.example.v1.ExampleEnum.Enum": "ENUM_VALUE2"},
				},
			},
		},
		{
			name: "examplev1.ExampleList",
			msg: &examplev1.ExampleList{
				Int64List:  []int64{1, 2, 3},
				StringList: []string{"1", "2", "3"},
				EnumList: []examplev1.ExampleList_Enum{
					examplev1.ExampleList_ENUM_UNSPECIFIED,
					examplev1.ExampleList_ENUM_VALUE1,
					examplev1.ExampleList_ENUM_VALUE2,
				},
				NestedList: []*examplev1.ExampleList_Nested{
					{StringList: []string{"1", "2", "3"}},
					{StringList: []string{"4", "5", "6"}},
				},
				FloatValueList: []*wrapperspb.FloatValue{
					nil,
					{Value: 1},
					{Value: 2},
					{Value: 3},
				},
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleList": map[string]interface{}{
					"int64_list": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{"long": int64(1)},
							map[string]interface{}{"long": int64(2)},
							map[string]interface{}{"long": int64(3)},
						},
					},
					"string_list": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{"string": "1"},
							map[string]interface{}{"string": "2"},
							map[string]interface{}{"string": "3"},
						},
					},
					"enum_list": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{"einride.avro.example.v1.ExampleList.Enum": "ENUM_UNSPECIFIED"},
							map[string]interface{}{"einride.avro.example.v1.ExampleList.Enum": "ENUM_VALUE1"},
							map[string]interface{}{"einride.avro.example.v1.ExampleList.Enum": "ENUM_VALUE2"},
						},
					},
					"nested_list": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{
								"einride.avro.example.v1.ExampleList.Nested": map[string]interface{}{
									"string_list": map[string]interface{}{
										"array": []interface{}{
											map[string]interface{}{"string": "1"},
											map[string]interface{}{"string": "2"},
											map[string]interface{}{"string": "3"},
										},
									},
								},
							},
							map[string]interface{}{
								"einride.avro.example.v1.ExampleList.Nested": map[string]interface{}{
									"string_list": map[string]interface{}{
										"array": []interface{}{
											map[string]interface{}{"string": "4"},
											map[string]interface{}{"string": "5"},
											map[string]interface{}{"string": "6"},
										},
									},
								},
							},
						},
					},
					"float_value_list": map[string]interface{}{
						"array": []interface{}{
							nil,
							map[string]interface{}{"float": float32(1)},
							map[string]interface{}{"float": float32(2)},
							map[string]interface{}{"float": float32(3)},
						},
					},
				},
			},
		},
		{
			name: "examplev1.ExampleMap",
			msg: &examplev1.ExampleMap{
				StringToString: map[string]string{
					"1": "a",
					"2": "b",
				},
				StringToNested: map[string]*examplev1.ExampleMap_Nested{
					"1": {
						StringToString: map[string]string{
							"1": "a",
							"2": "b",
						},
					},
					"3": {
						StringToString: map[string]string{
							"3": "c",
							"4": "d",
						},
					},
				},
				StringToEnum: map[string]examplev1.ExampleMap_Enum{
					"0": examplev1.ExampleMap_ENUM_UNSPECIFIED,
					"1": examplev1.ExampleMap_ENUM_VALUE1,
				},
				Int32ToString: map[int32]string{
					1: "a",
					2: "b",
				},
				Int64ToString: map[int64]string{
					1: "a",
					2: "b",
				},
				Uint32ToString: map[uint32]string{
					1: "a",
					2: "b",
				},
				BoolToString: map[bool]string{
					true:  "a",
					false: "b",
				},
				StringToFloatValue: map[string]*wrapperspb.FloatValue{
					"1": nil,
					"2": wrapperspb.Float(2),
				},
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleMap": map[string]interface{}{
					"string_to_string": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{
								"key":   map[string]interface{}{"string": "1"},
								"value": map[string]interface{}{"string": "a"},
							},
							map[string]interface{}{
								"key":   map[string]interface{}{"string": "2"},
								"value": map[string]interface{}{"string": "b"},
							},
						},
					},
					"string_to_nested": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{
								"key": map[string]interface{}{"string": "1"},
								"value": map[string]interface{}{
									"einride.avro.example.v1.ExampleMap.Nested": map[string]interface{}{
										"string_to_string": map[string]interface{}{
											"array": []interface{}{
												map[string]interface{}{
													"key":   map[string]interface{}{"string": "1"},
													"value": map[string]interface{}{"string": "a"},
												},
												map[string]interface{}{
													"key":   map[string]interface{}{"string": "2"},
													"value": map[string]interface{}{"string": "b"},
												},
											},
										},
									},
								},
							},
							map[string]interface{}{
								"key": map[string]interface{}{"string": "3"},
								"value": map[string]interface{}{
									"einride.avro.example.v1.ExampleMap.Nested": map[string]interface{}{
										"string_to_string": map[string]interface{}{
											"array": []interface{}{
												map[string]interface{}{
													"key":   map[string]interface{}{"string": "3"},
													"value": map[string]interface{}{"string": "c"},
												},
												map[string]interface{}{
													"key":   map[string]interface{}{"string": "4"},
													"value": map[string]interface{}{"string": "d"},
												},
											},
										},
									},
								},
							},
						},
					},
					"string_to_enum": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{
								"key":   map[string]interface{}{"string": "0"},
								"value": map[string]interface{}{"einride.avro.example.v1.ExampleMap.Enum": "ENUM_UNSPECIFIED"},
							},
							map[string]interface{}{
								"key":   map[string]interface{}{"string": "1"},
								"value": map[string]interface{}{"einride.avro.example.v1.ExampleMap.Enum": "ENUM_VALUE1"},
							},
						},
					},
					"int32_to_string": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{
								"key":   map[string]interface{}{"int": int32(1)},
								"value": map[string]interface{}{"string": "a"},
							},
							map[string]interface{}{
								"key":   map[string]interface{}{"int": int32(2)},
								"value": map[string]interface{}{"string": "b"},
							},
						},
					},
					"int64_to_string": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{
								"key":   map[string]interface{}{"long": int64(1)},
								"value": map[string]interface{}{"string": "a"},
							},
							map[string]interface{}{
								"key":   map[string]interface{}{"long": int64(2)},
								"value": map[string]interface{}{"string": "b"},
							},
						},
					},
					"uint32_to_string": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{
								"key":   map[string]interface{}{"int": int32(1)},
								"value": map[string]interface{}{"string": "a"},
							},
							map[string]interface{}{
								"key":   map[string]interface{}{"int": int32(2)},
								"value": map[string]interface{}{"string": "b"},
							},
						},
					},
					"bool_to_string": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{
								"key":   map[string]interface{}{"boolean": false},
								"value": map[string]interface{}{"string": "b"},
							},
							map[string]interface{}{
								"key":   map[string]interface{}{"boolean": true},
								"value": map[string]interface{}{"string": "a"},
							},
						},
					},
					"string_to_float_value": map[string]interface{}{
						"array": []interface{}{
							map[string]interface{}{
								"key":   map[string]interface{}{"string": "1"},
								"value": nil,
							},
							map[string]interface{}{
								"key":   map[string]interface{}{"string": "2"},
								"value": map[string]interface{}{"float": float32(2)},
							},
						},
					},
				},
			},
		},
		{
			name: "examplev1.ExampleOneof",
			msg: &examplev1.ExampleOneof{
				OneofFields_1: &examplev1.ExampleOneof_OneofEmptyMessage_1{
					OneofEmptyMessage_1: &examplev1.ExampleOneof_EmptyMessage{},
				},
				OneofFields_2: nil,
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleOneof": map[string]interface{}{
					// first oneof
					"oneof_bool_1": nil,
					"oneof_empty_message_1": map[string]interface{}{
						"einride.avro.example.v1.ExampleOneof.EmptyMessage": map[string]interface{}{},
					},
					// second oneof
					"oneof_empty_message_2": nil,
					"oneof_message":         nil,
				},
			},
		},
		{
			name: "examplev1.ExampleRecursive",
			msg: &examplev1.ExampleRecursive{
				Recursive: &examplev1.ExampleRecursive{
					Recursive: &examplev1.ExampleRecursive{},
				},
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleRecursive": map[string]interface{}{
					"recursive": map[string]interface{}{
						"einride.avro.example.v1.ExampleRecursive": map[string]interface{}{
							"recursive": map[string]interface{}{
								"einride.avro.example.v1.ExampleRecursive": map[string]interface{}{
									"recursive": nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "examplev1.ExampleStruct",
			msg: &examplev1.ExampleStruct{
				Struct: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"string":  structpb.NewStringValue("value"),
						"boolean": structpb.NewBoolValue(true),
					},
				},
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleStruct": map[string]interface{}{
					"struct": map[string]interface{}{
						"string": stableStruct(t, `{
	"boolean": true,
	"string": "value"
}`),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleTimeOfDay",
			msg: &examplev1.ExampleTimeOfDay{
				TimeOfDay: &timeofday.TimeOfDay{
					Hours:   19,
					Minutes: 42,
					Seconds: 31,
				},
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleTimeOfDay": map[string]interface{}{
					"time_of_day": map[string]interface{}{
						"long.time-micros": int64(70951000000),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleTimestamp",
			msg: &examplev1.ExampleTimestamp{
				Timestamp: timestamppb.New(time.Date(2021, 6, 27, 1, 39, 24, 0, time.UTC)),
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleTimestamp": map[string]interface{}{
					"timestamp": map[string]interface{}{
						"long.timestamp-micros": int64(1624757964000000),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleWrappers: empty",
			msg:  &examplev1.ExampleWrappers{},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleWrappers": map[string]interface{}{
					"float_value":  nil,
					"double_value": nil,
					"string_value": nil,
					"bytes_value":  nil,
					"int32_value":  nil,
					"int64_value":  nil,
					"uint32_value": nil,
					"uint64_value": nil,
					"bool_value":   nil,
				},
			},
		},
		{
			name: "examplev1.ExampleWrappers: full",
			msg: &examplev1.ExampleWrappers{
				FloatValue:  wrapperspb.Float(1),
				DoubleValue: wrapperspb.Double(2),
				StringValue: wrapperspb.String("3"),
				BytesValue:  wrapperspb.Bytes([]byte{4}),
				Int32Value:  wrapperspb.Int32(5),
				Int64Value:  wrapperspb.Int64(6),
				Uint32Value: wrapperspb.UInt32(7),
				Uint64Value: wrapperspb.UInt64(8),
				BoolValue:   wrapperspb.Bool(false),
			},
			expected: map[string]interface{}{
				"einride.avro.example.v1.ExampleWrappers": map[string]interface{}{
					"float_value":  map[string]interface{}{"float": float32(1)},
					"double_value": map[string]interface{}{"double": float64(2)},
					"string_value": map[string]interface{}{"string": "3"},
					"bytes_value":  map[string]interface{}{"bytes": []byte{4}},
					"int32_value":  map[string]interface{}{"int": int32(5)},
					"int64_value":  map[string]interface{}{"long": int64(6)},
					"uint32_value": map[string]interface{}{"int": int32(7)},
					"uint64_value": map[string]interface{}{"long": int64(8)},
					"bool_value":   map[string]interface{}{"boolean": false},
				},
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			got, err := encodeJSON(tt.msg)
			assert.NilError(t, err)
			assert.DeepEqual(t, tt.expected, got)

			// assert that it matches schema
			schema, err := InferSchema(tt.msg.ProtoReflect().Descriptor())
			assert.NilError(t, err)
			schemaBytes, err := json.Marshal(schema)
			assert.NilError(t, err)
			codec, err := goavro.NewCodec(string(schemaBytes))
			assert.NilError(t, err)
			_, err = codec.BinaryFromNative(nil, got)
			assert.NilError(t, err)

			next := proto.Clone(tt.msg)
			proto.Reset(next)
			assert.NilError(t, decodeJSON(got, next))
			assert.DeepEqual(t, tt.msg, next, protocmp.Transform())
		})
	}
}

func mustAny(t *testing.T, msg proto.Message) *anypb.Any {
	a, err := anypb.New(msg)
	assert.NilError(t, err)
	return a
}

func stableAny(t *testing.T, value string) string {
	t.Helper()
	var pb anypb.Any
	return stableProtoJSON(t, &pb, value)
}

func stableStruct(t *testing.T, value string) string {
	t.Helper()
	var pb structpb.Struct
	return stableProtoJSON(t, &pb, value)
}

// recommended way of comparing protojson output.
// relies on encoding being deterministic, but not stable.
// see: https://github.com/golang/protobuf/issues/1121#issuecomment-627554847
func stableProtoJSON(t *testing.T, pb proto.Message, value string) string {
	t.Helper()
	assert.NilError(t, protojson.Unmarshal([]byte(value), pb))
	next, err := protojson.Marshal(pb)
	assert.NilError(t, err)
	return string(next)
}
