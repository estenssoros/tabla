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

func (m mySQLType) ToGo(nulls bool, s *sqlparser.SQLVal) (gopher.GoType, error) {
	if nulls {
		return m.toGoNulls(s)
	}
	return m.toGoStandard(s)

}

func (m mySQLType) toGoStandard(s *sqlparser.SQLVal) (gopher.GoType, error) {
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

func (m mySQLType) toGoNulls(s *sqlparser.SQLVal) (gopher.GoType, error) {
	switch m {
	case intType:
		return gopher.NullsIntType, nil
	case varcharType:
		if string(s.Val) == "36" {
			return gopher.UUIDType, nil
		}
		return gopher.NullsStringType, nil
	case textType:
		return gopher.NullsStringType, nil
	case dateTimeType:
		return gopher.NullsTimeType, nil
	case boolType:
		return gopher.NullsBoolType, nil
	case floatType:
		return gopher.NullsFloatType, nil
	default:
		return "", errors.Errorf("unknown type %s", m)
	}
}
