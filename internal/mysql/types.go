package mysql

import (
	"github.com/estenssoros/tabla/internal/gopher"
	"github.com/pkg/errors"
	"github.com/xwb1989/sqlparser"
)

// SQLType mysql data types
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

// ToGoField conerts data type to a go field
func (m SQLType) ToGoField(nulls bool, c *sqlparser.ColumnDefinition) (*gopher.GoField, error) {
	field := &gopher.GoField{
		Name:    c.Name.String(),
		SQLType: string(m),
	}
	if c.Type.Length != nil {
		field.SQLExtra = string(c.Type.Length.Val)
	}

	if nulls {
		goType, err := m.toGoNulls(c.Type.Length)
		if err != nil {
			return nil, err
		}
		field.Type = goType
	} else {
		goType, err := m.toGoStandard(c.Type.Length)
		if err != nil {
			return nil, err
		}
		field.Type = goType
	}
	return field, nil

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
	case intType, bigIntType, smallIntType:
		return gopher.NullsIntType, nil
	case varcharType:
		if string(s.Val) == "36" {
			return gopher.UUIDType, nil
		}
		return gopher.NullsStringType, nil
	case textType, longTextType:
		return gopher.NullsStringType, nil
	case dateTimeType, dateType:
		return gopher.NullsTimeType, nil
	case boolType:
		return gopher.NullsBoolType, nil
	case floatType, doubleType, decimalType:
		return gopher.NullsFloatType, nil
	default:
		return "", errors.Errorf("unknown type %s", m)
	}
}

type converter struct{}

func (c converter) ColDefToGoField(colDef *sqlparser.ColumnDefinition, nulls bool) (*gopher.GoField, error) {
	field, err := SQLType(colDef.Type.Type).ToGoField(nulls, colDef)
	if err != nil {
		return nil, errors.Wrap(err, "myslq type to go")
	}
	return field, nil
}

func (c converter) PrepareStatment(sql string) string {
	return removeKeywords(sql)
}
