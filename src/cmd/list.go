package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gphorvath/grimoire/src/config"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [directory]",
	Short: "List files in the target directory",
	Args:  cobra.MaximumNArgs(1),
	Run:   runListCmd,
}

func runListCmd(cmd *cobra.Command, args []string) {
	dir := config.GetPromptDir()
	if len(args) > 0 {
		dir = args[0]
	}
	err := printFileTree(dir, "")
	if err != nil {
		log.Fatalf("Failed to list files in directory %s: %v", dir, err)
	}
}

func printFileTree(dir string, prefix string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for i, file := range files {
		if i == len(files)-1 {
			fmt.Printf("%s└── %s\n", prefix, file.Name())
			if file.IsDir() {
				printFileTree(dir+"/"+file.Name(), prefix+"    ")
			}
		} else {
			fmt.Printf("%s├── %s\n", prefix, file.Name())
			if file.IsDir() {
				printFileTree(dir+"/"+file.Name(), prefix+"│   ")
			}
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
