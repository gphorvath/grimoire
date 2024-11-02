// internal/clipboard/clipboard.go

package clipboard

import (
	"fmt"

	"github.com/atotto/clipboard"
)

// Service handles clipboard operations
type Service struct{}

// New creates a new clipboard service
func New() *Service {
	return &Service{}
}

// Copy writes text to the system clipboard
func (s *Service) Copy(text string) error {
	if err := clipboard.WriteAll(text); err != nil {
		return fmt.Errorf("failed to copy to clipboard: %w", err)
	}
	return nil
}

// Paste reads text from the system clipboard
func (s *Service) Paste() (string, error) {
	text, err := clipboard.ReadAll()
	if err != nil {
		return "", fmt.Errorf("failed to read from clipboard: %w", err)
	}
	return text, nil
}
