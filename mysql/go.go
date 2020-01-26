package mysql

import (
	"github.com/estenssoros/tabla/gopher"
	"github.com/pkg/errors"
)

// Go parses mysql to go struct text
func Go(sql string, nulls bool) (string, error) {
	goStruct, err := gopher.ParseSQLToGoStruct(sql, &converter{}, nulls)
	if err != nil {
		return "", errors.Wrap(err, "parse mysql to go struct")
	}
	return goStruct.ToGo()
}
