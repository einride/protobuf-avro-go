package protoavro

import (
	"testing"
	"time"

	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/timeofday"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"
)

func Test_WKT(t *testing.T) {
	for _, tt := range []proto.Message{
		// wrappers
		wrapperspb.Double(123),
		wrapperspb.Bool(true),
		wrapperspb.Bytes([]byte("123")),
		wrapperspb.Float(123),
		wrapperspb.Int32(123),
		wrapperspb.Int64(123),
		wrapperspb.String("123"),
		wrapperspb.UInt32(123),
		wrapperspb.UInt64(123),

		// any
		mustAny(t, &library.Book{Name: "shelves/1/books/1"}),

		// struct
		&structpb.Struct{
			Fields: map[string]*structpb.Value{
				"string":  structpb.NewStringValue("value"),
				"boolean": structpb.NewBoolValue(true),
			},
		},

		// timestamp
		timestamppb.New(time.Now().Truncate(time.Microsecond)), // nano second precision not expressible in avro
		timestamppb.New(time.Unix(-100, 0)),

		// date
		&date.Date{Year: 2021, Month: 7, Day: 26},
		&date.Date{Year: 1900, Month: 7, Day: 26},
		&date.Date{Year: -522, Month: 7, Day: 26},

		// time of day
		&timeofday.TimeOfDay{Hours: 23, Minutes: 59, Seconds: 59},
		&timeofday.TimeOfDay{Hours: 0, Minutes: 0, Seconds: 0, Nanos: 0},
		&timeofday.TimeOfDay{Hours: 36, Minutes: 0, Seconds: 0, Nanos: 0},
		&timeofday.TimeOfDay{Hours: 10, Minutes: 0, Seconds: 0, Nanos: 1000},

		// duration
		durationpb.New(time.Hour),
		durationpb.New(time.Microsecond),
		durationpb.New(time.Minute*59 + time.Second*43),
	} {
		tt := tt
		t.Run(string(tt.ProtoReflect().Descriptor().FullName()), func(t *testing.T) {
			encoded, err := SchemaOptions{}.encodeWKT(tt.ProtoReflect())
			assert.NilError(t, err)
			t.Log(encoded)
			decoded := tt.ProtoReflect().New()
			assert.NilError(t, decodeWKT(encoded, decoded))
			assert.DeepEqual(t, tt, decoded.Interface(), protocmp.Transform())
		})
	}
}

func Test_DecodeWKTErr(t *testing.T) {
	for _, tt := range []struct {
		name        string
		msg         proto.Message
		data        map[string]interface{}
		errContains string
	}{
		{
			name: "not any",
			msg:  &anypb.Any{},
			data: map[string]interface{}{
				"key": "value",
			},
			errContains: "google.protobuf.Any: expected key 'string'",
		},

		{
			name: "not encoded any",
			msg:  &anypb.Any{},
			data: map[string]interface{}{
				"string": "value",
			},
			errContains: "google.protobuf.Any: unmarshal",
		},

		{
			name: "wrong wrapper",
			msg:  &wrapperspb.DoubleValue{},
			data: map[string]interface{}{
				"string": "not a double",
			},
			errContains: "google.protobuf.DoubleValue: expected key 'double'",
		},
		{
			name: "not date",
			msg:  &date.Date{},
			data: map[string]interface{}{
				"key": "value",
			},
			errContains: "google.type.Date: expected key 'int.date'",
		},
		{
			name: "not duration",
			msg:  &durationpb.Duration{},
			data: map[string]interface{}{
				"key": "value",
			},
			errContains: "google.protobuf.Duration: expected key 'float'",
		},
		{
			name: "not timestamp",
			msg:  &timestamppb.Timestamp{},
			data: map[string]interface{}{
				"key": "value",
			},
			errContains: "google.protobuf.Timestamp: expected key 'long.timestamp-micros'",
		},
		{
			name: "not time of day",
			msg:  &timeofday.TimeOfDay{},
			data: map[string]interface{}{
				"key": "value",
			},
			errContains: "google.type.TimeOfDay: expected key 'long.time-micros'",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := decodeWKT(tt.data, tt.msg.ProtoReflect())
			assert.ErrorContains(t, err, tt.errContains)
		})
	}
}
