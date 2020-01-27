package mssql

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" //mssql
	"github.com/estenssoros/tabla/gopher"
	"github.com/pkg/errors"
)

// ConnectConfig config for database parameters
type ConnectConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
}

func connect(config *ConnectConfig) (*sql.DB, error) {
	connectionURL := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", config.User, config.Password, config.Host, config.Name)
	db, err := sql.Open("mssql", connectionURL)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(0)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func showTables(db *sql.DB) ([]*Table, error) {
	rows, err := db.Query("select name,object_id from sys.tables")
	if err != nil {
		return nil, errors.Wrap(err, "db query")
	}
	tables := []*Table{}
	for rows.Next() {
		var name string
		var objectID string
		if err := rows.Scan(&name, &objectID); err != nil {
			return nil, errors.Wrap(err, "rows scan")
		}
		tables = append(tables, &Table{Name: name, ObjectID: objectID})
	}
	return tables, nil
}

// Column sql server column information
type Column struct {
	Name   string
	Type   SQLType
	Length string
}

// Table sql server table information
type Table struct {
	Name     string
	ObjectID string
	Columns  []*Column
}

func (t Table) String() string {
	ju, _ := json.MarshalIndent(t, "", " ")
	return string(ju)
}

// GetColumns select columns from sql server
func (t *Table) GetColumns(db *sql.DB) error {
	q := `
	SELECT col.name
		, typ.name
		, typ.max_length
	FROM sys.columns col
	JOIN sys.types typ ON typ.user_type_id = col.user_type_id
	WHERE col.object_id = %s
	`
	rows, err := db.Query(fmt.Sprintf(q, t.ObjectID))
	if err != nil {
		return errors.Wrap(err, "query")
	}
	defer rows.Close()
	columns := []*Column{}
	for rows.Next() {
		col := &Column{}
		if err := rows.Scan(&col.Name, &col.Type, &col.Length); err != nil {
			return errors.Wrap(err, "rows scan")
		}
		columns = append(columns, col)
	}
	t.Columns = columns
	return nil
}

// ToGo converts a table definitions to go
func (t *Table) ToGo(nulls bool) (string, error) {
	fields := make([]*gopher.GoField, len(t.Columns))
	for i, col := range t.Columns {
		field, err := col.Type.ToGoField(nulls, col)
		if err != nil {
			return "", errors.Wrap(err, "to go field")
		}
		fields[i] = field
	}
	goStruct := &gopher.GoStruct{
		Name:   t.Name,
		Fields: fields,
	}
	return goStruct.ToGo()
}

// ShowTables wrapper for constructing tabel definitions
func ShowTables(config *ConnectConfig) ([]*Table, error) {
	db, err := connect(config)
	if err != nil {
		return nil, errors.Wrap(err, "mysql connect")
	}
	defer db.Close()
	tables, err := showTables(db)
	if err != nil {
		return nil, errors.Wrap(err, "show tables")
	}
	for _, table := range tables {
		if err := table.GetColumns(db); err != nil {
			return nil, errors.Wrap(err, "table get columns")
		}
	}
	return tables, nil
}
