package mysql

import (
	"github.com/estenssoros/tabla/gopher"
	"github.com/pkg/errors"
	"github.com/xwb1989/sqlparser"
)

// Go parses mysql to go struct text
func Go(sql string) (string, error) {
	goStruct, err := parseMySQLToGoStruct(sql)
	if err != nil {
		return "", errors.Wrap(err, "parse mysql to go struct")
	}
	return goStruct.ToGo()
}

func parseMySQLToGoStruct(sql string) (*gopher.GoStruct, error) {
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
	fields := []*gopher.GoField{}
	for _, c := range ddl.TableSpec.Columns {
		goType, err := mySQLType(c.Type.Type).ToGo(c.Type.Length)
		if err != nil {
			return nil, errors.Wrap(err, "mysql type to go")
		}
		field := &gopher.GoField{
			Name: c.Name.String(),
			Type: goType,
		}
		fields = append(fields, field)
	}
	return &gopher.GoStruct{
		Name:   ddl.NewName.Name.String(),
		Fields: fields,
	}, nil
}
