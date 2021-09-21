package protoavro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/linkedin/goavro/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type MarshalOptions struct {
	ExtraFields []ExtraField
}

type ExtraField struct {
	FieldName string
	Message proto.Message
}

func (o MarshalOptions) convertToSchemaOptions() SchemaOptions {
	extraFields := make([]SchemaExtraField, len(o.ExtraFields))
	for i, field := range o.ExtraFields {
		extraFields[i] = SchemaExtraField{
			FieldName: field.FieldName,
			Descriptor: field.Message.ProtoReflect().Descriptor(),
		}
	}
	return SchemaOptions{ExtraFields: extraFields}
}

func (o MarshalOptions) ConvertToUnmarshalOptions() *UnmarshalOptions {
	return &UnmarshalOptions{
		ExtraFields: o.ExtraFields,
		TrashCan: make([]proto.Message, 0),
	}
}

// NewMarshaler returns a new marshaler that writes protobuf messages to writer in
// Avro binary format.
func NewMarshaler(descriptor protoreflect.MessageDescriptor, writer io.Writer) (*Marshaler, error) {
	return NewMarshalerWithOptions(descriptor, writer, &MarshalOptions{})
}

// NewMarshalerWithOptions allows to specify a SchemaOptions to base the marshaler schema inferral from.
func NewMarshalerWithOptions(descriptor protoreflect.MessageDescriptor, writer io.Writer, options *MarshalOptions) (*Marshaler, error) {
	schema, err := options.convertToSchemaOptions().InferSchema(descriptor)
	if err != nil {
		return nil, fmt.Errorf("infer schema: %w", err)
	}
	schemaBytes, err := json.Marshal(schema)
	if err != nil {
		return nil, fmt.Errorf("json marshal schema: %w", err)
	}
	w, err := goavro.NewOCFWriter(goavro.OCFConfig{
		W:      writer,
		Schema: string(schemaBytes),
	})
	if err != nil {
		return nil, fmt.Errorf("new ocf writer: %w", err)
	}
	return &Marshaler{w: w, desc: descriptor, options: options}, nil
}

// Marshaler encodes and writes Avro binary encoded messages.
type Marshaler struct {
	desc protoreflect.MessageDescriptor
	w    *goavro.OCFWriter
	options *MarshalOptions
}

// Marshal encodes and writes messages to the writer.
func (m *Marshaler) Marshal(message proto.Message) error {
	data := make([]interface{}, 1)
	a := message.ProtoReflect().Descriptor().FullName()
	b := m.desc.FullName()
	if a != b {
		return fmt.Errorf("expected message '%s' but got '%s'", a, b)
	}
	j, err := encodeJSON(message, *m.options)
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	data[0] = j
	if err := m.w.Append(data); err != nil {
		return fmt.Errorf("append: %w", err)
	}
	return nil
}
