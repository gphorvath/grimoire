package prompt

import (
	"time"
)

// PromptMetadata represents the YAML front matter in prompt files
type PromptMetadata struct {
	Title       string    `yaml:"title"`
	Model       string    `yaml:"model"`
	Description string    `yaml:"description"`
	Input       string    `yaml:"input"`
	Output      string    `yaml:"output"`
	Version     string    `yaml:"version"`
	Updated     time.Time `yaml:"updated"`
	Author      string    `yaml:"author"`
	Email       string    `yaml:"email"`
	Tags        []string  `yaml:"tags"`
}

type FileMetadata struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type Prompt struct {
	FileMetadata   FileMetadata
	PromptMetadata PromptMetadata
	Content        string
}
