package protoavro

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// encodeJSON returns the Avro JSON encoding of message.
func (o SchemaOptions) encodeJSON(message proto.Message) (interface{}, error) {
	return o.messageJSON(message.ProtoReflect(), 0, true)
}

func (o SchemaOptions) unionValue(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: value,
	}
}

func (o SchemaOptions) messageJSON(message protoreflect.Message, recursiveIndex int, useUnion bool) (interface{}, error) {
	if !message.IsValid() {
		return nil, nil
	}
	if isWKT(message.Descriptor().FullName()) {
		value, err := o.encodeWKT(message, useUnion)
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
	if (o.OmitRootElement && recursiveIndex == 0) || !useUnion {
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
			fieldValue, err := o.fieldKindJSON(field, v, recursiveIndex, !o.NoNullArrayElements)
			if err != nil {
				return nil, err
			}

			if o.NoNullArrayElements {
				if fieldValue == nil {
					continue
				}
			}

			list = append(list, fieldValue)
		}

		return o.unionValue("array", list), nil
	}
	if field.IsMap() {
		return o.encodeMap(field, value.Map(), recursiveIndex)
	}
	return o.fieldKindJSON(field, value, recursiveIndex, true)
}

func (o SchemaOptions) maybeUnionValue(key string, value interface{}, selector bool) interface{} {
	if selector {
		return o.unionValue(key, value)
	}
	return value
}

func (o SchemaOptions) fieldKindJSON(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
	recursiveIndex int,
	useUnion bool,
) (interface{}, error) {
	switch field.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return o.messageJSON(value.Message(), recursiveIndex, useUnion)
	case protoreflect.EnumKind:
		if field.Enum().Values().ByNumber(value.Enum()) == nil {
			return o.maybeUnionValue(
				string(field.Enum().FullName()),
				string(field.Enum().Values().ByNumber(protoreflect.EnumNumber(0)).Name()),
				useUnion,
			), nil
		}
		return o.maybeUnionValue(
			string(field.Enum().FullName()),
			string(field.Enum().Values().ByNumber(value.Enum()).Name()), useUnion), nil
	case protoreflect.StringKind:
		return o.maybeUnionValue("string", value.String(), useUnion), nil
	case protoreflect.Int32Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sint32Kind:
		return o.maybeUnionValue("int", int32(value.Int()), useUnion), nil
	case protoreflect.Uint32Kind,
		protoreflect.Fixed32Kind:
		return o.maybeUnionValue("int", int32(value.Uint()), useUnion), nil
	case protoreflect.Int64Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Sint64Kind:
		return o.maybeUnionValue("long", value.Int(), useUnion), nil
	case protoreflect.Fixed64Kind,
		protoreflect.Uint64Kind:
		return o.maybeUnionValue("long", int64(value.Uint()), useUnion), nil
	case protoreflect.BoolKind:
		return o.maybeUnionValue("boolean", value.Bool(), useUnion), nil
	case protoreflect.BytesKind:
		return o.maybeUnionValue("bytes", value.Bytes(), useUnion), nil
	case protoreflect.DoubleKind:
		return o.maybeUnionValue("double", value.Float(), useUnion), nil
	case protoreflect.FloatKind:
		return o.maybeUnionValue("float", float32(value.Float()), useUnion), nil
	}
	return value.Interface(), nil
}
