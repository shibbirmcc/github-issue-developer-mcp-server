package prompts

import (
	"context"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestGitBestPracticesHandler(t *testing.T) {
	pm := NewPromptManager()
	ctx := context.Background()
	params := &mcp.GetPromptParams{}

	result, err := pm.gitBestPracticesHandler(ctx, nil, params)
	if err != nil {
		t.Fatalf("gitBestPracticesHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("gitBestPracticesHandler returned nil result")
	}

	if result.Description == "" {
		t.Error("Expected non-empty description")
	}

	if len(result.Messages) == 0 {
		t.Fatal("Expected at least one message")
	}

	message := result.Messages[0]
	if message.Role != "user" {
		t.Errorf("Expected role 'user', got %s", message.Role)
	}

	textContent, ok := message.Content.(*mcp.TextContent)
	if !ok {
		t.Fatal("Expected TextContent")
	}

	if textContent.Text == "" {
		t.Error("Expected non-empty text content")
	}

	// Check for key content
	expectedKeywords := []string{
		"Git best practices",
		"Commit Guidelines",
		"Branch Management",
		"atomic commits",
		"feature branches",
	}

	for _, keyword := range expectedKeywords {
		if !strings.Contains(textContent.Text, keyword) {
			t.Errorf("Expected text to contain %q", keyword)
		}
	}
}

func TestGithubWorkflowHandler(t *testing.T) {
	pm := NewPromptManager()
	ctx := context.Background()
	params := &mcp.GetPromptParams{}

	result, err := pm.githubWorkflowHandler(ctx, nil, params)
	if err != nil {
		t.Fatalf("githubWorkflowHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("githubWorkflowHandler returned nil result")
	}

	textContent, ok := result.Messages[0].Content.(*mcp.TextContent)
	if !ok {
		t.Fatal("Expected TextContent")
	}

	expectedKeywords := []string{
		"GitHub workflow",
		"Issue Management",
		"Pull Request",
		"Code Review",
		"Repository Management",
	}

	for _, keyword := range expectedKeywords {
		if !strings.Contains(textContent.Text, keyword) {
			t.Errorf("Expected text to contain %q", keyword)
		}
	}
}

func TestCodeReviewGuidelinesHandler(t *testing.T) {
	pm := NewPromptManager()
	ctx := context.Background()
	params := &mcp.GetPromptParams{}

	result, err := pm.codeReviewGuidelinesHandler(ctx, nil, params)
	if err != nil {
		t.Fatalf("codeReviewGuidelinesHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("codeReviewGuidelinesHandler returned nil result")
	}

	textContent, ok := result.Messages[0].Content.(*mcp.TextContent)
	if !ok {
		t.Fatal("Expected TextContent")
	}

	expectedKeywords := []string{
		"code review",
		"Correctness",
		"Performance",
		"Security",
		"Maintainability",
		"Testing",
		"Documentation",
	}

	for _, keyword := range expectedKeywords {
		if !strings.Contains(textContent.Text, keyword) {
			t.Errorf("Expected text to contain %q", keyword)
		}
	}
}

func TestCommitMessageFormatHandler(t *testing.T) {
	pm := NewPromptManager()
	ctx := context.Background()
	params := &mcp.GetPromptParams{}

	result, err := pm.commitMessageFormatHandler(ctx, nil, params)
	if err != nil {
		t.Fatalf("commitMessageFormatHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("commitMessageFormatHandler returned nil result")
	}

	textContent, ok := result.Messages[0].Content.(*mcp.TextContent)
	if !ok {
		t.Fatal("Expected TextContent")
	}

	expectedKeywords := []string{
		"commit messages",
		"Conventional Commit Format",
		"feat",
		"fix",
		"docs",
		"imperative mood",
	}

	for _, keyword := range expectedKeywords {
		if !strings.Contains(textContent.Text, keyword) {
			t.Errorf("Expected text to contain %q", keyword)
		}
	}
}

func TestBranchNamingConventionHandler(t *testing.T) {
	pm := NewPromptManager()
	ctx := context.Background()
	params := &mcp.GetPromptParams{}

	result, err := pm.branchNamingConventionHandler(ctx, nil, params)
	if err != nil {
		t.Fatalf("branchNamingConventionHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("branchNamingConventionHandler returned nil result")
	}

	textContent, ok := result.Messages[0].Content.(*mcp.TextContent)
	if !ok {
		t.Fatal("Expected TextContent")
	}

	expectedKeywords := []string{
		"Git branches",
		"Branch Naming Format",
		"feature/",
		"bugfix/",
		"hotfix/",
		"lowercase letters",
		"hyphens",
	}

	for _, keyword := range expectedKeywords {
		if !strings.Contains(textContent.Text, keyword) {
			t.Errorf("Expected text to contain %q", keyword)
		}
	}
}

func TestDevelopmentWorkflowHandler(t *testing.T) {
	pm := NewPromptManager()
	ctx := context.Background()
	params := &mcp.GetPromptParams{}

	result, err := pm.developmentWorkflowHandler(ctx, nil, params)
	if err != nil {
		t.Fatalf("developmentWorkflowHandler returned error: %v", err)
	}

	if result == nil {
		t.Fatal("developmentWorkflowHandler returned nil result")
	}

	textContent, ok := result.Messages[0].Content.(*mcp.TextContent)
	if !ok {
		t.Fatal("Expected TextContent")
	}

	expectedKeywords := []string{
		"development workflow",
		"PRE-PROCESSING CHECKS",
		"Git Repository Validation",
		"Uncommitted Changes Check",
		"Branch Protection Rule",
		"NEVER work on master/main branch",
		"100% coverage",
		"README.md FIRST",
		"gh CLI tool",
		"Pull Request",
		"CI/CD",
	}

	for _, keyword := range expectedKeywords {
		if !strings.Contains(textContent.Text, keyword) {
			t.Errorf("Expected text to contain %q", keyword)
		}
	}
}

func TestAllHandlersReturnValidStructure(t *testing.T) {
	pm := NewPromptManager()
	ctx := context.Background()
	params := &mcp.GetPromptParams{}

	handlers := []struct {
		name    string
		handler mcp.PromptHandler
	}{
		{"gitBestPracticesHandler", pm.gitBestPracticesHandler},
		{"githubWorkflowHandler", pm.githubWorkflowHandler},
		{"codeReviewGuidelinesHandler", pm.codeReviewGuidelinesHandler},
		{"commitMessageFormatHandler", pm.commitMessageFormatHandler},
		{"branchNamingConventionHandler", pm.branchNamingConventionHandler},
		{"developmentWorkflowHandler", pm.developmentWorkflowHandler},
	}

	for _, h := range handlers {
		t.Run(h.name, func(t *testing.T) {
			result, err := h.handler(ctx, nil, params)
			if err != nil {
				t.Fatalf("%s returned error: %v", h.name, err)
			}

			if result == nil {
				t.Fatalf("%s returned nil result", h.name)
			}

			if result.Description == "" {
				t.Errorf("%s: expected non-empty description", h.name)
			}

			if len(result.Messages) == 0 {
				t.Errorf("%s: expected at least one message", h.name)
			}

			for i, message := range result.Messages {
				if message.Role == "" {
					t.Errorf("%s: message %d has empty role", h.name, i)
				}

				if message.Content == nil {
					t.Errorf("%s: message %d has nil content", h.name, i)
				}

				if textContent, ok := message.Content.(*mcp.TextContent); ok {
					if textContent.Text == "" {
						t.Errorf("%s: message %d has empty text", h.name, i)
					}
				}
			}
		})
	}
}
