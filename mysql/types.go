package mysql

import (
	"github.com/estenssoros/tabla/gopher"
	"github.com/pkg/errors"
	"github.com/xwb1989/sqlparser"
)

type mySQLType string

var (
	intType      mySQLType = "int"
	varcharType  mySQLType = "varchar"
	textType     mySQLType = "text"
	dateTimeType mySQLType = "datetime"
	boolType     mySQLType = "tinyint"
	floatType    mySQLType = "float"
)

func (m mySQLType) ToGo(s *sqlparser.SQLVal) (gopher.GoType, error) {
	switch m {
	case intType:
		return gopher.IntType, nil
	case varcharType:
		if string(s.Val) == "36" {
			return gopher.UUIDType, nil
		}
		return gopher.StringType, nil
	case textType:
		return gopher.StringType, nil
	case dateTimeType:
		return gopher.TimeType, nil
	case boolType:
		return gopher.BoolType, nil
	case floatType:
		return gopher.FloatType, nil
	default:
		return "", errors.Errorf("unknown type %s", m)
	}
}

func (m mySQLType) ToGoNulls() gopher.GoType {
	switch m {
	case intType:
		return gopher.NullsIntType
	case varcharType, textType:
		return gopher.NullsStringType
	case dateTimeType:
		return gopher.NullsTimeType
	case boolType:
		return gopher.NullsBoolType
	case floatType:
		return gopher.NullsFloatType
	default:
		return ``
	}
}
