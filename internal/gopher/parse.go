package gopher

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

func checksrc(src string) string {
	if strings.HasPrefix(src, "package") {
		return src
	}
	return "package asdf\n" + src
}

type tags struct {
	rawTag reflect.StructTag
}

func (t tags) Get(key string) *tag {
	v := t.rawTag.Get(key)
	if v == "" {
		return &tag{valid: false}
	}
	var value string
	options := []string{}
	if strings.Contains(v, ",") {
		splits := strings.Split(v, ",")
		value = splits[0]
		options = append(options, splits[1:]...)
	} else {
		value = v
	}
	return &tag{
		key:     key,
		value:   value,
		options: options,
		valid:   true,
	}
}

type tag struct {
	key     string
	value   string
	options []string
	valid   bool
}

func parseTag(s string) *tags {
	return &tags{reflect.StructTag(s)}
}

func parseGoSrc(src string) (*Struct, error) {
	src = checksrc(src)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	if err != nil {
		return nil, errors.Wrap(err, "parser parse file")
	}
	typeDecl, ok := f.Decls[0].(*ast.GenDecl)
	if !ok {
		return nil, errors.New("could not find type delcaration")
	}
	structType, ok := typeDecl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return nil, errors.New("could not find type spec")
	}
	structDecl, ok := structType.Type.(*ast.StructType)
	if !ok {
		return nil, errors.New("could not find struct type")
	}
	goFields := []*Field{}
	fields := structDecl.Fields.List
	for _, field := range fields {

		typeExpr := field.Type
		goField := &Field{
			Name: field.Names[0].Name,
			Type: GoType(src[typeExpr.Pos()-1 : typeExpr.End()-1]),
		}
		if field.Tag != nil {
			tag := parseTag(field.Tag.Value[1 : len(field.Tag.Value)-1])

			if dbTag := tag.Get("db"); dbTag.valid {
				goField.Tag = dbTag.value
				if len(dbTag.options) > 0 {
					goField.SQLType = dbTag.options[0]
				}
				if len(dbTag.options) > 1 {
					goField.SQLExtra = dbTag.options[1]
				}
			}
		}
		goFields = append(goFields, goField)
	}
	return &Struct{
		Name:   structType.Name.Name,
		Fields: goFields,
	}, nil
}
