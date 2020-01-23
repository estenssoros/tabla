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

func parseSrc(src string) (*GoStruct, error) {
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
	goFields := []*GoField{}
	fields := structDecl.Fields.List
	for _, field := range fields {

		typeExpr := field.Type
		goField := &GoField{
			Name: field.Names[0].Name,
			Type: GoType(src[typeExpr.Pos()-1 : typeExpr.End()-1]),
		}
		if field.Tag != nil {
			tag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
			if dbTag := tag.Get("db"); dbTag != "" {
				goField.Tag = dbTag
			}
		}
		goFields = append(goFields, goField)
	}
	return &GoStruct{
		Name:   structType.Name.Name,
		Fields: goFields,
	}, nil
}
