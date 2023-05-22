package gopher

import (
	"encoding/json"
	"fmt"
	"go/format"
	"strings"

	"github.com/estenssoros/tabla/internal/helpers"
	"github.com/pkg/errors"
)

// Field field on a go struct
type Field struct {
	Name     string `json:"name"`
	Type     GoType `json:"type"`
	Tag      string `json:"tag"`
	SQLType  string
	SQLExtra string
}

// SnakeName converts a field name to snake
func (f *Field) SnakeName() string {
	if f.Tag != "" {
		return f.Tag
	}
	return helpers.ToSnake(f.Name)
}

// CamelName convert a field name to camel
func (f *Field) CamelName() string {
	if f.Name == "id" {
		return "ID"
	}
	return helpers.ToCamel(f.Name)
}

// ToGo converts to go fmt field
func (f *Field) ToGo() string {
	// if f.SQLType == "" {
	return fmt.Sprintf(
		"    %s %s `json:\"%s\" db:\"%s\"`",
		f.CamelName(),
		f.Type,
		f.SnakeName(),
		f.SnakeName(),
	)
	// }
	// if f.SQLExtra == "" {
	// 	return fmt.Sprintf(
	// 		"    %s %s `json:\"%s\" db:\"%s,%s\"`",
	// 		f.CamelName(),
	// 		f.Type,
	// 		f.SnakeName(),
	// 		f.SnakeName(),
	// 		f.SQLType,
	// 	)
	// }
	// return fmt.Sprintf(
	// 	"    %s %s `json:\"%s\" db:\"%s,%s,%s\"`",
	// 	f.CamelName(),
	// 	f.Type,
	// 	f.SnakeName(),
	// 	f.SnakeName(),
	// 	f.SQLType,
	// 	f.SQLExtra,
	// )
}

// Struct a go struct
type Struct struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"fields"`
}

// SnakeName converts a go struct name to snake
func (s *Struct) SnakeName() string {
	return helpers.ToSnake(s.Name)
}

// CamelName converts a go struct name to camel
func (s *Struct) CamelName() string {
	return helpers.ToCamel(s.Name)
}

func (s Struct) String() string {
	ju, _ := json.MarshalIndent(s, "", " ")
	return string(ju)
}

// ToGoFields creates the text for a go struct
func (s *Struct) ToGoFields() string {
	fields := []string{}
	for _, f := range s.Fields {
		fields = append(fields, f.ToGo())
	}
	return strings.Join(fields, "\n")
}

func (s *Struct) Stringer() string {
	expr := `func(%s %s) String() string {
		ju,_:=json.Marshal(%s)
		return string(ju)
	}
	`
	return fmt.Sprintf(expr, s.SnakeName()[:1], s.CamelName(), s.SnakeName()[:1])
}
func (s *Struct) TableName() string {
	expr := `
	func (%s %s) TableName() string{
		return "%s"
	}
	`
	return fmt.Sprintf(expr, s.SnakeName()[:1], s.CamelName(), s.SnakeName())
}

// ToGo converts go struct to text definition
func (s *Struct) ToGo() (string, error) {
	expr := fmt.Sprintf("// %s\ntype %s struct {\n%s\n}\n", s.CamelName(), s.CamelName(), s.ToGoFields())
	expr += s.Stringer()
	expr += s.TableName()
	b, err := format.Source([]byte(expr))
	if err != nil {
		return "", errors.Wrap(err, "format node")
	}
	return string(b), nil
}
