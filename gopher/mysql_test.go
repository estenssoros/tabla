package gopher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMySQLTables = []struct {
	in  string
	err bool
}{
	{
		"type Test struct {\n" +
			"id uuid.UUID\n" +
			"intField int `db:\"asdf\"`\n" +
			"stringField string\n" +
			"timeField time.Time\n" +
			"boolField bool\n" +
			"floatField float64\n" +
			"nullsIntField nulls.Int\n" +
			"nullsStringField nulls.String\n" +
			"nullsTimeField nulls.Time\n" +
			"nullsBoolField nulls.Bool\n" +
			"nullsFloatField nulls.Float64\n" +
			"}",
		false,
	},
	{
		"type Test struct {\n" +
			"id uuid.UUID\n" +
			"intField int `db:\"asdf\"`\n" +
			"stringField string\n" +
			"timeField time.Time\n" +
			"boolField bool\n" +
			"floatField float64\n" +
			"nullsIntField nulls.Int\n" +
			"nullsStringField nulls.String\n" +
			"nullsTimeField nulls.Time\n" +
			"nullsBoolField nulls.Bool\n" +
			"nullsFloatField map[string]int\n" +
			"}",
		true,
	},
	{
		"type Test struct {\n" +
			"id uuid.UUID\n" +
			"intField int `db:\"asdf\"`\n" +
			"stringField string\n" +
			"timeField time.Time\n" +
			"boolField bool\n" +
			"floatField float64\n" +
			"nullsIntField nulls.Int\n" +
			"nullsStringField nulls.String\n" +
			"nullsTimeField nulls.Time\n" +
			"nullsBoolField nulls.Bool\n" +
			"nullsFloatField map[string]int\n",
		true,
	},
}

func TestMySQL(t *testing.T) {
	for _, tt := range testMySQLTables {
		out, err := MySQL(tt.in)
		if tt.err {
			assert.NotNil(t, err)
			assert.Empty(t, out)
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, out)
		}
	}
}
