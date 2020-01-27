# tabla

convert a SQL create statement to a go struct and visa-versa

## suported dialects

- mysql 
- mssql

## go to SQL

```go
type Test struct {
	id               uuid.UUID
	intField         int `db:"tag_name"`
	stringField      string
	timeField        time.Time
	boolField        bool
	floatField       float64
	nullsIntField    nulls.Int
	nullsStringField nulls.String
	nullsTimeField   nulls.Time
	nullsBoolField   nulls.Bool
	nullsFloatField  nulls.Float64
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
    , `tag_name` INT
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

note:  

- table creates `{update}` areas where the user is supposed to input the length of a datatype
- struct tags with the key `db` replace field names in the generated SQL

## SQL to Go

```sql
CREATE TABLE `test` (
    `id` VARCHAR(36)
    , `int_field` INT
    , `string_field` VARCHAR(23)
    , `time_field` DATETIME
    , `bool_field` TINYINT(1)
    , `float_field` FLOAT
    , PRIMARY KEY(`id`)
);
```

copy the create statement to you clipboard

```bash
tabla mysql statement
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
}
```

```bash
tabla mysql database --host 127.0.0.1 -u root -p my_password -n my_datatbase
```

```go
package dev

// AuthGroup
type AuthGroup struct {
	ID   int    `json:"id" db:"id,int,11"`
	Name string `json:"name" db:"name,varchar,80"`
}

// AuthGroupPermissions
type AuthGroupPermissions struct {
	ID           int `json:"id" db:"id,int,11"`
	GroupId      int `json:"group_id" db:"group_id,int,11"`
	PermissionId int `json:"permission_id" db:"permission_id,int,11"`
}

// AuthPermission
type AuthPermission struct {
	ID            int    `json:"id" db:"id,int,11"`
	Name          string `json:"name" db:"name,varchar,255"`
	ContentTypeId int    `json:"content_type_id" db:"content_type_id,int,11"`
	Codename      string `json:"codename" db:"codename,varchar,100"`
}

// AuthUser
type AuthUser struct {
	ID          int       `json:"id" db:"id,int,11"`
	Password    string    `json:"password" db:"password,varchar,128"`
	LastLogin   time.Time `json:"last_login" db:"last_login,datetime,6"`
	IsSuperuser bool      `json:"is_superuser" db:"is_superuser,tinyint,1"`
	Username    string    `json:"username" db:"username,varchar,150"`
	FirstName   string    `json:"first_name" db:"first_name,varchar,30"`
	LastName    string    `json:"last_name" db:"last_name,varchar,150"`
	Email       string    `json:"email" db:"email,varchar,254"`
	IsStaff     bool      `json:"is_staff" db:"is_staff,tinyint,1"`
	IsActive    bool      `json:"is_active" db:"is_active,tinyint,1"`
	DateJoined  time.Time `json:"date_joined" db:"date_joined,datetime,6"`
}

// AuthUserGroups
type AuthUserGroups struct {
	ID      int `json:"id" db:"id,int,11"`
	UserId  int `json:"user_id" db:"user_id,int,11"`
	GroupId int `json:"group_id" db:"group_id,int,11"`
}

// AuthUserUserPermissions
type AuthUserUserPermissions struct {
	ID           int `json:"id" db:"id,int,11"`
	UserId       int `json:"user_id" db:"user_id,int,11"`
	PermissionId int `json:"permission_id" db:"permission_id,int,11"`
}
```