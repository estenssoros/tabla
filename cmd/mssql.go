package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/estenssoros/tabla/mssql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	mssqlCmd.AddCommand(mssqlStatementCmd)
}

var mssqlCmd = &cobra.Command{
	Use:   "mssql",
	Short: "converts mssql create statement to go struct",
}

var mssqlStatementCmd = &cobra.Command{
	Use:   "statement",
	Short: "reads a statement from the clipboard and returns a go stuct",
	RunE: func(cmd *cobra.Command, args []string) error {
		stmt, err := readClipboad()
		if err != nil {
			return errors.Wrap(err, "read clipboard")
		}
		out, err := mssql.Go(stmt, nulls)
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
