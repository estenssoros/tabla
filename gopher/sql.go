package gopher

import (
	"github.com/pkg/errors"
	"github.com/xwb1989/sqlparser"
)

// DropCreate parses go struct src into create statement
func DropCreate(src string, d Dialect) (string, error) {
	goStruct, err := parseGoSrc(src)
	if err != nil {
		return "", errors.Wrap(err, "parse go src")
	}
	create, err := d.Create(goStruct)
	if err != nil {
		return "", errors.Wrap(err, "dialect create")
	}
	return d.DropIfExists(goStruct) + create, nil
}

// ParseSQLToGoStruct parses raw sql into a go struct
func ParseSQLToGoStruct(sql string, dialect Converter, nulls bool) (*GoStruct, error) {
	sql = dialect.PrepareStatment(sql)
	stmt, err := sqlparser.ParseStrictDDL(sql)
	if err != nil {
		return nil, errors.Wrap(err, "sql parser parse")
	}
	ddl, ok := stmt.(*sqlparser.DDL)
	if !ok {
		return nil, errors.New("could not coerce to sqlparser.DDL")
	}

	if ddl.Action != "create" {
		return nil, errors.New("only create statements supported")
	}
	fields := []*GoField{}
	for _, c := range ddl.TableSpec.Columns {
		field, err := dialect.ColDefToGoField(c, nulls)
		if err != nil {
			return nil, errors.Wrap(err, "dialect coldef to go field")
		}
		fields = append(fields, field)
	}
	return &GoStruct{
		Name:   ddl.NewName.Name.String(),
		Fields: fields,
	}, nil
}
