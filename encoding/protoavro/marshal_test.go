package protoavro_test

import (
	"bytes"
	"testing"

	"github.com/einride/protobuf-avro-go/encoding/protoavro"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/testing/protocmp"
	"gotest.tools/v3/assert"
)

func Test_Marshaler(t *testing.T) {
	msgs := []*library.Book{
		{
			Name:   "shelves/1/books/1",
			Title:  "Harry Potter",
			Author: "J. K. Rowling",
		},
		{
			Name:   "shelves/1/books/2",
			Title:  "Lord of the Rings",
			Author: "J. R. R. Tolkien",
		},
	}

	var b bytes.Buffer

	// marshal messages
	marshaller, err := protoavro.NewMarshaler(msgs[0].ProtoReflect().Descriptor(), &b)
	assert.NilError(t, err)
	for _, msg := range msgs {
		assert.NilError(t, marshaller.Append(msg))
	}

	// unmarshal messages
	unmarshaler, err := protoavro.NewUnmarshaler(&b)
	assert.NilError(t, err)
	got := make([]*library.Book, 0, 2)
	for unmarshaler.Scan() {
		var msg library.Book
		assert.NilError(t, unmarshaler.Read(&msg))
		got = append(got, &msg)
	}

	assert.DeepEqual(t, msgs, got, protocmp.Transform())
}
