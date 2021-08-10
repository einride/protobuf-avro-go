package protoavro

import (
	"fmt"
	"strings"

	"go.einride.tech/protobuf-avro/avro"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// InferSchema returns the Avro schema for the protobuf message descriptor.
func InferSchema(desc protoreflect.MessageDescriptor) (avro.Schema, error) {
	return newSchemaInferrer().inferMessageSchema(desc)
}

type schemaInferrer struct {
	seen map[protoreflect.FullName]struct{}
}

func newSchemaInferrer() schemaInferrer {
	return schemaInferrer{seen: make(map[protoreflect.FullName]struct{})}
}

func (s schemaInferrer) inferMessageSchema(message protoreflect.MessageDescriptor) (avro.Schema, error) {
	if isWKT(message.FullName()) {
		return schemaWKT(message)
	}

	if _, ok := s.seen[message.FullName()]; ok {
		return avro.Nullable(avro.Reference(message.FullName())), nil
	}
	s.seen[message.FullName()] = struct{}{}
	doc := message.ParentFile().SourceLocations().ByDescriptor(message).LeadingComments
	record := avro.Record{
		Type:      avro.RecordType,
		Doc:       doc,
		Name:      string(message.Name()),
		Namespace: namespace(message),
		Fields:    make([]avro.Field, 0, message.Fields().Len()),
	}
	for i := 0; i < message.Fields().Len(); i++ {
		field := message.Fields().Get(i)
		fieldSchema, err := s.inferField(field)
		if err != nil {
			return nil, err
		}

		fieldSchema.Type = avro.Nullable(fieldSchema.Type)
		record.Fields = append(
			record.Fields,
			fieldSchema,
		)
	}
	if message.IsMapEntry() {
		return record, nil
	}
	return avro.Nullable(record), nil
}

func namespace(desc protoreflect.Descriptor) string {
	return strings.TrimSuffix(string(desc.FullName()), "."+string(desc.Name()))
}

func (s schemaInferrer) inferField(field protoreflect.FieldDescriptor) (avro.Field, error) {
	doc := field.ParentFile().SourceLocations().ByDescriptor(field).LeadingComments
	if field.IsMap() {
		mapType, err := s.inferMapSchema(field)
		if err != nil {
			return avro.Field{}, err
		}
		return avro.Field{
			Name: string(field.Name()),
			Doc:  doc,
			Type: mapType,
		}, nil
	}
	fieldKind, err := s.inferFieldKind(field)
	if err != nil {
		return avro.Field{}, err
	}
	if field.IsList() {
		return avro.Field{
			Name: string(field.Name()),
			Doc:  doc,
			Type: avro.Array{
				Type:  avro.ArrayType,
				Items: avro.Nullable(fieldKind),
			},
		}, nil
	}
	if oneof := field.ContainingOneof(); oneof != nil {
		return avro.Field{
			Name: string(field.Name()),
			Doc:  oneofDoc(doc, oneof),
			Type: avro.Nullable(fieldKind),
		}, nil
	}
	return avro.Field{
		Name: string(field.Name()),
		Doc:  doc,
		Type: fieldKind,
	}, nil
}

func oneofDoc(doc string, oneof protoreflect.OneofDescriptor) string {
	fieldNamesLi := make([]string, 0, oneof.Fields().Len())
	for i := 0; i < oneof.Fields().Len(); i++ {
		field := oneof.Fields().Get(i)
		fieldNamesLi = append(fieldNamesLi, fmt.Sprintf("* %s", field.Name()))
	}
	oneofDoc := fmt.Sprintf("At most one will be set:\n%s", strings.Join(fieldNamesLi, "\n"))
	if doc == "" {
		return oneofDoc
	}
	return fmt.Sprintf("%s\n\n%s", doc, oneofDoc)
}

func (s schemaInferrer) inferFieldKind(field protoreflect.FieldDescriptor) (avro.Schema, error) {
	switch field.Kind() {
	case protoreflect.DoubleKind:
		return avro.Double(), nil
	case protoreflect.FloatKind:
		return avro.Float(), nil
	case protoreflect.Int32Kind,
		protoreflect.Fixed32Kind,
		protoreflect.Uint32Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sint32Kind:
		return avro.Integer(), nil
	case protoreflect.Int64Kind,
		protoreflect.Uint64Kind,
		protoreflect.Fixed64Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Sint64Kind:
		return avro.Long(), nil
	case protoreflect.BoolKind:
		return avro.Boolean(), nil
	case protoreflect.BytesKind:
		return avro.Bytes(), nil
	case protoreflect.StringKind:
		return avro.String(), nil
	case protoreflect.EnumKind:
		return s.inferEnumSchema(field.Enum()), nil
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return s.inferMessageSchema(field.Message())
	}
	return nil, fmt.Errorf("unsupported field kind %s %s", field.Name(), field.Kind())
}

func (s schemaInferrer) inferEnumSchema(enum protoreflect.EnumDescriptor) avro.Schema {
	if _, ok := s.seen[enum.FullName()]; ok {
		return avro.Reference(enum.FullName())
	}
	s.seen[enum.FullName()] = struct{}{}
	doc := enum.ParentFile().SourceLocations().ByDescriptor(enum).LeadingComments
	e := avro.Enum{
		Type:      avro.EnumType,
		Doc:       doc,
		Name:      string(enum.Name()),
		Namespace: namespace(enum),
	}
	for i := 0; i < enum.Values().Len(); i++ {
		e.Symbols = append(e.Symbols, string(enum.Values().Get(i).Name()))
	}
	return e
}
