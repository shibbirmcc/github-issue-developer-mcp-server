# GitHub Issue Developer MCP Server

A Model Context Protocol (MCP) server that provides prompts to instruct LLM agents on GitHub development best practices. This server helps developers follow proper Git workflows, GitHub collaboration patterns, and code review guidelines.

## Features

The MCP server exposes the following prompts:

- **git-best-practices**: Provides Git best practices for development workflow
- **github-workflow**: Provides GitHub workflow best practices for collaborative development
- **code-review-guidelines**: Provides comprehensive code review guidelines and best practices
- **commit-message-format**: Provides commit message formatting guidelines using conventional commits
- **branch-naming-convention**: Provides branch naming convention guidelines for organized development
- **development-workflow**: Comprehensive development workflow with Git, GitHub, and CI/CD best practices including pre-processing checks, branch management, testing requirements, and PR workflows

## Architecture

The server is built using the [Go MCP SDK](https://github.com/modelcontextprotocol/go-sdk) and follows clean architecture principles:

```
├── main.go                     # Application entry point
├── internal/
│   ├── server/                 # MCP server implementation
│   │   ├── server.go          # Server setup and configuration
│   │   └── server_test.go     # Server tests
│   └── prompts/               # Prompt management
│       ├── manager.go         # Prompt manager
│       ├── handlers.go        # Prompt handlers implementation
│       ├── manager_test.go    # Manager unit tests
│       └── handlers_test.go   # Handler integration tests
```

## Installation

### Prerequisites

- Go 1.21 or later
- Git

### Build from Source

```bash
git clone https://github.com/shibbirmcc/github-issue-developer-mcp-server.git
cd github-issue-developer-mcp-server
go build -o github-issue-developer-mcp-server
```

## Usage

### Stdio Transport (Default)

Run the server with stdio transport for direct integration with MCP clients:

```bash
./github-issue-developer-mcp-server
```

### HTTP/SSE Transport

Set the `MCP_HTTP_ADDR` environment variable to use HTTP with Server-Sent Events:

```bash
export MCP_HTTP_ADDR=":8080"
./github-issue-developer-mcp-server
```

The server will be available at `http://localhost:8080` with SSE support.

## Development

### Running Tests

Run all tests with coverage:

```bash
go test ./... -v
```

Run tests with coverage report:

```bash
go test ./... -cover
```

### Code Structure

The codebase follows Go best practices:

- **Clean Architecture**: Separation of concerns with internal packages
- **Dependency Injection**: Loose coupling between components
- **Comprehensive Testing**: Unit and integration tests with 100% coverage
- **Error Handling**: Proper error handling throughout the codebase
- **Documentation**: Well-documented code with clear comments

### Adding New Prompts

To add a new prompt:

1. Add the prompt definition to `internal/prompts/manager.go`
2. Implement the handler in `internal/prompts/handlers.go`
3. Add tests in `internal/prompts/handlers_test.go`
4. Update this README

## MCP Integration

This server implements the Model Context Protocol specification and can be integrated with any MCP-compatible client. The server provides:

- **Prompts**: Pre-defined prompts for development best practices
- **SSE Transport**: Real-time communication support
- **Stdio Transport**: Direct integration support

### Example MCP Client Configuration

#### For Stdio Transport (Default)
```json
{
  "servers": {
    "git-issue-developer-mcp-server": {
      "command": "./github-issue-developer-mcp-server",
      "args": []
    }
  }
}
```

#### For Cline (SSE Transport)
```json
{
  "mcpServers": {
    "git-issue-developer-mcp-server": {
      "disabled": false,
      "timeout": 60,
      "type": "sse",
      "url": "http://127.0.0.1:8085/",
      "autoApprove": [
        "list_tools",
        "list_prompts"
      ]
    }
  }
}
```

**Note**: For Cline, this configuration is automatically managed in:
`~/.config/Code/User/globalStorage/saoudrizwan.claude-dev/settings/cline_mcp_settings.json`

## Testing with LLM Agents

Here are practical examples of how to test the MCP server with any MCP-compatible LLM agent:

### 1. Setup and Configuration

First, ensure the MCP server is configured in your LLM agent's MCP settings:

```json
{
  "mcpServers": {
    "github-issue-developer": {
      "command": "/path/to/github-issue-developer-mcp-server",
      "args": []
    }
  }
}
```

### 2. Testing Individual Prompts

You can test each prompt by asking your LLM agent to use them:

#### Test Git Best Practices
```
Can you use the git-best-practices prompt to help me understand proper Git workflow?
```

#### Test GitHub Workflow
```
Use the github-workflow prompt to show me how to properly manage pull requests and code reviews.
```

#### Test Development Workflow (Comprehensive)
```
Apply the development-workflow prompt to help me implement a new feature following all best practices.
```

### 3. Real-World Testing Scenarios

#### Scenario 1: New Feature Development
```
I need to add a new authentication feature to my project. Use the development-workflow prompt to guide me through the entire process from start to finish.
```

Expected behavior:
- The agent will check Git repository status
- Ask about uncommitted changes
- Ensure you're not on master/main branch
- Guide through branch creation
- Enforce 100% test coverage
- Monitor CI/CD pipeline
- Create and manage pull request

#### Scenario 2: Code Review Process
```
I have a pull request that needs review. Use the code-review-guidelines prompt to help me review this code properly.
```

#### Scenario 3: Commit Message Formatting
```
I need to commit my changes. Use the commit-message-format prompt to help me write a proper commit message.
```

### 4. Testing the Comprehensive Workflow

Create a test scenario to validate the full development workflow:

#### Step 1: Initialize Test Project
```
Create a new Go project and use the development-workflow prompt to set it up properly with all best practices.
```

#### Step 2: Add Feature with Full Workflow
```
Add a simple "Hello World" HTTP endpoint to this project. Follow the complete development-workflow prompt including:
- Git repository checks
- Branch management
- Test writing (100% coverage)
- CI/CD monitoring
- Pull request creation
```

#### Step 3: Test Documentation Standards
```
Add some documentation to the project. Use the development-workflow prompt to ensure it follows the README.md-first approach.
```

### 5. Advanced Testing Scenarios

#### Test Branch Protection
```
I'm currently on the master branch. Use the development-workflow prompt to help me implement a new feature.
```
Expected: The agent should immediately warn about working on master and guide you to create a feature branch.

#### Test Uncommitted Changes Handling
```
I have some uncommitted changes in my repository. Use the development-workflow prompt to help me add a new feature.
```
Expected: The agent should detect uncommitted changes and ask you to commit or stash them first.

#### Test CI/CD Integration
```
Implement a new feature and use the development-workflow prompt to monitor the GitHub Actions pipeline until it passes.
```
Expected: The agent should push changes and actively monitor CI/CD status using `gh` CLI commands.

### 6. Verification Commands

To verify the MCP server is working correctly with your LLM agent:

#### Check Available Prompts
```
What prompts are available from the github-issue-developer MCP server?
```

#### Test Prompt Access
```
Show me the content of the development-workflow prompt.
```

#### Validate Workflow Enforcement
```
Create a simple project and try to commit directly to master branch using the development-workflow prompt.
```
Expected: Should be prevented and guided to create a feature branch instead.

### 7. Troubleshooting

If prompts aren't working:

1. **Check MCP Server Status**: Ensure the server is running and accessible
2. **Verify Configuration**: Check that the server path is correct in your agent's settings
3. **Test Connection**: Ask your agent to list available MCP servers
4. **Check Logs**: Look for any error messages in the agent's MCP connection logs

### 8. Expected Behaviors

When using the `development-workflow` prompt, your LLM agent should:

- ✅ Always check Git repository status first
- ✅ Handle uncommitted changes appropriately
- ✅ Prevent work on master/main branches
- ✅ Ask for user preferences on branch management
- ✅ Enforce 100% test coverage requirements
- ✅ Monitor CI/CD pipeline status
- ✅ Create comprehensive pull requests
- ✅ Handle review cycles properly
- ✅ Prioritize README.md for documentation
- ✅ Follow proper commit message formatting

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests for your changes
5. Ensure all tests pass (`go test ./...`)
6. Commit your changes (`git commit -m 'feat: add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Model Context Protocol](https://modelcontextprotocol.io/) for the specification
- [Go MCP SDK](https://github.com/modelcontextprotocol/go-sdk) for the implementation framework
