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
	mysqlCmd.Flags().BoolVarP(&nulls, "nulls", "", false, "create go struct with nulls")
}

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "converts a mysql create statement to a go struct",
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
