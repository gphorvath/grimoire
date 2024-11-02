package prompt

import (
	"time"
)

// Metadata represents the YAML front matter in prompt files
type Metadata struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
}

// Prompt represents a complete prompt with its metadata and content
type Prompt struct {
	Metadata  Metadata
	Content   string
	FilePath  string
	UpdatedAt time.Time
}
