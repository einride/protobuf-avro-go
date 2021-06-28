package protoavro_test

import (
	"encoding/json"
	"fmt"

	"github.com/einride/protobuf-avro-go/encoding/protoavro"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func ExampleEncodeJSON() {
	msg := &library.Book{
		Name:   "books/1",
		Author: "J. K. Rowling",
		Title:  "Harry Potter",
		Read:   true,
	}
	encoded, err := protoavro.EncodeJSON(msg)
	if err != nil {
		panic(err)
	}
	var decoded library.Book
	if err := protoavro.DecodeJSON(encoded, &decoded); err != nil {
		panic(err)
	}

	encodedBytes, err := json.MarshalIndent(encoded, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encodedBytes))
	fmt.Println(cmp.Equal(msg, &decoded, protocmp.Transform()))
	// Output:
	// {
	//   "google.example.library.v1.Book": {
	//     "author": "J. K. Rowling",
	//     "name": "books/1",
	//     "read": true,
	//     "title": "Harry Potter"
	//   }
	// }
	// true
}
