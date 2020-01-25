package gopher

import (
	"encoding/json"
	"fmt"
	"go/format"
	"strings"

	"github.com/estenssoros/tabla/helpers"
	"github.com/pkg/errors"
)

// GoField field on a go struct
type GoField struct {
	Name     string `json:"name"`
	Type     GoType `json:"type"`
	Tag      string `json:"tag"`
	SQLType  string
	SQLExtra string
}

// SnakeName converts a field name to snake
func (f *GoField) SnakeName() string {
	if f.Tag != "" {
		return f.Tag
	}
	return helpers.ToSnake(f.Name)
}

// CamelName convert a field name to camel
func (f *GoField) CamelName() string {
	if f.Name == "id" {
		return "ID"
	}
	return helpers.ToCamel(f.Name)
}

// ToGo converts to go fmt field
func (f *GoField) ToGo() string {
	if f.SQLType == "" {
		return fmt.Sprintf(
			"    %s %s `json:\"%s\" db:\"%s\"`",
			f.CamelName(),
			f.Type,
			f.SnakeName(),
			f.SnakeName(),
		)
	}
	if f.SQLExtra == "" {
		return fmt.Sprintf(
			"    %s %s `json:\"%s\" db:\"%s,%s\"`",
			f.CamelName(),
			f.Type,
			f.SnakeName(),
			f.SnakeName(),
			f.SQLType,
		)
	}
	return fmt.Sprintf(
		"    %s %s `json:\"%s\" db:\"%s,%s,%s\"`",
		f.CamelName(),
		f.Type,
		f.SnakeName(),
		f.SnakeName(),
		f.SQLType,
		f.SQLExtra,
	)
}

// GoStruct a go struct
type GoStruct struct {
	Name   string     `json:"name"`
	Fields []*GoField `json:"fields"`
}

// SnakeName converts a go struct name to snake
func (s *GoStruct) SnakeName() string {
	return helpers.ToSnake(s.Name)
}

// CamelName converts a go struct name to camel
func (s *GoStruct) CamelName() string {
	return helpers.ToCamel(s.Name)
}

func (s GoStruct) String() string {
	ju, _ := json.MarshalIndent(s, "", " ")
	return string(ju)
}

// ToGoFields creates the text for a go struct
func (s *GoStruct) ToGoFields() string {
	fields := []string{}
	for _, f := range s.Fields {
		fields = append(fields, f.ToGo())
	}
	return strings.Join(fields, "\n")
}

// ToGo converts go struct to text definition
func (s *GoStruct) ToGo() (string, error) {
	expr := fmt.Sprintf("// %s\ntype %s struct {\n%s\n}\n", s.CamelName(), s.CamelName(), s.ToGoFields())
	b, err := format.Source([]byte(expr))
	if err != nil {
		return "", errors.Wrap(err, "format node")
	}
	return string(b), nil
}
