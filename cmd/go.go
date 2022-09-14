package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/estenssoros/tabla/internal/gopher"
	"github.com/estenssoros/tabla/internal/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	goCmd.AddCommand(goMySQLCmd)
}

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "converts a go struct to a create table statement",
}

var goMySQLCmd = &cobra.Command{
	Use:     "mysql",
	Short:   "create statement in mysql format",
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
