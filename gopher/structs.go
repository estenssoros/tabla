package gopher

import (
	"encoding/json"
	"fmt"
	"go/format"
	"strings"

	"github.com/estenssoros/tabla/helpers"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

// GoField field on a go struct
type GoField struct {
	Name string `json:"name"`
	Type GoType `json:"type"`
	Tag  string `json:"tag"`
}

func (f *GoField) snakeName() string {
	if f.Tag != "" {
		return f.Tag
	}
	return helpers.ToSnake(f.Name)
}
func (f *GoField) camelName() string {
	if f.Name == "id" {
		return "ID"
	}
	return strcase.ToCamel(f.Name)
}

// GoStruct a go struct
type GoStruct struct {
	Name   string     `json:"name"`
	Fields []*GoField `json:"fields"`
}

func (s *GoStruct) snakeName() string {
	return helpers.ToSnake(s.Name)
}

func (s *GoStruct) camelName() string {
	return strcase.ToCamel(s.Name)
}

func (s GoStruct) String() string {
	ju, _ := json.MarshalIndent(s, "", " ")
	return string(ju)
}

// ToGoFields creates the text for a go struct
func (s *GoStruct) ToGoFields() string {
	fields := []string{}
	for _, f := range s.Fields {
		field := fmt.Sprintf(
			"    %s %s `json:\"%s\" db:\"%s\"`",
			f.camelName(),
			f.Type,
			f.snakeName(),
			f.snakeName(),
		)
		fields = append(fields, field)
	}
	return strings.Join(fields, "\n")
}

// ToGo converts go struct to text definition
func (s *GoStruct) ToGo() (string, error) {
	expr := fmt.Sprintf("type %s struct {\n%s\n}", s.camelName(), s.ToGoFields())

	b, err := format.Source([]byte(expr))
	if err != nil {
		return "", errors.Wrap(err, "format node")
	}

	return string(b), nil
}
