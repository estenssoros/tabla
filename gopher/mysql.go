package gopher

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type mysql struct{}

func (m mysql) DropIfExists(s *GoStruct) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS `%s`;\n", s.snakeName())
}

func (m mysql) Fields(goFields []*GoField) (string, error) {
	var hasID bool
	fields := []string{}
	for _, goField := range goFields {
		if strings.ToLower(goField.Name) == "id" {
			hasID = true
		}
		dataType, err := m.FieldToDataType(goField.Type)
		if err != nil {
			return "", errors.Wrap(err, "field to datatype")
		}
		field := fmt.Sprintf(
			"`%s` %s",
			goField.snakeName(),
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

func (m mysql) Create(s *GoStruct) (string, error) {
	fields, err := m.Fields(s.Fields)
	if err != nil {
		return "", errors.Wrap(err, "mysql fields")
	}
	stmt := fmt.Sprintf(
		"CREATE TABLE `%s` (\n%s\n);",
		s.snakeName(),
		fields,
	)
	return stmt, nil
}

func (m mysql) FieldToDataType(goType GoType) (string, error) {
	switch goType {
	case IntType, NullsIntType:
		return `INT`, nil
	case StringType, NullsStringType:
		return `VARCHAR({update})`, nil
	case TimeType, NullsTimeType:
		return `DATETIME`, nil
	case BoolType, NullsBoolType:
		return `TINYINT(1)`, nil
	case FloatType, NullsFloatType:
		return `FLOAT`, nil
	case UUIDType:
		return `VARCHAR(36)`, nil
	default:
		return "", errors.Errorf("unknown type: %s", goType)
	}
}

// MySQL parses go struct src into MySQL create statement
func MySQL(src string) (string, error) {
	goStruct, err := parseSrc(src)
	if err != nil {
		return "", errors.Wrap(err, "parse src")
	}
	m := mysql{}
	drop := m.DropIfExists(goStruct)
	create, err := m.Create(goStruct)
	if err != nil {
		return "", errors.Wrap(err, "mysql create")
	}
	return drop + create, nil
}
