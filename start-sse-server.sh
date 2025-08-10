#!/bin/bash

# GitHub Issue Developer MCP Server - SSE Mode Startup Script

echo "Starting GitHub Issue Developer MCP Server in SSE mode..."

# Kill any existing server instances
pkill -f github-issue-developer-mcp-server 2>/dev/null

# Build the server (in case there were changes)
echo "Building server..."
go build -o github-issue-developer-mcp-server

if [ $? -ne 0 ]; then
    echo "❌ Build failed!"
    exit 1
fi

# Start the server with SSE transport on port 8085
echo "Starting server on http://127.0.0.1:8085/ with SSE transport..."
MCP_HTTP_ADDR=":8085" ./github-issue-developer-mcp-server &

SERVER_PID=$!
echo "✅ Server started with PID: $SERVER_PID"
echo "🌐 Server URL: http://127.0.0.1:8085/"
echo "📋 Available prompts:"
echo "   - git-best-practices"
echo "   - github-workflow"
echo "   - code-review-guidelines"
echo "   - commit-message-format"
echo "   - branch-naming-convention"
echo "   - development-workflow"
echo ""
echo "💡 To test prompts, ask Cline:"
echo "   'Use the git-best-practices prompt to help me with Git workflow'"
echo ""
echo "🛑 To stop the server: pkill -f github-issue-developer-mcp-server"
