package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/estenssoros/tabla/internal/mssql"
	"github.com/estenssoros/tabla/internal/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var mySQLNulls bool

func init() {
	sqlCmd.AddCommand(
		sqlMySqlCmd,
		sqlMsSqlCmd,
	)
}

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "start with sql",
}

func init() {
	sqlMySqlCmd.AddCommand(
		sqlMySqlClipboardCmd,
	)
	sqlMySqlCmd.PersistentFlags().BoolVarP(&mySQLNulls, "nulls", "", false, "create go struct with nulls")
}

var sqlMySqlCmd = &cobra.Command{
	Use:     "mysql",
	Short:   "mysql to go",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return nil },
}

var sqlMySqlClipboardCmd = &cobra.Command{
	Use:     "clipboard",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
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
var msSQLNulls bool

func init() {
	sqlMsSqlCmd.AddCommand(
		sqlMsSqlClipboardCmd,
	)
	sqlMsSqlCmd.PersistentFlags().BoolVarP(&msSQLNulls, "nulls", "", false, "create go struct with nulls")
}

var sqlMsSqlCmd = &cobra.Command{
	Use:   "mssql",
	Short: "mssql to go",
}

var sqlMsSqlClipboardCmd = &cobra.Command{
	Use:     "clipboard",
	Short:   "reads a statement from the clipboard and returns a go struct",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
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
