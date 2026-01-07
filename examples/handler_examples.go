package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/scarymovie/scarylog"
)

func main() {
	fmt.Println("=== Handler Configuration Examples ===")

	// Example 1: JSON Handler (default)
	fmt.Println("\n1. JSON Handler (default):")
	jsonLogger := scarylog.NewLogger()
	jsonLogger.Info("This uses the default JSON handler", "format", "json")

	// Example 2: Text Handler
	fmt.Println("\n2. Text Handler:")
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	textLogger := scarylog.NewLogger(scarylog.WithHandler(textHandler))
	textLogger.Info("This uses a text handler", "format", "text")
	textLogger.Warn("This is a warning in text format")

	// Example 3: Custom attributes with different handlers
	fmt.Println("\n3. Logger with Custom Attributes:")
	customLogger := scarylog.NewLogger(
		scarylog.WithHandler(textHandler),
		scarylog.WithDefaultAttrs("app", "demo", "version", "1.0"),
	)
	customLogger.Info("Application started", "port", 8080)

	// Example 4: Debug level logging
	fmt.Println("\n4. Debug Level Logging:")
	debugHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	debugLogger := scarylog.NewLogger(
		scarylog.WithHandler(debugHandler),
		scarylog.WithDefaultAttrs("component", "debug-demo"),
	)
	debugLogger.Debug("Debug message", "details", "verbose information")
	debugLogger.Info("Info message")
	debugLogger.Warn("Warning message")

	fmt.Println("\n=== Handler Examples Completed ===")
}
