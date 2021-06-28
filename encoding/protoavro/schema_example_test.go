package protoavro_test

import (
	"fmt"

	"github.com/einride/protobuf-avro-go/avro"
	"github.com/einride/protobuf-avro-go/encoding/protoavro"
	"github.com/google/go-cmp/cmp"
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
			{Name: "name", Type: avro.String()},
			{Name: "author", Type: avro.String()},
			{Name: "title", Type: avro.String()},
			{Name: "read", Type: avro.Boolean()},
		},
	})
	fmt.Println(cmp.Equal(expected, schema))
	// Output: true
}
