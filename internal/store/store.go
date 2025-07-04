package store

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/AngeloMihaelle/CodeStash/internal/snippet"
)

var storagePath = filepath.Join(os.Getenv("HOME"), ".codestash", "snippets.json")

func LoadSnippets() ([]snippet.Snippet, error) {
	file, err := os.ReadFile(storagePath)
	if errors.Is(err, os.ErrNotExist) {
		return []snippet.Snippet{}, nil
	}
	if err != nil {
		return nil, err
	}
	var snippets []snippet.Snippet
	if err := json.Unmarshal(file, &snippets); err != nil {
		return nil, err
	}
	return snippets, nil
}

func SaveSnippets(snippets []snippet.Snippet) error {
	os.MkdirAll(filepath.Dir(storagePath), os.ModePerm)
	data, err := json.MarshalIndent(snippets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(storagePath, data, 0644)
}
