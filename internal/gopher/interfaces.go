package gopher

import "github.com/xwb1989/sqlparser"

// Dialect interface for sql dialects
type Dialect interface {
	DropIfExists(*Struct) string
	Create(*Struct) (string, error)
}

// Converter interface for converting from SQL to goField
type Converter interface {
	ColDefToGoField(*sqlparser.ColumnDefinition, bool) (*Field, error)
	PrepareStatment(string) string
}
