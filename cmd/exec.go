package cmd

import (
	"fmt"

	"github.com/AngeloMihaelle/CodeStash/internal/store"
	"github.com/spf13/cobra"
)

var forceExec bool

var execCmd = &cobra.Command{
	Use:   "exec [snippet-id-or-title]",
	Short: "Execute a snippet (executable snippets only)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := store.LoadSnippets()
		if err != nil {
			fmt.Println("❌ Failed to load snippets:", err)
			return
		}

		// Find snippet by ID or title
		targetSnippet := findSnippet(snippets, args[0])
		if targetSnippet == nil {
			fmt.Printf("❌ Snippet '%s' not found\n", args[0])
			return
		}

		// Check if snippet is marked as executable (unless --force is used)
		if !targetSnippet.Executable && !forceExec {
			fmt.Printf("❌ Snippet '%s' is not marked as executable\n", targetSnippet.Title)
			fmt.Println("💡 Use 'codestash edit' to mark it as executable, or use 'codestash exec --force' to force execution")
			return
		}

		// Show warning if forcing execution of non-executable snippet
		if !targetSnippet.Executable && forceExec {
			fmt.Printf("⚠️  Forcing execution of non-executable snippet '%s'\n", targetSnippet.Title)
		}

		// Update usage stats
		updateUsageStats(targetSnippet)

		// Save updated stats
		if err := store.SaveSnippets(snippets); err != nil {
			fmt.Println("⚠️  Failed to update usage stats:", err)
		}

		// Execute the snippet
		if err := executeSnippet(targetSnippet); err != nil {
			fmt.Println("❌ Failed to execute snippet:", err)
			return
		}
	},
}

func init() {
	execCmd.Flags().BoolVarP(&forceExec, "force", "f", false, "Force execution even if snippet is not marked as executable")
}
