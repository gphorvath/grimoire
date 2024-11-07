package cmd

import (
	"context"

	"github.com/gphorvath/grimoire/src/cmd/say"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grimoire",
	Short: "Grimoire - A CLI tool for managing and using GenAI prompts",
	Long:  `Grimoire helps you manage and quickly access a collection of GenAI prompts.`,
}

func Execute(ctx context.Context) error {
	rootCmd.AddCommand(say.Command())
	return rootCmd.Execute()
}
