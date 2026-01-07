# ScaryLog Examples

This repository contains examples demonstrating the usage of the `scarylog` logging library, which is a wrapper around Go's standard `slog` package.

## Overview

ScaryLog provides a convenient wrapper around Go's structured logger (`slog`) with additional features:

- Simple and easy-to-use API
- Colorized console output for improved readability
- Flexible configuration through functional options
- Default attributes for consistent context
- Context integration for distributed tracing

## Installation

```bash
go get github.com/scarymovie/scarylog
```

## Examples

This repository contains several examples of using the scarylog library:

### Basic Examples

The `main.go` file contains comprehensive basic examples of using the scarylog library:

#### 1. Basic Initialization and Usage

```go
logger := scarylog.NewLogger()
logger.Info("Application started")
logger.Warn("This is a warning message")
logger.Error("An error occurred", err)
```

#### 2. Configuring Log Levels

```go
// Enable debug logging
debugLogger := scarylog.NewLogger(scarylog.WithLevel(slog.LevelDebug))
debugLogger.Debug("This is a debug message", "key", "value")
```

#### 3. Context Integration

```go
ctx := context.Background()
ctx = scarylog.ToContext(ctx, logger)
scarylog.FromContext(ctx).Info("Message using logger from context")
```

#### 4. Structured Logging with Fields

```go
logger.With("user_id", 123, "action", "login").Info("User logged in")
logger.With("method", "GET", "url", "/api/users", "status", 200).Info("API request completed")
```

#### 5. Grouped Fields

```go
userLogger := logger.Group("user-service")
userLogger.With("user_id", 456).Info("Processing user request")
```

#### 6. Error Handling

```go
err := fmt.Errorf("sample error for demonstration")
logger.Error("An error occurred during processing", err)
```

#### 7. Default Attributes

```go
customLogger := scarylog.NewLogger(scarylog.WithDefaultAttrs(
    "service", "user-service",
    "version", "1.0.0"
))
customLogger.Info("Log message with custom attributes")
```

### Advanced Examples

More advanced examples can be found in the `examples/` directory, including:

- Service-oriented logging patterns
- Middleware implementations with contextual logging
- Different output formats
- Environment-based configurations

## Features Demonstrated

- **Basic logging**: Info, Warn, Error, and Debug levels
- **Structured logging**: Key-value pairs for better context
- **Context integration**: Storing and retrieving loggers from context.Context
- **Groups**: Organizing related fields together
- **Error handling**: Proper error logging with stack traces
- **Custom attributes**: Default attributes for consistent logging
- **Timing operations**: Logging duration of operations
- **Testing**: How to test code that uses the logging library

## Running the Examples

### Basic Examples
```bash
go run main.go
```

### Advanced Examples
```bash
cd examples/
go run advanced_example.go
```

### Running Tests
```bash
go test -v
```

This will run all the examples and show the output in JSON format (the default for scarylog which uses slog underneath).