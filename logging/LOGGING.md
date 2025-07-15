# Logging

The `logging` package provides a simple, flexible, and structured logging interface for your Go application. It supports multiple log levels, configurable output formats (TEXT or JSON), and optional caller information.

## Features

- Log levels: `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`
- Output formats: plain text or JSON
- Custom log fields
- Optional caller (file:line) output
- Global logger for convenience
- Prefix support for consistent context
- Extensible configuration

## Usage

### Import

```go
import "github.com/secnex/kit/logging"
```

### Quick Start

```go
logging.SetGlobalLevel(logging.DEBUG)
logging.SetGlobalFormat(logging.TEXT)
logging.SetGlobalPrefix("MyApp")
logging.ShowGlobalCaller(true)

logging.Info("Application started")
logging.ErrorWithFieldsAndErr("Failed to load config", map[string]interface{}{
    "filename": "config.yaml",
}, fmt.Errorf("file not found"))
```

## Configuration

| Method                       | Description                               |
| ---------------------------- | ----------------------------------------- |
| `SetGlobalLevel(level)`      | Sets the minimum log level                |
| `SetGlobalFormat(format)`    | Sets log output format (`TEXT` or `JSON`) |
| `SetGlobalPrefix(prefix)`    | Adds a prefix to each log message         |
| `SetGlobalOutput(io.Writer)` | Redirects output (e.g. to file or buffer) |
| `ShowGlobalCaller(bool)`     | Enables/disables file and line info       |

## Log Levels

| Level | Description                       |
| ----- | --------------------------------- |
| DEBUG | Verbose development information   |
| INFO  | General operational messages      |
| WARN  | Potential issues or anomalies     |
| ERROR | Errors that do not stop execution |
| FATAL | Critical issues – exits program   |

## Output Formats

### TEXT (default)

```
2025-07-15 10:42:30 [INFO ] [MyApp] Application started (main.go:25)
2025-07-15 10:42:30 [ERROR] [MyApp] Failed to load config (main.go:26) - Error: file not found | filename=config.yaml
```

### JSON

```json
{
	"timestamp": "2025-07-15T10:42:30.123Z",
	"level": 3,
	"message": "[MyApp] Failed to load config",
	"caller": "main.go:26",
	"error": "file not found",
	"fields": {
		"filename": "config.yaml"
	}
}
```

## Advanced Logging

```go
logger := logging.NewLogger().
    SetLevel(logging.DEBUG).
    SetFormat(logging.JSON).
    SetPrefix("Worker").
    ShowCaller(true)

logger.WarnWithFields("Task delayed", map[string]interface{}{
    "task_id": 42,
    "retry":   true,
})
```

## Exit on Fatal

`FATAL` level logs will automatically call `os.Exit(1)` after logging the message.
