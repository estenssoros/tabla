package cmd

import "github.com/spf13/cobra"

var copy bool

func init() {
	rootCmd.AddCommand(goCmd)
	rootCmd.AddCommand(mysqlCmd)
	rootCmd.PersistentFlags().BoolVarP(&copy, "copy", "c", false, "send output to clipboard")
}

var rootCmd = &cobra.Command{
	Use:   "tabla",
	Short: "",
}

func Execute() error {
	return rootCmd.Execute()
}
