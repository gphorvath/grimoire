package cmd

import (
	"fmt"
	"os"

	"github.com/gphorvath/grimoire/src/cmd/common"
	"github.com/gphorvath/grimoire/src/config"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var copyCmd = &cobra.Command{
	Use:   "copy [prompt name]",
	Short: "Copy the prompt to the clipboard",
	Args:  cobra.ExactArgs(1),
	Run:   runCopyCmd,
}

func init() {
	rootCmd.AddCommand(copyCmd)
}

func runCopyCmd(cmd *cobra.Command, args []string) {
	baseDir := config.GetPromptDir()
	filename := args[0] + ".md"
	filePath, err := common.FindAndJoin(baseDir, filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Init returns an error if the package is not ready for use.
	err = clipboard.Init()
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	clipboard.Write(clipboard.FmtText, content)
}
