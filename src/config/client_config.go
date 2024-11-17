package config

type ClientConfig struct {
	OllamaConfig OllamaConfig `json:"ollama"`
}

type OllamaConfig struct {
	Host   string `json:"host"`
	Model  string `json:"model"`
	Stream bool   `json:"stream"`
}

func NewClientConfig() *ClientConfig {
	return &ClientConfig{
		OllamaConfig: OllamaConfig{
			Host:   getEnvAsString("OLLAMA_HOST", "http://localhost:11434/"),
			Model:  getEnvAsString("OLLAMA_MODEL", "llama3"),
			Stream: getEnvAsBool("OLLAMA_STREAM", true),
		},
	}
}
