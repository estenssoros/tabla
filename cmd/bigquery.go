package cmd

import "github.com/spf13/cobra"

var bigqueryCmd = &cobra.Command{
	Use:     "bigquery",
	Short:   "converts a go struct to a create table statement",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return nil },
}
