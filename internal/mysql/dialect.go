package mysql

import (
	"fmt"
	"strings"

	"github.com/estenssoros/tabla/internal/gopher"
	"github.com/pkg/errors"
)

// Dialect mysql dialect
type Dialect struct{}

// DropIfExists drop if exists statement
func (d Dialect) DropIfExists(s *gopher.GoStruct) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS `%s`;\n", s.SnakeName())
}

// Fields construct the column field definitions
func (d Dialect) Fields(goFields []*gopher.GoField) (string, error) {
	var hasID bool
	fields := []string{}
	for _, goField := range goFields {
		if !hasID && strings.ToLower(goField.Name) == "id" {
			hasID = true
		}
		field, err := d.Field(goField)
		if err != nil {
			return "", errors.Wrap(err, "dialect field")
		}
		fields = append(fields, field)
	}
	if hasID {
		fields = append(fields, "PRIMARY KEY(`id`)")
	}
	if len(fields) > 0 {
		fields[0] = "    " + fields[0]
	}
	return strings.Join(fields, "\n    , "), nil
}

// Field converts a go field to SQL syntax
func (d Dialect) Field(goField *gopher.GoField) (string, error) {
	if goField.SQLType == "" {
		dataType, err := d.FieldToDataType(goField.Type)
		if err != nil {
			return "", errors.Wrap(err, "field to datatype")
		}
		return fmt.Sprintf(
			"`%s` %s",
			goField.SnakeName(),
			dataType,
		), nil
	}
	if goField.SQLExtra == "" {
		return fmt.Sprintf(
			"`%s` %s",
			goField.SnakeName(),
			goField.SQLType,
		), nil
	}
	return fmt.Sprintf(
		"`%s` %s(%s)",
		goField.SnakeName(),
		goField.SQLType,
		goField.SQLExtra,
	), nil
}

// Create for a create table statement
func (d Dialect) Create(s *gopher.GoStruct) (string, error) {
	fieldsFormatted, err := d.Fields(s.Fields)
	if err != nil {
		return "", errors.Wrap(err, "mysql fields")
	}
	stmt := fmt.Sprintf(
		"CREATE TABLE `%s` (\n%s\n);",
		s.SnakeName(),
		fieldsFormatted,
	)
	return stmt, nil
}

// FieldToDataType converts a field to a datatype
func (d Dialect) FieldToDataType(goType gopher.GoType) (string, error) {
	switch goType {
	case gopher.IntType, gopher.NullsIntType:
		return `INT`, nil
	case gopher.StringType, gopher.NullsStringType:
		return `VARCHAR({update})`, nil
	case gopher.TimeType, gopher.NullsTimeType:
		return `DATETIME`, nil
	case gopher.BoolType, gopher.NullsBoolType:
		return `TINYINT(1)`, nil
	case gopher.FloatType, gopher.NullsFloatType:
		return `FLOAT`, nil
	case gopher.UUIDType:
		return `VARCHAR(36)`, nil
	default:
		return "", errors.Errorf("unknown type: %s", goType)
	}
}
