package protoavro

import (
	"fmt"
	"io"

	"github.com/linkedin/goavro/v2"
	"google.golang.org/protobuf/proto"
)

// NewUnmarshaler returns a new unmarshaler that reads protobuf messages from reader in
// Avro binary format.
func NewUnmarshaler(reader io.Reader) (*Unmarshaler, error) {
	r, err := goavro.NewOCFReader(reader)
	if err != nil {
		return nil, fmt.Errorf("new ocf writer: %w", err)
	}
	return &Unmarshaler{r: r}, nil
}

// NewUnmarshaler returns a new unmarshaler that reads protobuf messages from reader in
// Avro binary format.
func (o SchemaOptions) NewUnmarshaler(reader io.Reader) (*Unmarshaler, error) {
	r, err := goavro.NewOCFReader(reader)
	if err != nil {
		return nil, fmt.Errorf("new ocf writer: %w", err)
	}
	return &Unmarshaler{opts: o, r: r}, nil
}

// Unmarshaler reads and decodes Avro binary encoded messages.
type Unmarshaler struct {
	opts SchemaOptions
	r    *goavro.OCFReader
}

// Scan returns true when there is at least one more
// message to be read. Scan should be called prior to calling Unmarshal.
func (m *Unmarshaler) Scan() bool {
	return m.r.Scan()
}

// Unmarshal consumes one message from the reader and places it in message.
func (m *Unmarshaler) Unmarshal(message proto.Message) error {
	data, err := m.r.Read()
	if err != nil {
		return fmt.Errorf("read message: %w", err)
	}
	if err := m.opts.decodeJSON(data, message); err != nil {
		return fmt.Errorf("decode message: %w", err)
	}
	return nil
}
