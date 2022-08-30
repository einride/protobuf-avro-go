package protoavro

// SchemaOptions contains configuration options for Avro schema inference.
// OmitRootElement is used to determine whether the root element of a message should be omitted, when writing to Avro.
type SchemaOptions struct {
	OmitRootElement bool
}
