package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gphorvath/grimoire/src/config"
	"github.com/spf13/cobra"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

var (
	completeCmd = &cobra.Command{
		Use:   "complete [prompt]",
		Short: "Get completion from Ollama",
		Args:  cobra.ExactArgs(1),
		RunE:  runComplete,
	}
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

func runComplete(cmd *cobra.Command, args []string) error {
	prompt := args[0]

	// Create request body
	reqBody := OllamaRequest{
		Model:  config.OllamaModel,
		Prompt: prompt,
		Stream: config.OllamaStream,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	url := config.OllamaURL + "/generate"

	// Make POST request to Ollama API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle streaming response
	decoder := json.NewDecoder(resp.Body)
	for {
		var ollamaResp OllamaResponse
		if err := decoder.Decode(&ollamaResp); err != nil {
			if err == io.EOF {
				fmt.Println()
				break
			}
			return err
		}
		fmt.Print(ollamaResp.Response)
	}

	return nil
}
