package main

import (
	"context"
	"log"
	"os"

	"github.com/shibbirmcc/github-issue-developer-mcp-server/internal/server"
)

func main() {
	ctx := context.Background()

	// Create and start the MCP server
	mcpServer := server.NewMCPServer()

	if err := mcpServer.Start(ctx); err != nil {
		log.Printf("Error starting MCP server: %v", err)
		os.Exit(1)
	}
}
