package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/AngeloMihaelle/CodeStash/internal/snippet"
	"github.com/AngeloMihaelle/CodeStash/internal/store"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print [snippet-id-or-title]",
	Short: "Print a snippet to the terminal",
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

		// Update usage stats
		updateUsageStats(targetSnippet)

		// Save updated stats
		if err := store.SaveSnippets(snippets); err != nil {
			fmt.Println("⚠️  Failed to update usage stats:", err)
		}

		// Print the snippet
		fmt.Printf("📄 %s\n", targetSnippet.Title)
		fmt.Printf("📝 %s\n", targetSnippet.Description)
		fmt.Println("─────────────────────────────────────")
		fmt.Println(targetSnippet.Code)
		fmt.Println("─────────────────────────────────────")
	},
}

func findSnippet(snippets []snippet.Snippet, query string) *snippet.Snippet {
	for i, s := range snippets {
		if s.ID == query || strings.EqualFold(s.Title, query) {
			return &snippets[i]
		}
	}
	return nil
}

func updateUsageStats(s *snippet.Snippet) {
	s.UsageCount++
	s.LastUsed = time.Now().UTC().Format(time.RFC3339)
}
