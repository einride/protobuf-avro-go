package protoavro

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type encoder struct {
	opts  SchemaOptions
	names nameStack
}

func (s encoder) namespace(desc protoreflect.Descriptor) string {
	if s.opts.UniqueNames {
		return s.names.current()
	}
	return strings.TrimSuffix(string(desc.FullName()), "."+string(desc.Name()))
}

func newEncoder(opts SchemaOptions) encoder {
	return encoder{opts: opts, names: newNameStack()}
}

// encodeJSON returns the Avro JSON encoding of message.
func (o SchemaOptions) encodeJSON(message proto.Message) (interface{}, error) {
	encoder := newEncoder(o)
	return encoder.encodeJSON(message)
}

func (o encoder) encodeJSON(message proto.Message) (interface{}, error) {
	return o.messageJSON(message.ProtoReflect(), 0, true)
}

func (o SchemaOptions) unionValue(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: value,
	}
}

func (o encoder) messageJSON(message protoreflect.Message, recursiveIndex int, useUnion bool) (interface{}, error) {
	if !message.IsValid() {
		return nil, nil
	}
	if isWKT(message.Descriptor().FullName()) {
		value, err := o.opts.encodeWKT(message, useUnion)
		if err != nil {
			return nil, err
		}
		return value, nil
	}

	n := string(message.Descriptor().Name())
	ns := o.namespace(message.Descriptor())
	fullName := fmt.Sprintf("%s.%s", ns, n)

	o.names.push(n)
	defer o.names.pop()

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
	if (o.opts.OmitRootElement && recursiveIndex == 0) || !useUnion {
		return record, nil
	}
	return map[string]interface{}{
		string(fullName): record,
	}, nil
}

func (o encoder) fieldJSON(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
	recursiveIndex int,
) (interface{}, error) {
	if field.IsList() {
		list := make([]interface{}, 0, value.List().Len())
		for i := 0; i < value.List().Len(); i++ {
			v := value.List().Get(i)
			fieldValue, err := o.fieldKindJSON(field, v, recursiveIndex, !o.opts.NoNullArrayElements)
			if err != nil {
				return nil, err
			}

			if o.opts.NoNullArrayElements {
				if fieldValue == nil {
					continue
				}
			}

			list = append(list, fieldValue)
		}

		if o.opts.NoNullArray {
			return list, nil
		}

		return o.opts.unionValue("array", list), nil
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

func (o encoder) fieldKindJSON(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
	recursiveIndex int,
	useUnion bool,
) (interface{}, error) {

	o.names.push(string(field.Name()))
	defer o.names.pop()

	switch field.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return o.messageJSON(value.Message(), recursiveIndex, useUnion)
	case protoreflect.EnumKind:
		if field.Enum().Values().ByNumber(value.Enum()) == nil {
			return o.opts.maybeUnionValue(
				string(field.Enum().FullName()),
				string(field.Enum().Values().ByNumber(protoreflect.EnumNumber(0)).Name()),
				useUnion,
			), nil
		}
		return o.opts.maybeUnionValue(
			string(field.Enum().FullName()),
			string(field.Enum().Values().ByNumber(value.Enum()).Name()), useUnion), nil
	case protoreflect.StringKind:
		return o.opts.maybeUnionValue("string", value.String(), useUnion), nil
	case protoreflect.Int32Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sint32Kind:
		return o.opts.maybeUnionValue("int", int32(value.Int()), useUnion), nil
	case protoreflect.Uint32Kind,
		protoreflect.Fixed32Kind:
		return o.opts.maybeUnionValue("int", int32(value.Uint()), useUnion), nil
	case protoreflect.Int64Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Sint64Kind:
		return o.opts.maybeUnionValue("long", value.Int(), useUnion), nil
	case protoreflect.Fixed64Kind,
		protoreflect.Uint64Kind:
		return o.opts.maybeUnionValue("long", int64(value.Uint()), useUnion), nil
	case protoreflect.BoolKind:
		return o.opts.maybeUnionValue("boolean", value.Bool(), useUnion), nil
	case protoreflect.BytesKind:
		return o.opts.maybeUnionValue("bytes", value.Bytes(), useUnion), nil
	case protoreflect.DoubleKind:
		return o.opts.maybeUnionValue("double", value.Float(), useUnion), nil
	case protoreflect.FloatKind:
		return o.opts.maybeUnionValue("float", float32(value.Float()), useUnion), nil
	}
	return value.Interface(), nil
}
