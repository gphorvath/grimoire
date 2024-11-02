// cmd/grimoire/main.go

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gphorvath/grimoire/internal/clipboard"
	"github.com/gphorvath/grimoire/internal/prompt"
	"github.com/spf13/cobra"
)

var (
	clipSvc   *clipboard.Service
	promptMgr *prompt.Manager
	version   = "dev" // This will be set during build
)

func handleError(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "grimoire",
	Short: "Grimoire - A CLI tool for managing and using GenAI prompts",
	Long: `Grimoire helps you manage and quickly access a collection of GenAI prompts.
Copy prompts directly to your clipboard with simple commands.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("grimoire version %s\n", version)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available prompts",
	Run: func(cmd *cobra.Command, args []string) {
		prompts, err := promptMgr.List()
		handleError(err, "Failed to list prompts")

		fmt.Println("Available prompts:")
		for _, p := range prompts {
			fmt.Printf("- %s: %s [%s]\n",
				filepath.Base(p.FilePath),
				p.Metadata.Title,
				strings.Join(p.Metadata.Tags, ", "))
		}
	},
}

var copyCmd = &cobra.Command{
	Use:   "copy [prompt-name]",
	Short: "Copy a prompt to clipboard",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p, err := promptMgr.Load(args[0])
		handleError(err, "Failed to load prompt")

		err = clipSvc.Copy(p.Content)
		handleError(err, "Failed to copy to clipboard")
		fmt.Printf("Copied prompt '%s' to clipboard\n", p.Metadata.Title)
	},
}

var pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "Read current clipboard content",
	Run: func(cmd *cobra.Command, args []string) {
		text, err := clipSvc.Paste()
		handleError(err, "Failed to read clipboard")
		fmt.Println("Clipboard content:")
		fmt.Println("---")
		fmt.Println(text)
		fmt.Println("---")
	},
}

func init() {
	var err error
	clipSvc = clipboard.New()
	promptMgr, err = prompt.New("")
	if err != nil {
		handleError(err, "Failed to initialize prompt manager")
	}

	clipSvc = clipboard.New()
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(copyCmd)
	rootCmd.AddCommand(pasteCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
