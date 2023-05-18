package protoavro

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// encodeJSON returns the Avro JSON encoding of message.
func (o SchemaOptions) encodeJSON(message proto.Message) (interface{}, error) {
	return o.messageJSON(message.ProtoReflect(), 0)
}

func (o SchemaOptions) unionValue(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: value,
	}
}

func (o SchemaOptions) messageJSON(message protoreflect.Message, recursiveIndex int) (interface{}, error) {
	if !message.IsValid() {
		return nil, nil
	}
	if isWKT(message.Descriptor().FullName()) {
		value, err := o.encodeWKT(message)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	desc := message.Descriptor()
	record := make(map[string]interface{}, desc.Fields().Len())
	for i := 0; i < desc.Fields().Len(); i++ {
		field := desc.Fields().Get(i)
		if field.ContainingOneof() != nil {
			if !message.Has(field) {
				// dont populate scalar fields belonging to
				// a oneof (.Get returns the default value)
				record[string(field.Name())] = nil
			} else {
				value := message.Get(field)
				jsonValue, err := o.fieldJSON(field, value, recursiveIndex+1)
				if err != nil {
					return nil, err
				}
				record[string(field.Name())] = jsonValue
			}
			continue
		}
		value := message.Get(field)
		jsonValue, err := o.fieldJSON(field, value, recursiveIndex+1)
		if err != nil {
			return nil, err
		}
		record[string(field.Name())] = jsonValue
	}
	if o.OmitRootElement && recursiveIndex == 0 {
		return record, nil
	}
	return map[string]interface{}{
		string(desc.FullName()): record,
	}, nil
}

func (o SchemaOptions) fieldJSON(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
	recursiveIndex int,
) (interface{}, error) {
	if field.IsList() {
		list := make([]interface{}, 0, value.List().Len())
		for i := 0; i < value.List().Len(); i++ {
			v := value.List().Get(i)
			fieldValue, err := o.fieldKindJSON(field, v, recursiveIndex)
			if err != nil {
				return nil, err
			}
			list = append(list, fieldValue)
		}
		return o.unionValue("array", list), nil
	}
	if field.IsMap() {
		return o.encodeMap(field, value.Map(), recursiveIndex)
	}
	return o.fieldKindJSON(field, value, recursiveIndex)
}

func (o SchemaOptions) fieldKindJSON(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
	recursiveIndex int,
) (interface{}, error) {
	switch field.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return o.messageJSON(value.Message(), recursiveIndex)
	case protoreflect.EnumKind:
		return o.unionValue(
			string(field.Enum().FullName()),
			string(field.Enum().Values().ByNumber(value.Enum()).Name()),
		), nil
	case protoreflect.StringKind:
		return o.unionValue("string", value.String()), nil
	case protoreflect.Int32Kind,
		protoreflect.Fixed32Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sint32Kind:
		return o.unionValue("int", int32(value.Int())), nil
	case protoreflect.Uint32Kind:
		return o.unionValue("int", int32(value.Uint())), nil
	case protoreflect.Int64Kind,
		protoreflect.Fixed64Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Sint64Kind:
		return o.unionValue("long", value.Int()), nil
	case protoreflect.Uint64Kind:
		return o.unionValue("long", int64(value.Uint())), nil
	case protoreflect.BoolKind:
		return o.unionValue("boolean", value.Bool()), nil
	case protoreflect.BytesKind:
		return o.unionValue("bytes", value.Bytes()), nil
	case protoreflect.DoubleKind:
		return o.unionValue("double", value.Float()), nil
	case protoreflect.FloatKind:
		return o.unionValue("float", float32(value.Float())), nil
	}
	return value.Interface(), nil
}
