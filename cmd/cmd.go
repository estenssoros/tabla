package cmd

import (
	"github.com/atotto/clipboard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var copy bool

func init() {
	Cmd.AddCommand(
		goCmd,
		sqlCmd,
	)
	Cmd.PersistentFlags().BoolVarP(&copy, "copy", "c", false, "send output to clipboard")
}

var Cmd = &cobra.Command{
	Use:   "tabla",
	Short: "pbridge SQL and go",
}

// Execute main entry point for command
func Execute() error {
	return Cmd.Execute()
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
