package common

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gphorvath/grimoire/src/config"
)

// FileExists reports if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// CreateIfNotExists creates a file if it does not exist
// and returns an error if it fails to create the file
func CreateIfNotExists(filePath string) error {
	if !FileExists(filePath) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

// FindAndJoin walks the directory tree rooted at baseDir
// and returns the path of the first file with the given filename
// and an error if the file is not found
func FindAndJoin(baseDir, filename string) (string, error) {
	var foundPath string
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == filename {
			foundPath = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if foundPath == "" {
		return "", os.ErrNotExist
	}
	return foundPath, nil
}

// OpenInEditor opens a file in the configured editor
// and returns an error if it fails to open the file
func OpenInEditor(path string) error {
	if !FileExists(path) {
		return errors.New("file does not exist")
	}

	cmd := exec.Command(config.Editor, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
