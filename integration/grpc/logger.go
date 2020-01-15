/*
Package grpc provides a V2 gRPC logger.

gRPC operates with a globally configured logger that implements the google.golang.org/grpc/grpclog.LoggerV2 interface.

With logur you can easily wire the logging library of your choice into gRPC:

	package main

	import (
		"logur.dev/logur"
		grpcintegration "logur.dev/logur/integration/grpc"
		"google.golang.org/grpc/grpclog"
	)

	func main() {
		logger := logur.NewNoopLogger() // choose an actual implementation
		grpclog.SetLoggerV2(grpcintegration.New(logger))
	}
*/
package grpc

import (
	"fmt"
	"strings"

	"logur.dev/logur"
)

// Logger is a V2 gRPC logger.
type Logger struct {
	logger       logur.Logger
	levelEnabler logur.LevelEnabler
}

// New returns a new V2 gRPC logger.
func New(logger logur.Logger) *Logger {
	l := &Logger{
		logger: logger,
	}

	if levelEnabler, ok := logger.(logur.LevelEnabler); ok {
		l.levelEnabler = levelEnabler
	}

	return l
}

// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}

// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Infoln(args ...interface{}) {
	l.logger.Info(strings.TrimSuffix(fmt.Sprintln(args...), "\n"))
}

// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}

// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Warning(args ...interface{}) {
	l.logger.Warn(fmt.Sprint(args...))
}

// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Warningln(args ...interface{}) {
	l.logger.Warn(strings.TrimSuffix(fmt.Sprintln(args...), "\n"))
}

// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...))
}

// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Errorln(args ...interface{}) {
	l.logger.Error(strings.TrimSuffix(fmt.Sprintln(args...), "\n"))
}

// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...))
}

// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *Logger) Fatalln(args ...interface{}) {
	l.logger.Error(strings.TrimSuffix(fmt.Sprintln(args...), "\n"))
}

// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

// V reports whether verbosity level l is at least the requested verbose level.
func (l *Logger) V(level int) bool {
	if l.levelEnabler == nil {
		return true
	}

	if level == 3 { // fatal level
		level = 2
	}

	// grpc log doesn't have trace and debug levels
	return l.levelEnabler.LevelEnabled(logur.Level(level + 2))
}
