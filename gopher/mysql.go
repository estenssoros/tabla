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

func (m mysql) Fields(goFields []*GoField) string {
	var hasID bool
	fields := []string{}
	for _, goField := range goFields {
		if strings.ToLower(goField.Name) == "id" {
			hasID = true
		}
		field := fmt.Sprintf(
			"`%s` %s",
			goField.snakeName(),
			m.FieldToDataType(goField.Type),
		)
		fields = append(fields, field)
	}
	if hasID {
		fields = append(fields, "PRIMARY KEY(`id`)")
	}
	if len(fields) > 0 {
		fields[0] = "    " + fields[0]
	}
	return strings.Join(fields, "\n    , ")
}

func (m mysql) Create(s *GoStruct) string {
	stmt := fmt.Sprintf(
		"CREATE TABLE `%s` (\n%s\n);",
		s.snakeName(),
		m.Fields(s.Fields),
	)
	return stmt
}

func (m mysql) FieldToDataType(goType GoType) string {
	switch goType {
	case IntType, NullsIntType:
		return `INT`
	case StringType, NullsStringType:
		return `VARCHAR({update})`
	case TimeType, NullsTimeType:
		return `DATETIME`
	case BoolType, NullsBoolType:
		return `TINYINT(1)`
	case FloatType, NullsFloatType:
		return `FLOAT`
	case UuidType:
		return `VARCHAR(36)`
	default:
		return ``
	}
}

func MySQL(src string) (string, error) {
	var out string
	goStruct, err := parseSrc(src)
	if err != nil {
		return "", errors.Wrap(err, "parse src")
	}
	m := mysql{}
	out += m.DropIfExists(goStruct)
	out += m.Create(goStruct)
	return out, nil
}
