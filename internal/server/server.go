package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/shibbirmcc/github-issue-developer-mcp-server/internal/prompts"
)

// MCPServer represents the MCP server instance
type MCPServer struct {
	server *mcp.Server
}

// NewMCPServer creates a new MCP server instance
func NewMCPServer() *MCPServer {
	return &MCPServer{}
}

// Start initializes and starts the MCP server
func (s *MCPServer) Start(ctx context.Context) error {
	// Create server with proper implementation
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "github-issue-developer",
		Version: "1.0.0",
	}, nil)

	// Register prompt handlers
	s.registerPrompts(server)

	s.server = server

	log.Println("Starting GitHub Issue Developer MCP Server...")
	log.Println("Server Name: github-issue-developer")
	log.Println("Version: 1.0.0")

	// Check if HTTP address is provided via environment variable
	httpAddr := os.Getenv("MCP_HTTP_ADDR")
	if httpAddr != "" {
		// Start HTTP server with SSE support
		handler := mcp.NewSSEHandler(func(*http.Request) *mcp.Server {
			return server
		})
		log.Printf("MCP server listening at %s", httpAddr)
		return http.ListenAndServe(httpAddr, handler)
	} else {
		// Use stdio transport
		transport := mcp.NewStdioTransport()
		if err := server.Run(ctx, transport); err != nil {
			return fmt.Errorf("failed to start server: %w", err)
		}
	}

	return nil
}

// registerPrompts registers all available prompts with the server
func (s *MCPServer) registerPrompts(server *mcp.Server) {
	promptManager := prompts.NewPromptManager()

	// Register all prompts
	for _, prompt := range promptManager.GetAllPrompts() {
		server.AddPrompt(&mcp.Prompt{
			Name:        prompt.Name,
			Description: prompt.Description,
		}, prompt.Handler)
		log.Printf("Registered prompt: %s - %s", prompt.Name, prompt.Description)
	}
}
