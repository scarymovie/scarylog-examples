package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/scarymovie/scarylog"
)

// Example of a service that uses logging throughout its lifecycle
type UserService struct {
	logger *scarylog.Logger
}

func NewUserService() *UserService {
	// Create a logger with default attributes for this service
	serviceLogger := scarylog.NewLogger(
		scarylog.WithDefaultAttrs("service", "user-service"),
		scarylog.WithLevel(slog.LevelDebug),
		scarylog.WithGroup("ctx"),
	)

	return &UserService{
		logger: serviceLogger,
	}
}

func (us *UserService) CreateUser(ctx context.Context, userID int, username string) error {
	// Create a child logger with user-specific context
	userLogger := us.logger.With("user_id", userID, "username", username)
	userLogger.Info("Creating new user")

	// Simulate some work
	time.Sleep(10 * time.Millisecond)

	// Log successful creation
	userLogger.Info("User created successfully")

	// Simulate potential error scenario
	if userID < 0 {
		err := fmt.Errorf("invalid user ID: %d", userID)
		userLogger.Error("Failed to create user", err)
		return err
	}

	return nil
}

func (us *UserService) GetUser(ctx context.Context, userID int) (map[string]interface{}, error) {
	logger := us.logger.With("user_id", userID)
	logger.Info("Fetching user details")

	// Simulate some work
	time.Sleep(5 * time.Millisecond)

	// Simulate finding or not finding the user
	if userID%2 == 0 {
		userData := map[string]interface{}{
			"id":      userID,
			"name":    fmt.Sprintf("User%d", userID),
			"email":   fmt.Sprintf("user%d@example.com", userID),
			"created": time.Now(),
		}
		logger.Info("User found", "has_data", true)
		return userData, nil
	} else {
		err := fmt.Errorf("user not found: %d", userID)
		logger.Error("User not found", err)
		return nil, err
	}
}

// Example of middleware that adds request-scoped logging
func LoggingMiddleware(next func(context.Context)) func(context.Context) {
	return func(ctx context.Context) {
		requestID := fmt.Sprintf("req-%d", time.Now().Unix())

		// Create a logger for this request and store it in context
		reqLogger := scarylog.NewLogger(
			scarylog.WithDefaultAttrs(
				"request_id", requestID,
				"service", "api-gateway",
			),
		)

		// Store logger in context
		ctx = scarylog.ToContext(ctx, reqLogger)

		reqLogger.Info("Request started")

		// Call the next handler
		next(ctx)

		reqLogger.Info("Request completed")
	}
}

func main() {
	fmt.Println("=== Advanced ScaryLog Examples ===")

	// Example 1: Service with structured logging
	fmt.Println("\n1. Service with Structured Logging:")
	userService := NewUserService()

	ctx := context.Background()

	// Create a few users
	usersToCreate := []int{1001, 1002, -1} // -1 will trigger an error

	for _, userID := range usersToCreate {
		err := userService.CreateUser(ctx, userID, fmt.Sprintf("user_%d", userID))
		if err != nil {
			fmt.Printf("  Error creating user %d: %v\n", userID, err)
		}
	}

	// Get some users
	fmt.Println("\n2. Retrieving Users:")
	usersToGet := []int{1002, 1003}

	for _, userID := range usersToGet {
		userData, err := userService.GetUser(ctx, userID)
		if err != nil {
			fmt.Printf("  Error getting user %d: %v\n", userID, err)
		} else {
			fmt.Printf("  Retrieved user %d with email: %s\n",
				userID, userData["email"])
		}
	}

	// Example 3: Context-based logging with middleware pattern
	fmt.Println("\n3. Context-Based Logging with Middleware:")

	// Define a handler function
	handlerFunc := func(ctx context.Context) {
		logger := scarylog.FromContext(ctx)
		logger.Info("Processing business logic")

		// Simulate some work
		time.Sleep(8 * time.Millisecond)

		logger.Info("Business logic completed", "result", "success")
	}

	// Wrap the handler with logging middleware
	wrappedHandler := LoggingMiddleware(handlerFunc)

	// Execute the wrapped handler
	wrappedHandler(ctx)

	// Example 4: Different output formats
	fmt.Println("\n4. Different Logger Configurations:")

	// Console-style logger (text format)
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	textLogger := scarylog.NewLogger(
		scarylog.WithHandler(consoleHandler),
		scarylog.WithDefaultAttrs("example", "text-format"),
	)

	textLogger.Info("This is a text-formatted log message")
	textLogger.Warn("This warning appears in text format")

	// Example 5: Conditional logging based on environment
	fmt.Println("\n5. Environment-Based Configuration:")

	// In a real app, this would come from environment variables
	isProduction := false

	var appLogger *scarylog.Logger
	if isProduction {
		// Production logger - only errors and warnings
		appLogger = scarylog.NewLogger(
			scarylog.WithLevel(slog.LevelWarn),
			scarylog.WithDefaultAttrs("env", "production"),
		)
	} else {
		// Development logger - all levels
		appLogger = scarylog.NewLogger(
			scarylog.WithLevel(slog.LevelDebug),
			scarylog.WithDefaultAttrs("env", "development"),
		)
	}

	appLogger.Info("Application started in development mode")
	appLogger.Debug("Detailed debug information")

	fmt.Println("\n=== Advanced Examples Completed ===")
}
