package gopher

import "github.com/xwb1989/sqlparser"

// Dialect interface for sql dialects
type Dialect interface {
	DropIfExists(*GoStruct) string
	Create(*GoStruct) (string, error)
}

// Converter interface for converting from SQL to goField
type Converter interface {
	ColDefToGoField(*sqlparser.ColumnDefinition, bool) (*GoField, error)
	PrepareStatment(string) string
}
