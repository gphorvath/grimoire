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
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [prompt name]",
	Short: "Delete the prompt",
	Args:  cobra.ExactArgs(1),
	Run:   runDeleteCmd,
}

func runDeleteCmd(cmd *cobra.Command, args []string) {
	baseDir := config.GetPromptDir()
	filename := args[0] + ".md"

	dir, err := common.FindFile(baseDir, filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if dir == "" {
		fmt.Fprintf(os.Stderr, "Error: prompt not found\n")
		os.Exit(1)
	}

	filePath := filepath.Join(dir, filename)
	if err := os.Remove(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Deleted prompt file %s\n", filePath)
}
