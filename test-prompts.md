# Testing MCP Prompts

Your MCP server is now running and connected to Cline! Here's how to test the prompts:

## Server Status
✅ Server is running on http://127.0.0.1:8085/
✅ Server name matches Cline configuration: `git-issue-developer-mcp-server`
✅ All 6 prompts are registered and available
✅ MCP connection is established (confirmed by proper MCP error responses)

## Available Prompts

1. **git-best-practices** - Git best practices for development workflow
2. **github-workflow** - GitHub workflow best practices for collaborative development
3. **code-review-guidelines** - Comprehensive code review guidelines and best practices
4. **commit-message-format** - Commit message formatting guidelines using conventional commits
5. **branch-naming-convention** - Branch naming convention guidelines for organized development
6. **development-workflow** - Comprehensive development workflow with Git, GitHub, and CI/CD best practices

## How to Test

Try asking Cline to use the prompts with these example requests:

### Test 1: Git Best Practices
```
Can you use the git-best-practices prompt to help me understand proper Git workflow?
```

### Test 2: Development Workflow (Most Comprehensive)
```
Use the development-workflow prompt to guide me through implementing a new feature following all best practices.
```

### Test 3: Code Review Guidelines
```
Apply the code-review-guidelines prompt to help me review some code properly.
```

### Test 4: Commit Message Format
```
Use the commit-message-format prompt to help me write a proper commit message.
```

### Test 5: GitHub Workflow
```
Use the github-workflow prompt to show me how to properly manage pull requests.
```

### Test 6: Branch Naming
```
Apply the branch-naming-convention prompt to help me name my branches properly.
```

## Troubleshooting

If prompts still don't work:

1. **Restart Cline Extension**: 
   - Press Ctrl+Shift+P
   - Type "Developer: Reload Window"
   - Press Enter

2. **Check Server Status**:
   ```bash
   netstat -tlnp | grep 8085
   ```

3. **Restart Server**:
   ```bash
   pkill -f github-issue-developer-mcp-server
   MCP_HTTP_ADDR=":8085" ./github-issue-developer-mcp-server &
   ```

## Expected Behavior

When you use a prompt, Cline should:
- Recognize the prompt name
- Fetch the prompt content from the MCP server
- Apply the prompt instructions to guide its responses
- Follow the specific guidelines provided in each prompt

The prompts contain detailed instructions for Git workflows, GitHub best practices, code review processes, and comprehensive development workflows.
