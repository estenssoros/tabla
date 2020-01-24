package mysql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGo(t *testing.T) {
	stmt := "CREATE TABLE `test` (" +
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
		", PRIMARY KEY(`id`)" +
		")"
	out, err := Go(stmt, false)
	assert.Nil(t, err)
	fmt.Println(out)
}
