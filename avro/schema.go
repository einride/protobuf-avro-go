// Package avro provides types for working with Avro schemas according
// to spec at http://avro.apache.org/docs/current/spec.html.
package avro

// Schema describes an Avro schema.
// JSON encoding of a Schema value matches the specification
// for a schema declaration.
type Schema interface {
	isSchema()
}

// Type matches Avro primitive and complex types.
// See:
// http://avro.apache.org/docs/current/spec.html#schema_primitive
type Type string

const (
	NullType    Type = "null"
	BooleanType Type = "boolean"
	IntType     Type = "int"
	LongType    Type = "long"
	FloatType   Type = "float"
	DoubleType  Type = "double"
	BytesType   Type = "bytes"
	StringType  Type = "string"
	RecordType  Type = "record"
	EnumType    Type = "enum"
	ArrayType   Type = "array"
)

// LogicalType is an Avro primitive or complex type with extra attributes to represent a derived type.
// See: http://avro.apache.org/docs/current/spec.html#Logical+Types
type LogicalType string

const (
	DateLogicalType            LogicalType = "date"
	TimeMicrosLogicalType      LogicalType = "time-micros"
	TimestampMicrosLogicalType LogicalType = "timestamp-micros"
)

type Reference string

func (r Reference) isSchema() {}

type Union []Schema

func (e Union) isSchema() {}

type Primitive struct {
	Type        Type        `json:"type"`
	LogicalType LogicalType `json:"logicalType,omitempty"`
}

func (p Primitive) isSchema() {}

func Null() Primitive {
	return Primitive{Type: NullType}
}

func Integer() Primitive {
	return Primitive{Type: IntType}
}

func Long() Primitive {
	return Primitive{Type: LongType}
}

func Float() Primitive {
	return Primitive{Type: FloatType}
}

func Double() Primitive {
	return Primitive{Type: DoubleType}
}

func String() Primitive {
	return Primitive{Type: StringType}
}

func Bytes() Primitive {
	return Primitive{Type: BytesType}
}

func Boolean() Primitive {
	return Primitive{Type: BooleanType}
}

type Record struct {
	Type      Type    `json:"type"`
	Namespace string  `json:"namespace,omitempty"`
	Doc       string  `json:"doc,omitempty"`
	Name      string  `json:"name"`
	Fields    []Field `json:"fields"`
}

func (p Record) isSchema() {}

type Field struct {
	Name string `json:"name"`
	Doc  string `json:"doc,omitempty"`
	Type Schema `json:"type"`
}

type Enum struct {
	Type      Type     `json:"type"`
	Namespace string   `json:"namespace,omitempty"`
	Doc       string   `json:"doc,omitempty"`
	Name      string   `json:"name"`
	Symbols   []string `json:"symbols"`
}

func (e Enum) isSchema() {}

type Array struct {
	Type  Type   `json:"type"`
	Items Schema `json:"items"`
}

func (e Array) isSchema() {}

type Fixed struct {
	Type      Type   `json:"type"`
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	Size      int    `json:"size"`
}

func (e Fixed) isSchema() {}

func Date() Primitive {
	return Primitive{
		Type:        IntType,
		LogicalType: DateLogicalType,
	}
}

func TimeMicros() Primitive {
	return Primitive{
		Type:        LongType,
		LogicalType: TimeMicrosLogicalType,
	}
}

func TimestampMicros() Primitive {
	return Primitive{
		Type:        LongType,
		LogicalType: TimestampMicrosLogicalType,
	}
}

func Nullable(schema Schema) Union {
	if union, ok := schema.(Union); ok {
		var found bool
		for _, v := range union {
			if v == Null() {
				found = true
			}
		}
		if !found {
			return append(Union{Null()}, union)
		}
		return union
	}

	return Union{
		Null(),
		schema,
	}
}
