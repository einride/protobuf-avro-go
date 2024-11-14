package protoavro_test

import (
	"bytes"
	"testing"
	"time"

	"go.einride.tech/protobuf-avro/encoding/protoavro"
	examplev1 "go.einride.tech/protobuf-avro/internal/examples/proto/gen/einride/avro/example/v1"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/timeofday"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
)

func Test_MarshalUnmarshal(t *testing.T) {
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
		assert.NilError(t, marshaller.Marshal(msg))
	}

	// unmarshal messages
	unmarshaler, err := protoavro.NewUnmarshaler(&b)
	assert.NilError(t, err)
	got := make([]*library.Book, 0, 2)
	for unmarshaler.Scan() {
		var msg library.Book
		assert.NilError(t, unmarshaler.Unmarshal(&msg))
		got = append(got, &msg)
	}

	assert.DeepEqual(t, msgs, got, protocmp.Transform())
}

func Test_MarshalSymmetric(t *testing.T) {
	// when `goavro` decodes a file, it will not read back exactly what was
	// written. For example timestamps are returned as `time.Time`. These
	// tests assert that those conversions are handled properly.
	for _, tt := range []struct {
		name string
		msg  proto.Message
	}{
		{
			name: "examplev1.ExampleTimestamp",
			msg: &examplev1.ExampleTimestamp{
				Timestamp: timestamppb.New(time.Now().Truncate(time.Microsecond)), // nano second precision not expressible in avro
			},
		},
		{
			name: "examplev1.ExampleDate",
			msg: &examplev1.ExampleDate{
				Date: &date.Date{Year: 2021, Month: 1, Day: 1},
			},
		},
		{
			name: "examplev1.ExampleDuration",
			msg: &examplev1.ExampleDuration{
				Duration: durationpb.New(time.Hour),
			},
		},
		{
			name: "examplev1.TimeOfDay",
			msg: &examplev1.ExampleTimeOfDay{
				TimeOfDay: &timeofday.TimeOfDay{Hours: 20},
			},
		},
		{
			name: "examplev1.ExampleNumbers",
			msg: &examplev1.ExampleNumbers{
				Valint32:  1,
				Valint64:  1,
				Valfloat:  1.0,
				Valdouble: 1.0,
				Valuint32: 1,
				Valuint64: 1,
				Valsint32: 1,
				Valsint64: 1,
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer

			// marshal messages
			marshaller, err := protoavro.NewMarshaler(tt.msg.ProtoReflect().Descriptor(), &b)
			assert.NilError(t, err)
			assert.NilError(t, marshaller.Marshal(tt.msg))

			// unmarshal messages
			unmarshaler, err := protoavro.NewUnmarshaler(&b)
			assert.NilError(t, err)
			got := make([]proto.Message, 0, 2)
			for unmarshaler.Scan() {
				msg := proto.Clone(tt.msg)
				proto.Reset(msg)

				assert.NilError(t, unmarshaler.Unmarshal(msg))
				got = append(got, msg)
			}
			assert.Equal(t, len(got), 1)
			assert.DeepEqual(t, tt.msg, got[0], protocmp.Transform())
		})
	}
}
