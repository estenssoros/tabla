package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/estenssoros/tabla/internal/gopher"
	"github.com/estenssoros/tabla/internal/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var mySQLNulls bool

func init() {
	mysqlCmd.PersistentFlags().BoolVarP(&mySQLNulls, "nulls", "", false, "create go struct with nulls")
	mysqlCmd.AddCommand(
		mysqlStatementCmd,
		mysqlDatabaseCmd,
		mysqlStructCmd,
	)
}

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "converts a mysql create statement to a go struct",
}

var mysqlStatementCmd = &cobra.Command{
	Use:   "statement",
	Short: "reads a statement from the clipboard and returns a go struct",
	RunE: func(cmd *cobra.Command, args []string) error {
		stmt, err := readClipboard()
		if err != nil {
			return errors.Wrap(err, "read clipboard")
		}
		out, err := mysql.Go(stmt, mySQLNulls)
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
	mysqlHost, mysqlUser, mysqlPassword, mysqlPort, mysqlName string
)

func init() {
	mysqlDatabaseCmd.Flags().StringVarP(&mysqlHost, "host", "", "", "database host")
	mysqlDatabaseCmd.Flags().StringVarP(&mysqlUser, "user", "u", "", "database user")
	mysqlDatabaseCmd.Flags().StringVarP(&mysqlPassword, "password", "p", "", "database password")
	mysqlDatabaseCmd.Flags().StringVarP(&mysqlPort, "port", "", "3306", "database port")
	mysqlDatabaseCmd.Flags().StringVarP(&mysqlName, "name", "n", "", "database name")
}

var mysqlDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "converts database into structs",
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
		if copy {
			clipboard.WriteAll(out)
		}
		return nil
	},
}

var mysqlStructCmd = &cobra.Command{
	Use:     "struct",
	Short:   "converts a mysql statement to a go struct",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		stmt, err := clipboard.ReadAll()
		if err != nil {
			return errors.Wrap(err, "clipboard readall")
		}
		if stmt == "" {
			return errors.New("no stmt in clipboard")
		}
		out, err := gopher.DropCreate(stmt, mysql.Dialect{})
		if err != nil {
			return errors.Wrap(err, "gopher mysql")
		}
		if copy {
			return errors.Wrap(clipboard.WriteAll(out), "clipboard write")
		}
		fmt.Println(out)
		return nil
	},
}
