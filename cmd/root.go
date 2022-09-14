package cmd

import (
	"github.com/atotto/clipboard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var copy bool

func init() {
	rootCmd.AddCommand(
		goCmd,
		mysqlCmd,
		mssqlCmd,
		bigqueryCmd,
	)
	rootCmd.PersistentFlags().BoolVarP(&copy, "copy", "c", false, "send output to clipboard")
}

var rootCmd = &cobra.Command{
	Use:   "tabla",
	Short: "parses SQL and Go to bridge the languages",
}

// Execute main entry point for command
func Execute() error {
	return rootCmd.Execute()
}

func readClipboard() (string, error) {
	stmt, err := clipboard.ReadAll()
	if err != nil {
		return "", errors.Wrap(err, "clipboard readall")
	}
	if stmt == "" {
		return "", errors.New("no stmt in clipboard")
	}
	return stmt, nil
}
