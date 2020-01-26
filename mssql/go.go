package mssql

import (
	"github.com/estenssoros/tabla/gopher"
	"github.com/pkg/errors"
)

func Go(sql string, nulls bool) (string, error) {
	goStruct, err := gopher.ParseSQLToGoStruct(sql, &converter{}, nulls)
	if err != nil {
		return "", errors.Wrap(err, "parse mysql to go struct")
	}
	return goStruct.ToGo()
}
