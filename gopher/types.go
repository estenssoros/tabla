package gopher

type GoType string

var (
	IntType         GoType = "int"
	StringType      GoType = "string"
	TimeType        GoType = "time.Time"
	BoolType        GoType = "bool"
	FloatType       GoType = "float64"
	NullsIntType    GoType = "nulls.Int"
	NullsStringType GoType = "nulls.String"
	NullsTimeType   GoType = "nulls.Time"
	NullsBoolType   GoType = "nulls.Bool"
	NullsFloatType  GoType = "nulls.Float64"
	UuidType        GoType = "uuid.UUID"
)
