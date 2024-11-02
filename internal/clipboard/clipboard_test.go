// internal/clipboard/clipboard_test.go

package clipboard

import (
	"strings"
	"testing"
)

func TestClipboardService(t *testing.T) {
	s := New()

	tests := []struct {
		name      string
		content   string
		wantErr   bool
		testPaste bool
	}{
		{
			name:      "basic text",
			content:   "hello world",
			wantErr:   false,
			testPaste: true,
		},
		{
			name:      "empty string",
			content:   "",
			wantErr:   false,
			testPaste: true,
		},
		{
			name:      "long text",
			content:   strings.Repeat("long ", 1000),
			wantErr:   false,
			testPaste: true,
		},
		{
			name:      "special characters",
			content:   "!@#$%^&*()_+-=[]{}|;:,.<>?",
			wantErr:   false,
			testPaste: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Copy
			err := s.Copy(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Test Paste if specified
			if tt.testPaste {
				got, err := s.Paste()
				if (err != nil) != tt.wantErr {
					t.Errorf("Paste() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.content {
					t.Errorf("Paste() = %v, want %v", got, tt.content)
				}
			}
		})
	}
}
