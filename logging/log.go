package logging

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// LogLevel represents the severity level of a log entry
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// String returns the string representation of the log level
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// LogFormat represents the output format for logs
type LogFormat int

const (
	TEXT LogFormat = iota
	JSON
)

// LogEntry represents a single log entry
type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     LogLevel               `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Caller    string                 `json:"caller,omitempty"`
	Error     string                 `json:"error,omitempty"`
}

// Logger represents the main logger instance
type Logger struct {
	level      LogLevel
	format     LogFormat
	output     io.Writer
	showCaller bool
	prefix     string
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		level:      INFO,
		format:     TEXT,
		output:     os.Stdout,
		showCaller: false,
		prefix:     "",
	}
}

// NewLoggerWithConfig creates a new logger with specific configuration
func NewLoggerWithConfig(level LogLevel, format LogFormat, output io.Writer) *Logger {
	return &Logger{
		level:      level,
		format:     format,
		output:     output,
		showCaller: false,
		prefix:     "",
	}
}

// SetLevel sets the minimum log level
func (l *Logger) SetLevel(level LogLevel) *Logger {
	l.level = level
	return l
}

// SetFormat sets the output format (TEXT or JSON)
func (l *Logger) SetFormat(format LogFormat) *Logger {
	l.format = format
	return l
}

// SetOutput sets the output writer
func (l *Logger) SetOutput(output io.Writer) *Logger {
	l.output = output
	return l
}

// ShowCaller enables/disables caller information in logs
func (l *Logger) ShowCaller(show bool) *Logger {
	l.showCaller = show
	return l
}

// SetPrefix sets a prefix for all log messages
func (l *Logger) SetPrefix(prefix string) *Logger {
	l.prefix = prefix
	return l
}

// getCaller returns the caller information
func (l *Logger) getCaller() string {
	if !l.showCaller {
		return ""
	}

	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return "unknown"
	}

	// Get just the filename, not the full path
	parts := strings.Split(file, "/")
	filename := parts[len(parts)-1]

	return fmt.Sprintf("%s:%d", filename, line)
}

// log is the internal logging method
func (l *Logger) log(level LogLevel, message string, fields map[string]interface{}, err error) {
	if level < l.level {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Fields:    fields,
		Caller:    l.getCaller(),
	}

	if err != nil {
		entry.Error = err.Error()
	}

	if l.prefix != "" {
		entry.Message = fmt.Sprintf("[%s] %s", l.prefix, entry.Message)
	}

	l.write(entry)

	// Exit on fatal
	if level == FATAL {
		os.Exit(1)
	}
}

// write outputs the log entry in the specified format
func (l *Logger) write(entry LogEntry) {
	switch l.format {
	case JSON:
		l.writeJSON(entry)
	default:
		l.writeText(entry)
	}
}

// writeText outputs the log entry in text format
func (l *Logger) writeText(entry LogEntry) {
	timestamp := entry.Timestamp.Format("2006-01-02 15:04:05")
	level := fmt.Sprintf("%-5s", entry.Level.String())

	message := fmt.Sprintf("%s [%s] %s", timestamp, level, entry.Message)

	if entry.Caller != "" {
		message = fmt.Sprintf("%s (%s)", message, entry.Caller)
	}

	if entry.Error != "" {
		message = fmt.Sprintf("%s - Error: %s", message, entry.Error)
	}

	if len(entry.Fields) > 0 {
		fieldsStr := ""
		for k, v := range entry.Fields {
			fieldsStr += fmt.Sprintf(" %s=%v", k, v)
		}
		message = fmt.Sprintf("%s |%s", message, fieldsStr)
	}

	fmt.Fprintln(l.output, message)
}

// writeJSON outputs the log entry in JSON format
func (l *Logger) writeJSON(entry LogEntry) {
	data, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Error marshaling log entry: %v", err)
		return
	}
	fmt.Fprintln(l.output, string(data))
}

// Debug logs a debug message
func (l *Logger) Debug(message string) {
	l.log(DEBUG, message, nil, nil)
}

// DebugWithFields logs a debug message with additional fields
func (l *Logger) DebugWithFields(message string, fields map[string]interface{}) {
	l.log(DEBUG, message, fields, nil)
}

// Info logs an info message
func (l *Logger) Info(message string) {
	l.log(INFO, message, nil, nil)
}

// InfoWithFields logs an info message with additional fields
func (l *Logger) InfoWithFields(message string, fields map[string]interface{}) {
	l.log(INFO, message, fields, nil)
}

// Warn logs a warning message
func (l *Logger) Warn(message string) {
	l.log(WARN, message, nil, nil)
}

// WarnWithFields logs a warning message with additional fields
func (l *Logger) WarnWithFields(message string, fields map[string]interface{}) {
	l.log(WARN, message, fields, nil)
}

// Error logs an error message
func (l *Logger) Error(message string) {
	l.log(ERROR, message, nil, nil)
}

// ErrorWithFields logs an error message with additional fields
func (l *Logger) ErrorWithFields(message string, fields map[string]interface{}) {
	l.log(ERROR, message, fields, nil)
}

// ErrorWithErr logs an error message with an error object
func (l *Logger) ErrorWithErr(message string, err error) {
	l.log(ERROR, message, nil, err)
}

// ErrorWithFieldsAndErr logs an error message with fields and an error object
func (l *Logger) ErrorWithFieldsAndErr(message string, fields map[string]interface{}, err error) {
	l.log(ERROR, message, fields, err)
}

// Fatal logs a fatal message and exits the program
func (l *Logger) Fatal(message string) {
	l.log(FATAL, message, nil, nil)
}

// FatalWithFields logs a fatal message with additional fields and exits
func (l *Logger) FatalWithFields(message string, fields map[string]interface{}) {
	l.log(FATAL, message, fields, nil)
}

// FatalWithErr logs a fatal message with an error object and exits
func (l *Logger) FatalWithErr(message string, err error) {
	l.log(FATAL, message, nil, err)
}

// Global logger instance
var globalLogger = NewLogger()

// SetGlobalLevel sets the log level for the global logger
func SetGlobalLevel(level LogLevel) {
	globalLogger.SetLevel(level)
}

// SetGlobalFormat sets the format for the global logger
func SetGlobalFormat(format LogFormat) {
	globalLogger.SetFormat(format)
}

// SetGlobalOutput sets the output for the global logger
func SetGlobalOutput(output io.Writer) {
	globalLogger.SetOutput(output)
}

// SetGlobalPrefix sets the prefix for the global logger
func SetGlobalPrefix(prefix string) {
	globalLogger.SetPrefix(prefix)
}

// ShowGlobalCaller enables/disables caller info for the global logger
func ShowGlobalCaller(show bool) {
	globalLogger.ShowCaller(show)
}

// Global convenience functions
func Debug(message string) {
	globalLogger.Debug(message)
}

func DebugWithFields(message string, fields map[string]interface{}) {
	globalLogger.DebugWithFields(message, fields)
}

func Info(message string) {
	globalLogger.Info(message)
}

func InfoWithFields(message string, fields map[string]interface{}) {
	globalLogger.InfoWithFields(message, fields)
}

func Warn(message string) {
	globalLogger.Warn(message)
}

func WarnWithFields(message string, fields map[string]interface{}) {
	globalLogger.WarnWithFields(message, fields)
}

func Error(message string) {
	globalLogger.Error(message)
}

func ErrorWithFields(message string, fields map[string]interface{}) {
	globalLogger.ErrorWithFields(message, fields)
}

func ErrorWithErr(message string, err error) {
	globalLogger.ErrorWithErr(message, err)
}

func ErrorWithFieldsAndErr(message string, fields map[string]interface{}, err error) {
	globalLogger.ErrorWithFieldsAndErr(message, fields, err)
}

func Fatal(message string) {
	globalLogger.Fatal(message)
}

func FatalWithFields(message string, fields map[string]interface{}) {
	globalLogger.FatalWithFields(message, fields)
}

func FatalWithErr(message string, err error) {
	globalLogger.FatalWithErr(message, err)
}
