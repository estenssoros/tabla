package gopher

import (
	"github.com/pkg/errors"
)

// DropCreate parses go struct src into create statement
func DropCreate(src string, d Dialect) (string, error) {
	goStruct, err := parseGoSrc(src)
	if err != nil {
		return "", errors.Wrap(err, "parse go src")
	}
	create, err := d.Create(goStruct)
	if err != nil {
		return "", errors.Wrap(err, "mysql create")
	}
	return d.DropIfExists(goStruct) + create, nil
}
