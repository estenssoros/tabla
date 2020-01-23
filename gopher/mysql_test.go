package gopher

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMySQLTables = []struct {
	in string
}{
	{
		`type Test struct {
		id uuid.UUID
		intField int
		stringField string
		timeField time.Time
		boolField bool
		floatField float64
		nullsIntField nulls.Int
		nullsStringField nulls.String
		nullsTimeField nulls.Time
		nullsBoolField nulls.Bool
		nullsFloatField nulls.Float64
	}`,
	},
}

func TestMySQL(t *testing.T) {
	for _, tt := range testMySQLTables {
		// fmt.Println(tt.in)
		out, err := MySQL(tt.in)
		assert.Nil(t, err)
		fmt.Println(out)
	}
}
