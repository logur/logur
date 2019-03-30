/*
Package logrlog provides a github.com/go-logr/logr logger.

With logur you can easily wire the logging library of your choice into any logr logger compatible library:

	package main

	import (
		"github.com/goph/logur"
		"github.com/goph/logur/integrations/logrlog"
	)

	func main() {
		logger := logur.NewNoopLogger() // choose an actual implementation
		ilogger := logrlog.New(logger)

		// inject the logger somewhere
	}
*/
package logrlog

import (
	"fmt"

	"github.com/go-logr/logr"

	"github.com/goph/logur"
	"github.com/goph/logur/internal/keyvals"
)

// Logger is a logr Logger.
type Logger struct {
	logger       logur.Logger
	levelEnabler logur.LevelEnabler
	name         string
}

// New returns a new logr logger.
func New(logger logur.Logger) *Logger {
	l := &Logger{
		logger: logger,
	}

	if levelEnabler, ok := logger.(logur.LevelEnabler); ok {
		l.levelEnabler = levelEnabler
	}

	return l
}

// Info logs a non-error message with the given key/value pairs as context.
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	if len(keysAndValues) > 0 {
		l.logger.Info(msg, keyvals.ToMap(keysAndValues))

		return
	}

	l.logger.Info(msg)
}

// Enabled tests whether this InfoLogger is enabled.
func (l *Logger) Enabled() bool {
	if l.levelEnabler == nil {
		return true
	}

	return l.levelEnabler.LevelEnabled(logur.Info)
}

// Error logs an error, with the given message and key/value pairs as context.
func (l *Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	if len(keysAndValues) > 0 {
		l.logger.Error(msg, keyvals.ToMap(keysAndValues))

		return
	}

	l.logger.Error(msg)
}

// V returns an InfoLogger value for a specific verbosity level.
//
// Currently this function just returns the logger as is.
func (l *Logger) V(level int) logr.InfoLogger {
	// V is not properly implemented for the moment
	return l
}

// WithValues adds some key-value pairs of context to a logger.
func (l *Logger) WithValues(keysAndValues ...interface{}) logr.Logger {
	if len(keysAndValues) == 0 {
		return l
	}

	return &Logger{
		logger:       logur.WithFields(l.logger, keyvals.ToMap(keysAndValues)),
		levelEnabler: l.levelEnabler,
		name:         l.name,
	}
}

// WithName adds a new element to the logger's name.
func (l *Logger) WithName(name string) logr.Logger {
	return &Logger{
		logger:       l.logger,
		levelEnabler: l.levelEnabler,
		name:         fmt.Sprintf("%s-%s", l.name, name),
	}
}
