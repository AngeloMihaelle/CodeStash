package snippet

import (
	"time"
)

type Snippet struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Code        string   `json:"code"`
	Tags        []string `json:"tags"`
	Executable  bool     `json:"executable"`
	Language    string   `json:"language"`
	Description string   `json:"description"`
	UsageCount  int      `json:"usage_count"`
	LastUsed    string   `json:"last_used"`
	CreatedAt   string   `json:"created_at"`
}

func NewSnippet(title, code, desc, lang string, tags []string, executable bool) *Snippet {
	now := time.Now().UTC().Format(time.RFC3339)
	return &Snippet{
		ID:          generateID(),
		Title:       title,
		Code:        code,
		Tags:        tags,
		Executable:  executable,
		Language:    lang,
		Description: desc,
		UsageCount:  0,
		LastUsed:    "",
		CreatedAt:   now,
	}
}
