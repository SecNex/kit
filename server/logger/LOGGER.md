# HTTP Logger

The `logger` package provides middleware to log HTTP requests in a structured and asynchronous way. It supports file and database logging with a background goroutine for performance.

## Features

- Logs HTTP requests asynchronously
- Customizable log file path
- Logs to both file and database
- Logs response size, duration, status code, and more
- Supports Nginx-compatible log format
- Built-in middleware for easy integration with `http.Handler`

---

## Usage

### Import

```go
import "github.com/secnex/kit/server/logger"
```

### Initialization

```go
db := &database.DatabaseConnection{ /* ... */ }
httpLogger := logger.NewHTTPLogger(db, "./logs/http.log")
defer httpLogger.Stop()
```

---

### Middleware Integration

```go
mux := http.NewServeMux()
mux.Handle("/", httpLogger.LogHTTPRequest(http.HandlerFunc(yourHandler)))
```

---

### Example Output

**Console (pretty log)**

```
🔎 127.0.0.1 - - "GET /api/status HTTP/1.1" 200 512 "-" "PostmanRuntime/7.36.0" 0.005 (localhost:8080)
```

**Log file (Nginx format)**

```
127.0.0.1 - - [15/Jul/2025:10:42:30 +0000] "GET /api/status HTTP/1.1" 200 512 "-" "PostmanRuntime/7.36.0"
```

---

## Log Entry Fields

| Field          | Description                        |
| -------------- | ---------------------------------- |
| `RemoteAddr`   | Client IP address                  |
| `RequestTime`  | Time the request started           |
| `Method`       | HTTP method (GET, POST, etc.)      |
| `Path`         | URL path                           |
| `Protocol`     | HTTP version (e.g. HTTP/1.1)       |
| `StatusCode`   | HTTP response status code          |
| `ResponseSize` | Number of bytes sent in response   |
| `ResponseTime` | Duration it took to handle request |
| `UserAgent`    | Client's user agent                |
| `Referer`      | HTTP Referer header                |
| `Host`         | The requested host name            |

---

## Database Integration

Logs are stored into a database using GORM. You must define a model like:

```go
type HTTPLog struct {
	ID         uint
	IPAddress  string
	Method     string
	Path       string
	Proto      string
	Duration   time.Duration
	UserAgent  string
	Status     int
	CreatedAt  time.Time
}
```

Make sure your `database.DatabaseConnection` struct has a valid `DB *gorm.DB`.

---

## Graceful Shutdown

Use `httpLogger.Stop()` to properly terminate the logger and avoid missing logs:

```go
defer httpLogger.Stop()
```
