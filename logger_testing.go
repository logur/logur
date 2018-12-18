package logur

import (
	"sync"
)

// LogEvent represents a single log event recorded by the test logger.
type LogEvent struct {
	Line   string
	Level  Level
	Fields map[string]interface{}
}

// TestLogger is a simple stub for the logger interface.
//
// The TestLogger is safe for concurrent use.
type TestLogger struct {
	events []LogEvent
	fields map[string]interface{}
	mu     sync.RWMutex

	parent *TestLogger
}

// NewTestLogger returns a new TestLogger.
func NewTestLogger() *TestLogger {
	return &TestLogger{}
}

// Count returns the number of events recorded in the logger.
func (l *TestLogger) Count() int {
	if l.parent != nil {
		return l.parent.Count()
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.events == nil {
		return 0
	}

	return len(l.events)
}

// LastEvent returns the last recorded event in the logger (if any).
func (l *TestLogger) LastEvent() *LogEvent {
	if l.parent != nil {
		return l.parent.LastEvent()
	}

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
	if l.parent != nil {
		return l.parent.Events()
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.events == nil {
		return []LogEvent{}
	}

	return l.events[:len(l.events)]
}

func (l *TestLogger) recordEvent(event LogEvent) {
	if l.parent != nil {
		l.parent.recordEvent(event)

		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.events = append(l.events, event)
}

func (l *TestLogger) record(level Level, msg string) {
	l.recordEvent(LogEvent{
		Line:   msg,
		Level:  level,
		Fields: l.fields,
	})
}

// Trace records a Trace level event.
func (l *TestLogger) Trace(msg string) {
	l.record(Trace, msg)
}

// Debug records a Debug level event.
func (l *TestLogger) Debug(msg string) {
	l.record(Debug, msg)
}

// Info records a Info level event.
func (l *TestLogger) Info(msg string) {
	l.record(Info, msg)
}

// Warn records a Warn level event.
func (l *TestLogger) Warn(msg string) {
	l.record(Warn, msg)
}

// Error records a Error level event.
func (l *TestLogger) Error(msg string) {
	l.record(Error, msg)
}

// WithFields returns a new TestLogger with the appended fields.
func (l *TestLogger) WithFields(fields map[string]interface{}) Logger {
	var f = l.fields

	if f == nil {
		f = make(Fields)
	}

	for key, value := range fields {
		f[key] = value
	}

	return &TestLogger{
		fields: f,
		parent: l,
	}
}
