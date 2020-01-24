# tabla

convert a SQL create statement to a go struct and visa-versa

## go to SQL

```go
type Test struct {
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
}
```

copy the struct to you clipboard

```bash
tabla go mysql
```
 
output:

```SQL
DROP TABLE IF EXISTS `test`;
CREATE TABLE `test` (
    `id` VARCHAR(36)
    , `int_field` INT
    , `string_field` VARCHAR({update})
    , `time_field` DATETIME
    , `bool_field` TINYINT(1)
    , `float_field` FLOAT
    , `nulls_int_field` INT
    , `nulls_string_field` VARCHAR({update})
    , `nulls_time_field` DATETIME
    , `nulls_bool_field` TINYINT(1)
    , `nulls_float_field` FLOAT
    , PRIMARY KEY(`id`)
);
```

note:  table creates `{update}` areas where the user is supposed to input the length of a datatype

## SQL to Go

```sql
CREATE TABLE `test` (
    `id` VARCHAR(36)
    , `int_field` INT
    , `string_field` VARCHAR(23)
    , `time_field` DATETIME
    , `bool_field` TINYINT(1)
    , `float_field` FLOAT
    , `nulls_int_field` INT
    , `nulls_string_field` VARCHAR(48)
    , `nulls_time_field` DATETIME
    , `nulls_bool_field` TINYINT(1)
    , `nulls_float_field` FLOAT
    , PRIMARY KEY(`id`)
);
```

copy the create statement to you clipboard

```bash
tabla mysql go
```
 
output:
 
```go
type Test struct {
	ID               uuid.UUID `json:"id" db:"id"`
	IntField         int       `json:"int_field" db:"int_field"`
	StringField      string    `json:"string_field" db:"string_field"`
	TimeField        time.Time `json:"time_field" db:"time_field"`
	BoolField        bool      `json:"bool_field" db:"bool_field"`
	FloatField       float64   `json:"float_field" db:"float_field"`
	NullsIntField    int       `json:"nulls_int_field" db:"nulls_int_field"`
	NullsStringField string    `json:"nulls_string_field" db:"nulls_string_field"`
	NullsTimeField   time.Time `json:"nulls_time_field" db:"nulls_time_field"`
	NullsBoolField   bool      `json:"nulls_bool_field" db:"nulls_bool_field"`
	NullsFloatField  float64   `json:"nulls_float_field" db:"nulls_float_field"`
}
```