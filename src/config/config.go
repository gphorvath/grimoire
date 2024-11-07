package config

import (
	"os"
	"path/filepath"
)

const (
	GitHubProject = "gphorvath/grimoire"
)

func GetConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, ".grimoire")
}

func GetPromptDir() string {
	return filepath.Join(GetConfigDir(), "prompts")
}
