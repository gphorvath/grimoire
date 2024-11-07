package say

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
	return `
┌─────────────────────────────┐
│     ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄     │
│   ⫷ █ G R I M O I R E █ ⫸   │
│     ▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀     │
│ ╔═══*.·:·.☽✧    ✦☾.·:·.*══╗ │
│ ║ ⚡         ✴        ⚡  ║ │
│ ║    Gen AI Companion     ║ │
│ ║     ⚝    Prompt   ⚝     ║ │
│ ║     ⚝    Prompt   ⚝     ║ │
│ ║     ⚝    Prompt   ⚝     ║ │
│ ║     ⚝    Prompt   ⚝     ║ │
│ ║       ☽    ❈    ☾       ║ │
│ ╚═════*.·:·.☽✧✦☾.·:·.*════╝ │
└─────────────────────────────┘
`
}
