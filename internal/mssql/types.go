package mssql

import (
	"github.com/estenssoros/tabla/internal/gopher"
	"github.com/pkg/errors"
	"github.com/xwb1989/sqlparser"
)

// SQLType sql server data types
type SQLType string

var (
	bigIntType         SQLType = "bigint"
	binaryType         SQLType = "binary"
	bitType            SQLType = "bit"
	charType           SQLType = "char"
	dateType           SQLType = "date"
	dateTimeType       SQLType = "datetime"
	dateTime2Type      SQLType = "datetime2"
	dateTimeOffsetType SQLType = "datetimeoffset"
	decimalType        SQLType = "decimal"
	floatType          SQLType = "float"
	intType            SQLType = "int"
	ncharType          SQLType = "nchar"
	ntextType          SQLType = "ntext"
	numericType        SQLType = "numeric"
	nvarcharType       SQLType = "nvarchar"
	smallDatetimeType  SQLType = "smalldatetime"
	smallIntType       SQLType = "smallint"
	textType           SQLType = "text"
	timeType           SQLType = "time"
	timestampType      SQLType = "timestamp"
	tinyIntType        SQLType = "tinyint"
	varcharType        SQLType = "varchar"
)

// ToGoField converts a data type to a go field
func (m SQLType) ToGoField(nulls bool, c *Column) (*gopher.Field, error) {
	field := &gopher.Field{
		Name:    c.Name,
		SQLType: string(m),
	}
	if c.Length != "" {
		field.SQLExtra = c.Length
	}

	if nulls {
		goType, err := m.toGoNulls(c.Length)
		if err != nil {
			return nil, err
		}
		field.Type = goType
	} else {
		goType, err := m.toGoStandard(c.Length)
		if err != nil {
			return nil, err
		}
		field.Type = goType
	}
	return field, nil
}

func (m SQLType) toGoStandard(s string) (gopher.GoType, error) {
	switch m {
	case intType, bigIntType, smallIntType, smallIntType:
		return gopher.IntType, nil
	case varcharType:
		if s == "36" {
			return gopher.UUIDType, nil
		}
		return gopher.StringType, nil
	case varcharType, charType, textType, ncharType, ntextType, nvarcharType:
		return gopher.StringType, nil
	case timestampType, dateTimeType, dateType, dateTime2Type, dateTimeOffsetType, smallDatetimeType, timeType:
		return gopher.TimeType, nil
	case binaryType, bitType, tinyIntType:
		return gopher.BoolType, nil
	case floatType, decimalType, numericType:
		return gopher.FloatType, nil
	default:
		return "", errors.Errorf("unknown type %s", m)
	}
}

func (m SQLType) toGoNulls(s string) (gopher.GoType, error) {
	switch m {
	case intType, bigIntType, smallIntType, smallIntType:
		return gopher.NullsIntType, nil
	case varcharType:
		if s == "36" {
			return gopher.UUIDType, nil
		}
		return gopher.NullsStringType, nil
	case varcharType, charType, textType, ncharType, ntextType, nvarcharType:
		return gopher.NullsStringType, nil
	case timestampType, dateTimeType, dateType, dateTime2Type, dateTimeOffsetType, smallDatetimeType, timeType:
		return gopher.NullsTimeType, nil
	case binaryType, bitType, tinyIntType:
		return gopher.NullsBoolType, nil
	case floatType, decimalType, numericType:
		return gopher.NullsFloatType, nil
	default:
		return "", errors.Errorf("unknown type %s", m)
	}
}

type converter struct{}

func (c converter) ColDefToGoField(colDef *sqlparser.ColumnDefinition, nulls bool) (*gopher.Field, error) {
	col := &Column{
		Name: colDef.Name.String(),
		Type: SQLType(colDef.Type.Type),
	}
	if colDef.Type.Length != nil {
		col.Length = string(colDef.Type.Length.Val)
	}
	field, err := SQLType(colDef.Type.Type).ToGoField(nulls, col)
	if err != nil {
		return nil, errors.Wrap(err, "myslq type to go")
	}
	return field, nil
}

func (c converter) PrepareStatment(sql string) string {
	return sql
}
