package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/AngeloMihaelle/CodeStash/internal/snippet"
	"github.com/AngeloMihaelle/CodeStash/internal/store"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [snippet-id-or-title]",
	Short: "Print, copy, or execute a snippet",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := store.LoadSnippets()
		if err != nil {
			fmt.Println("❌ Failed to load snippets:", err)
			return
		}

		var targetSnippet *snippet.Snippet
		query := args[0]

		for i, s := range snippets {
			if s.ID == query || strings.EqualFold(s.Title, query) {
				targetSnippet = &snippets[i]
				break
			}
		}

		if targetSnippet == nil {
			fmt.Printf("❌ Snippet '%s' not found\n", query)
			return
		}

		targetSnippet.UsageCount++
		targetSnippet.LastUsed = time.Now().UTC().Format(time.RFC3339)

		if err := store.SaveSnippets(snippets); err != nil {
			fmt.Println("⚠️  Failed to update usage stats:", err)
		}

		copy, _ := cmd.Flags().GetBool("copy")
		execute, _ := cmd.Flags().GetBool("execute")
		force, _ := cmd.Flags().GetBool("force")

		if copy {
			if err := copyToClipboard(targetSnippet.Code); err != nil {
				fmt.Println("❌ Failed to copy to clipboard:", err)
				return
			}
			fmt.Printf("📋 Copied '%s' to clipboard\n", targetSnippet.Title)
		} else if execute {
			if !targetSnippet.Executable && !force {
				fmt.Printf("❌ Snippet '%s' is not marked as executable\n", targetSnippet.Title)
				fmt.Println("💡 Use --force to execute anyway, or mark the snippet as executable")
				return
			}

			if !targetSnippet.Executable && force {
				fmt.Printf("⚠️  Forcing execution of non-executable snippet '%s'\n", targetSnippet.Title)
			}

			if err := executeSnippet(targetSnippet); err != nil {
				fmt.Println("❌ Failed to execute snippet:", err)
				return
			}
		} else {
			fmt.Printf("📄 %s\n", targetSnippet.Title)
			fmt.Printf("📝 %s\n", targetSnippet.Description)
			if targetSnippet.Executable {
				fmt.Println("🚀 This snippet is executable")
			}
			fmt.Println("─────────────────────────────────────")
			fmt.Println(targetSnippet.Code)
			fmt.Println("─────────────────────────────────────")
		}
	},
}

func copyToClipboard(text string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		} else {
			return fmt.Errorf("clipboard access requires xclip or xsel on Linux")
		}
	case "windows":
		cmd = exec.Command("clip")
	default:
		return fmt.Errorf("unsupported platform for clipboard operations: %s", runtime.GOOS)
	}

	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func executeSnippet(s *snippet.Snippet) error {
	shellLangs := []string{"shell", "bash", "sh", "zsh", "fish", "powershell", "ps1", "cmd", "bat"}
	isShell := false
	for _, lang := range shellLangs {
		if strings.EqualFold(s.Language, lang) {
			isShell = true
			break
		}
	}

	if !s.Executable && !isShell {
		return fmt.Errorf("snippet '%s' is not marked as executable or shell-compatible", s.Title)
	}

	fmt.Printf("🚀 Executing '%s'...\n", s.Title)
	fmt.Println("─────────────────────────────────────")

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		switch {
		case strings.EqualFold(s.Language, "powershell"), strings.EqualFold(s.Language, "ps1"):
			path, err := writeTempScript(s.Code, ".ps1")
			if err != nil {
				return err
			}
			defer os.Remove(path)
			cmd = exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", path)

		case strings.EqualFold(s.Language, "cmd"), strings.EqualFold(s.Language, "bat"), strings.EqualFold(s.Language, "batch"):
			path, err := writeTempScript(s.Code, ".bat")
			if err != nil {
				return err
			}
			defer os.Remove(path)
			cmd = exec.Command("cmd", "/C", path)

		default:
			if strings.ContainsAny(s.Code, "\n\r\"") {
				path, err := writeTempScript(s.Code, ".bat")
				if err != nil {
					return err
				}
				defer os.Remove(path)
				cmd = exec.Command("cmd", "/C", path)
			} else {
				cmd = exec.Command("cmd", "/C", s.Code)
			}
		}

	default:
		shell := "/bin/sh"
		if strings.EqualFold(s.Language, "bash") {
			if _, err := exec.LookPath("bash"); err == nil {
				shell = "bash"
			}
		} else if strings.EqualFold(s.Language, "zsh") {
			if _, err := exec.LookPath("zsh"); err == nil {
				shell = "zsh"
			}
		} else if strings.EqualFold(s.Language, "fish") {
			if _, err := exec.LookPath("fish"); err == nil {
				shell = "fish"
			}
		}
		cmd = exec.Command(shell, "-c", s.Code)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func writeTempScript(code, extension string) (string, error) {
	tmpFile, err := os.CreateTemp("", "*"+extension)
	if err != nil {
		return "", fmt.Errorf("failed to create temp script: %v", err)
	}
	if _, err := tmpFile.WriteString(code); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to write to temp script: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to close temp script: %v", err)
	}
	return tmpFile.Name(), nil
}

func init() {
	useCmd.Flags().BoolP("copy", "c", false, "Copy snippet to clipboard")
	useCmd.Flags().BoolP("execute", "x", false, "Execute snippet")
	useCmd.Flags().BoolP("force", "f", false, "Force execution even if not marked as executable")
}

func parseCommand(command string) (string, []string) {
	var parts []string
	var current strings.Builder
	inQuotes := false
	escapeNext := false

	for _, char := range command {
		if escapeNext {
			current.WriteRune(char)
			escapeNext = false
			continue
		}
		if char == '\\' {
			escapeNext = true
			continue
		}
		if char == '"' {
			inQuotes = !inQuotes
			continue
		}
		if char == ' ' && !inQuotes {
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
		} else {
			current.WriteRune(char)
		}
	}
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}
	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}
