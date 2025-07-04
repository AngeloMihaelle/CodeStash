package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codestash",
	Short: "🧰 CodeStash - Your local code snippet manager",
	Long:  "CodeStash is a local-first CLI tool to manage and execute code snippets efficiently.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(useCmd)
	rootCmd.AddCommand(deleteCmd)
	// 	rootCmd.AddCommand(tagCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(copyCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(statsCmd)
}
