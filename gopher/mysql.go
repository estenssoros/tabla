package gopher

import (
	"github.com/pkg/errors"
)

// DropCreate parses go struct src into create statement
func DropCreate(src string, d Dialect) (string, error) {
	goStruct, err := parseSrc(src)
	if err != nil {
		return "", errors.Wrap(err, "parse src")
	}
	drop := d.DropIfExists(goStruct)
	create, err := d.Create(goStruct)
	if err != nil {
		return "", errors.Wrap(err, "mysql create")
	}
	return drop + create, nil
}
