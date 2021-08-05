package protoavro_test

import (
	"bytes"

	"github.com/einride/protobuf-avro-go/encoding/protoavro"
	"google.golang.org/genproto/googleapis/example/library/v1"
)

func ExampleMarshaler() {
	var msg library.Book
	var b bytes.Buffer
	marshaller, err := protoavro.NewMarshaler(msg.ProtoReflect().Descriptor(), &b)
	if err != nil {
		panic(err)
	}
	if err := marshaller.Marshal(
		&library.Book{
			Name:   "shelves/1/books/1",
			Title:  "Harry Potter",
			Author: "J. K. Rowling",
		},
	); err != nil {
		panic(err)
	}
	// Output:
}
