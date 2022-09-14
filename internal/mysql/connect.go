package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
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
	connectionURL := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Name)
	db, err := sql.Open("mysql", connectionURL)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(0)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func showTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, errors.Wrap(err, "db query")
	}
	tables := []string{}
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, errors.Wrap(err, "rows scan")
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func showCreateTable(db *sql.DB, table string) (string, error) {
	row := db.QueryRow("show create table " + table)
	var name string
	var create string
	if err := row.Scan(&name, &create); err != nil {
		return "", errors.Wrap(err, "row scan")
	}
	return create, nil
}

// ShowCreateTables returns the table definitions for a database
func ShowCreateTables(config *ConnectConfig) ([]string, error) {
	db, err := connect(config)
	if err != nil {
		return nil, errors.Wrap(err, "mysql connect")
	}
	defer db.Close()
	tables, err := showTables(db)
	if err != nil {
		return nil, errors.Wrap(err, "show tables")
	}
	creates := []string{}
	for _, t := range tables {
		create, err := showCreateTable(db, t)
		if err != nil {
			return nil, errors.Wrapf(err, "show create table %s", t)
		}
		creates = append(creates, create)
	}
	return creates, nil
}
