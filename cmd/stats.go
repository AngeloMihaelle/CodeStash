package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/AngeloMihaelle/CodeStash/internal/snippet"
	"github.com/AngeloMihaelle/CodeStash/internal/store"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show usage statistics and analytics",
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

		// Get detailed flag
		detailed, _ := cmd.Flags().GetBool("detailed")

		displayStats(snippets, detailed)
	},
}

func displayStats(snippets []snippet.Snippet, detailed bool) {
	fmt.Printf("📊 CodeStash Statistics\n")
	fmt.Printf("═══════════════════════════════════════\n\n")

	// Basic stats
	totalSnippets := len(snippets)
	totalUsage := 0
	executableCount := 0
	languageCount := make(map[string]int)
	tagCount := make(map[string]int)

	for _, s := range snippets {
		totalUsage += s.UsageCount
		if s.Executable {
			executableCount++
		}
		languageCount[s.Language]++
		for _, tag := range s.Tags {
			tagCount[tag]++
		}
	}

	fmt.Printf("📚 Total Snippets: %d\n", totalSnippets)
	fmt.Printf("🚀 Executable Snippets: %d\n", executableCount)
	fmt.Printf("📋 Non-executable Snippets: %d\n", totalSnippets-executableCount)
	fmt.Printf("📈 Total Usage: %d times\n", totalUsage)

	if totalSnippets > 0 {
		avgUsage := float64(totalUsage) / float64(totalSnippets)
		fmt.Printf("📊 Average Usage: %.1f times per snippet\n", avgUsage)
	}

	// Most used snippets
	fmt.Printf("\n🏆 Top 5 Most Used Snippets:\n")
	fmt.Printf("───────────────────────────────────────\n")

	sortedByUsage := make([]snippet.Snippet, len(snippets))
	copy(sortedByUsage, snippets)
	sort.Slice(sortedByUsage, func(i, j int) bool {
		return sortedByUsage[i].UsageCount > sortedByUsage[j].UsageCount
	})

	topCount := 5
	if len(sortedByUsage) < 5 {
		topCount = len(sortedByUsage)
	}

	for i := 0; i < topCount; i++ {
		s := sortedByUsage[i]
		fmt.Printf("%d. %s — used %d times", i+1, s.Title, s.UsageCount)
		if s.LastUsed != "" {
			if lastUsed, err := time.Parse(time.RFC3339, s.LastUsed); err == nil {
				fmt.Printf(" (last used: %s)", formatTimeAgo(lastUsed))
			}
		}
		fmt.Println()
	}

	// Most popular languages
	fmt.Printf("\n💻 Top Languages:\n")
	fmt.Printf("───────────────────────────────────────\n")

	type languageStat struct {
		name  string
		count int
	}

	var languageStats []languageStat
	for lang, count := range languageCount {
		languageStats = append(languageStats, languageStat{lang, count})
	}

	sort.Slice(languageStats, func(i, j int) bool {
		return languageStats[i].count > languageStats[j].count
	})

	topLangCount := 5
	if len(languageStats) < 5 {
		topLangCount = len(languageStats)
	}

	for i := 0; i < topLangCount; i++ {
		lang := languageStats[i]
		percentage := float64(lang.count) / float64(totalSnippets) * 100
		fmt.Printf("%d. %s — %d snippets (%.1f%%)\n", i+1, lang.name, lang.count, percentage)
	}

	// Most popular tags
	fmt.Printf("\n🏷️  Top Tags:\n")
	fmt.Printf("───────────────────────────────────────\n")

	type tagStat struct {
		name  string
		count int
	}

	var tagStats []tagStat
	for tag, count := range tagCount {
		tagStats = append(tagStats, tagStat{tag, count})
	}

	sort.Slice(tagStats, func(i, j int) bool {
		return tagStats[i].count > tagStats[j].count
	})

	topTagCount := 5
	if len(tagStats) < 5 {
		topTagCount = len(tagStats)
	}

	for i := 0; i < topTagCount; i++ {
		tag := tagStats[i]
		fmt.Printf("%d. %s — %d snippets\n", i+1, tag.name, tag.count)
	}

	// Recently created snippets
	fmt.Printf("\n🆕 Recently Created:\n")
	fmt.Printf("───────────────────────────────────────\n")

	sortedByCreated := make([]snippet.Snippet, len(snippets))
	copy(sortedByCreated, snippets)
	sort.Slice(sortedByCreated, func(i, j int) bool {
		return sortedByCreated[i].CreatedAt > sortedByCreated[j].CreatedAt
	})

	recentCount := 3
	if len(sortedByCreated) < 3 {
		recentCount = len(sortedByCreated)
	}

	for i := 0; i < recentCount; i++ {
		s := sortedByCreated[i]
		if created, err := time.Parse(time.RFC3339, s.CreatedAt); err == nil {
			fmt.Printf("• %s — created %s\n", s.Title, formatTimeAgo(created))
		} else {
			fmt.Printf("• %s — recently created\n", s.Title)
		}
	}

	// Recently used snippets
	fmt.Printf("\n🕒 Recently Used:\n")
	fmt.Printf("───────────────────────────────────────\n")

	var recentlyUsed []snippet.Snippet
	for _, s := range snippets {
		if s.LastUsed != "" {
			recentlyUsed = append(recentlyUsed, s)
		}
	}

	sort.Slice(recentlyUsed, func(i, j int) bool {
		return recentlyUsed[i].LastUsed > recentlyUsed[j].LastUsed
	})

	recentUsedCount := 3
	if len(recentlyUsed) < 3 {
		recentUsedCount = len(recentlyUsed)
	}

	if recentUsedCount == 0 {
		fmt.Println("• No snippets used yet")
	} else {
		for i := 0; i < recentUsedCount; i++ {
			s := recentlyUsed[i]
			if lastUsed, err := time.Parse(time.RFC3339, s.LastUsed); err == nil {
				fmt.Printf("• %s — used %s\n", s.Title, formatTimeAgo(lastUsed))
			} else {
				fmt.Printf("• %s — recently used\n", s.Title)
			}
		}
	}

	// Detailed stats if requested
	if detailed {
		fmt.Printf("\n📋 Detailed Statistics:\n")
		fmt.Printf("───────────────────────────────────────\n")

		// Unused snippets
		var unusedSnippets []snippet.Snippet
		for _, s := range snippets {
			if s.UsageCount == 0 {
				unusedSnippets = append(unusedSnippets, s)
			}
		}

		if len(unusedSnippets) > 0 {
			fmt.Printf("😴 Unused Snippets (%d):\n", len(unusedSnippets))
			for _, s := range unusedSnippets {
				fmt.Printf("   • %s (%s)\n", s.Title, s.Language)
			}
			fmt.Println()
		}

		// All languages breakdown
		fmt.Printf("💻 All Languages:\n")
		for _, lang := range languageStats {
			fmt.Printf("   • %s: %d snippets\n", lang.name, lang.count)
		}
		fmt.Println()

		// All tags breakdown
		if len(tagStats) > 0 {
			fmt.Printf("🏷️  All Tags:\n")
			for _, tag := range tagStats {
				fmt.Printf("   • %s: %d snippets\n", tag.name, tag.count)
			}
		}
	}
}

func formatTimeAgo(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)

	if duration < time.Hour {
		minutes := int(duration.Minutes())
		if minutes <= 1 {
			return "just now"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	}

	if duration < 24*time.Hour {
		hours := int(duration.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	}

	if duration < 7*24*time.Hour {
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	}

	if duration < 30*24*time.Hour {
		weeks := int(duration.Hours() / (24 * 7))
		if weeks == 1 {
			return "1 week ago"
		}
		return fmt.Sprintf("%d weeks ago", weeks)
	}

	months := int(duration.Hours() / (24 * 30))
	if months == 1 {
		return "1 month ago"
	}
	return fmt.Sprintf("%d months ago", months)
}

func init() {
	statsCmd.Flags().BoolP("detailed", "d", false, "Show detailed statistics including unused snippets")
}
