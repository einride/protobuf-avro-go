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

// Unmarshaler reads and decodes Avro binary formatted messages.
type Unmarshaler struct {
	r *goavro.OCFReader
}

func (m *Unmarshaler) Scan() bool {
	return m.r.Scan()
}

func (m *Unmarshaler) Read(msg proto.Message) error {
	data, err := m.r.Read()
	if err != nil {
		return fmt.Errorf("read message: %w", err)
	}
	if err := DecodeJSON(data, msg); err != nil {
		return fmt.Errorf("decode message: %w", err)
	}
	return nil
}
