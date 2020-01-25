package mysql

import (
	"github.com/estenssoros/tabla/gopher"
	"github.com/pkg/errors"
	"github.com/xwb1989/sqlparser"
)

// Go parses mysql to go struct text
func Go(sql string, nulls bool) (string, error) {
	goStruct, err := parseMySQLToGoStruct(sql, nulls)
	if err != nil {
		return "", errors.Wrap(err, "parse mysql to go struct")
	}
	return goStruct.ToGo()
}

func parseMySQLToGoStruct(sql string, nulls bool) (*gopher.GoStruct, error) {
	sql = removeKeywords(sql)
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
		field, err := SQLType(c.Type.Type).ToGoField(nulls, c)
		if err != nil {
			return nil, errors.Wrap(err, "mysql type to go")
		}
		fields = append(fields, field)
	}
	return &gopher.GoStruct{
		Name:   ddl.NewName.Name.String(),
		Fields: fields,
	}, nil
}
