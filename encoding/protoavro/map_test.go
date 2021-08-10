package protoavro

import (
	"testing"

	"go.einride.tech/protobuf-avro/avro"
	examplev1 "go.einride.tech/protobuf-avro/internal/examples/proto/gen/einride/avro/example/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"gotest.tools/v3/assert"
)

func Test_MapSchema(t *testing.T) {
	for _, tt := range []struct {
		name      string
		msg       proto.Message
		fieldName protoreflect.Name
		expected  avro.Schema
	}{
		{
			name:      "string to string",
			msg:       &examplev1.ExampleMap{},
			fieldName: "string_to_string",
			expected: avro.Nullable(avro.Array{
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
			name:      "nested map",
			msg:       &examplev1.ExampleMap{},
			fieldName: "string_to_nested",
			expected: avro.Nullable(avro.Array{
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
			name:      "enum value",
			msg:       &examplev1.ExampleMap{},
			fieldName: "string_to_enum",
			expected: avro.Nullable(avro.Array{
				Type: avro.ArrayType,
				Items: avro.Record{
					Type:      avro.RecordType,
					Name:      "StringToEnumEntry",
					Namespace: "einride.avro.example.v1.ExampleMap",
					Fields: []avro.Field{
						{Name: "key", Type: avro.Nullable(avro.String())},
						{
							Name: "value", Type: avro.Nullable(avro.Enum{
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
			name:      "int32 key",
			msg:       &examplev1.ExampleMap{},
			fieldName: "int32_to_string",
			expected: avro.Nullable(avro.Array{
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
			name:      "int64 key",
			msg:       &examplev1.ExampleMap{},
			fieldName: "int64_to_string",
			expected: avro.Nullable(avro.Array{
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
			name:      "uint32 key",
			msg:       &examplev1.ExampleMap{},
			fieldName: "uint32_to_string",
			expected: avro.Nullable(avro.Array{
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
			name:      "bool key",
			msg:       &examplev1.ExampleMap{},
			fieldName: "bool_to_string",
			expected: avro.Nullable(avro.Array{
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
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			schema, err := newSchemaInferrer().inferMapSchema(tt.msg.ProtoReflect().Descriptor().Fields().ByName(tt.fieldName))
			assert.NilError(t, err)
			assert.DeepEqual(t, schema, tt.expected)
		})
	}
}

func Test_MapEncode(t *testing.T) {
	for _, tt := range []struct {
		name      string
		msg       proto.Message
		fieldName protoreflect.Name
		expected  interface{}
	}{
		{
			name: "string to string",
			msg: &examplev1.ExampleMap{
				StringToString: map[string]string{
					"1": "a",
					"2": "b",
					"3": "c",
				},
			},
			fieldName: "string_to_string",
			expected: map[string]interface{}{
				"array": []interface{}{
					map[string]interface{}{
						"key":   map[string]interface{}{"string": "1"},
						"value": map[string]interface{}{"string": "a"},
					},
					map[string]interface{}{
						"key":   map[string]interface{}{"string": "2"},
						"value": map[string]interface{}{"string": "b"},
					},
					map[string]interface{}{
						"key":   map[string]interface{}{"string": "3"},
						"value": map[string]interface{}{"string": "c"},
					},
				},
			},
		},
		{
			name: "int32 key",
			msg: &examplev1.ExampleMap{
				Int32ToString: map[int32]string{
					1: "a",
					2: "b",
					3: "c",
				},
			},
			fieldName: "int32_to_string",
			expected: map[string]interface{}{
				"array": []interface{}{
					map[string]interface{}{
						"key":   map[string]interface{}{"int": int32(1)},
						"value": map[string]interface{}{"string": "a"},
					},
					map[string]interface{}{
						"key":   map[string]interface{}{"int": int32(2)},
						"value": map[string]interface{}{"string": "b"},
					},
					map[string]interface{}{
						"key":   map[string]interface{}{"int": int32(3)},
						"value": map[string]interface{}{"string": "c"},
					},
				},
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			desc := tt.msg.ProtoReflect().Descriptor().Fields().ByName(tt.fieldName)
			val := tt.msg.ProtoReflect().Get(desc)
			got, err := encodeMap(desc, val.Map())
			assert.NilError(t, err)
			assert.DeepEqual(t, got, tt.expected)
		})
	}
}

func Test_MapDecode(t *testing.T) {
	for _, tt := range []struct {
		name      string
		msg       proto.Message
		fieldName protoreflect.Name
		data      interface{}
		expected  proto.Message
		expectErr string
	}{
		{
			name:      "string to string",
			msg:       &examplev1.ExampleMap{},
			fieldName: "string_to_string",
			data: []interface{}{
				map[string]interface{}{"key": "1", "value": "a"},
				map[string]interface{}{"key": "2", "value": "b"},
				map[string]interface{}{"key": "3", "value": "c"},
			},
			expected: &examplev1.ExampleMap{
				StringToString: map[string]string{
					"1": "a",
					"2": "b",
					"3": "c",
				},
			},
		},
		{
			name:      "invalid type",
			msg:       &examplev1.ExampleMap{},
			fieldName: "string_to_string",
			data:      map[string]int32{},
			expectErr: "expected list-like, got map[]",
		},
		{
			name:      "invalid element type",
			msg:       &examplev1.ExampleMap{},
			fieldName: "string_to_string",
			data: []interface{}{
				1,
			},
			expectErr: "expected map entry, got int for 'string_to_string'",
		},
		{
			name:      "missing key field",
			msg:       &examplev1.ExampleMap{},
			fieldName: "string_to_string",
			data: []interface{}{
				map[string]interface{}{"value": "a"},
			},
			expectErr: "missing 'key' in map entry for 'string_to_string'",
		},
		{
			name:      "missing value field",
			msg:       &examplev1.ExampleMap{},
			fieldName: "string_to_string",
			data: []interface{}{
				map[string]interface{}{"key": "1"},
			},
			expectErr: "missing 'value' in map entry for 'string_to_string'",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			desc := tt.msg.ProtoReflect().Descriptor().Fields().ByName(tt.fieldName)
			val := tt.msg.ProtoReflect().Mutable(desc)
			err := decodeMap(tt.data, desc, val.Map())
			if tt.expectErr != "" {
				assert.ErrorContains(t, err, tt.expectErr)
				return
			}
			assert.NilError(t, err)
			assert.DeepEqual(t, tt.msg, tt.expected, protocmp.Transform())
		})
	}
}
