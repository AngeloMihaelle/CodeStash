package cmd

import (
	"fmt"
	"strings"

	"github.com/AngeloMihaelle/CodeStash/internal/snippet"
	"github.com/AngeloMihaelle/CodeStash/internal/store"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search snippets by title, description, tags, or content",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := store.LoadSnippets()
		if err != nil {
			fmt.Println("❌ Failed to load snippets:", err)
			return
		}

		query := strings.ToLower(args[0])
		var matches []snippet.Snippet

		// Get filter flags
		executable, _ := cmd.Flags().GetBool("executable")

		for _, s := range snippets {
			if matchesQuery(s, query) {
				// Apply executable filter if specified
				if executable && !s.Executable {
					continue
				}
				matches = append(matches, s)
			}
		}

		if len(matches) == 0 {
			fmt.Printf("🔍 No snippets found matching '%s'\n", args[0])
			return
		}

		// Get expanded flag
		expanded, _ := cmd.Flags().GetBool("expanded")

		fmt.Printf("🔍 Found %d snippet(s) matching '%s':\n\n", len(matches), args[0])

		for _, s := range matches {
			fmt.Printf("🔹 ID: %s\n", s.ID)
			fmt.Printf("   Title: %s\n", s.Title)
			fmt.Printf("   Language: %s\n", s.Language)
			fmt.Printf("   Tags: %s\n", strings.Join(s.Tags, ", "))
			fmt.Printf("   Description: %s\n", s.Description)
			fmt.Printf("   Used: %d times\n", s.UsageCount)

			// Show executable status
			if s.Executable {
				fmt.Println("   🚀 Executable: Yes")
			} else {
				fmt.Println("   📄 Executable: No")
			}

			if expanded {
				fmt.Println("   Code:")
				fmt.Println("   ─────────────────────────────────────")
				// Indent each line of code
				codeLines := strings.Split(s.Code, "\n")
				for _, line := range codeLines {
					fmt.Printf("   %s\n", line)
				}
				fmt.Println("   ─────────────────────────────────────")
			} else {
				// Show code preview if it matches (only when not expanded)
				if strings.Contains(strings.ToLower(s.Code), query) {
					preview := getCodePreview(s.Code, query)
					fmt.Printf("   Preview: %s\n", preview)
				}
			}
			fmt.Println()
		}
	},
}

func init() {
	searchCmd.Flags().BoolP("expanded", "e", false, "Show full code content for each snippet")
	searchCmd.Flags().BoolP("executable", "x", false, "Show only executable snippets")
}

func matchesQuery(s snippet.Snippet, query string) bool {
	// Check ID
	if strings.Contains(strings.ToLower(s.ID), query) {
		return true
	}

	// Check title
	if strings.Contains(strings.ToLower(s.Title), query) {
		return true
	}

	// Check description
	if strings.Contains(strings.ToLower(s.Description), query) {
		return true
	}

	// Check tags
	for _, tag := range s.Tags {
		if strings.Contains(strings.ToLower(tag), query) {
			return true
		}
	}

	// Check code content
	if strings.Contains(strings.ToLower(s.Code), query) {
		return true
	}

	// Check language
	if strings.Contains(strings.ToLower(s.Language), query) {
		return true
	}

	return false
}

func getCodePreview(code, query string) string {
	lines := strings.Split(code, "\n")
	query = strings.ToLower(query)

	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), query) {
			trimmed := strings.TrimSpace(line)
			if len(trimmed) > 60 {
				return trimmed[:60] + "..."
			}
			return trimmed
		}
	}
	// If no specific line matches, return first few chars
	if len(code) > 60 {
		return strings.TrimSpace(code[:60]) + "..."
	}
	return strings.TrimSpace(code)
}
