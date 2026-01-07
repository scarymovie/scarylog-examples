package main

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"

	"github.com/scarymovie/scarylog"
)

// TestBasicLogging tests basic logging functionality
func TestBasicLogging(t *testing.T) {
	var buf bytes.Buffer

	// Create a logger that writes to our buffer
	handler := slog.NewJSONHandler(&buf, nil)
	logger := scarylog.NewLogger(scarylog.WithHandler(handler))

	// Perform logging
	logger.Info("Test message", "key", "value")

	output := buf.String()
	if !strings.Contains(output, "Test message") {
		t.Errorf("Expected output to contain 'Test message', got: %s", output)
	}
	if !strings.Contains(output, "key") || !strings.Contains(output, "value") {
		t.Errorf("Expected output to contain key-value pair, got: %s", output)
	}
}

// TestContextLogging tests logging with context
func TestContextLogging(t *testing.T) {
	var buf bytes.Buffer

	handler := slog.NewJSONHandler(&buf, nil)
	logger := scarylog.NewLogger(scarylog.WithHandler(handler))

	// Add logger to context
	ctx := scarylog.ToContext(context.Background(), logger)

	// Retrieve and use logger from context
	ctxLogger := scarylog.FromContext(ctx)
	ctxLogger.Warn("Warning from context logger")

	output := buf.String()
	if !strings.Contains(output, "Warning from context logger") {
		t.Errorf("Expected output to contain 'Warning from context logger', got: %s", output)
	}
}

// TestErrorLogging tests error logging functionality
func TestErrorLogging(t *testing.T) {
	var buf bytes.Buffer

	handler := slog.NewJSONHandler(&buf, nil)
	logger := scarylog.NewLogger(scarylog.WithHandler(handler))

	err := &CustomError{"test error"}
	logger.Error("Error occurred", err)

	output := buf.String()
	if !strings.Contains(output, "Error occurred") {
		t.Errorf("Expected output to contain 'Error occurred', got: %s", output)
	}
	if !strings.Contains(output, "test error") {
		t.Errorf("Expected output to contain error message, got: %s", output)
	}
}

// TestWithFields tests logging with additional fields
func TestWithFields(t *testing.T) {
	var buf bytes.Buffer

	handler := slog.NewJSONHandler(&buf, nil)
	logger := scarylog.NewLogger(scarylog.WithHandler(handler))

	childLogger := logger.With("user_id", 123, "action", "login")
	childLogger.Info("User action")

	output := buf.String()
	if !strings.Contains(output, "user_id") || !strings.Contains(output, "123") {
		t.Errorf("Expected output to contain user_id field, got: %s", output)
	}
	if !strings.Contains(output, "action") || !strings.Contains(output, "login") {
		t.Errorf("Expected output to contain action field, got: %s", output)
	}
}

// CustomError is a test error type that implements error interface
type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}
