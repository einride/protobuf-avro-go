package protoavro_test

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"go.einride.tech/protobuf-avro/avro"
	"go.einride.tech/protobuf-avro/encoding/protoavro"
	"google.golang.org/genproto/googleapis/example/library/v1"
)

func ExampleInferSchema() {
	msg := &library.Book{}
	schema, err := protoavro.InferSchema(msg.ProtoReflect().Descriptor())
	if err != nil {
		panic(err)
	}
	expected := avro.Nullable(avro.Record{
		Type:      avro.RecordType,
		Name:      "Book",
		Namespace: "google.example.library.v1",
		Fields: []avro.Field{
			{Name: "name", Type: avro.Nullable(avro.String())},
			{Name: "author", Type: avro.Nullable(avro.String())},
			{Name: "title", Type: avro.Nullable(avro.String())},
			{Name: "read", Type: avro.Nullable(avro.Boolean())},
		},
	})
	fmt.Println(cmp.Equal(expected, schema))
	// Output: true
}
