# GitHub Issue Developer MCP Server

A Model Context Protocol (MCP) server that provides prompts to instruct LLM agents on GitHub development best practices. This server helps developers follow proper Git workflows, GitHub collaboration patterns, and code review guidelines.

## Features

The MCP server exposes the following prompts:

- **git-best-practices**: Provides Git best practices for development workflow
- **github-workflow**: Provides GitHub workflow best practices for collaborative development
- **code-review-guidelines**: Provides comprehensive code review guidelines and best practices
- **commit-message-format**: Provides commit message formatting guidelines using conventional commits
- **branch-naming-convention**: Provides branch naming convention guidelines for organized development

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

```json
{
  "servers": {
    "github-issue-developer": {
      "command": "./github-issue-developer-mcp-server",
      "args": []
    }
  }
}
```

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
