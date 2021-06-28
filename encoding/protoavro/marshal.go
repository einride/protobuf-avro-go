package protoavro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/linkedin/goavro/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// NewMarshaler returns a new marshaler that writes protobuf messages to writer in
// Avro binary format.
func NewMarshaler(descriptor protoreflect.MessageDescriptor, writer io.Writer) (*Marshaler, error) {
	schema, err := InferSchema(descriptor)
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
	return &Marshaler{w: w, desc: descriptor}, nil
}

// Marshaler encodes and writes Avro binary formatted messages.
type Marshaler struct {
	desc protoreflect.MessageDescriptor
	w    *goavro.OCFWriter
}

func (m *Marshaler) Append(messages ...proto.Message) error {
	data := make([]interface{}, 0, len(messages))
	for _, message := range messages {
		a := message.ProtoReflect().Descriptor().FullName()
		b := m.desc.FullName()
		if a != b {
			return fmt.Errorf("expected message '%s' but got '%s'", a, b)
		}
		m, err := EncodeJSON(message)
		if err != nil {
			return fmt.Errorf("encode json: %w", err)
		}
		data = append(data, m)
	}
	if err := m.w.Append(data); err != nil {
		return fmt.Errorf("append: %w", err)
	}
	return nil
}
