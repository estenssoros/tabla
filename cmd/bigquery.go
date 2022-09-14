package cmd

import (
	"github.com/spf13/cobra"
)

var bigQueryNulls bool

func init() {
	bigqueryCmd.PersistentFlags().BoolVarP(&bigQueryNulls, "nulls", "", false, "create go struct with nulls")
	bigqueryCmd.AddCommand(
		bigqueryStatementCmd,
	)
}

var bigqueryCmd = &cobra.Command{
	Use:   "bigquery",
	Short: "converts a go struct to a create table statement",
}

var bigqueryStatementCmd = &cobra.Command{
	Use:     "statement",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
