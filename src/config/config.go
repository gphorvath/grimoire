package config

import (
	"os"
	"path/filepath"
)

const (
	Logo = `
┌─────────────────────────────┐
│     ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄     │
│   ⫷ █ G R I M O I R E █ ⫸   │
│     ▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀     │
│ ╔═══*.·:·.☽✧    ✦☾.·:·.*══╗ │
│ ║ ⚡         ✴        ⚡  ║ │
│ ║     Gen AI Companion    ║ │
│ ║     ⚝    Prompt   ⚝     ║ │
│ ║     ⚝    Prompt   ⚝     ║ │
│ ║     ⚝    Prompt   ⚝     ║ │
│ ║     ⚝    Prompt   ⚝     ║ │
│ ║       ☽    ❈    ☾       ║ │
│ ╚═════*.·:·.☽✧✦☾.·:·.*════╝ │
└─────────────────────────────┘
`
)

var (
	OllamaModel  = getEnv("OLLAMA_MODEL", "llama3.2")
	OllamaURL    = getEnv("OLLAMA_URL", "http://localhost:11434/api")
	OllamaStream = getEnvAsBool("OLLAMA_STREAM", true)
	Editor       = getEnv("EDITOR", "vim")
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return value == "true"
	}
	return defaultValue
}

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
