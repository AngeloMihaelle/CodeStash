package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AngeloMihaelle/CodeStash/internal/snippet"
	"github.com/AngeloMihaelle/CodeStash/internal/store"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new snippet",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("📝 Title: ")
		title, _ := reader.ReadString('\n')

		fmt.Print("🧾 Description: ")
		description, _ := reader.ReadString('\n')

		fmt.Print("💻 Language: ")
		language, _ := reader.ReadString('\n')

		fmt.Print("🏷️ Tags (comma separated): ")
		tagsRaw, _ := reader.ReadString('\n')

		// Ask if the snippet is executable
		fmt.Print("🚀 Is this snippet executable? (y/N): ")
		executableRaw, _ := reader.ReadString('\n')
		executable := strings.ToLower(strings.TrimSpace(executableRaw)) == "y" || strings.ToLower(strings.TrimSpace(executableRaw)) == "yes"

		fmt.Println("📋 Enter code (end with 'EOF' on a new line):")
		var lines []string
		for {
			line, _ := reader.ReadString('\n')
			if strings.TrimSpace(line) == "EOF" {
				break
			}
			lines = append(lines, line)
		}

		s := snippet.NewSnippet(
			strings.TrimSpace(title),
			strings.Join(lines, ""),
			strings.TrimSpace(description),
			strings.TrimSpace(language),
			parseTags(tagsRaw),
			executable,
		)

		snippets, err := store.LoadSnippets()
		if err != nil {
			fmt.Println("❌ Failed to load snippets:", err)
			return
		}

		snippets = append(snippets, *s)

		if err := store.SaveSnippets(snippets); err != nil {
			fmt.Println("❌ Failed to save snippet:", err)
			return
		}

		fmt.Println("✅ Snippet added successfully!")
		if executable {
			fmt.Println("🚀 This snippet is marked as executable and can be run with 'codestash exec'")
		}
	},
}

func parseTags(input string) []string {
	parts := strings.Split(input, ",")
	var tags []string
	for _, t := range parts {
		if tag := strings.TrimSpace(t); tag != "" {
			tags = append(tags, tag)
		}
	}
	return tags
}
