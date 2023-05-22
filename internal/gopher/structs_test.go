package gopher

import (
	"fmt"
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
	for i, tt := range testSnakeNameTables {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			gField := &Field{Name: tt.in}
			assert.Equal(t, gField.SnakeName(), tt.out)
		})
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

	for i, tt := range testGoFieldCamelNameTables {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			gField := &Field{Name: tt.in}
			assert.Equal(t, gField.CamelName(), tt.out)
		})
	}
}

var testGoStructCamelNameTables = []struct {
	in  string
	out string
}{
	{"asdf_asdf", "AsdfAsdf"},
}

func TestGoStructCamelName(t *testing.T) {
	for i, tt := range testGoStructCamelNameTables {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			gStruct := &Struct{Name: tt.in}
			assert.Equal(t, gStruct.CamelName(), tt.out)
			assert.NotEmpty(t, gStruct.String())
		})
	}
}

var testGoStruct = &Struct{
	Name: "test",
	Fields: []*Field{
		&Field{
			Name: "adsf",
			Type: IntType,
		},
	},
}

var testGoStructToGoTables = []struct {
	in  *Struct
	err bool
}{
	{testGoStruct, false},
}

func TestGoStructToGo(t *testing.T) {
	for i, tt := range testGoStructToGoTables {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			out, err := tt.in.ToGo()
			if tt.err {
				assert.NotNil(t, err)
				assert.Nil(t, out)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, out)
			}
		})
	}
}
