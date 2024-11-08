package cmd

import (
	"fmt"
	"os"

	"github.com/gphorvath/grimoire/src/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(Command())
}

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "say",
		Short: "Print Grimoire logo",
		Long:  `Print out the Grimoire logo.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintln(os.Stderr, logo())
			return err
		},
	}

	return cmd
}

func logo() string {
	return config.Logo
}
