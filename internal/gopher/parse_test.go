package gopher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCheckSrcTables = []struct {
	in  string
	out string
}{
	{"asdf", "package asdf\nasdf"},
	{"package asdf", "package asdf"},
}

func TestCheckSrc(t *testing.T) {
	for _, tt := range testCheckSrcTables {
		out := checksrc(tt.in)
		assert.Equal(t, out, tt.out)
	}
}

var testParseSrcErrTables = []struct {
	in  string
	err bool
}{
	{"type asdf struct {}", false},
	{"type asdf struct {", true},
	{"var asdf string", true},
	{"type asdf string", true},
	{"func asfd() {}", true},
}

func TestParseSrcErr(t *testing.T) {
	for _, tt := range testParseSrcErrTables {
		out, err := parseGoSrc(tt.in)
		if tt.err {
			assert.NotNil(t, err)
			assert.Empty(t, out)
		} else {
			assert.Nil(t, err)
			assert.NotEmpty(t, out)
		}
	}
}

func TestParseTag(t *testing.T) {
	tag := `db:"asdf" json:"asdf,int,11"`
	tags := parseTag(tag)
	dbTag := tags.Get("db")
	assert.Equal(t, true, dbTag.valid)
	assert.Equal(t, 0, len(dbTag.options))
	jsonTag := tags.Get("json")
	assert.Equal(t, true, jsonTag.valid)
	assert.Equal(t, 2, len(jsonTag.options))
	noTag := tags.Get("asdf")
	assert.Equal(t, false, noTag.valid)
}
