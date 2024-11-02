package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grimoire",
	Short: "Grimoire - A CLI tool for managing and using GenAI prompts",
	Long: `Grimoire helps you manage and quickly access a collection of GenAI prompts.
Copy prompts directly to your clipboard with simple commands.`,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available prompts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available prompts will be listed here")
	},
}

var copyCmd = &cobra.Command{
	Use:   "copy [prompt-name]",
	Short: "Copy a prompt to clipboard",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Will copy prompt: %s to clipboard\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(copyCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
