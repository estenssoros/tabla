package gopher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSnakeNameTables = []struct {
	in  string
	out string
}{
	{"in", "in"},
	{"In", "in"},
	{"IN", "in"},
	{"AsdfAsdf", "asdf_asdf"},
	{"asdfAsdf", "asdf_asdf"},
}

func TestSnakeName(t *testing.T) {
	for _, tt := range testSnakeNameTables {
		gField := &GoField{Name: tt.in}
		assert.Equal(t, gField.snakeName(), tt.out)
	}
}

var testGoFieldCamelNameTables = []struct {
	in  string
	out string
}{
	{"id", "ID"},
	{"asdf_asdf", "AsdfAsdf"},
}

func TestGoFieldCamelName(t *testing.T) {
	for _, tt := range testGoFieldCamelNameTables {
		gField := &GoField{Name: tt.in}
		assert.Equal(t, gField.camelName(), tt.out)
	}
}

var testGoStructCamelNameTables = []struct {
	in  string
	out string
}{
	{"asdf_asdf", "AsdfAsdf"},
}

func TestGoStructCamelName(t *testing.T) {
	for _, tt := range testGoStructCamelNameTables {
		gStruct := &GoStruct{Name: tt.in}
		assert.Equal(t, gStruct.camelName(), tt.out)
		assert.NotEmpty(t, gStruct.String())
	}
}

var testGoStruct = &GoStruct{
	Name: "test",
	Fields: []*GoField{
		&GoField{
			Name: "adsf",
			Type: IntType,
		},
	},
}

var testGoStructToGoTables = []struct {
	in  *GoStruct
	err bool
}{
	{testGoStruct, false},
}

func TestGoStructToGo(t *testing.T) {
	for _, tt := range testGoStructToGoTables {
		out, err := tt.in.ToGo()
		if tt.err {
			assert.NotNil(t, err)
			assert.Nil(t, out)
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, out)
		}
	}
}