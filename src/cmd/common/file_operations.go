package common

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gphorvath/grimoire/src/config"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func CreateFileIfNotExists(path string) error {
	if !FileExists(path) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func FindFile(baseDir, filename string) (string, error) {
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

func OpenFileInEditor(path string) error {
	if !FileExists(path) {
		return errors.New("file does not exist")
	}

	cmd := exec.Command(config.Editor, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
