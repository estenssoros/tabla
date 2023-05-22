package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/estenssoros/tabla/internal/gopher"
	"github.com/estenssoros/tabla/internal/mssql"
	"github.com/estenssoros/tabla/internal/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	goCmd.AddCommand(
		goErdCmd,
		goMySQLCmd,
		goMsSqlCmd,
		goPostgresQLCmd,
		goBigqueryCmd,
	)
}

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "start with go",
}

var goErdCmd = &cobra.Command{
	Use:     "erd",
	Short:   "create an ERD for a collection of go structs",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return nil },
}
var (
	mysqlHost, mysqlUser, mysqlPassword, mysqlPort, mysqlName string
	mssqlHost, mssqlUser, mssqlPassword, mssqlPort, mssqlName string
)

func init() {
	goMySQLCmd.AddCommand(
		goMySQLClipboardCmd,
		goMySQLDatabaseCmd,
	)
	goMySQLDatabaseCmd.Flags().StringVarP(&mysqlHost, "host", "", "", "database host")
	goMySQLDatabaseCmd.Flags().StringVarP(&mysqlUser, "user", "u", "", "database user")
	goMySQLDatabaseCmd.Flags().StringVarP(&mysqlPassword, "password", "p", "", "database password")
	goMySQLDatabaseCmd.Flags().StringVarP(&mysqlPort, "port", "", "3306", "database port")
	goMySQLDatabaseCmd.Flags().StringVarP(&mysqlName, "name", "n", "", "database name")
}

var goMySQLCmd = &cobra.Command{
	Use:   "mysql",
	Short: "convert go struct to mysql create statement",
}

var goMySQLClipboardCmd = &cobra.Command{
	Use:     "clipboard",
	Short:   "reads from clipboard",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		stmt, err := clipboard.ReadAll()
		if err != nil {
			return errors.Wrap(err, "clipboard.ReadAll")
		}
		if stmt == "" {
			return errors.New("no stmt in clipboard")
		}
		out, err := gopher.DropCreate(stmt, mysql.Dialect{})
		if err != nil {
			return errors.Wrap(err, "gopher mysql")
		}
		fmt.Println(out)
		return nil
	},
}

var goMySQLDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "reads from database",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if mysqlHost == "" {
			return errors.New("host cannot be blank")
		}
		if mysqlUser == "" {
			return errors.New("user cannot be blank")
		}
		if mysqlPassword == "" {
			return errors.New("password cannot be blank")
		}
		if mysqlName == "" {
			return errors.New("name cannot be blank")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		config := &mysql.ConnectConfig{
			Host:     mysqlHost,
			User:     mysqlUser,
			Password: mysqlPassword,
			Port:     mysqlPort,
			Name:     mysqlName,
		}
		tables, err := mysql.ShowCreateTables(config)
		if err != nil {
			return errors.Wrap(err, "show create tables")
		}
		var out = fmt.Sprintf("package %s\n", mysqlName)
		fmt.Println(out)
		for _, table := range tables {
			goStructString, err := mysql.Go(table, mySQLNulls)
			if err != nil {
				return errors.Wrap(err, "mysql go")
			}
			fmt.Println(goStructString)
			out += "\n" + goStructString
		}
		return nil
	},
}

func init() {
	goMsSqlCmd.AddCommand(
		goMsSqlClipboardCmd,
		goMsSqlDatabaseCmd,
	)
}

var goMsSqlCmd = &cobra.Command{
	Use:   "mssql",
	Short: "convert go struct to mssql create statement",
}

var goMsSqlClipboardCmd = &cobra.Command{
	Use:     "clipboard",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("not implemented")
	},
}

func init() {
	goMsSqlDatabaseCmd.Flags().StringVarP(&mssqlHost, "host", "", "", "database host")
	goMsSqlDatabaseCmd.Flags().StringVarP(&mssqlUser, "user", "u", "", "database user")
	goMsSqlDatabaseCmd.Flags().StringVarP(&mssqlPassword, "password", "p", "", "database password")
	goMsSqlDatabaseCmd.Flags().StringVarP(&mssqlPort, "port", "", "1433", "database port")
	goMsSqlDatabaseCmd.Flags().StringVarP(&mssqlName, "name", "n", "", "database name")
}

var goMsSqlDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "database to structs",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if mssqlHost == "" {
			return errors.New("host cannot be blank")
		}
		if mssqlUser == "" {
			return errors.New("user cannot be blank")
		}
		if mssqlPassword == "" {
			return errors.New("password cannot be blank")
		}
		if mssqlName == "" {
			return errors.New("name cannot be blank")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		config := &mssql.ConnectConfig{
			Host:     mssqlHost,
			User:     mssqlUser,
			Password: mssqlPassword,
			Port:     mssqlPort,
			Name:     mssqlName,
		}
		tables, err := mssql.ShowTables(config)
		if err != nil {
			return errors.Wrap(err, "mssql show tables")
		}
		var out = fmt.Sprintf("package %s\n", mssqlName)
		fmt.Println(out)
		for _, table := range tables {
			goStructString, err := table.ToGo(msSQLNulls)
			if err != nil {
				return errors.Wrap(err, "mysql go")
			}
			fmt.Println(goStructString)
			out += "\n" + goStructString
		}
		if copy {
			clipboard.WriteAll(out)
		}
		return nil
	},
}

var goPostgresQLCmd = &cobra.Command{
	Use:     "postgresql",
	Short:   "convert go to postgresql create statement",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return nil },
}

var goBigqueryCmd = &cobra.Command{
	Use:     "bigquery",
	Short:   "convert go to bigquery create statement",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return nil },
}
