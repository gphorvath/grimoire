// internal/prompt/manager_test.go

package prompt

import (
	"os"
	"path/filepath"
	"testing"
)

func TestManager(t *testing.T) {
	// Create temporary directory for test prompts
	tmpDir, err := os.MkdirTemp("", "grimoire-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test prompt
	testPrompt := `---
title: Test Prompt
description: A test prompt
tags: [test]
---
This is a test prompt content`

	promptPath := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(promptPath, []byte(testPrompt), 0644); err != nil {
		t.Fatalf("Failed to write test prompt: %v", err)
	}

	// Create manager
	manager, err := New(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create prompt manager: %v", err)
	}

	// Test loading prompt
	t.Run("Load prompt", func(t *testing.T) {
		prompt, err := manager.Load("test")
		if err != nil {
			t.Fatalf("Failed to load prompt: %v", err)
		}

		if prompt.Metadata.Title != "Test Prompt" {
			t.Errorf("Expected title 'Test Prompt', got '%s'", prompt.Metadata.Title)
		}

		if !contains(prompt.Metadata.Tags, "test") {
			t.Error("Expected tags to contain 'test'")
		}
	})

	// Test listing prompts
	t.Run("List prompts", func(t *testing.T) {
		prompts, err := manager.List()
		if err != nil {
			t.Fatalf("Failed to list prompts: %v", err)
		}

		if len(prompts) != 1 {
			t.Errorf("Expected 1 prompt, got %d", len(prompts))
		}
	})
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
