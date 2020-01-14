package logur

import (
	"context"
	"fmt"
	"sync"
)

// LogEvent represents a single log event recorded by the test logger.
type LogEvent struct {
	Line   string
	Level  Level
	Fields map[string]interface{}
}

// LogEventsEqual asserts that two LogEvents are identical and returns an error with detailed information if not.
func LogEventsEqual(expected LogEvent, actual LogEvent) error {
	if expected.Level != actual.Level {
		return fmt.Errorf("expected log levels to be equal\ngot:  %s\nwant: %s", actual.Level, expected.Level)
	}

	if expected.Line != actual.Line {
		return fmt.Errorf("expected log lines to be equal\ngot:  %q\nwant: %q", actual.Line, expected.Line)
	}

	if len(expected.Fields) != len(actual.Fields) {
		return fmt.Errorf("expected log fields to be equal\ngot:  %v\nwant: %v", actual.Fields, expected.Fields)
	}

	for key, value := range expected.Fields {
		if actual.Fields[key] != value {
			return fmt.Errorf("expected log fields to be equal\ngot:  %v\nwant: %v", actual.Fields, expected.Fields)
		}
	}

	return nil
}

// TestLogger is a Logger recording every log event.
//
// Useful when you want to test behavior with an Logger, but not with LoggerContext.
// In every other cases TestLoggerFacade should be the default choice of test logger.
//
// The TestLogger is safe for concurrent use.
type TestLogger struct {
	events []LogEvent
	mu     sync.RWMutex
}

// NewTestLogger returns a new TestLogger.
//
// Deprecated: use TestLogger.
func NewTestLogger() *TestLogger {
	return &TestLogger{}
}

// Count returns the number of events recorded in the logger.
func (l *TestLogger) Count() int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return len(l.events)
}

// LastEvent returns the last recorded event in the logger (if any).
func (l *TestLogger) LastEvent() *LogEvent {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.events == nil || len(l.events) < 1 {
		return nil
	}

	event := l.events[len(l.events)-1]

	return &event
}

// Events returns all recorded events in the logger.
func (l *TestLogger) Events() []LogEvent {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.events == nil {
		return []LogEvent{}
	}

	return l.events[:len(l.events)]
}

func (l *TestLogger) recordEvent(event LogEvent) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.events = append(l.events, event)
}

func (l *TestLogger) record(level Level, msg string, varfields []map[string]interface{}) {
	var fields map[string]interface{}
	if len(varfields) > 0 {
		fields = varfields[0]
	}

	l.recordEvent(LogEvent{
		Line:   msg,
		Level:  level,
		Fields: fields,
	})
}

// Trace records a Trace level event.
func (l *TestLogger) Trace(msg string, fields ...map[string]interface{}) {
	l.record(Trace, msg, fields)
}

// Debug records a Debug level event.
func (l *TestLogger) Debug(msg string, fields ...map[string]interface{}) {
	l.record(Debug, msg, fields)
}

// Info records an Info level event.
func (l *TestLogger) Info(msg string, fields ...map[string]interface{}) {
	l.record(Info, msg, fields)
}

// Warn records a Warn level event.
func (l *TestLogger) Warn(msg string, fields ...map[string]interface{}) {
	l.record(Warn, msg, fields)
}

// Error records an Error level event.
func (l *TestLogger) Error(msg string, fields ...map[string]interface{}) {
	l.record(Error, msg, fields)
}

// TestLoggerContext is a LoggerContext recording every log event.
//
// Useful when you want to test behavior with an LoggerContext, but not with Logger.
// In every other cases TestLoggerFacade should be the default choice of test logger.
//
// The TestLoggerContext is safe for concurrent use.
type TestLoggerContext struct {
	events []LogEvent
	mu     sync.RWMutex
}

// Count returns the number of events recorded in the logger.
func (l *TestLoggerContext) Count() int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return len(l.events)
}

// LastEvent returns the last recorded event in the logger (if any).
func (l *TestLoggerContext) LastEvent() *LogEvent {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.events == nil || len(l.events) < 1 {
		return nil
	}

	event := l.events[len(l.events)-1]

	return &event
}

// Events returns all recorded events in the logger.
func (l *TestLoggerContext) Events() []LogEvent {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.events == nil {
		return []LogEvent{}
	}

	return l.events[:len(l.events)]
}

func (l *TestLoggerContext) recordEvent(event LogEvent) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.events = append(l.events, event)
}

func (l *TestLoggerContext) recordCtx(_ context.Context, level Level, msg string, varfields []map[string]interface{}) {
	var fields map[string]interface{}
	if len(varfields) > 0 {
		fields = varfields[0]
	}

	l.recordEvent(LogEvent{
		Line:   msg,
		Level:  level,
		Fields: fields,
	})
}

// TraceContext records a Trace level event.
func (l *TestLoggerContext) TraceContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.recordCtx(ctx, Trace, msg, fields)
}

// DebugContext records a Debug level event.
func (l *TestLoggerContext) DebugContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.recordCtx(ctx, Debug, msg, fields)
}

// InfoContext records an Info level event.
func (l *TestLoggerContext) InfoContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.recordCtx(ctx, Info, msg, fields)
}

// WarnContext records a Warn level event.
func (l *TestLoggerContext) WarnContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.recordCtx(ctx, Warn, msg, fields)
}

// ErrorContext records an Error level event.
func (l *TestLoggerContext) ErrorContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.recordCtx(ctx, Error, msg, fields)
}

// TestLoggerFacade is a LoggerFacade recording every log event.
//
// The TestLoggerFacade is safe for concurrent use.
type TestLoggerFacade struct {
	events []LogEvent
	mu     sync.RWMutex
}

// Count returns the number of events recorded in the logger.
func (l *TestLoggerFacade) Count() int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return len(l.events)
}

// LastEvent returns the last recorded event in the logger (if any).
func (l *TestLoggerFacade) LastEvent() *LogEvent {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.events == nil || len(l.events) < 1 {
		return nil
	}

	event := l.events[len(l.events)-1]

	return &event
}

// Events returns all recorded events in the logger.
func (l *TestLoggerFacade) Events() []LogEvent {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.events == nil {
		return []LogEvent{}
	}

	return l.events[:len(l.events)]
}

func (l *TestLoggerFacade) recordEvent(event LogEvent) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.events = append(l.events, event)
}

func (l *TestLoggerFacade) record(level Level, msg string, varfields []map[string]interface{}) {
	var fields map[string]interface{}
	if len(varfields) > 0 {
		fields = varfields[0]
	}

	l.recordEvent(LogEvent{
		Line:   msg,
		Level:  level,
		Fields: fields,
	})
}

// Trace records a Trace level event.
func (l *TestLoggerFacade) Trace(msg string, fields ...map[string]interface{}) {
	l.record(Trace, msg, fields)
}

// Debug records a Debug level event.
func (l *TestLoggerFacade) Debug(msg string, fields ...map[string]interface{}) {
	l.record(Debug, msg, fields)
}

// Info records an Info level event.
func (l *TestLoggerFacade) Info(msg string, fields ...map[string]interface{}) {
	l.record(Info, msg, fields)
}

// Warn records a Warn level event.
func (l *TestLoggerFacade) Warn(msg string, fields ...map[string]interface{}) {
	l.record(Warn, msg, fields)
}

// Error records an Error level event.
func (l *TestLoggerFacade) Error(msg string, fields ...map[string]interface{}) {
	l.record(Error, msg, fields)
}

func (l *TestLoggerFacade) recordCtx(_ context.Context, level Level, msg string, varfields []map[string]interface{}) {
	var fields map[string]interface{}
	if len(varfields) > 0 {
		fields = varfields[0]
	}

	l.recordEvent(LogEvent{
		Line:   msg,
		Level:  level,
		Fields: fields,
	})
}

// TraceContext records a Trace level event.
func (l *TestLoggerFacade) TraceContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.recordCtx(ctx, Trace, msg, fields)
}

// DebugContext records a Debug level event.
func (l *TestLoggerFacade) DebugContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.recordCtx(ctx, Debug, msg, fields)
}

// InfoContext records an Info level event.
func (l *TestLoggerFacade) InfoContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.recordCtx(ctx, Info, msg, fields)
}

// WarnContext records a Warn level event.
func (l *TestLoggerFacade) WarnContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.recordCtx(ctx, Warn, msg, fields)
}

// ErrorContext records an Error level event.
func (l *TestLoggerFacade) ErrorContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.recordCtx(ctx, Error, msg, fields)
}
