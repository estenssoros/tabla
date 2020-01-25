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

var testParseSrcTables = []struct {
	in  string
	err bool
}{
	{"type asdf struct {}", false},
	{"type asdf struct {", true},
	{"var asdf string", true},
	{"type asdf string", true},
	{"func asfd() {}", true},
}

func TestParseSrc(t *testing.T) {
	for _, tt := range testParseSrcTables {
		out, err := parseSrc(tt.in)
		if tt.err {
			assert.NotNil(t, err)
			assert.Empty(t, out)
		} else {
			assert.Nil(t, err)
			assert.NotEmpty(t, out)
		}
	}
}
