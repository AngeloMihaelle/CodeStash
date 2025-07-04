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

var editCmd = &cobra.Command{
	Use:   "edit [snippet-id-or-title]",
	Short: "Edit an existing snippet",
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

		// Get field flag
		field, _ := cmd.Flags().GetString("field")

		if field != "" {
			// Edit specific field
			if err := editField(targetSnippet, field); err != nil {
				fmt.Printf("❌ Failed to edit field '%s': %v\n", field, err)
				return
			}
		} else {
			// Interactive edit of all fields
			if err := editSnippetInteractive(targetSnippet); err != nil {
				fmt.Printf("❌ Failed to edit snippet: %v\n", err)
				return
			}
		}

		// Save updated snippets
		if err := store.SaveSnippets(snippets); err != nil {
			fmt.Println("❌ Failed to save snippet:", err)
			return
		}

		fmt.Printf("✅ Snippet '%s' updated successfully!\n", targetSnippet.Title)
	},
}

func editSnippetInteractive(snippet *snippet.Snippet) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("📝 Editing snippet: %s\n", snippet.Title)
	fmt.Println("─────────────────────────────────────")
	fmt.Println("Press Enter to keep current value, or type new value:")
	fmt.Println()

	// Edit title
	fmt.Printf("📝 Title [%s]: ", snippet.Title)
	newTitle, _ := reader.ReadString('\n')
	newTitle = strings.TrimSpace(newTitle)
	if newTitle != "" {
		snippet.Title = newTitle
	}

	// Edit description
	fmt.Printf("🧾 Description [%s]: ", snippet.Description)
	newDescription, _ := reader.ReadString('\n')
	newDescription = strings.TrimSpace(newDescription)
	if newDescription != "" {
		snippet.Description = newDescription
	}

	// Edit language
	fmt.Printf("💻 Language [%s]: ", snippet.Language)
	newLanguage, _ := reader.ReadString('\n')
	newLanguage = strings.TrimSpace(newLanguage)
	if newLanguage != "" {
		snippet.Language = newLanguage
	}

	// Edit tags
	currentTags := strings.Join(snippet.Tags, ", ")
	fmt.Printf("🏷️ Tags [%s]: ", currentTags)
	newTagsRaw, _ := reader.ReadString('\n')
	newTagsRaw = strings.TrimSpace(newTagsRaw)
	if newTagsRaw != "" {
		snippet.Tags = parseTags(newTagsRaw)
	}

	// Edit executable status
	executableStatus := "No"
	if snippet.Executable {
		executableStatus = "Yes"
	}
	fmt.Printf("🚀 Executable [%s] (y/n): ", executableStatus)
	newExecutableRaw, _ := reader.ReadString('\n')
	newExecutableRaw = strings.TrimSpace(newExecutableRaw)
	if newExecutableRaw != "" {
		snippet.Executable = strings.ToLower(newExecutableRaw) == "y" || strings.ToLower(newExecutableRaw) == "yes"
	}

	// Edit code
	fmt.Printf("📋 Edit code? (y/N): ")
	editCodeRaw, _ := reader.ReadString('\n')
	editCodeRaw = strings.TrimSpace(editCodeRaw)
	if strings.ToLower(editCodeRaw) == "y" || strings.ToLower(editCodeRaw) == "yes" {
		fmt.Println("📋 Enter new code (end with 'EOF' on a new line):")
		fmt.Println("Current code:")
		fmt.Println("─────────────────────────────────────")
		fmt.Println(snippet.Code)
		fmt.Println("─────────────────────────────────────")
		fmt.Println()

		var lines []string
		for {
			line, _ := reader.ReadString('\n')
			if strings.TrimSpace(line) == "EOF" {
				break
			}
			lines = append(lines, line)
		}

		if len(lines) > 0 {
			snippet.Code = strings.Join(lines, "")
		}
	}

	return nil
}

func editField(snippet *snippet.Snippet, field string) error {
	reader := bufio.NewReader(os.Stdin)

	switch strings.ToLower(field) {
	case "title":
		fmt.Printf("📝 Current title: %s\n", snippet.Title)
		fmt.Print("📝 New title: ")
		newTitle, _ := reader.ReadString('\n')
		newTitle = strings.TrimSpace(newTitle)
		if newTitle == "" {
			return fmt.Errorf("title cannot be empty")
		}
		snippet.Title = newTitle

	case "description":
		fmt.Printf("🧾 Current description: %s\n", snippet.Description)
		fmt.Print("🧾 New description: ")
		newDescription, _ := reader.ReadString('\n')
		newDescription = strings.TrimSpace(newDescription)
		if newDescription != "" {
			snippet.Description = newDescription
		}

	case "language":
		fmt.Printf("💻 Current language: %s\n", snippet.Language)
		fmt.Print("💻 New language: ")
		newLanguage, _ := reader.ReadString('\n')
		newLanguage = strings.TrimSpace(newLanguage)
		if newLanguage != "" {
			snippet.Language = newLanguage
		}

	case "tags":
		currentTags := strings.Join(snippet.Tags, ", ")
		fmt.Printf("🏷️ Current tags: %s\n", currentTags)
		fmt.Print("🏷️ New tags (comma separated): ")
		newTagsRaw, _ := reader.ReadString('\n')
		newTagsRaw = strings.TrimSpace(newTagsRaw)
		if newTagsRaw != "" {
			snippet.Tags = parseTags(newTagsRaw)
		}

	case "executable":
		executableStatus := "No"
		if snippet.Executable {
			executableStatus = "Yes"
		}
		fmt.Printf("🚀 Current executable status: %s\n", executableStatus)
		fmt.Print("🚀 New executable status (y/n): ")
		newExecutableRaw, _ := reader.ReadString('\n')
		newExecutableRaw = strings.TrimSpace(newExecutableRaw)
		if newExecutableRaw != "" {
			snippet.Executable = strings.ToLower(newExecutableRaw) == "y" || strings.ToLower(newExecutableRaw) == "yes"
		}

	case "code":
		fmt.Println("📋 Current code:")
		fmt.Println("─────────────────────────────────────")
		fmt.Println(snippet.Code)
		fmt.Println("─────────────────────────────────────")
		fmt.Println("📋 Enter new code (end with 'EOF' on a new line):")

		var lines []string
		for {
			line, _ := reader.ReadString('\n')
			if strings.TrimSpace(line) == "EOF" {
				break
			}
			lines = append(lines, line)
		}

		if len(lines) > 0 {
			snippet.Code = strings.Join(lines, "")
		}

	default:
		return fmt.Errorf("unknown field '%s'. Valid fields: title, description, language, tags, executable, code", field)
	}

	return nil
}

func init() {
	editCmd.Flags().StringP("field", "f", "", "Edit specific field (title, description, language, tags, executable, code)")
}
