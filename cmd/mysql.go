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

var mysqlDatabaseCmd = &cobra.Command{
	Use:     "database",
	Short:   "converts database into structs",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return nil },
}
