package cmd

import (
	"fmt"
	"strings"

	"github.com/AngeloMihaelle/CodeStash/internal/snippet"
	"github.com/AngeloMihaelle/CodeStash/internal/store"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snippets",
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := store.LoadSnippets()
		if err != nil {
			fmt.Println("❌ Failed to load snippets:", err)
			return
		}

		if len(snippets) == 0 {
			fmt.Println("📭 No snippets found. Use 'codestash add' to create your first snippet!")
			return
		}

		// Get filter flags
		language, _ := cmd.Flags().GetString("language")
		tag, _ := cmd.Flags().GetString("tag")

		// Filter snippets
		var filteredSnippets []snippet.Snippet
		for _, s := range snippets {
			if language != "" && s.Language != language {
				continue
			}
			if tag != "" && !contains(s.Tags, tag) {
				continue
			}
			filteredSnippets = append(filteredSnippets, s)
		}

		if len(filteredSnippets) == 0 {
			fmt.Println("📭 No snippets match your filters.")
			return
		}

		// Get expanded flag
		expanded, _ := cmd.Flags().GetBool("expanded")

		fmt.Printf("📚 Found %d snippet(s):\n\n", len(filteredSnippets))

		for _, s := range filteredSnippets {
			fmt.Printf("🔹 ID: %s\n", s.ID)
			fmt.Printf("   Title: %s\n", s.Title)
			fmt.Printf("   Language: %s\n", s.Language)
			fmt.Printf("   Tags: %s\n", strings.Join(s.Tags, ", "))
			fmt.Printf("   Description: %s\n", s.Description)
			fmt.Printf("   Executable: %t\n", s.Executable)
			fmt.Printf("   Used: %d times\n", s.UsageCount)

			if expanded {
				fmt.Println("   Code:")
				fmt.Println("   ─────────────────────────────────────")
				// Indent each line of code
				codeLines := strings.Split(s.Code, "\n")
				for _, line := range codeLines {
					fmt.Printf("   %s\n", line)
				}
				fmt.Println("   ─────────────────────────────────────")
				fmt.Println("   last used:", s.LastUsed)
				fmt.Println("   created at:", s.CreatedAt)

			}
			fmt.Println()
		}
	},
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func init() {
	listCmd.Flags().StringP("language", "l", "", "Filter by language")
	listCmd.Flags().StringP("tag", "t", "", "Filter by tag")
	listCmd.Flags().BoolP("expanded", "e", false, "Show code content for each snippet")
}
