package prompts

import (
	"testing"
)

func TestNewPromptManager(t *testing.T) {
	pm := NewPromptManager()

	if pm == nil {
		t.Fatal("NewPromptManager() returned nil")
	}

	prompts := pm.GetAllPrompts()
	if len(prompts) == 0 {
		t.Fatal("Expected prompts to be initialized, got empty slice")
	}

	expectedPrompts := []string{
		"git-best-practices",
		"github-workflow",
		"code-review-guidelines",
		"commit-message-format",
		"branch-naming-convention",
	}

	if len(prompts) != len(expectedPrompts) {
		t.Fatalf("Expected %d prompts, got %d", len(expectedPrompts), len(prompts))
	}

	// Check that all expected prompts are present
	promptMap := make(map[string]bool)
	for _, prompt := range prompts {
		promptMap[prompt.Name] = true
	}

	for _, expected := range expectedPrompts {
		if !promptMap[expected] {
			t.Errorf("Expected prompt %s not found", expected)
		}
	}
}

func TestGetAllPrompts(t *testing.T) {
	pm := NewPromptManager()
	prompts := pm.GetAllPrompts()

	// Verify each prompt has required fields
	for _, prompt := range prompts {
		if prompt.Name == "" {
			t.Error("Prompt name should not be empty")
		}

		if prompt.Description == "" {
			t.Error("Prompt description should not be empty")
		}

		if prompt.Handler == nil {
			t.Errorf("Prompt %s handler should not be nil", prompt.Name)
		}
	}
}

func TestPromptDescriptions(t *testing.T) {
	pm := NewPromptManager()
	prompts := pm.GetAllPrompts()

	expectedDescriptions := map[string]string{
		"git-best-practices":       "Provides Git best practices for development workflow",
		"github-workflow":          "Provides GitHub workflow best practices",
		"code-review-guidelines":   "Provides code review guidelines and best practices",
		"commit-message-format":    "Provides commit message formatting guidelines",
		"branch-naming-convention": "Provides branch naming convention guidelines",
	}

	for _, prompt := range prompts {
		expectedDesc, exists := expectedDescriptions[prompt.Name]
		if !exists {
			t.Errorf("Unexpected prompt: %s", prompt.Name)
			continue
		}

		if prompt.Description != expectedDesc {
			t.Errorf("Prompt %s: expected description %q, got %q",
				prompt.Name, expectedDesc, prompt.Description)
		}
	}
}
