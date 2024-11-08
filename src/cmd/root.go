package cmd

import (
	"context"

	"github.com/gphorvath/grimoire/src/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{

	Use:   "grimoire",
	Short: "Grimoire - A CLI tool for managing and using GenAI prompts",
	Long:  getBanner(),
}

func getBanner() string {
	return config.Logo + "\n Grimoire is your magical companion to manage and quickly access a collection of GenAI prompts."
}

func Execute(ctx context.Context) error {
	return rootCmd.Execute()
}
