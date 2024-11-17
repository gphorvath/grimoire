package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gphorvath/grimoire/src/cmd/common"
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
	promptFile         string
	editBeforeGenerate bool

	generateCmd = &cobra.Command{
		Use:   "generate [flags] [input...]",
		Short: "Get generation from Ollama",
		Long:  "Requests generation from Ollama while optionally prepending a prompt",
		Args: func(cmd *cobra.Command, args []string) error {
			editFlag, _ := cmd.Flags().GetBool("edit")
			if !editFlag && len(args) < 1 {
				return fmt.Errorf("requires at least 1 arg when not using edit flag")
			}
			return nil
		},
		RunE: runGenerate,
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&promptFile, "prompt", "p", "", "Prompt file to prepend to input")
	generateCmd.Flags().BoolVarP(&editBeforeGenerate, "edit", "e", false, "Edit input before generating")
}
func runGenerate(cmd *cobra.Command, args []string) error {
	// Join all arguments as the input text
	input := strings.Join(args, " ")
	var finalPrompt string

	if promptFile != "" {
		// Load the prompt file if specified
		baseDir := config.GetPromptDir()
		filename := promptFile + ".md"

		filepath, err := common.FindAndJoin(baseDir, filename)
		if err != nil {
			return err
		}

		content, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}

		finalPrompt = string(content) + "\n" + input
	} else {
		finalPrompt = input
	}

	if editBeforeGenerate {
		// Create temporary file for editing
		tmpFile, err := os.CreateTemp("", "grimoire-*.txt")
		if err != nil {
			return err
		}
		defer os.Remove(tmpFile.Name())

		// Write prompt to temp file
		if _, err := tmpFile.WriteString(finalPrompt); err != nil {
			return err
		}
		tmpFile.Close()

		// Open in editor
		if err := common.OpenInEditor(tmpFile.Name()); err != nil {
			return err
		}

		// Read back edited content
		content, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			return err
		}
		finalPrompt = string(content)
	}

	cfg := config.NewClientConfig()

	// Create request body
	reqBody := OllamaRequest{
		Model:  cfg.OllamaConfig.Model,
		Prompt: finalPrompt,
		Stream: cfg.OllamaConfig.Stream,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	endpoint := "/api/generate"

	ollamaURL, err := url.JoinPath(cfg.OllamaConfig.Host, endpoint)
	if err != nil {
		return err
	}

	// Make POST request to Ollama API
	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(jsonData))
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
