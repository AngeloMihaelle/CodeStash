package cmd

import (
	"fmt"
	"strings"

	"github.com/AngeloMihaelle/CodeStash/internal/store"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [snippet-id-or-title]",
	Short: "Delete a snippet",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := store.LoadSnippets()
		if err != nil {
			fmt.Println("❌ Failed to load snippets:", err)
			return
		}

		// Find snippet by ID or title
		query := args[0]
		targetIndex := -1
		var targetSnippet *string

		for i, s := range snippets {
			if s.ID == query || strings.EqualFold(s.Title, query) {
				targetIndex = i
				targetSnippet = &s.Title
				break
			}
		}

		if targetIndex == -1 {
			fmt.Printf("❌ Snippet '%s' not found\n", query)
			return
		}

		// Get confirmation flag
		force, _ := cmd.Flags().GetBool("force")

		// Ask for confirmation unless --force is used
		if !force {
			fmt.Printf("⚠️  Are you sure you want to delete '%s'? [y/N]: ", *targetSnippet)
			var response string
			fmt.Scanln(&response)

			if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
				fmt.Println("❌ Deletion cancelled")
				return
			}
		}

		// Remove snippet from slice
		snippets = append(snippets[:targetIndex], snippets[targetIndex+1:]...)

		// Save updated snippets
		if err := store.SaveSnippets(snippets); err != nil {
			fmt.Println("❌ Failed to save snippets:", err)
			return
		}

		fmt.Printf("✅ Deleted snippet '%s'\n", *targetSnippet)
	},
}

func init() {
	deleteCmd.Flags().BoolP("force", "f", false, "Delete without confirmation")
}
