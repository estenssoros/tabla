package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/estenssoros/tabla/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var nulls bool

func init() {
	mysqlCmd.PersistentFlags().BoolVarP(&nulls, "nulls", "", false, "create go struct with nulls")
	mysqlCmd.AddCommand(mysqlStatementCmd)
	mysqlCmd.AddCommand(mysqlDatabaseCmd)
}

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "converts a mysql create statement to a go struct",
}

var mysqlStatementCmd = &cobra.Command{
	Use:     "statement",
	Short:   "reads the clipboard for a statement and returns a go struct",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		stmt, err := clipboard.ReadAll()
		if err != nil {
			return errors.Wrap(err, "clipboard readall")
		}
		if stmt == "" {
			return errors.New("no stmt in clipboard")
		}
		out, err := mysql.Go(stmt, nulls)
		if err != nil {
			return errors.Wrap(err, "mySQLToGo")
		}
		if copy {
			return errors.Wrap(clipboard.WriteAll(out), "clipboard write")
		}
		fmt.Println(out)
		return nil
	},
}

var (
	host, user, password, port, name string
)

func init() {
	mysqlDatabaseCmd.Flags().StringVarP(&host, "host", "", "", "database host")
	mysqlDatabaseCmd.Flags().StringVarP(&user, "user", "u", "", "database user")
	mysqlDatabaseCmd.Flags().StringVarP(&password, "password", "p", "", "database password")
	mysqlDatabaseCmd.Flags().StringVarP(&port, "port", "", "3306", "database port")
	mysqlDatabaseCmd.Flags().StringVarP(&name, "name", "n", "", "database name")
}

var mysqlDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "converts database into structs",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if host == "" {
			return errors.New("host cannot be blank")
		}
		if user == "" {
			return errors.New("user cannot be blank")
		}
		if password == "" {
			return errors.New("password cannot be blank")
		}
		if name == "" {
			return errors.New("name cannot be blank")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		config := &mysql.ConnectConfig{
			Host:     host,
			User:     user,
			Password: password,
			Port:     port,
			Name:     name,
		}
		tables, err := mysql.ShowCreateTables(config)
		if err != nil {
			return errors.Wrap(err, "show create tables")
		}
		var out = fmt.Sprintf("package %s\n", name)
		fmt.Println(out)
		for _, table := range tables {
			goStructString, err := mysql.Go(table, nulls)
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
