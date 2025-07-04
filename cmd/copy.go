package cmd

import (
	"fmt"

	"github.com/AngeloMihaelle/CodeStash/internal/store"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy [snippet-id-or-title]",
	Short: "Copy a snippet to the clipboard",
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

		// Copy to clipboard
		if err := copyToClipboard(targetSnippet.Code); err != nil {
			fmt.Println("❌ Failed to copy to clipboard:", err)
			return
		}

		fmt.Printf("📋 Copied '%s' to clipboard\n", targetSnippet.Title)
	},
}

// func copyToClipboard(text string) error {
// 	var cmd *exec.Cmd

// 	switch runtime.GOOS {
// 	case "darwin":
// 		cmd = exec.Command("pbcopy")
// 	case "linux":
// 		// Try xclip first, then xsel as fallback
// 		if _, err := exec.LookPath("xclip"); err == nil {
// 			cmd = exec.Command("xclip", "-selection", "clipboard")
// 		} else if _, err := exec.LookPath("xsel"); err == nil {
// 			cmd = exec.Command("xsel", "--clipboard", "--input")
// 		} else {
// 			return fmt.Errorf("clipboard access requires xclip or xsel on Linux")
// 		}
// 	case "windows":
// 		cmd = exec.Command("clip")
// 	default:
// 		return fmt.Errorf("unsupported platform for clipboard operations: %s", runtime.GOOS)
// 	}

// 	cmd.Stdin = strings.NewReader(text)
// 	return cmd.Run()
// }
