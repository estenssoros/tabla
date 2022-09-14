package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/estenssoros/tabla/internal/mssql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var msSQLNulls bool

func init() {
	mssqlCmd.PersistentFlags().BoolVarP(&msSQLNulls, "nulls", "", false, "create go struct with nulls")
	mssqlCmd.AddCommand(mssqlStatementCmd)
	mssqlCmd.AddCommand(mssqlDatabaseCmd)
}

var mssqlCmd = &cobra.Command{
	Use:   "mssql",
	Short: "converts mssql create statement to go struct",
}

var mssqlStatementCmd = &cobra.Command{
	Use:   "statement",
	Short: "reads a statement from the clipboard and returns a go struct",
	RunE: func(cmd *cobra.Command, args []string) error {
		stmt, err := readClipboard()
		if err != nil {
			return errors.Wrap(err, "read clipboard")
		}
		out, err := mssql.Go(stmt, msSQLNulls)
		if err != nil {
			return errors.Wrap(err, "mssql to go")
		}
		if copy {
			return errors.Wrap(clipboard.WriteAll(out), "clipboard write")
		}
		fmt.Println(out)
		return nil
	},
}

var (
	mssqlHost, mssqlUser, mssqlPassword, mssqlPort, mssqlName string
)

func init() {
	mssqlDatabaseCmd.Flags().StringVarP(&mssqlHost, "host", "", "", "database host")
	mssqlDatabaseCmd.Flags().StringVarP(&mssqlUser, "user", "u", "", "database user")
	mssqlDatabaseCmd.Flags().StringVarP(&mssqlPassword, "password", "p", "", "database password")
	mssqlDatabaseCmd.Flags().StringVarP(&mssqlPort, "port", "", "1433", "database port")
	mssqlDatabaseCmd.Flags().StringVarP(&mssqlName, "name", "n", "", "database name")
}

var mssqlDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "converts database into structs",
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
