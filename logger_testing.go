package logur

import (
	"fmt"
	"sync"
)

// LogEvent represents a single log event recorded by the test logger.
type LogEvent struct {
	Line    string
	RawLine []interface{}
	Level   Level
	Fields  Fields
}

// TestLogger is a simple stub for the logger interface.
//
// The TestLogger is safe for concurrent use.
type TestLogger struct {
	events []LogEvent
	fields Fields
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

func (l *TestLogger) record(level Level, args []interface{}) {
	l.recordEvent(LogEvent{
		Line:    fmt.Sprint(args...),
		RawLine: args,
		Level:   level,
		Fields:  l.fields,
	})
}

// Trace records a Trace level event.
func (l *TestLogger) Trace(args ...interface{}) {
	l.record(TraceLevel, args)
}

// Debug records a Debug level event.
func (l *TestLogger) Debug(args ...interface{}) {
	l.record(DebugLevel, args)
}

// Info records a Info level event.
func (l *TestLogger) Info(args ...interface{}) {
	l.record(InfoLevel, args)
}

// Warn records a Warn level event.
func (l *TestLogger) Warn(args ...interface{}) {
	l.record(WarnLevel, args)
}

// Error records a Error level event.
func (l *TestLogger) Error(args ...interface{}) {
	l.record(ErrorLevel, args)
}

func (l *TestLogger) recordln(level Level, args []interface{}) {
	l.recordEvent(LogEvent{
		Line:    fmt.Sprintln(args...),
		RawLine: args,
		Level:   level,
		Fields:  l.fields,
	})
}

// Traceln records a Trace level event using fmt.Println characteristics.
func (l *TestLogger) Traceln(args ...interface{}) {
	l.recordln(TraceLevel, args)
}

// Debugln records a Debug level event using fmt.Println characteristics.
func (l *TestLogger) Debugln(args ...interface{}) {
	l.recordln(DebugLevel, args)
}

// Infoln records a Info level event using fmt.Println characteristics.
func (l *TestLogger) Infoln(args ...interface{}) {
	l.recordln(InfoLevel, args)
}

// Warnln records a Warn level event using fmt.Println characteristics.
func (l *TestLogger) Warnln(args ...interface{}) {
	l.recordln(WarnLevel, args)
}

// Errorln records a Error level event using fmt.Println characteristics.
func (l *TestLogger) Errorln(args ...interface{}) {
	l.recordln(ErrorLevel, args)
}

func (l *TestLogger) recordf(level Level, format string, args []interface{}) {
	l.recordEvent(LogEvent{
		Line:    fmt.Sprintf(format, args...),
		RawLine: append([]interface{}{format}, args...),
		Level:   level,
		Fields:  l.fields,
	})
}

// Tracef records a Trace level event with a formatted message.
func (l *TestLogger) Tracef(format string, args ...interface{}) {
	l.recordf(TraceLevel, format, args)
}

// Debugf records a Debug level event with a formatted message.
func (l *TestLogger) Debugf(format string, args ...interface{}) {
	l.recordf(DebugLevel, format, args)
}

// Infof records a Info level event with a formatted message.
func (l *TestLogger) Infof(format string, args ...interface{}) {
	l.recordf(InfoLevel, format, args)
}

// Warnf records a Warn level event with a formatted message.
func (l *TestLogger) Warnf(format string, args ...interface{}) {
	l.recordf(WarnLevel, format, args)
}

// Errorf records a Error level event with a formatted message.
func (l *TestLogger) Errorf(format string, args ...interface{}) {
	l.recordf(ErrorLevel, format, args)
}

// WithFields returns a new TestLogger with the appended fields.
func (l *TestLogger) WithFields(fields Fields) Logger {
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
