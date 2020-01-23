package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(goCmd)
	rootCmd.AddCommand(mysqlCmd)
}

var rootCmd = &cobra.Command{
	Use:   "tabla",
	Short: "",
}

func Execute() error {
	return rootCmd.Execute()
}
