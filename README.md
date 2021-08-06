# Protobuf + Avro

Functionality for converting between [Protocol Buffers][protobuf] and [Avro][avro].
This can for example be used to bulk load protobuf messages to BigQuery.

[protobuf]: https://developers.google.com/protocol-buffers/
[avro]: https://avro.apache.org/

## Examples

Examples use the following protobuf message:
```proto
message Book {
  string name = 1;
  string author = 2;
  string title = 3;
  bool read = 4;
}
```

### `protoavro.InferSchema`

Avro schema inference for arbitrary protobuf messages.

```go
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
}
```

### `protoavro.Marshaler`

Writes protobuf messages to an [Object Container File][ocr].

[ocr]: https://avro.apache.org/docs/current/spec.html#Object+Container+Files

```go
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
        &library.Book{
            Name:   "shelves/1/books/2",
            Title:  "Lord of the Rings",
            Author: "J. R. R. Tolkien",
        },
    ); err != nil {
        panic(err)
    }
}
```

### `protoavro.Unmarshaler`

Reads protobuf messages from a [Object Container File][ocr].

[ocr]: https://avro.apache.org/docs/current/spec.html#Object+Container+Files

```go
func ExampleUnmarshaler() {
	var reader io.Reader
	unmarshaller, err := protoavro.NewUnmarshaler(reader)
	if err != nil {
		panic(err)
	}
	for unmarshaller.Scan() {
		var msg library.Book
		if err := unmarshaller.Unmarshal(&msg); err != nil {
			panic(err)
		}
	}
}
```

### Mapping

**Messages** are mapped as nullable records in Avro. Fields will have the same casing as in the protobuf descriptor.

**One of**s are mapped to nullable fields in Avro, where at most one field will be set at a time.

**Maps** are mapped as a list of records with two fields, `key` and `value`. Order of map entries is undefined.

**Enums** are mapped as enums of string values in Avro.

Some **well known types** have a special mapping:

| Protobuf                                      | Avro                                          |
| --------------------------------------------- | --------------------------------------------- |
| wrappers (ex google.protobuf.DoubleValue)     | Nullable scalars (ex `[null, double]`)        |
| google.protobuf.Any                           | string containing JSON encoding of `Any`      |
| google.protobuf.Struct                        | string containing JSON encoding of `Struct`   |
| google.protobuf.Timestamp                     | `long.timestamp-micros`                       |
| google.protobuf.Duration                      | `float` (seconds)                             |
| google.type.Date                              | `int.date`                                    |
| google.type.TimeOfDay                         | `long.time-micros`                            |


### Limitations

Avro does not have a native type for timestamps with nanosecond precision. 
`google.protobuf.Timestamp` and `google.type.TimeOfDay` are truncated to 
microsecond precision when encoded as Avro.
