package logur

import (
	"fmt"
	"sync"
)

// LogEvent represents a single log event recorded by the test logger.
type LogEvent struct {
	Line   string
	Level  Level
	Fields map[string]interface{}
}

// AssertLogEventsEqual asserts that two LogEvents are identical and returns an error with detailed information if not.
func AssertLogEventsEqual(expected LogEvent, actual LogEvent) error {
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

// TestLogger is a simple stub for the logger interface.
//
// The TestLogger is safe for concurrent use.
type TestLogger struct {
	events []LogEvent
	mu     sync.RWMutex
}

// NewTestLogger returns a new TestLogger.
func NewTestLogger() *TestLogger {
	return &TestLogger{}
}

// Count returns the number of events recorded in the logger.
func (l *TestLogger) Count() int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.events == nil {
		return 0
	}

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

func (l *TestLogger) record(level Level, msg string, fields map[string]interface{}) {
	l.recordEvent(LogEvent{
		Line:   msg,
		Level:  level,
		Fields: fields,
	})
}

// Trace records a Trace level event.
func (l *TestLogger) Trace(msg string, fields map[string]interface{}) {
	l.record(Trace, msg, fields)
}

// Debug records a Debug level event.
func (l *TestLogger) Debug(msg string, fields map[string]interface{}) {
	l.record(Debug, msg, fields)
}

// Info records a Info level event.
func (l *TestLogger) Info(msg string, fields map[string]interface{}) {
	l.record(Info, msg, fields)
}

// Warn records a Warn level event.
func (l *TestLogger) Warn(msg string, fields map[string]interface{}) {
	l.record(Warn, msg, fields)
}

// Error records a Error level event.
func (l *TestLogger) Error(msg string, fields map[string]interface{}) {
	l.record(Error, msg, fields)
}
