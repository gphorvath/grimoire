package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gphorvath/grimoire/src/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit [prompt name]",
	Short: "Edit the prompt in the chosen editor",
	Args:  cobra.ExactArgs(1),
	Run:   runEditCmd,
}

func runEditCmd(cmd *cobra.Command, args []string) {
	baseDir := config.GetPromptDir()
	filename := args[0] + ".md"

	dir, err := findFileDir(baseDir, filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if dir == "" {
		dir = baseDir
		filePath := filepath.Join(dir, filename)
		if err := createNewFile(filePath); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Created new prompt file %s\n", filePath)
	}

	filePath := filepath.Join(dir, filename)
	if err := openEditor(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func findFileDir(baseDir, filename string) (string, error) {
	var dir string
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == filename {
			dir = filepath.Dir(path)
			return filepath.SkipDir
		}
		return nil
	})
	return dir, err
}

func createNewFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("A new prompt!")
	return err
}

func openEditor(filePath string) error {
	editor := config.Editor

	editCmd := exec.Command(editor, filePath)
	editCmd.Stdin = os.Stdin
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stderr

	return editCmd.Run()
}
