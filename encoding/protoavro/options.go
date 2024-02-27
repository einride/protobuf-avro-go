package protoavro

import "google.golang.org/protobuf/reflect/protoreflect"

type GetDocCallback func(protoreflect.Descriptor) string

// SchemaOptions contains configuration options for Avro schema inference.
// OmitRootElement is used to determine whether the root element of a message should be omitted, when writing to Avro.
// DocCallback is used to determine the documentation for a field or message.
type SchemaOptions struct {
	OmitRootElement     bool
	DocCallback         GetDocCallback
	NoNullArrayElements bool // don't nullify array elements
	UniqueNames         bool // use unique names for records
}
