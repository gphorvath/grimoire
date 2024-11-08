package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func GenerateDocs(rootCmd *cobra.Command) error {
	// Generate markdown files
	return doc.GenMarkdownTree(rootCmd, "./docs/content")
}

func init() {
	rootCmd.AddCommand(docsCmd)
}

var docsCmd = &cobra.Command{
	Use:    "docs",
	Short:  "Generate markdown documentation",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return GenerateDocs(rootCmd)
	},
}
