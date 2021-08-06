package protoavro

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// decodeJSON decodes the JSON encoded avro data and places the
// result in msg.
func decodeJSON(data interface{}, msg proto.Message) error {
	return decodeMessage(data, msg.ProtoReflect())
}

func decodeMessage(data interface{}, msg protoreflect.Message) error {
	if data == nil {
		return nil
	}
	d, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("expected message encoded as map[string]interface{}, got %T", data)
	}

	if isWKT(msg.Descriptor().FullName()) {
		return decodeWKT(d, msg)
	}
	// unwrap union
	desc := msg.Descriptor()
	if msgData, ok := d[string(desc.FullName())]; len(d) == 1 && ok {
		return decodeMessage(msgData, msg)
	}
	for fieldName, fieldValue := range d {
		fd, ok := findField(desc, fieldName)
		if !ok {
			return fmt.Errorf("unexpected field %s", fieldName)
		}
		if err := decodeField(fieldValue, msg, fd); err != nil {
			return err
		}
	}
	return nil
}

func decodeField(data interface{}, val protoreflect.Message, f protoreflect.FieldDescriptor) error {
	if data == nil {
		return nil
	}
	switch {
	case f.IsMap():
		mp := val.NewField(f).Map()
		if err := decodeMap(data, f, mp); err != nil {
			return err
		}
		val.Set(f, protoreflect.ValueOfMap(mp))
		return nil
	case f.IsList():
		listData, ok := data.([]interface{})
		if !ok {
			return fmt.Errorf("expected list, got %T", data)
		}
		list := val.NewField(f).List()
		for _, el := range listData {
			if el == nil {
				list.Append(list.NewElement())
				continue
			}
			fieldValue, err := decodeFieldKind(el, list.NewElement(), f)
			if err != nil {
				return err
			}
			list.Append(fieldValue)
		}
		val.Set(f, protoreflect.ValueOfList(list))
		return nil
	default:
		fieldValue, err := decodeFieldKind(data, val.NewField(f), f)
		if err != nil {
			return err
		}
		val.Set(f, fieldValue)
	}
	return nil
}

func decodeFieldKind(data interface{}, mutable protoreflect.Value, f protoreflect.FieldDescriptor) (protoreflect.Value, error) {
	switch f.Kind() {
	case protoreflect.StringKind:
		str, ok := data.(string)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("field %s: expected string, got %T", f.Name(), data)
		}
		return protoreflect.ValueOfString(str), nil
	case protoreflect.BoolKind:
		bo, ok := data.(bool)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("field %s: expected bool, got %T", f.Name(), data)
		}
		return protoreflect.ValueOfBool(bo), nil
	case protoreflect.MessageKind, protoreflect.GroupKind:
		if err := decodeMessage(data, mutable.Message()); err != nil {
			return protoreflect.Value{}, err
		}
		return mutable, nil
	case protoreflect.Int32Kind, protoreflect.Sfixed32Kind, protoreflect.Sint32Kind:
		i, ok := data.(int32)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("field %s: expected int32, got %T", f.Name(), data)
		}
		return protoreflect.ValueOfInt32(i), nil
	case protoreflect.Int64Kind, protoreflect.Sfixed64Kind, protoreflect.Sint64Kind:
		i, ok := data.(int64)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("field %s: expected int64, got %T", f.Name(), data)
		}
		return protoreflect.ValueOfInt64(i), nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		i, ok := data.(int32)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("field %s: expected int32, got %T", f.Name(), data)
		}
		return protoreflect.ValueOfUint32(uint32(i)), nil
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		i, ok := data.(int64)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("field %s: expected int64, got %T", f.Name(), data)
		}
		return protoreflect.ValueOfUint64(uint64(i)), nil
	case protoreflect.BytesKind:
		bs, ok := data.([]byte)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("field %s: expected []byte, got %T", f.Name(), data)
		}
		return protoreflect.ValueOfBytes(bs), nil
	case protoreflect.EnumKind:
		str, ok := data.(string)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("field %s: expected string, got %T", f.Name(), data)
		}
		if v := f.Enum().Values().ByName(protoreflect.Name(str)); v != nil {
			return protoreflect.ValueOfEnum(v.Number()), nil
		} else {
			return protoreflect.ValueOfEnum(0), nil
		}
	case protoreflect.DoubleKind:
		dbl, ok := data.(float64)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("field %s: expected float64, got %T", f.Name(), data)
		}
		return protoreflect.ValueOfFloat64(dbl), nil
	case protoreflect.FloatKind:
		flt, ok := data.(float32)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf("field %s: expected float32, got %T", f.Name(), data)
		}
		return protoreflect.ValueOfFloat32(flt), nil

	}
	return protoreflect.Value{}, fmt.Errorf("unexpected kind %s", f.Kind())
}

func findField(desc protoreflect.MessageDescriptor, name string) (protoreflect.FieldDescriptor, bool) {
	if fd := desc.Fields().ByJSONName(name); fd != nil {
		return fd, true
	}
	if fd := desc.Fields().ByTextName(name); fd != nil {
		return fd, true
	}
	return nil, false
}
