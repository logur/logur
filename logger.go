package logur

import (
	"fmt"
	"strings"
)

// Logger is a unified interface for various logging use cases and practices, including:
// 		- leveled logging
// 		- leveled formatters
// 		- structured logging
type Logger interface {
	// Trace logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	Trace(msg string, fields ...map[string]interface{})

	// Debug logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	Debug(msg string, fields ...map[string]interface{})

	// Info logs an Info event.
	//
	// General information about what's happening inside the system.
	Info(msg string, fields ...map[string]interface{})

	// Warn logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	Warn(msg string, fields ...map[string]interface{})

	// Error logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	Error(msg string, fields ...map[string]interface{})
}

// Fields is used to define structured fields which are appended to log events.
// It can be used as a shorthand for map[string]interface{}.
type Fields map[string]interface{}

// LogFunc is a function recording a log event.
type LogFunc func(msg string, fields ...map[string]interface{})

type noopLogger struct{}

// NewNoopLogger creates a no-op logger that discards all received log events.
// Useful in examples and as a fallback logger.
func NewNoopLogger() Logger {
	return &noopLogger{}
}

func (*noopLogger) Trace(msg string, fields ...map[string]interface{}) {}
func (*noopLogger) Debug(msg string, fields ...map[string]interface{}) {}
func (*noopLogger) Info(msg string, fields ...map[string]interface{})  {}
func (*noopLogger) Warn(msg string, fields ...map[string]interface{})  {}
func (*noopLogger) Error(msg string, fields ...map[string]interface{}) {}

// WithFields returns a new logger instance that attaches the given fields to every subsequent log call.
func WithFields(logger Logger, fields map[string]interface{}) Logger {
	if len(fields) == 0 {
		return logger
	}

	// Do not add a new layer
	// Create a new logger instead with the parent fields
	if l, ok := logger.(*fieldLogger); ok && len(l.fields) > 0 {
		_fields := make(map[string]interface{}, len(l.fields)+len(fields))

		for key, value := range l.fields {
			_fields[key] = value
		}

		for key, value := range fields {
			_fields[key] = value
		}

		fields = _fields
		logger = l.logger
	}

	l := &fieldLogger{logger: logger, fields: fields}

	if levelEnabler, ok := logger.(LevelEnabler); ok {
		l.levelEnabler = levelEnabler
	}

	return l
}

// fieldLogger holds a context and passes it to the underlying logger when a log event is recorded.
type fieldLogger struct {
	logger       Logger
	fields       map[string]interface{}
	levelEnabler LevelEnabler
}

// Trace implements the logur.Logger interface.
func (l *fieldLogger) Trace(msg string, fields ...map[string]interface{}) {
	l.log(Trace, l.logger.Trace, msg, fields)
}

// Debug implements the logur.Logger interface.
func (l *fieldLogger) Debug(msg string, fields ...map[string]interface{}) {
	l.log(Debug, l.logger.Debug, msg, fields)
}

// Info implements the logur.Logger interface.
func (l *fieldLogger) Info(msg string, fields ...map[string]interface{}) {
	l.log(Info, l.logger.Info, msg, fields)
}

// Warn implements the logur.Logger interface.
func (l *fieldLogger) Warn(msg string, fields ...map[string]interface{}) {
	l.log(Warn, l.logger.Warn, msg, fields)
}

// Error implements the logur.Logger interface.
func (l *fieldLogger) Error(msg string, fields ...map[string]interface{}) {
	l.log(Error, l.logger.Error, msg, fields)
}

// log deduplicates some field logger code.
func (l *fieldLogger) log(level Level, logFunc LogFunc, msg string, fields []map[string]interface{}) {
	if !l.levelEnabled(level) {
		return
	}

	var f = l.fields
	if len(fields) > 0 {
		f = l.mergeFields(fields[0])
	}

	logFunc(msg, f)
}

func (l *fieldLogger) mergeFields(fields map[string]interface{}) map[string]interface{} {
	if len(fields) == 0 { // Not having any fields passed to the log function has a higher chance
		return l.fields
	}

	if len(l.fields) == 0 { // This is possible too, but has a much lower probability
		return fields
	}

	f := make(map[string]interface{}, len(fields)+len(l.fields))

	for key, value := range l.fields {
		f[key] = value
	}

	for key, value := range fields {
		f[key] = value
	}

	return f
}

func (l *fieldLogger) levelEnabled(level Level) bool {
	if l.levelEnabler != nil {
		return l.levelEnabler.LevelEnabled(level)
	}

	return true
}

// PrintLogger logs messages with fmt.Print* function semantics.
type PrintLogger struct {
	logger LogFunc
}

// NewPrintLogger returns a new PrintLogger.
func NewPrintLogger(logger LogFunc) *PrintLogger {
	return &PrintLogger{
		logger: logger,
	}
}

// NewErrorPrintLogger returns a new PrintLogger that logs everything on error level.
func NewErrorPrintLogger(logger Logger) *PrintLogger {
	return NewPrintLogger(LevelFunc(logger, Error))
}

// Print logs a line with fmt.Print semantics.
func (l *PrintLogger) Print(v ...interface{}) {
	l.logger(fmt.Sprint(v...))
}

// Println logs a line with fmt.Println semantics.
func (l *PrintLogger) Println(v ...interface{}) {
	l.logger(strings.TrimSuffix(fmt.Sprintln(v...), "\n"))
}

// Printf logs a line with fmt.Printf semantics.
func (l *PrintLogger) Printf(format string, args ...interface{}) {
	l.logger(fmt.Sprintf(format, args...))
}

// MessageLogger simplifies the Logger interface by removing the second context parameter.
// Useful when there is no need for contextual logging.
type MessageLogger struct {
	logger Logger
}

// NewMessageLogger returns a new MessageLogger instance.
func NewMessageLogger(logger Logger) *MessageLogger {
	return &MessageLogger{
		logger: logger,
	}
}

// Trace logs a Trace level event.
func (l *MessageLogger) Trace(msg string) {
	l.logger.Trace(msg)
}

// Debug logs a Debug level event.
func (l *MessageLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

// Info logs a Info level event.
func (l *MessageLogger) Info(msg string) {
	l.logger.Info(msg)
}

// Warn logs a Warn level event.
func (l *MessageLogger) Warn(msg string) {
	l.logger.Warn(msg)
}

// Error logs a Error level event.
func (l *MessageLogger) Error(msg string) {
	l.logger.Error(msg)
}
