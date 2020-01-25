package mysql

import (
	"fmt"
	"strings"

	"github.com/estenssoros/tabla/gopher"
	"github.com/pkg/errors"
)

type Dialect struct{}

func (d Dialect) DropIfExists(s *gopher.GoStruct) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS `%s`;\n", s.SnakeName())
}

func (d Dialect) Fields(goFields []*gopher.GoField) (string, error) {
	var hasID bool
	fields := []string{}
	for _, goField := range goFields {
		if strings.ToLower(goField.Name) == "id" {
			hasID = true
		}
		dataType, err := d.FieldToDataType(goField.Type)
		if err != nil {
			return "", errors.Wrap(err, "field to datatype")
		}
		field := fmt.Sprintf(
			"`%s` %s",
			goField.SnakeName(),
			dataType,
		)
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

func (d Dialect) Create(s *gopher.GoStruct) (string, error) {
	fields, err := d.Fields(s.Fields)
	if err != nil {
		return "", errors.Wrap(err, "mysql fields")
	}
	stmt := fmt.Sprintf(
		"CREATE TABLE `%s` (\n%s\n);",
		s.SnakeName(),
		fields,
	)
	return stmt, nil
}

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
