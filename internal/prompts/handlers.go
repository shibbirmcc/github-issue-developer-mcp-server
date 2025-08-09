package prompts

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// gitBestPracticesHandler provides Git best practices guidance
func (pm *PromptManager) gitBestPracticesHandler(ctx context.Context, ss *mcp.ServerSession, params *mcp.GetPromptParams) (*mcp.GetPromptResult, error) {
	return &mcp.GetPromptResult{
		Description: "Git best practices for development workflow",
		Messages: []*mcp.PromptMessage{
			{
				Role: "user",
				Content: &mcp.TextContent{
					Text: `You are working on a software development project. Follow these Git best practices:

## Commit Guidelines:
1. Make atomic commits - each commit should represent a single logical change
2. Write clear, descriptive commit messages in present tense
3. Use conventional commit format: type(scope): description
4. Keep commits small and focused
5. Test your changes before committing

## Branch Management:
1. Use feature branches for new development
2. Keep main/master branch stable and deployable
3. Use descriptive branch names (feature/add-user-auth, bugfix/fix-login-error)
4. Regularly sync with main branch to avoid conflicts
5. Delete merged branches to keep repository clean

## Workflow Best Practices:
1. Pull latest changes before starting new work
2. Create feature branch from main
3. Make incremental commits with clear messages
4. Push regularly to backup your work
5. Create pull request when feature is complete
6. Review code thoroughly before merging
7. Use squash merge for clean history when appropriate

## Code Quality:
1. Run tests before committing
2. Use pre-commit hooks for code formatting and linting
3. Write meaningful commit messages that explain the "why"
4. Include relevant issue numbers in commit messages
5. Keep commits focused on a single concern

Apply these practices consistently to maintain a clean, professional development workflow.`,
				},
			},
		},
	}, nil
}

// githubWorkflowHandler provides GitHub workflow best practices
func (pm *PromptManager) githubWorkflowHandler(ctx context.Context, ss *mcp.ServerSession, params *mcp.GetPromptParams) (*mcp.GetPromptResult, error) {
	return &mcp.GetPromptResult{
		Description: "GitHub workflow best practices for collaborative development",
		Messages: []*mcp.PromptMessage{
			{
				Role: "user",
				Content: &mcp.TextContent{
					Text: `You are collaborating on a GitHub project. Follow these GitHub workflow best practices:

## Issue Management:
1. Create detailed issues with clear descriptions and acceptance criteria
2. Use issue templates for consistency
3. Label issues appropriately (bug, feature, enhancement, etc.)
4. Assign issues to team members
5. Link issues to pull requests using keywords (fixes #123, closes #456)
6. Use project boards for tracking progress

## Pull Request Best Practices:
1. Create pull requests from feature branches
2. Write descriptive PR titles and descriptions
3. Include screenshots for UI changes
4. Reference related issues in PR description
5. Request reviews from appropriate team members
6. Respond to review comments promptly and professionally
7. Keep PRs focused and reasonably sized
8. Update PR branch with latest main before merging

## Code Review Guidelines:
1. Review code thoroughly for logic, style, and potential issues
2. Provide constructive feedback with suggestions
3. Approve PRs only when confident in the changes
4. Use GitHub's review features (comments, suggestions, approvals)
5. Test the changes locally when necessary
6. Check that CI/CD pipelines pass

## Repository Management:
1. Use branch protection rules for main branch
2. Require PR reviews before merging
3. Enable status checks and require them to pass
4. Use semantic versioning for releases
5. Maintain a clear README with setup instructions
6. Use GitHub Actions for CI/CD automation
7. Keep repository organized with proper folder structure

## Communication:
1. Use @mentions to notify relevant team members
2. Keep discussions focused and professional
3. Document decisions in issues and PRs
4. Use GitHub Discussions for broader topics
5. Update issue status regularly

Follow these practices to maintain an efficient and collaborative development environment.`,
				},
			},
		},
	}, nil
}

// codeReviewGuidelinesHandler provides code review guidelines
func (pm *PromptManager) codeReviewGuidelinesHandler(ctx context.Context, ss *mcp.ServerSession, params *mcp.GetPromptParams) (*mcp.GetPromptResult, error) {
	return &mcp.GetPromptResult{
		Description: "Comprehensive code review guidelines and best practices",
		Messages: []*mcp.PromptMessage{
			{
				Role: "user",
				Content: &mcp.TextContent{
					Text: `You are conducting a code review. Follow these comprehensive guidelines:

## What to Look For:
1. **Correctness**: Does the code do what it's supposed to do?
2. **Performance**: Are there any obvious performance issues?
3. **Security**: Are there potential security vulnerabilities?
4. **Maintainability**: Is the code easy to understand and modify?
5. **Testing**: Are there adequate tests for the changes?
6. **Documentation**: Is the code properly documented?

## Code Quality Checklist:
- [ ] Code follows project coding standards and style guidelines
- [ ] Variable and function names are descriptive and meaningful
- [ ] Code is properly structured and organized
- [ ] No code duplication (DRY principle)
- [ ] Functions are focused and do one thing well
- [ ] Error handling is appropriate and comprehensive
- [ ] No hardcoded values where configuration should be used
- [ ] Memory leaks and resource management are handled properly

## Review Process:
1. **Understand the Context**: Read the PR description and linked issues
2. **Check the Big Picture**: Does the approach make sense?
3. **Review Line by Line**: Look for bugs, style issues, and improvements
4. **Test Considerations**: Verify test coverage and quality
5. **Documentation**: Ensure code is well-documented
6. **Performance**: Consider performance implications
7. **Security**: Look for potential security issues

## Providing Feedback:
1. **Be Constructive**: Focus on the code, not the person
2. **Explain Why**: Provide reasoning for your suggestions
3. **Offer Solutions**: Don't just point out problems, suggest fixes
4. **Prioritize Issues**: Distinguish between critical issues and nice-to-haves
5. **Use Examples**: Show better alternatives when possible
6. **Be Respectful**: Maintain a professional and helpful tone

## Common Issues to Watch For:
- Null pointer exceptions and boundary conditions
- Race conditions in concurrent code
- SQL injection and other security vulnerabilities
- Memory leaks and resource management
- Inefficient algorithms or data structures
- Missing error handling
- Inconsistent coding style
- Lack of input validation
- Hardcoded configuration values
- Missing or inadequate tests

## Approval Criteria:
Only approve a PR when:
- All critical issues are resolved
- Code meets quality standards
- Tests are adequate and passing
- Documentation is sufficient
- You're confident the changes won't break existing functionality

Remember: Code review is a collaborative process aimed at improving code quality and sharing knowledge.`,
				},
			},
		},
	}, nil
}

// commitMessageFormatHandler provides commit message formatting guidelines
func (pm *PromptManager) commitMessageFormatHandler(ctx context.Context, ss *mcp.ServerSession, params *mcp.GetPromptParams) (*mcp.GetPromptResult, error) {
	return &mcp.GetPromptResult{
		Description: "Commit message formatting guidelines using conventional commits",
		Messages: []*mcp.PromptMessage{
			{
				Role: "user",
				Content: &mcp.TextContent{
					Text: "You are writing commit messages. Follow these formatting guidelines:\n\n" +
						"## Conventional Commit Format:\n" +
						"<type>[optional scope]: <description>\n\n" +
						"[optional body]\n\n" +
						"[optional footer(s)]\n\n" +
						"## Commit Types:\n" +
						"- **feat**: A new feature for the user\n" +
						"- **fix**: A bug fix for the user\n" +
						"- **docs**: Documentation changes\n" +
						"- **style**: Code style changes (formatting, missing semicolons, etc.)\n" +
						"- **refactor**: Code refactoring without changing functionality\n" +
						"- **perf**: Performance improvements\n" +
						"- **test**: Adding or updating tests\n" +
						"- **chore**: Maintenance tasks, dependency updates\n" +
						"- **ci**: CI/CD configuration changes\n" +
						"- **build**: Build system or external dependency changes\n" +
						"- **revert**: Reverting a previous commit\n\n" +
						"## Examples:\n" +
						"feat(auth): add user login functionality\n" +
						"fix(api): resolve null pointer exception in user service\n" +
						"docs(readme): update installation instructions\n" +
						"style(components): format code according to style guide\n" +
						"refactor(utils): extract common validation logic\n" +
						"test(auth): add unit tests for login service\n" +
						"chore(deps): update dependencies to latest versions\n\n" +
						"## Best Practices:\n" +
						"1. **Use imperative mood**: \"add feature\" not \"added feature\"\n" +
						"2. **Keep subject line under 50 characters**\n" +
						"3. **Capitalize the subject line**\n" +
						"4. **Don't end subject line with a period**\n" +
						"5. **Use body to explain what and why, not how**\n" +
						"6. **Separate subject from body with blank line**\n" +
						"7. **Wrap body at 72 characters**\n" +
						"8. **Reference issues and PRs in footer**\n\n" +
						"Follow these guidelines to maintain a clean, professional commit history.",
				},
			},
		},
	}, nil
}

// branchNamingConventionHandler provides branch naming guidelines
func (pm *PromptManager) branchNamingConventionHandler(ctx context.Context, ss *mcp.ServerSession, params *mcp.GetPromptParams) (*mcp.GetPromptResult, error) {
	return &mcp.GetPromptResult{
		Description: "Branch naming convention guidelines for organized development",
		Messages: []*mcp.PromptMessage{
			{
				Role: "user",
				Content: &mcp.TextContent{
					Text: "You are creating Git branches. Follow these naming conventions:\n\n" +
						"## Branch Naming Format:\n" +
						"<type>/<short-description>\n" +
						"<type>/<issue-number>-<short-description>\n" +
						"<type>/<scope>/<short-description>\n\n" +
						"## Branch Types:\n" +
						"- **feature/**: New features or enhancements\n" +
						"- **bugfix/**: Bug fixes\n" +
						"- **hotfix/**: Critical fixes for production\n" +
						"- **release/**: Release preparation branches\n" +
						"- **chore/**: Maintenance tasks, refactoring\n" +
						"- **docs/**: Documentation updates\n" +
						"- **test/**: Test-related changes\n" +
						"- **experiment/**: Experimental or proof-of-concept work\n\n" +
						"## Naming Rules:\n" +
						"1. Use lowercase letters\n" +
						"2. Use hyphens (-) to separate words, not underscores or spaces\n" +
						"3. Keep names concise but descriptive\n" +
						"4. Include issue numbers when applicable\n" +
						"5. Avoid special characters except hyphens\n" +
						"6. Use present tense verbs\n\n" +
						"Choose a convention that works for your team and stick to it consistently.",
				},
			},
		},
	}, nil
}
