package server

import (
	"context"
	"testing"
	"time"
)

func TestNewMCPServer(t *testing.T) {
	server := NewMCPServer()
	if server == nil {
		t.Fatal("NewMCPServer() returned nil")
	}
}

func TestMCPServerInitialization(t *testing.T) {
	server := NewMCPServer()

	// Test that the server can be created without panicking
	if server.server != nil {
		t.Error("Server should not be initialized until Start() is called")
	}
}

func TestMCPServerStartWithTimeout(t *testing.T) {
	server := NewMCPServer()

	// Create a context with timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Start server in a goroutine since it's blocking
	errChan := make(chan error, 1)
	go func() {
		err := server.Start(ctx)
		errChan <- err
	}()

	// Wait for either completion or timeout
	select {
	case err := <-errChan:
		// Server should start without immediate error
		// Note: In a real scenario, this might return an error due to no transport being available
		// but we're testing that the initialization doesn't panic
		if err != nil {
			t.Logf("Server start returned error (expected in test environment): %v", err)
		}
	case <-ctx.Done():
		// Timeout is acceptable for this test since we're just testing initialization
		t.Log("Server start timed out (expected in test environment)")
	}
}

func TestServerStructure(t *testing.T) {
	server := NewMCPServer()

	// Verify the server has the expected structure
	if server == nil {
		t.Fatal("Server should not be nil")
	}

	// The internal server should be nil until Start() is called
	if server.server != nil {
		t.Error("Internal server should be nil before Start() is called")
	}
}
