package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	promptFile string

	generateCmd = &cobra.Command{
		Use:   "generate [flags] [input...]",
		Short: "Get generation from Ollama.",
		Long:  "Requests generation from Ollama while optionally prepending a prompt.",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runGenerate,
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&promptFile, "prompt", "p", "", "Prompt file to prepend to input")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	// Join all arguments as the input text
	input := strings.Join(args, " ")
	var finalPrompt string

	if promptFile != "" {
		// Load the prompt file if specified
		baseDir := config.GetPromptDir()
		filename := promptFile + ".md"

		dir, err := findFileDir(baseDir, filename)
		if err != nil {
			return err
		}

		if dir == "" {
			return fmt.Errorf("prompt not found")
		}

		filePath := filepath.Join(dir, filename)
		content, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		finalPrompt = string(content) + "\n" + input
	} else {
		finalPrompt = input
	}

	// Create request body
	reqBody := OllamaRequest{
		Model:  config.OllamaModel,
		Prompt: finalPrompt,
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
