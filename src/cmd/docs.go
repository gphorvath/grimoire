package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func generateHugoDoc(cmd *cobra.Command, baseDir string, weight int) error {
	// Create directory for the command
	cmdPath := strings.ReplaceAll(cmd.CommandPath(), " ", "/")
	cmdDir := filepath.Join(baseDir, cmdPath)
	if err := os.MkdirAll(cmdDir, 0755); err != nil {
		return err
	}

	// Create _index.md in the command's directory
	filename := filepath.Join(cmdDir, "_index.md")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write Hugo front matter
	fmt.Fprintf(f, `---
title: "%s"
description: "%s"
weight: %d
date: %s
draft: false
---

`, cmd.Name(), cmd.Short, weight, time.Now().Format("2006-01-02"))

	// Write command description
	fmt.Fprintf(f, "# %s\n\n", cmd.CommandPath())
	if cmd.Long != "" {
		fmt.Fprintf(f, "%s\n\n", cmd.Long)
	} else {
		fmt.Fprintf(f, "%s\n\n", cmd.Short)
	}

	// Write usage
	fmt.Fprintf(f, "## Usage\n\n```bash\n%s\n```\n\n", cmd.UseLine())

	// Write examples if any
	if cmd.Example != "" {
		fmt.Fprintf(f, "## Examples\n\n```bash\n%s\n```\n\n", cmd.Example)
	}

	// Write flags if any
	if cmd.HasAvailableFlags() {
		fmt.Fprintf(f, "## Flags\n\n")
		fmt.Fprintf(f, "| Flag | Description | Default |\n")
		fmt.Fprintf(f, "|------|-------------|----------|\n")
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			defaultValue := flag.DefValue
			if defaultValue == "" {
				defaultValue = "-"
			}
			fmt.Fprintf(f, "| `--%s` | %s | `%s` |\n", flag.Name, flag.Usage, defaultValue)
		})
		fmt.Fprintf(f, "\n")
	}

	// Write subcommands if any
	if cmd.HasSubCommands() {
		fmt.Fprintf(f, "## Subcommands\n\n")
		for _, subCmd := range cmd.Commands() {
			if !subCmd.Hidden {
				fmt.Fprintf(f, "* [%s](./%s/) - %s\n",
					subCmd.Name(),
					strings.ReplaceAll(subCmd.Name(), " ", "-"),
					subCmd.Short)
			}
		}
		fmt.Fprintf(f, "\n")
	}

	// Generate docs for each subcommand
	for i, subCmd := range cmd.Commands() {
		if !subCmd.Hidden {
			if err := generateHugoDoc(subCmd, baseDir, weight+i+1); err != nil {
				return err
			}
		}
	}

	return nil
}

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation for grimoire",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create main commands directory and index
		docsDir := "./docs/content/"
		if err := os.MkdirAll(docsDir, 0755); err != nil {
			return err
		}

		// Create main index file
		index, err := os.Create(filepath.Join(docsDir, "_index.md"))
		if err != nil {
			return err
		}
		defer index.Close()

		// Write main index front matter
		fmt.Fprintf(index, `---
title: "Command Reference"
description: "Complete reference for Grimoire CLI commands"
weight: 1
date: %s
draft: false
---

# Command Reference

This section contains the complete reference for all Grimoire CLI commands.

## Available Commands

`, time.Now().Format("2006-01-02"))

		// List all top-level commands in the index
		for _, cmd := range rootCmd.Commands() {
			if !cmd.Hidden {
				fmt.Fprintf(index, "* [%s](grimoire/%s/) - %s\n",
					cmd.Name(),
					cmd.Name(),
					cmd.Short)
			}
		}

		// Generate docs for root command and all its children
		return generateHugoDoc(rootCmd, docsDir, 1)
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
