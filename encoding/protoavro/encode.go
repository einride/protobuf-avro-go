package protoavro

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// encodeJSON returns the Avro JSON encoding of message.
func encodeJSON(message proto.Message, options MarshalOptions) (interface{}, error) {
	return topLevelMessageJSON(message.ProtoReflect(), options)
}

// topLevelMessageJSON encodes the proto message including potential extra fields from MarshalOptions.
// For nested fields in the proto message it will recursively call messageJSON function to convert
// to the avro format.
func topLevelMessageJSON(message protoreflect.Message, options MarshalOptions) (interface{}, error) {
	if !message.IsValid() {
		return nil, nil
	}
	if isWKT(message.Descriptor().FullName()) && len(options.ExtraFields) == 0{
		value, err := encodeWKT(message)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	desc := message.Descriptor()
	record := make(map[string]interface{}, desc.Fields().Len() + len(options.ExtraFields))
	for i := 0; i < desc.Fields().Len(); i++ {
		field := desc.Fields().Get(i)
		if field.ContainingOneof() != nil {
			if !message.Has(field) {
				// dont populate scalar fields belonging to
				// a oneof (.Get returns the default value)
				record[string(field.Name())] = nil
			} else {
				value := message.Get(field)
				jsonValue, err := fieldJSON(field, value)
				if err != nil {
					return nil, err
				}
				record[string(field.Name())] = jsonValue
			}
			continue
		}
		value := message.Get(field)
		jsonValue, err := fieldJSON(field, value)
		if err != nil {
			return nil, err
		}
		record[string(field.Name())] = jsonValue
	}
	for _, field := range options.ExtraFields {
		jsonValue, err := messageJSON(field.Message.ProtoReflect())
		if err != nil {
			return nil, err
		}
		record[field.FieldName] = jsonValue
	}
	return map[string]interface{}{
		string(desc.FullName()): record,
	}, nil
}

func unionValue(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: value,
	}
}

// messageJSON encodes a proto message to the binary avro format.
// It complies to the avro nullability by returning it on the format of map[string]interface{}.
// For nested fields in the proto message it will recursively call messageJSON function.
func messageJSON(message protoreflect.Message) (interface{}, error) {
	if !message.IsValid() {
		return nil, nil
	}
	if isWKT(message.Descriptor().FullName()) {
		value, err := encodeWKT(message)
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
				jsonValue, err := fieldJSON(field, value)
				if err != nil {
					return nil, err
				}
				record[string(field.Name())] = jsonValue
			}
			continue
		}
		value := message.Get(field)
		jsonValue, err := fieldJSON(field, value)
		if err != nil {
			return nil, err
		}
		record[string(field.Name())] = jsonValue
	}
	return map[string]interface{}{
		string(desc.FullName()): record,
	}, nil
}

func fieldJSON(field protoreflect.FieldDescriptor, value protoreflect.Value) (interface{}, error) {
	if field.IsList() {
		list := make([]interface{}, 0, value.List().Len())
		for i := 0; i < value.List().Len(); i++ {
			v := value.List().Get(i)
			fieldValue, err := fieldKindJSON(field, v)
			if err != nil {
				return nil, err
			}
			list = append(list, fieldValue)
		}
		return unionValue("array", list), nil
	}
	if field.IsMap() {
		return encodeMap(field, value.Map())
	}
	return fieldKindJSON(field, value)
}

func fieldKindJSON(field protoreflect.FieldDescriptor, value protoreflect.Value) (interface{}, error) {
	switch field.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return messageJSON(value.Message())
	case protoreflect.EnumKind:
		return unionValue(
			string(field.Enum().FullName()),
			string(field.Enum().Values().Get(int(value.Enum())).Name()),
		), nil
	case protoreflect.StringKind:
		return unionValue("string", value.String()), nil
	case protoreflect.Int32Kind,
		protoreflect.Fixed32Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sint32Kind:
		return unionValue("int", int32(value.Int())), nil
	case protoreflect.Uint32Kind:
		return unionValue("int", int32(value.Uint())), nil
	case protoreflect.Int64Kind,
		protoreflect.Fixed64Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Sint64Kind:
		return unionValue("long", value.Int()), nil
	case protoreflect.Uint64Kind:
		return unionValue("long", int64(value.Uint())), nil
	case protoreflect.BoolKind:
		return unionValue("boolean", value.Bool()), nil
	case protoreflect.BytesKind:
		return unionValue("bytes", value.Bytes()), nil
	case protoreflect.DoubleKind:
		return unionValue("double", value.Float()), nil
	case protoreflect.FloatKind:
		return unionValue("float", float32(value.Float())), nil
	}
	return value.Interface(), nil
}
