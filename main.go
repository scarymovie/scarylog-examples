package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/scarymovie/scarylog"
)

func main() {
	fmt.Println("=== ScaryLog Logger Examples ===")

	// Example 1: Basic initialization and usage
	fmt.Println("\n1. Basic Logger Usage:")
	logger := scarylog.NewLogger()
	logger.Info("Application started")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("An error occurred", fmt.Errorf("sample error"))

	// Debug messages require setting the debug level
	debugLogger := scarylog.NewLogger(scarylog.WithLevel(slog.LevelDebug))
	debugLogger.Debug("This is a debug message", "debug_key", "debug_value")

	// Example 2: Logger with different log levels
	fmt.Println("\n2. Different Log Levels:")
	infoLogger := scarylog.NewLogger(scarylog.WithLevel(slog.LevelInfo))
	infoLogger.Info("This is an info message")
	infoLogger.Warn("This is a warning message")
	infoLogger.Error("This is an error message", fmt.Errorf("sample error"))

	// Example 3: Logger with context
	fmt.Println("\n3. Logger with Context:")
	ctx := context.Background()
	ctx = scarylog.ToContext(ctx, logger)
	scarylog.FromContext(ctx).Info("Message using logger from context")

	// Example 4: Logger with fields (structured logging)
	fmt.Println("\n4. Structured Logging with Fields:")
	logger.With("user_id", 123, "action", "login").Info("User logged in")
	logger.With("method", "GET", "url", "/api/users", "status", 200).Info("API request completed")

	// Example 5: Nested fields using groups
	fmt.Println("\n5. Grouped Fields Example:")
	userLogger := logger.Group("user-service")
	userLogger.With("user_id", 456).Info("Processing user request")
	userLogger.With("user_id", 789).Error("Failed to process user request", fmt.Errorf("validation failed"))

	// Example 6: Error logging with stack traces
	fmt.Println("\n6. Error Logging:")
	err := fmt.Errorf("sample error for demonstration")
	logger.Error("An error occurred during processing", err)

	// Example 7: Timing operations
	fmt.Println("\n7. Timing Operations:")
	start := time.Now()
	time.Sleep(100 * time.Millisecond) // Simulate some work
	logger.With("duration", time.Since(start)).Info("Operation completed")

	// Example 8: Conditional logging based on environment
	fmt.Println("\n8. Environment-based Configuration:")
	// In a real application, you might configure differently based on environment
	productionLogger := scarylog.NewLogger()
	productionLogger.Info("Running in production mode")

	// Example 9: Using logger in a function
	fmt.Println("\n9. Logger in Function:")
	processUser(logger, 1001)

	// Example 10: Logger with custom attributes
	fmt.Println("\n10. Logger with Custom Attributes:")
	customLogger := scarylog.NewLogger(scarylog.WithDefaultAttrs("service", "user-service", "version", "1.0.0"))
	customLogger.Info("Log message with custom attributes")

	fmt.Println("\n=== All Examples Completed ===")
}

// Example function demonstrating logger usage in different contexts
func processUser(logger *scarylog.Logger, userID int) {
	logger.With("user_id", userID).Info("Starting user processing")

	// Simulate some processing
	time.Sleep(50 * time.Millisecond)

	if userID%2 == 0 {
		logger.With("user_id", userID).Info("User processed successfully")
	} else {
		logger.With("user_id", userID).Warn("User requires additional validation")
	}
}
