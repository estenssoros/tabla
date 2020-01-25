package mysql

import (
	"github.com/estenssoros/tabla/gopher"
	"github.com/pkg/errors"
	"github.com/xwb1989/sqlparser"
)

type SQLType string

var (
	intType      SQLType = "int"
	varcharType  SQLType = "varchar"
	textType     SQLType = "text"
	dateTimeType SQLType = "datetime"
	boolType     SQLType = "tinyint"
	floatType    SQLType = "float"
	bigIntType   SQLType = "bigint"
	doubleType   SQLType = "double"
	dateType     SQLType = "date"
	longTextType SQLType = "longtext"
	smallIntType SQLType = "smallint"
	decimalType  SQLType = "decimal"
)

func (m SQLType) ToGo(nulls bool, s *sqlparser.SQLVal) (gopher.GoType, error) {
	if nulls {
		return m.toGoNulls(s)
	}
	return m.toGoStandard(s)

}

func (m SQLType) toGoStandard(s *sqlparser.SQLVal) (gopher.GoType, error) {
	switch m {
	case intType, bigIntType, smallIntType:
		return gopher.IntType, nil
	case varcharType:
		if string(s.Val) == "36" {
			return gopher.UUIDType, nil
		}
		return gopher.StringType, nil
	case textType, longTextType:
		return gopher.StringType, nil
	case dateTimeType, dateType:
		return gopher.TimeType, nil
	case boolType:
		return gopher.BoolType, nil
	case floatType, doubleType, decimalType:
		return gopher.FloatType, nil
	default:
		return "", errors.Errorf("unknown type %s", m)
	}
}

func (m SQLType) toGoNulls(s *sqlparser.SQLVal) (gopher.GoType, error) {
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
