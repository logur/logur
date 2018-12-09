// Package simplelogadapter exposes a simpler interface than logur.Logger and makes logging library integration easier.
package simplelogadapter

import (
	"fmt"

	"github.com/goph/logur"
)

// Logger is a subset of logur.Logger to make logging library integration easier.
// The rest of the methods are implemented by the adapter.
type Logger interface {
	// Trace logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	Trace(args ...interface{})

	// Debug logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	Debug(args ...interface{})

	// Info logs an Info event.
	//
	// General information about what's happening inside the system.
	Info(args ...interface{})

	// Warn logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	Warn(args ...interface{})

	// Error logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	Error(args ...interface{})

	// WithFields appends structured fields to a new (child) logger instance.
	WithFields(fields Fields) Logger
}

// Fields is used to define structured fields which are appended to log events.
type Fields map[string]interface{}

type adapter struct {
	Logger
}

func New(logger Logger) logur.Logger {
	return &adapter{logger}
}

func (a *adapter) Traceln(args ...interface{}) {
	a.Trace(fmt.Sprintln(args...))
}

func (a *adapter) Debugln(args ...interface{}) {
	a.Debug(fmt.Sprintln(args...))
}

func (a *adapter) Infoln(args ...interface{}) {
	a.Info(fmt.Sprintln(args...))
}

func (a *adapter) Warnln(args ...interface{}) {
	a.Warn(fmt.Sprintln(args...))
}

func (a *adapter) Errorln(args ...interface{}) {
	a.Error(fmt.Sprintln(args...))
}

func (a *adapter) Tracef(format string, args ...interface{}) {
	a.Trace(fmt.Sprintf(format, args...))
}

func (a *adapter) Debugf(format string, args ...interface{}) {
	a.Debug(fmt.Sprintf(format, args...))
}

func (a *adapter) Infof(format string, args ...interface{}) {
	a.Info(fmt.Sprintf(format, args...))
}

func (a *adapter) Warnf(format string, args ...interface{}) {
	a.Warn(fmt.Sprintf(format, args...))
}

func (a *adapter) Errorf(format string, args ...interface{}) {
	a.Error(fmt.Sprintf(format, args...))
}

func (a *adapter) WithFields(fields logur.Fields) logur.Logger {
	return &adapter{a.Logger.WithFields(Fields(fields))}
}
