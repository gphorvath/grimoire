package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gphorvath/grimoire/src/cmd/common"
	"github.com/gphorvath/grimoire/src/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit [prompt name]",
	Short: "Edit the prompt in the chosen editor",
	Args:  cobra.ExactArgs(1),
	Run:   runEditCmd,
}

func runEditCmd(cmd *cobra.Command, args []string) {
	baseDir := config.GetPromptDir()
	filename := args[0] + ".md"

	// Try to find existing file
	filePath, err := common.FindAndJoin(baseDir, filename)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// If file doesn't exist, create it in the base directory
	if filePath == "" {
		filePath = filepath.Join(baseDir, filename)
		if err := common.CreateIfNotExists(filePath); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		// Write example prompt to the new file
		if err := os.WriteFile(filePath, []byte(config.ExamplePrompt), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Created new prompt file %s\n", filePath)
	}

	if err := common.OpenInEditor(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
