// internal/prompt/manager.go

package prompt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Manager handles prompt storage and retrieval
type Manager struct {
	promptDir string
}

// New creates a new prompt manager
func New(promptDir string) (*Manager, error) {
	if promptDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get home directory: %w", err)
		}
		promptDir = filepath.Join(homeDir, ".grimoire", "prompts")
	}

	// Ensure prompt directory exists
	if err := os.MkdirAll(promptDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create prompt directory: %w", err)
	}

	return &Manager{promptDir: promptDir}, nil
}

// Load reads a prompt file and returns its contents
func (m *Manager) Load(name string) (*Prompt, error) {
	// Try exact path first
	path := filepath.Join(m.promptDir, name)
	if !strings.HasSuffix(path, ".md") {
		path += ".md"
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read prompt file: %w", err)
	}

	// Split front matter and content
	parts := strings.Split(string(content), "---\n")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid prompt format: missing front matter")
	}

	var metadata Metadata
	if err := yaml.Unmarshal([]byte(parts[1]), &metadata); err != nil {
		return nil, fmt.Errorf("failed to parse front matter: %w", err)
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	return &Prompt{
		Metadata:  metadata,
		Content:   strings.Join(parts[2:], "---\n"),
		FilePath:  path,
		UpdatedAt: fileInfo.ModTime(),
	}, nil
}

// List returns all available prompts
func (m *Manager) List() ([]Prompt, error) {
	var prompts []Prompt

	err := filepath.Walk(m.promptDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		relPath, err := filepath.Rel(m.promptDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		prompt, err := m.Load(relPath)
		if err != nil {
			return fmt.Errorf("failed to load prompt %s: %w", relPath, err)
		}

		prompts = append(prompts, *prompt)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list prompts: %w", err)
	}

	return prompts, nil
}
