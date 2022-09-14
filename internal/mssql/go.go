package mssql

import (
	"github.com/estenssoros/tabla/internal/gopher"
	"github.com/pkg/errors"
)

// Go parses sql into go struct
func Go(sql string, nulls bool) (string, error) {
	goStruct, err := gopher.ParseSQLToGoStruct(sql, &converter{}, nulls)
	if err != nil {
		return "", errors.Wrap(err, "parse mysql to go struct")
	}
	return goStruct.ToGo()
}
