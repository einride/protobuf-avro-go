package protoavro

import (
	"testing"

	"go.einride.tech/protobuf-avro/avro"
	examplev1 "go.einride.tech/protobuf-avro/internal/examples/proto/gen/einride/avro/example/v1"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/proto"
	"gotest.tools/v3/assert"
)

func TestInferSchema(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name     string
		msg      proto.Message
		expected avro.Schema
	}{
		{
			name: "library.Book",
			msg:  &library.Book{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "Book",
				Namespace: "google.example.library.v1",
				Fields: []avro.Field{
					{Name: "name", Type: avro.Nullable(avro.String())},
					{Name: "author", Type: avro.Nullable(avro.String())},
					{Name: "title", Type: avro.Nullable(avro.String())},
					{Name: "read", Type: avro.Nullable(avro.Boolean())},
				},
			}),
		},
		{
			name: "library.UpdateBookRequest",
			msg:  &library.UpdateBookRequest{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "UpdateBookRequest",
				Namespace: "google.example.library.v1",
				Fields: []avro.Field{
					{
						Name: "book",
						Type: avro.Nullable(avro.Record{
							Type:      avro.RecordType,
							Name:      "Book",
							Namespace: "google.example.library.v1",
							Fields: []avro.Field{
								{Name: "name", Type: avro.Nullable(avro.String())},
								{Name: "author", Type: avro.Nullable(avro.String())},
								{Name: "title", Type: avro.Nullable(avro.String())},
								{Name: "read", Type: avro.Nullable(avro.Boolean())},
							},
						}),
					},
					{
						Name: "update_mask",
						Type: avro.Nullable(avro.Record{
							Type:      avro.RecordType,
							Name:      "FieldMask",
							Namespace: "google.protobuf",
							Fields: []avro.Field{
								{
									Name: "paths",
									Type: avro.Nullable(avro.Array{
										Type:  avro.ArrayType,
										Items: avro.Nullable(avro.String()),
									}),
								},
							},
						}),
					},
				},
			}),
		},
		{
			name: "examplev1.ExampleAny",
			msg:  &examplev1.ExampleAny{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleAny",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "any",
						Type: avro.Nullable(
							avro.String(),
						),
					},
				},
			}),
		},
		{
			name: "examplev1.ExampleBytes",
			msg:  &examplev1.ExampleBytes{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleBytes",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "bytes",
						Type: avro.Nullable(avro.Bytes()),
					},
				},
			}),
		},
		{
			name: "examplev1.ExampleDate",
			msg:  &examplev1.ExampleDate{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleDate",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{Name: "date", Type: avro.Nullable(avro.Date())},
				},
			}),
		},
		{
			name: "examplev1.ExampleDateTime",
			msg:  &examplev1.ExampleDateTime{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleDateTime",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "date_time",
						Type: avro.Nullable(avro.Record{
							Type:      avro.RecordType,
							Name:      "DateTime",
							Namespace: "google.type",
							Fields: []avro.Field{
								{Name: "year", Type: avro.Nullable(avro.Integer())},
								{Name: "month", Type: avro.Nullable(avro.Integer())},
								{Name: "day", Type: avro.Nullable(avro.Integer())},
								{Name: "hours", Type: avro.Nullable(avro.Integer())},
								{Name: "minutes", Type: avro.Nullable(avro.Integer())},
								{Name: "seconds", Type: avro.Nullable(avro.Integer())},
								{Name: "nanos", Type: avro.Nullable(avro.Integer())},
								{
									Name: "utc_offset",
									Doc:  "At most one will be set:\n* utc_offset\n* time_zone",
									Type: avro.Nullable(avro.Float()),
								},
								{
									Name: "time_zone",
									Doc:  "At most one will be set:\n* utc_offset\n* time_zone",
									Type: avro.Nullable(avro.Record{
										Type:      avro.RecordType,
										Name:      "TimeZone",
										Namespace: "google.type",
										Fields: []avro.Field{
											{Name: "id", Type: avro.Nullable(avro.String())},
											{Name: "version", Type: avro.Nullable(avro.String())},
										},
									}),
								},
							},
						}),
					},
				},
			}),
		},
		{
			name: "examplev1.ExampleDuration",
			msg:  &examplev1.ExampleDuration{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleDuration",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{Name: "duration", Type: avro.Nullable(avro.Float())},
				},
			}),
		},
		{
			name: "examplev1.ExampleEnum",
			msg:  &examplev1.ExampleEnum{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleEnum",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "enum_value",
						Type: avro.Nullable(avro.Enum{
							Type:      avro.EnumType,
							Name:      "Enum",
							Namespace: "einride.avro.example.v1.ExampleEnum",
							Symbols: []string{
								"ENUM_UNSPECIFIED",
								"ENUM_VALUE1",
								"ENUM_VALUE2",
								"ENUM_VALUE3",
							},
						}),
					},
				},
			}),
		},
		{
			name: "examplev1.ExampleList",
			msg:  &examplev1.ExampleList{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleList",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "int64_list",
						Type: avro.Nullable(avro.Array{
							Type:  avro.ArrayType,
							Items: avro.Nullable(avro.Long()),
						}),
					},
					{
						Name: "string_list",
						Type: avro.Nullable(avro.Array{
							Type:  avro.ArrayType,
							Items: avro.Nullable(avro.String()),
						}),
					},
					{
						Name: "enum_list",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Nullable(avro.Enum{
								Type:      avro.EnumType,
								Name:      "Enum",
								Namespace: "einride.avro.example.v1.ExampleList",
								Symbols: []string{
									"ENUM_UNSPECIFIED",
									"ENUM_VALUE1",
									"ENUM_VALUE2",
								},
							}),
						}),
					},
					{
						Name: "nested_list",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Nullable(avro.Record{
								Type:      avro.RecordType,
								Name:      "Nested",
								Namespace: "einride.avro.example.v1.ExampleList",
								Fields: []avro.Field{
									{
										Name: "string_list",
										Type: avro.Nullable(avro.Array{
											Type:  avro.ArrayType,
											Items: avro.Nullable(avro.String()),
										}),
									},
								},
							}),
						}),
					},
					{
						Name: "float_value_list",
						Type: avro.Nullable(avro.Array{
							Type:  avro.ArrayType,
							Items: avro.Nullable(avro.Float()),
						}),
					},
				},
			}),
		},
		{
			name: "examplev1.ExampleMap",
			msg:  &examplev1.ExampleMap{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleMap",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "string_to_string",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "StringToStringEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.String())},
									{Name: "value", Type: avro.Nullable(avro.String())},
								},
							},
						}),
					},
					{
						Name: "string_to_nested",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "StringToNestedEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.String())},
									{
										Name: "value",
										Type: avro.Nullable(avro.Record{
											Type:      avro.RecordType,
											Name:      "Nested",
											Namespace: "einride.avro.example.v1.ExampleMap",
											Fields: []avro.Field{
												{
													Name: "string_to_string",
													Type: avro.Nullable(avro.Array{
														Type: avro.ArrayType,
														Items: avro.Record{
															Type:      avro.RecordType,
															Name:      "StringToStringEntry",
															Namespace: "einride.avro.example.v1.ExampleMap.Nested",
															Fields: []avro.Field{
																{Name: "key", Type: avro.Nullable(avro.String())},
																{Name: "value", Type: avro.Nullable(avro.String())},
															},
														},
													}),
												},
											},
										}),
									},
								},
							},
						}),
					},
					{
						Name: "string_to_enum",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "StringToEnumEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.String())},
									{
										Name: "value",
										Type: avro.Nullable(avro.Enum{
											Type:      avro.EnumType,
											Name:      "Enum",
											Namespace: "einride.avro.example.v1.ExampleMap",
											Symbols: []string{
												"ENUM_UNSPECIFIED",
												"ENUM_VALUE1",
												"ENUM_VALUE2",
											},
										}),
									},
								},
							},
						}),
					},
					{
						Name: "int32_to_string",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "Int32ToStringEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.Integer())},
									{Name: "value", Type: avro.Nullable(avro.String())},
								},
							},
						}),
					},
					{
						Name: "int64_to_string",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "Int64ToStringEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.Long())},
									{Name: "value", Type: avro.Nullable(avro.String())},
								},
							},
						}),
					},
					{
						Name: "uint32_to_string",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "Uint32ToStringEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.Integer())},
									{Name: "value", Type: avro.Nullable(avro.String())},
								},
							},
						}),
					},
					{
						Name: "bool_to_string",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "BoolToStringEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.Boolean())},
									{Name: "value", Type: avro.Nullable(avro.String())},
								},
							},
						}),
					},
					{
						Name: "string_to_float_value",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "StringToFloatValueEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.String())},
									{Name: "value", Type: avro.Nullable(avro.Float())},
								},
							},
						}),
					},
				},
			}),
		},
		{
			name: "examplev1.ExampleOneof",
			msg:  &examplev1.ExampleOneof{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleOneof",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "oneof_empty_message_1",
						Doc:  "At most one will be set:\n* oneof_empty_message_1\n* oneof_bool_1",
						Type: avro.Nullable(avro.Record{
							Type:      avro.RecordType,
							Name:      "EmptyMessage",
							Namespace: "einride.avro.example.v1.ExampleOneof",
							Fields:    []avro.Field{},
						}),
					},
					{
						Name: "oneof_bool_1",
						Doc:  "At most one will be set:\n* oneof_empty_message_1\n* oneof_bool_1",
						Type: avro.Nullable(avro.Boolean()),
					},
					{
						Name: "oneof_empty_message_2",
						Doc:  "At most one will be set:\n* oneof_empty_message_2\n* oneof_message",
						Type: avro.Nullable(avro.Reference("einride.avro.example.v1.ExampleOneof.EmptyMessage")),
					},
					{
						Name: "oneof_message",
						Doc:  "At most one will be set:\n* oneof_empty_message_2\n* oneof_message",
						Type: avro.Nullable(avro.Record{
							Type:      avro.RecordType,
							Name:      "Message",
							Namespace: "einride.avro.example.v1.ExampleOneof",
							Fields: []avro.Field{
								{Name: "string_value", Type: avro.Nullable(avro.String())},
							},
						}),
					},
				},
			}),
		},
		{
			name: "examplev1.ExampleRecursive",
			msg:  &examplev1.ExampleRecursive{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleRecursive",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "recursive",
						Type: avro.Nullable(
							avro.Reference("einride.avro.example.v1.ExampleRecursive"),
						),
					},
				},
			}),
		},
		{
			name: "examplev1.ExampleStruct",
			msg:  &examplev1.ExampleStruct{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleStruct",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "struct",
						Type: avro.Nullable(
							avro.String(),
						),
					},
				},
			}),
		},
		{
			name: "examplev1.ExampleTimestamp",
			msg:  &examplev1.ExampleTimestamp{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleTimestamp",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{Name: "timestamp", Type: avro.Nullable(avro.TimestampMicros())},
				},
			}),
		},
		{
			name: "examplev1.ExampleTimeOfDay",
			msg:  &examplev1.ExampleTimeOfDay{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleTimeOfDay",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{Name: "time_of_day", Type: avro.Nullable(avro.TimeMicros())},
				},
			}),
		},
		{
			name: "examplev1.ExampleWrappers",
			msg:  &examplev1.ExampleWrappers{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleWrappers",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{Name: "float_value", Type: avro.Nullable(avro.Float())},
					{Name: "double_value", Type: avro.Nullable(avro.Double())},
					{Name: "string_value", Type: avro.Nullable(avro.String())},
					{Name: "bytes_value", Type: avro.Nullable(avro.Bytes())},
					{Name: "int32_value", Type: avro.Nullable(avro.Integer())},
					{Name: "int64_value", Type: avro.Nullable(avro.Long())},
					{Name: "uint32_value", Type: avro.Nullable(avro.Integer())},
					{Name: "uint64_value", Type: avro.Nullable(avro.Long())},
					{Name: "bool_value", Type: avro.Nullable(avro.Boolean())},
				},
			}),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := InferSchema(tt.msg.ProtoReflect().Descriptor())
			assert.NilError(t, err)
			assert.DeepEqual(t, tt.expected, got)
		})
	}
}

func TestInferRootSchema(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name     string
		msg      proto.Message
		expected avro.Schema
	}{
		{
			name: "library.Book",
			msg:  &library.Book{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "Book",
				Namespace: "google.example.library.v1",
				Fields: []avro.Field{
					{Name: "name", Type: avro.Nullable(avro.String())},
					{Name: "author", Type: avro.Nullable(avro.String())},
					{Name: "title", Type: avro.Nullable(avro.String())},
					{Name: "read", Type: avro.Nullable(avro.Boolean())},
				},
			},
		},
		{
			name: "library.UpdateBookRequest",
			msg:  &library.UpdateBookRequest{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "UpdateBookRequest",
				Namespace: "google.example.library.v1",
				Fields: []avro.Field{
					{
						Name: "book",
						Type: avro.Nullable(avro.Record{
							Type:      avro.RecordType,
							Name:      "Book",
							Namespace: "google.example.library.v1",
							Fields: []avro.Field{
								{Name: "name", Type: avro.Nullable(avro.String())},
								{Name: "author", Type: avro.Nullable(avro.String())},
								{Name: "title", Type: avro.Nullable(avro.String())},
								{Name: "read", Type: avro.Nullable(avro.Boolean())},
							},
						}),
					},
					{
						Name: "update_mask",
						Type: avro.Nullable(avro.Record{
							Type:      avro.RecordType,
							Name:      "FieldMask",
							Namespace: "google.protobuf",
							Fields: []avro.Field{
								{
									Name: "paths",
									Type: avro.Nullable(avro.Array{
										Type:  avro.ArrayType,
										Items: avro.Nullable(avro.String()),
									}),
								},
							},
						}),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleAny",
			msg:  &examplev1.ExampleAny{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleAny",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "any",
						Type: avro.Nullable(
							avro.String(),
						),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleBytes",
			msg:  &examplev1.ExampleBytes{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleBytes",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "bytes",
						Type: avro.Nullable(avro.Bytes()),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleDate",
			msg:  &examplev1.ExampleDate{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleDate",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{Name: "date", Type: avro.Nullable(avro.Date())},
				},
			},
		},
		{
			name: "examplev1.ExampleDateTime",
			msg:  &examplev1.ExampleDateTime{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleDateTime",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "date_time",
						Type: avro.Nullable(avro.Record{
							Type:      avro.RecordType,
							Name:      "DateTime",
							Namespace: "google.type",
							Fields: []avro.Field{
								{Name: "year", Type: avro.Nullable(avro.Integer())},
								{Name: "month", Type: avro.Nullable(avro.Integer())},
								{Name: "day", Type: avro.Nullable(avro.Integer())},
								{Name: "hours", Type: avro.Nullable(avro.Integer())},
								{Name: "minutes", Type: avro.Nullable(avro.Integer())},
								{Name: "seconds", Type: avro.Nullable(avro.Integer())},
								{Name: "nanos", Type: avro.Nullable(avro.Integer())},
								{
									Name: "utc_offset",
									Doc:  "At most one will be set:\n* utc_offset\n* time_zone",
									Type: avro.Nullable(avro.Float()),
								},
								{
									Name: "time_zone",
									Doc:  "At most one will be set:\n* utc_offset\n* time_zone",
									Type: avro.Nullable(avro.Record{
										Type:      avro.RecordType,
										Name:      "TimeZone",
										Namespace: "google.type",
										Fields: []avro.Field{
											{Name: "id", Type: avro.Nullable(avro.String())},
											{Name: "version", Type: avro.Nullable(avro.String())},
										},
									}),
								},
							},
						}),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleDuration",
			msg:  &examplev1.ExampleDuration{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleDuration",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{Name: "duration", Type: avro.Nullable(avro.Float())},
				},
			},
		},
		{
			name: "examplev1.ExampleEnum",
			msg:  &examplev1.ExampleEnum{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleEnum",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "enum_value",
						Type: avro.Nullable(avro.Enum{
							Type:      avro.EnumType,
							Name:      "Enum",
							Namespace: "einride.avro.example.v1.ExampleEnum",
							Symbols: []string{
								"ENUM_UNSPECIFIED",
								"ENUM_VALUE1",
								"ENUM_VALUE2",
								"ENUM_VALUE3",
							},
						}),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleList",
			msg:  &examplev1.ExampleList{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleList",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "int64_list",
						Type: avro.Nullable(avro.Array{
							Type:  avro.ArrayType,
							Items: avro.Nullable(avro.Long()),
						}),
					},
					{
						Name: "string_list",
						Type: avro.Nullable(avro.Array{
							Type:  avro.ArrayType,
							Items: avro.Nullable(avro.String()),
						}),
					},
					{
						Name: "enum_list",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Nullable(avro.Enum{
								Type:      avro.EnumType,
								Name:      "Enum",
								Namespace: "einride.avro.example.v1.ExampleList",
								Symbols: []string{
									"ENUM_UNSPECIFIED",
									"ENUM_VALUE1",
									"ENUM_VALUE2",
								},
							}),
						}),
					},
					{
						Name: "nested_list",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Nullable(avro.Record{
								Type:      avro.RecordType,
								Name:      "Nested",
								Namespace: "einride.avro.example.v1.ExampleList",
								Fields: []avro.Field{
									{
										Name: "string_list",
										Type: avro.Nullable(avro.Array{
											Type:  avro.ArrayType,
											Items: avro.Nullable(avro.String()),
										}),
									},
								},
							}),
						}),
					},
					{
						Name: "float_value_list",
						Type: avro.Nullable(avro.Array{
							Type:  avro.ArrayType,
							Items: avro.Nullable(avro.Float()),
						}),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleMap",
			msg:  &examplev1.ExampleMap{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleMap",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "string_to_string",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "StringToStringEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.String())},
									{Name: "value", Type: avro.Nullable(avro.String())},
								},
							},
						}),
					},
					{
						Name: "string_to_nested",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "StringToNestedEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.String())},
									{
										Name: "value",
										Type: avro.Nullable(avro.Record{
											Type:      avro.RecordType,
											Name:      "Nested",
											Namespace: "einride.avro.example.v1.ExampleMap",
											Fields: []avro.Field{
												{
													Name: "string_to_string",
													Type: avro.Nullable(avro.Array{
														Type: avro.ArrayType,
														Items: avro.Record{
															Type:      avro.RecordType,
															Name:      "StringToStringEntry",
															Namespace: "einride.avro.example.v1.ExampleMap.Nested",
															Fields: []avro.Field{
																{Name: "key", Type: avro.Nullable(avro.String())},
																{Name: "value", Type: avro.Nullable(avro.String())},
															},
														},
													}),
												},
											},
										}),
									},
								},
							},
						}),
					},
					{
						Name: "string_to_enum",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "StringToEnumEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.String())},
									{
										Name: "value",
										Type: avro.Nullable(avro.Enum{
											Type:      avro.EnumType,
											Name:      "Enum",
											Namespace: "einride.avro.example.v1.ExampleMap",
											Symbols: []string{
												"ENUM_UNSPECIFIED",
												"ENUM_VALUE1",
												"ENUM_VALUE2",
											},
										}),
									},
								},
							},
						}),
					},
					{
						Name: "int32_to_string",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "Int32ToStringEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.Integer())},
									{Name: "value", Type: avro.Nullable(avro.String())},
								},
							},
						}),
					},
					{
						Name: "int64_to_string",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "Int64ToStringEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.Long())},
									{Name: "value", Type: avro.Nullable(avro.String())},
								},
							},
						}),
					},
					{
						Name: "uint32_to_string",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "Uint32ToStringEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.Integer())},
									{Name: "value", Type: avro.Nullable(avro.String())},
								},
							},
						}),
					},
					{
						Name: "bool_to_string",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "BoolToStringEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.Boolean())},
									{Name: "value", Type: avro.Nullable(avro.String())},
								},
							},
						}),
					},
					{
						Name: "string_to_float_value",
						Type: avro.Nullable(avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "StringToFloatValueEntry",
								Namespace: "einride.avro.example.v1.ExampleMap",
								Fields: []avro.Field{
									{Name: "key", Type: avro.Nullable(avro.String())},
									{Name: "value", Type: avro.Nullable(avro.Float())},
								},
							},
						}),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleOneof",
			msg:  &examplev1.ExampleOneof{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleOneof",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "oneof_empty_message_1",
						Doc:  "At most one will be set:\n* oneof_empty_message_1\n* oneof_bool_1",
						Type: avro.Nullable(avro.Record{
							Type:      avro.RecordType,
							Name:      "EmptyMessage",
							Namespace: "einride.avro.example.v1.ExampleOneof",
							Fields:    []avro.Field{},
						}),
					},
					{
						Name: "oneof_bool_1",
						Doc:  "At most one will be set:\n* oneof_empty_message_1\n* oneof_bool_1",
						Type: avro.Nullable(avro.Boolean()),
					},
					{
						Name: "oneof_empty_message_2",
						Doc:  "At most one will be set:\n* oneof_empty_message_2\n* oneof_message",
						Type: avro.Nullable(avro.Reference("einride.avro.example.v1.ExampleOneof.EmptyMessage")),
					},
					{
						Name: "oneof_message",
						Doc:  "At most one will be set:\n* oneof_empty_message_2\n* oneof_message",
						Type: avro.Nullable(avro.Record{
							Type:      avro.RecordType,
							Name:      "Message",
							Namespace: "einride.avro.example.v1.ExampleOneof",
							Fields: []avro.Field{
								{Name: "string_value", Type: avro.Nullable(avro.String())},
							},
						}),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleRecursive",
			msg:  &examplev1.ExampleRecursive{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleRecursive",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "recursive",
						Type: avro.Nullable(
							avro.Reference("einride.avro.example.v1.ExampleRecursive"),
						),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleStruct",
			msg:  &examplev1.ExampleStruct{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleStruct",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "struct",
						Type: avro.Nullable(
							avro.String(),
						),
					},
				},
			},
		},
		{
			name: "examplev1.ExampleTimestamp",
			msg:  &examplev1.ExampleTimestamp{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleTimestamp",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{Name: "timestamp", Type: avro.Nullable(avro.TimestampMicros())},
				},
			},
		},
		{
			name: "examplev1.ExampleTimeOfDay",
			msg:  &examplev1.ExampleTimeOfDay{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleTimeOfDay",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{Name: "time_of_day", Type: avro.Nullable(avro.TimeMicros())},
				},
			},
		},
		{
			name: "examplev1.ExampleWrappers",
			msg:  &examplev1.ExampleWrappers{},
			expected: avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleWrappers",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{Name: "float_value", Type: avro.Nullable(avro.Float())},
					{Name: "double_value", Type: avro.Nullable(avro.Double())},
					{Name: "string_value", Type: avro.Nullable(avro.String())},
					{Name: "bytes_value", Type: avro.Nullable(avro.Bytes())},
					{Name: "int32_value", Type: avro.Nullable(avro.Integer())},
					{Name: "int64_value", Type: avro.Nullable(avro.Long())},
					{Name: "uint32_value", Type: avro.Nullable(avro.Integer())},
					{Name: "uint64_value", Type: avro.Nullable(avro.Long())},
					{Name: "bool_value", Type: avro.Nullable(avro.Boolean())},
				},
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := SchemaOptions{OmitRootElement: true}.InferSchema(tt.msg.ProtoReflect().Descriptor())
			assert.NilError(t, err)
			assert.DeepEqual(t, tt.expected, got)
		})
	}
}

func TestInferSchemaOptions(t *testing.T) {
	for _, tt := range []struct {
		name     string
		msg      proto.Message
		expected avro.Schema
		opt      SchemaOptions
	}{
		{
			name: "examplev1.ExampleList",
			msg:  &examplev1.ExampleList{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Name:      "ExampleList",
				Namespace: "einride.avro.example.v1",
				Fields: []avro.Field{
					{
						Name: "int64_list",
						Type: avro.Array{
							Type:  avro.ArrayType,
							Items: avro.Long(),
						},
					},
					{
						Name: "string_list",
						Type: avro.Array{
							Type:  avro.ArrayType,
							Items: avro.String(),
						},
					},
					{
						Name: "enum_list",
						Type: avro.Array{
							Type: avro.ArrayType,
							Items: avro.Enum{
								Type:      avro.EnumType,
								Name:      "Enum",
								Namespace: "einride.avro.example.v1.ExampleList",
								Symbols: []string{
									"ENUM_UNSPECIFIED",
									"ENUM_VALUE1",
									"ENUM_VALUE2",
								},
							},
						},
					},
					{
						Name: "nested_list",
						Type: avro.Array{
							Type: avro.ArrayType,
							Items: avro.Record{
								Type:      avro.RecordType,
								Name:      "Nested",
								Namespace: "einride.avro.example.v1.ExampleList",
								Fields: []avro.Field{
									{
										Name: "string_list",
										Type: avro.Array{
											Type:  avro.ArrayType,
											Items: avro.String(),
										},
									},
								},
							},
						},
					},
					{
						Name: "float_value_list",
						Type: avro.Array{
							Type:  avro.ArrayType,
							Items: avro.Float(),
						},
					},
				},
			}),
			opt: SchemaOptions{
				NoNullArrayElements: true,
				NoNullArray:         true,
			},
		},
		{
			name: "examplev1.ExampleSeen",
			msg:  &examplev1.ExampleSeen{},
			expected: avro.Nullable(avro.Record{
				Type:      avro.RecordType,
				Namespace: "root",
				Name:      "ExampleSeen",
				Fields: []avro.Field{
					{
						Name: "left",
						Type: avro.Union{
							avro.Primitive{Type: "null"},
							avro.Record{
								Type:      "record",
								Namespace: "root.exampleseen.left",
								Name:      "ExampleData",
								Fields: []avro.Field{
									{
										Name: "value",
										Type: avro.Nullable(avro.String()),
									},
								},
							},
						},
					},
					{
						Name: "right",
						Type: avro.Union{
							avro.Primitive{Type: "null"},
							avro.Record{
								Type:      "record",
								Namespace: "root.exampleseen.right",
								Name:      "ExampleData",
								Fields: []avro.Field{
									{
										Name: "value",
										Type: avro.Nullable(avro.String()),
									},
								},
							},
						},
					},
				},
			}),
			opt: SchemaOptions{
				UniqueNames: true,
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.opt.InferSchema(tt.msg.ProtoReflect().Descriptor())
			assert.NilError(t, err)
			assert.DeepEqual(t, tt.expected, got)
		})
	}
}
