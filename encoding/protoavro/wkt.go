package protoavro

import (
	"fmt"
	"time"

	"cloud.google.com/go/civil"
	"go.einride.tech/protobuf-avro/avro"
	"go.einride.tech/protobuf-avro/internal/wkt"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/timeofday"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func isWKT(name protoreflect.FullName) bool {
	switch name {
	case wkt.DoubleValue,
		wkt.FloatValue,
		wkt.Int32Value,
		wkt.UInt32Value,
		wkt.Int64Value,
		wkt.UInt64Value,
		wkt.BoolValue,
		wkt.StringValue,
		wkt.BytesValue,
		wkt.Struct,
		wkt.Any,
		wkt.Timestamp,
		wkt.Duration,
		wkt.Date,
		wkt.TimeOfDay:
		return true
	}
	return false
}

func schemaWKT(message protoreflect.MessageDescriptor) (avro.Schema, error) {
	switch message.FullName() {
	case wkt.DoubleValue,
		wkt.FloatValue,
		wkt.Int32Value,
		wkt.UInt32Value,
		wkt.Int64Value,
		wkt.UInt64Value,
		wkt.BoolValue,
		wkt.StringValue,
		wkt.BytesValue:
		schema, err := schemaWrapper(string(message.FullName()))
		if err != nil {
			return nil, err
		}
		return schema, nil
	case wkt.Struct:
		return schemaStruct(), nil
	case wkt.Any:
		return schemaAny(), nil
	case wkt.Timestamp:
		return schemaTimestamp(), nil
	case wkt.Duration:
		return schemaDuration(), nil
	case wkt.Date:
		return schemaDate(), nil
	case wkt.TimeOfDay:
		return schemaTimeOfDay(), nil
	}
	return nil, fmt.Errorf("uknown wellknown type %s", message.FullName())
}

func encodeWKT(message protoreflect.Message) (map[string]interface{}, error) {
	desc := message.Descriptor()
	switch desc.FullName() {
	case wkt.DoubleValue,
		wkt.FloatValue,
		wkt.Int32Value,
		wkt.UInt32Value,
		wkt.Int64Value,
		wkt.UInt64Value,
		wkt.BoolValue,
		wkt.StringValue,
		wkt.BytesValue:
		value, err := encodeWrapper(message)
		if err != nil {
			return nil, err
		}
		return value, nil
	case wkt.Struct:
		value, err := encodeStruct(message.Interface().(*structpb.Struct))
		if err != nil {
			return nil, err
		}
		return value, nil
	case wkt.Any:
		value, err := encodeAny(message.Interface().(*anypb.Any))
		if err != nil {
			return nil, err
		}
		return value, nil
	case wkt.Timestamp:
		return encodeTimestamp(message.Interface().(*timestamppb.Timestamp)), nil
	case wkt.Duration:
		return encodeDuration(message.Interface().(*durationpb.Duration)), nil
	case wkt.Date:
		return encodeDate(message.Interface().(*date.Date)), nil
	case wkt.TimeOfDay:
		return encodeTimeOfDay(message.Interface().(*timeofday.TimeOfDay)), nil
	default:
		return nil, fmt.Errorf("unknown wellknown type %s", desc.FullName())
	}
}

func decodeWKT(data map[string]interface{}, msg protoreflect.Message) error {
	desc := msg.Descriptor()
	var value proto.Message
	var err error
	switch desc.FullName() {
	case wkt.Any:
		value, err = decodeAny(data)
	case wkt.Date:
		value, err = decodeDate(data)
	case wkt.Struct:
		value, err = decodeStruct(data)
	case wkt.TimeOfDay:
		value, err = decodeTimeOfDay(data)
	case wkt.Duration:
		value, err = decodeDuration(data)
	case wkt.Timestamp:
		value, err = decodeTimestamp(data)
	case wkt.FloatValue,
		wkt.DoubleValue,
		wkt.UInt32Value,
		wkt.UInt64Value,
		wkt.Int32Value,
		wkt.Int64Value,
		wkt.BytesValue,
		wkt.StringValue,
		wkt.BoolValue:
		value, err = decodeWrapper(string(desc.FullName()), data)
	default:
		return fmt.Errorf("unknown wellknown type %s", desc.FullName())
	}
	if err != nil {
		return err
	}
	proto.Merge(msg.Interface(), value)
	return nil
}

func schemaWrapper(w string) (avro.Schema, error) {
	switch w {
	case wkt.DoubleValue:
		return avro.Nullable(avro.Double()), nil
	case wkt.FloatValue:
		return avro.Nullable(avro.Float()), nil
	case wkt.Int32Value, wkt.UInt32Value:
		return avro.Nullable(avro.Integer()), nil
	case wkt.Int64Value, wkt.UInt64Value:
		return avro.Nullable(avro.Long()), nil
	case wkt.BoolValue:
		return avro.Nullable(avro.Boolean()), nil
	case wkt.StringValue:
		return avro.Nullable(avro.String()), nil
	case wkt.BytesValue:
		return avro.Nullable(avro.Bytes()), nil
	default:
		return nil, fmt.Errorf("unknown wrapper type %s", w)
	}
}

func encodeWrapper(msg protoreflect.Message) (map[string]interface{}, error) {
	if msg == nil {
		return nil, nil
	}
	switch msg.Descriptor().FullName() {
	case wkt.DoubleValue:
		return unionValue("double", msg.Interface().(*wrapperspb.DoubleValue).GetValue()), nil
	case wkt.FloatValue:
		return unionValue("float", msg.Interface().(*wrapperspb.FloatValue).GetValue()), nil
	case wkt.Int32Value:
		return unionValue("int", msg.Interface().(*wrapperspb.Int32Value).GetValue()), nil
	case wkt.UInt32Value:
		return unionValue("int", int32(msg.Interface().(*wrapperspb.UInt32Value).GetValue())), nil
	case wkt.Int64Value:
		return unionValue("long", msg.Interface().(*wrapperspb.Int64Value).GetValue()), nil
	case wkt.UInt64Value:
		return unionValue("long", int64(msg.Interface().(*wrapperspb.UInt64Value).GetValue())), nil
	case wkt.BoolValue:
		return unionValue("boolean", msg.Interface().(*wrapperspb.BoolValue).GetValue()), nil
	case wkt.StringValue:
		return unionValue("string", msg.Interface().(*wrapperspb.StringValue).GetValue()), nil
	case wkt.BytesValue:
		return unionValue("bytes", msg.Interface().(*wrapperspb.BytesValue).GetValue()), nil
	default:
		return nil, fmt.Errorf("unknown wrapper type %s", msg.Descriptor().FullName())
	}
}

func decodeWrapper(w string, v map[string]interface{}) (proto.Message, error) {
	if v == nil {
		return nil, nil
	}
	switch w {
	case wkt.DoubleValue:
		f, err := decodeFloatLike(v, "double")
		if err != nil {
			return nil, fmt.Errorf("google.protobuf.DoubleValue: %w", err)
		}
		return wrapperspb.Double(f), nil
	case wkt.FloatValue:
		f, err := decodeFloatLike(v, "float")
		if err != nil {
			return nil, fmt.Errorf("google.protobuf.FloatValue: %w", err)
		}
		return wrapperspb.Float(float32(f)), nil
	case wkt.UInt32Value:
		i, err := decodeIntLike(v, "int")
		if err != nil {
			return nil, fmt.Errorf("google.protobuf.UInt32Value: %w", err)
		}
		return wrapperspb.UInt32(uint32(i)), nil
	case wkt.UInt64Value:
		i, err := decodeIntLike(v, "long")
		if err != nil {
			return nil, fmt.Errorf("google.protobuf.UInt32Value: %w", err)
		}
		return wrapperspb.UInt64(uint64(i)), nil
	case wkt.Int32Value:
		i, err := decodeIntLike(v, "int")
		if err != nil {
			return nil, fmt.Errorf("google.protobuf.Int32Value: %w", err)
		}
		return wrapperspb.Int32(int32(i)), nil
	case wkt.Int64Value:
		i, err := decodeIntLike(v, "long")
		if err != nil {
			return nil, fmt.Errorf("google.protobuf.Int32Value: %w", err)
		}
		return wrapperspb.Int64(int64(i)), nil
	case wkt.BytesValue:
		b, err := decodeBytes(v, "bytes")
		if err != nil {
			return nil, fmt.Errorf("google.protobuf.BytesValue: %w", err)
		}
		return wrapperspb.Bytes(b), nil
	case wkt.StringValue:
		s, err := decodeString(v, "string")
		if err != nil {
			return nil, fmt.Errorf("google.protobuf.StringValue: %w", err)
		}
		return wrapperspb.String(s), nil
	case wkt.BoolValue:
		b, err := decodeBool(v, "boolean")
		if err != nil {
			return nil, fmt.Errorf("google.protobuf.BoolValue: %w", err)
		}
		return wrapperspb.Bool(b), nil
	default:
		return nil, fmt.Errorf("unknown wrapper type %s", w)
	}
}

func schemaDate() avro.Schema {
	return avro.Nullable(avro.Date())
}

func encodeDate(d *date.Date) map[string]interface{} {
	civilDate := civil.Date{
		Year:  int(d.Year),
		Month: time.Month(d.Month),
		Day:   int(d.Day),
	}
	epoch := civil.Date{
		Year:  1970,
		Month: time.January,
		Day:   1,
	}
	return unionValue("int.date", int32(civilDate.DaysSince(epoch)))
}

func decodeDate(v map[string]interface{}) (*date.Date, error) {
	if v == nil {
		return nil, nil
	}
	if tm, ok := tryDecodeTime(v, "int.date"); ok {
		d := civil.DateOf(tm)
		return dateFromCivil(d), nil
	}
	i, err := decodeIntLike(v, "int.date")
	if err != nil {
		return nil, fmt.Errorf("google.type.Date: %w", err)
	}
	d := civil.Date{Year: 1970, Month: time.January, Day: 1}.AddDays(int(i))
	return dateFromCivil(d), nil
}

func dateFromCivil(c civil.Date) *date.Date {
	return &date.Date{
		Year:  int32(c.Year),
		Month: int32(c.Month),
		Day:   int32(c.Day),
	}
}

func schemaAny() avro.Schema {
	return avro.Nullable(avro.String()) // EncodeJSON string
}

func encodeAny(a *anypb.Any) (map[string]interface{}, error) {
	data, err := protojson.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("google.protobuf.Any: marshal: %w", err)
	}
	return unionValue("string", string(data)), nil
}

func decodeAny(v map[string]interface{}) (*anypb.Any, error) {
	if v == nil {
		return nil, nil
	}
	str, err := decodeString(v, "string")
	if err != nil {
		return nil, fmt.Errorf("google.protobuf.Any: %w", err)
	}
	var any anypb.Any
	if err := protojson.Unmarshal([]byte(str), &any); err != nil {
		return nil, fmt.Errorf("google.protobuf.Any: unmarshal: %w", err)
	}
	return &any, nil
}

func schemaStruct() avro.Schema {
	return avro.Nullable(avro.String()) // EncodeJSON string
}

func encodeStruct(a *structpb.Struct) (map[string]interface{}, error) {
	data, err := protojson.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("google.protobuf.Struct: marshal: %w", err)
	}
	return unionValue("string", string(data)), nil
}

func decodeStruct(v map[string]interface{}) (*structpb.Struct, error) {
	if v == nil {
		return nil, nil
	}
	str, err := decodeString(v, "string")
	if err != nil {
		return nil, fmt.Errorf("google.protobuf.Struct: %w", err)
	}
	var strct structpb.Struct
	if err := protojson.Unmarshal([]byte(str), &strct); err != nil {
		return nil, fmt.Errorf("google.protobuf.Struct: unmarshal: %w", err)
	}
	return &strct, nil
}

func schemaTimeOfDay() avro.Schema {
	return avro.Nullable(avro.TimeMicros())
}

func encodeTimeOfDay(t *timeofday.TimeOfDay) map[string]interface{} {
	d := time.Hour*time.Duration(t.Hours) +
		time.Minute*time.Duration(t.Minutes) +
		time.Second*time.Duration(t.Seconds) +
		time.Nanosecond*time.Duration(t.Nanos)
	return unionValue("long.time-micros", d.Microseconds())
}

func decodeTimeOfDay(v map[string]interface{}) (*timeofday.TimeOfDay, error) {
	if v == nil {
		return nil, nil
	}
	if dur, ok := tryDecodeDuration(v, "long.time-micros"); ok {
		return timeOfDayFromDuration(dur), nil
	}
	micro, err := decodeIntLike(v, "long.time-micros")
	if err != nil {
		return nil, fmt.Errorf("google.type.TimeOfDay: %w", err)
	}
	dur := time.Microsecond * time.Duration(micro)
	return timeOfDayFromDuration(dur), nil
}

func timeOfDayFromDuration(dur time.Duration) *timeofday.TimeOfDay {
	hours := dur.Truncate(time.Hour)
	dur -= hours
	minutes := dur.Truncate(time.Minute)
	dur -= minutes
	seconds := dur.Truncate(time.Second)
	dur -= seconds
	nanos := dur.Truncate(time.Nanosecond)
	dur -= nanos
	return &timeofday.TimeOfDay{
		Hours:   int32(hours.Hours()),
		Minutes: int32(minutes.Minutes()),
		Seconds: int32(seconds.Seconds()),
		Nanos:   int32(nanos.Nanoseconds()),
	}
}

func schemaDuration() avro.Schema {
	return avro.Nullable(avro.Float())
}

func encodeDuration(dur *durationpb.Duration) map[string]interface{} {
	return unionValue("float", dur.AsDuration().Seconds())
}

func decodeDuration(v map[string]interface{}) (*durationpb.Duration, error) {
	seconds, err := decodeFloatLike(v, "float")
	if err != nil {
		return nil, fmt.Errorf("google.protobuf.Duration: %w", err)
	}
	// prevent downcasting float64 to int64 when passing to time.Duration
	micros := seconds / time.Microsecond.Seconds()
	return durationpb.New(time.Microsecond * time.Duration(micros)), nil
}

func schemaTimestamp() avro.Schema {
	return avro.Nullable(avro.TimestampMicros())
}

func encodeTimestamp(t *timestamppb.Timestamp) map[string]interface{} {
	return unionValue("long.timestamp-micros", t.AsTime().UnixNano()/1e3)
}

func decodeTimestamp(v map[string]interface{}) (*timestamppb.Timestamp, error) {
	if tm, ok := tryDecodeTime(v, "long.timestamp-micros"); ok {
		return timestamppb.New(tm), nil
	}
	micros, err := decodeIntLike(v, "long.timestamp-micros")
	if err != nil {
		return nil, fmt.Errorf("google.protobuf.Timestamp: %w", err)
	}
	t := time.Unix(0, 0).Add(time.Microsecond * time.Duration(micros))
	return timestamppb.New(t), nil
}

func decodeIntLike(v map[string]interface{}, key string) (int64, error) {
	maybeInt, ok := v[key]
	if !ok {
		return 0, fmt.Errorf("expected key '%s'", key)
	}
	switch i := maybeInt.(type) {
	case int:
		return int64(i), nil
	case int32:
		return int64(i), nil
	case int64:
		return i, nil
	default:
		return 0, fmt.Errorf("expected int-like, got %T", maybeInt)
	}
}

func tryDecodeTime(v map[string]interface{}, key string) (time.Time, bool) {
	maybeTime, ok := v[key]
	if !ok {
		return time.Time{}, false
	}
	if tm, ok := maybeTime.(time.Time); ok {
		return tm, true
	}
	return time.Time{}, false
}

func tryDecodeDuration(v map[string]interface{}, key string) (time.Duration, bool) {
	maybeDuration, ok := v[key]
	if !ok {
		return time.Duration(0), false
	}
	if dur, ok := maybeDuration.(time.Duration); ok {
		return dur, true
	}
	return time.Duration(0), false
}

func decodeFloatLike(v map[string]interface{}, key string) (float64, error) {
	maybeFloat, ok := v[key]
	if !ok {
		return 0, fmt.Errorf("expected key '%s'", key)
	}
	switch i := maybeFloat.(type) {
	case float32:
		return float64(i), nil
	case float64:
		return i, nil
	default:
		return 0, fmt.Errorf("expected float-like, got %T", maybeFloat)
	}
}

func decodeString(v map[string]interface{}, key string) (string, error) {
	maybeString, ok := v[key]
	if !ok {
		return "", fmt.Errorf("expected key '%s'", key)
	}
	switch i := maybeString.(type) {
	case string:
		return i, nil
	default:
		return "", fmt.Errorf("expected string, got %T", maybeString)
	}
}

func decodeBytes(v map[string]interface{}, key string) ([]byte, error) {
	maybeByte, ok := v[key]
	if !ok {
		return nil, fmt.Errorf("expected key '%s'", key)
	}
	switch b := maybeByte.(type) {
	case []byte:
		return b, nil
	default:
		return nil, fmt.Errorf("expected []byte, got %T", maybeByte)
	}
}

func decodeBool(v map[string]interface{}, key string) (bool, error) {
	maybeBool, ok := v[key]
	if !ok {
		return false, fmt.Errorf("expected key '%s'", key)
	}
	switch b := maybeBool.(type) {
	case bool:
		return b, nil
	default:
		return false, fmt.Errorf("expected bool, got %T", maybeBool)
	}
}
