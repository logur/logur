package example

import "github.com/goph/logur"

// Logger is the fundamental interface for all log operations.
type Logger interface {
	Trace(msg string, fields map[string]interface{})
	Debug(msg string, fields map[string]interface{})
	Info(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})

	// WithFields annotates a logger with some context.
	WithFields(fields map[string]interface{}) Logger
}

// Adapter wraps a logur logger and exposes it under a custom interface.
type Adapter struct {
	logger logur.Logger
}

// NewAdapter returns a new Logger instance.
func NewAdapter(logger logur.Logger) *Adapter {
	return &Adapter{
		logger: logger,
	}
}

// Trace logs a trace event.
func (l *Adapter) Trace(msg string, fields map[string]interface{}) {
	l.logger.Trace(msg, fields)
}

// Debug logs a debug event.
func (l *Adapter) Debug(msg string, fields map[string]interface{}) {
	l.logger.Debug(msg, fields)
}

// Info logs an info event.
func (l *Adapter) Info(msg string, fields map[string]interface{}) {
	l.logger.Info(msg, fields)
}

// Warn logs a warning event.
func (l *Adapter) Warn(msg string, fields map[string]interface{}) {
	l.logger.Warn(msg, fields)
}

// Error logs an error event.
func (l *Adapter) Error(msg string, fields map[string]interface{}) {
	l.logger.Error(msg, fields)
}

// WithFields annotates a logger with some context.
func (l *Adapter) WithFields(fields map[string]interface{}) Logger {
	return &Adapter{logger: logur.WithFields(l.logger, fields)}
}
