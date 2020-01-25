package mysql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testGoTables = []struct {
	in  string
	err bool
}{
	{
		"CREATE TABLE `test` (" +
			"`id` VARCHAR(36)" +
			", `int_field` INT" +
			", `string_field` TEXT " +
			", `time_field` DATETIME" +
			", `bool_field` TINYINT(1)" +
			", `float_field` FLOAT" +
			", `char_field` VARCHAR(100)" +
			", `nulls_int_field` INT" +
			", `nulls_string_field` TEXT" +
			", `nulls_time_field` DATETIME" +
			", `nulls_bool_field` TINYINT(1)" +
			", `nulls_float_field` FLOAT" +
			", PRIMARY KEY(`id`)" +
			")",
		false,
	},
	{
		"CREATE TABLE `test` (" +
			"`id` VARCHAR(36)" +
			", `int_field` INT" +
			", `string_field` TEXT " +
			", `time_field` DATETIME" +
			", `bool_field` TINYINT(1)" +
			", `float_field` FLOAT" +
			", `nulls_int_field` INT" +
			", `nulls_string_field` TEXT" +
			", `nulls_time_field` DATETIME" +
			", `nulls_bool_field` TINYINT(1)" +
			", `nulls_float_field` FLOAT" +
			", PRIMARY KEY(`id`)",
		true,
	},
	{
		"DROP TABLE `test`",
		true,
	},
	{
		"INSERT INTO `test` (id) VALUES (1)",
		true,
	},
	{
		"CREATE TABLE `test` (" +
			"`id` VARCHAR(36)" +
			", `int_field` INT" +
			", `string_field` TEXT " +
			", `time_field` DATETIME" +
			", `bool_field` TINYINT(1)" +
			", `float_field` FLOAT" +
			", `nulls_int_field` INT" +
			", `nulls_string_field` TEXT" +
			", `nulls_time_field` DATE" +
			", `nulls_bool_field` TINYINT(1)" +
			", `nulls_float_field` FLOAT" +
			", PRIMARY KEY(`id`)" +
			")",
		false,
	},
}

func TestGo(t *testing.T) {
	for i, tt := range testGoTables {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			out, err := Go(tt.in, false)
			if tt.err {
				assert.NotNil(t, err)
				assert.Empty(t, out)
			} else {
				assert.Nil(t, err)
				assert.NotEmpty(t, out)
			}
		})
	}
}

func TestGoNulls(t *testing.T) {
	for i, tt := range testGoTables {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			out, err := Go(tt.in, true)
			if tt.err {
				assert.NotNil(t, err)
				assert.Empty(t, out)
			} else {
				assert.Nil(t, err)
				assert.NotEmpty(t, out)
			}
		})
	}
}
