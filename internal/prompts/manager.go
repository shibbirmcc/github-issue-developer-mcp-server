package prompts

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Prompt represents a single prompt with its handler
type Prompt struct {
	Name        string
	Description string
	Handler     mcp.PromptHandler
}

// PromptManager manages all available prompts
type PromptManager struct {
	prompts []Prompt
}

// NewPromptManager creates a new prompt manager with all available prompts
func NewPromptManager() *PromptManager {
	pm := &PromptManager{}
	pm.initializePrompts()
	return pm
}

// GetAllPrompts returns all registered prompts
func (pm *PromptManager) GetAllPrompts() []Prompt {
	return pm.prompts
}

// initializePrompts sets up all available prompts
func (pm *PromptManager) initializePrompts() {
	pm.prompts = []Prompt{
		{
			Name:        "git-best-practices",
			Description: "Provides Git best practices for development workflow",
			Handler:     pm.gitBestPracticesHandler,
		},
		{
			Name:        "github-workflow",
			Description: "Provides GitHub workflow best practices",
			Handler:     pm.githubWorkflowHandler,
		},
		{
			Name:        "code-review-guidelines",
			Description: "Provides code review guidelines and best practices",
			Handler:     pm.codeReviewGuidelinesHandler,
		},
		{
			Name:        "commit-message-format",
			Description: "Provides commit message formatting guidelines",
			Handler:     pm.commitMessageFormatHandler,
		},
		{
			Name:        "branch-naming-convention",
			Description: "Provides branch naming convention guidelines",
			Handler:     pm.branchNamingConventionHandler,
		},
		{
			Name:        "development-workflow",
			Description: "Comprehensive development workflow with Git, GitHub, and CI/CD best practices",
			Handler:     pm.developmentWorkflowHandler,
		},
	}
}
