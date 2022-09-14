package gopher

// GoType different supported types
type GoType string

var (
	// IntType represents an int
	IntType GoType = "int"
	// StringType represents a string
	StringType GoType = "string"
	// TimeType represents a time
	TimeType GoType = "time.Time"
	// BoolType represents a bool
	BoolType GoType = "bool"
	// FloatType represents a float
	FloatType GoType = "float64"
	// NullsIntType represents a null int
	NullsIntType GoType = "nulls.Int"
	// NullsStringType represents a null string
	NullsStringType GoType = "nulls.String"
	// NullsTimeType represents a null time
	NullsTimeType GoType = "nulls.Time"
	// NullsBoolType represents a null bool
	NullsBoolType GoType = "nulls.Bool"
	// NullsFloatType represents a float
	NullsFloatType GoType = "nulls.Float64"
	// UUIDType represents a uuid
	UUIDType GoType = "uuid.UUID"
)
