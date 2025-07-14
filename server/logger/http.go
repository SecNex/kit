package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/secnex/kit/database"
	"github.com/secnex/kit/models"
)

type HTTPLogger struct {
	Database *database.DatabaseConnection
	Path     string
	logChan  chan HTTPLogEntry
	done     chan bool
}

type HTTPLogEntry struct {
	Host         string
	RemoteAddr   string
	RequestTime  time.Time
	Method       string
	Path         string
	Protocol     string
	StatusCode   int
	ResponseSize int64
	ResponseTime time.Duration
	UserAgent    string
	Referer      string
}

func NewHTTPLogger(database *database.DatabaseConnection, path string) *HTTPLogger {
	logger := &HTTPLogger{
		Database: database,
		Path:     path,
		logChan:  make(chan HTTPLogEntry, 100), // Buffered channel
		done:     make(chan bool),
	}

	// Start the background logging goroutine
	go logger.processLogs()

	return logger
}

// processLogs runs in a separate goroutine and processes all log entries
func (l *HTTPLogger) processLogs() {
	for {
		select {
		case entry := <-l.logChan:
			l.logToFileSync(entry)
			l.logToDatabaseSync(entry)
		case <-l.done:
			return
		}
	}
}

// Stop gracefully stops the logger
func (l *HTTPLogger) Stop() {
	close(l.done)
}

// logToFileSync is the synchronous version for internal use
func (l *HTTPLogger) logToFileSync(entry HTTPLogEntry) {
	file, err := os.OpenFile(l.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(NginxFormatHTTPLog(entry) + "\n")
	if err != nil {
		log.Printf("Error writing to log file: %v", err)
	}
}

// logToDatabaseSync is the synchronous version for internal use
func (l *HTTPLogger) logToDatabaseSync(entry HTTPLogEntry) {
	err := l.Database.DB.Create(&models.HTTPLog{
		IPAddress: entry.RemoteAddr,
		Method:    entry.Method,
		Path:      entry.Path,
		Proto:     entry.Protocol,
		Duration:  entry.ResponseTime,
		UserAgent: entry.UserAgent,
		Status:    entry.StatusCode,
	}).Error

	if err != nil {
		log.Printf("Error logging to database: %v", err)
	}
}

func FormatHTTPLog(entry HTTPLogEntry) string {
	return fmt.Sprintf("🔎 %s - - \"%s %s %s\" %d %d \"%s\" \"%s\" %.3f (%s)",
		entry.RemoteAddr,
		entry.Method,
		entry.Path,
		entry.Protocol,
		entry.StatusCode,
		entry.ResponseSize,
		entry.Referer,
		entry.UserAgent,
		entry.ResponseTime.Seconds(),
		entry.Host,
	)
}

func NginxFormatHTTPLog(entry HTTPLogEntry) string {
	timestamp := entry.RequestTime.Format("02/Jan/2006:15:04:05 -0700")

	referer := entry.Referer
	if referer == "" {
		referer = "-"
	}

	return fmt.Sprintf("%s - - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\"",
		entry.RemoteAddr,
		timestamp,
		entry.Method,
		entry.Path,
		entry.Protocol,
		entry.StatusCode,
		entry.ResponseSize,
		referer,
		entry.UserAgent,
	)
}

func (l *HTTPLogger) LogHTTPRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		entry := HTTPLogEntry{
			RemoteAddr:   r.RemoteAddr,
			RequestTime:  start,
			Method:       r.Method,
			Path:         r.URL.Path,
			Protocol:     r.Proto,
			StatusCode:   rw.statusCode,
			ResponseSize: rw.size,
			ResponseTime: time.Since(start),
			UserAgent:    r.UserAgent(),
			Referer:      r.Referer(),
			Host:         r.Host,
		}

		fmt.Println(FormatHTTPLog(entry))

		l.logChan <- entry
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int64
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += int64(size)
	return size, err
}
