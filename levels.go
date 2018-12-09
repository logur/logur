package logur

import (
	"fmt"
	"strings"
)

// Level represents the same levels as defined in Logger.
type Level uint32

// Levels as defined in Logger.
const (
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	TraceLevel Level = iota

	// A verbose series of information events.
	// They are useful when debugging the system.
	DebugLevel

	// General information about what's happening inside the system.
	InfoLevel

	// Non-critical events that should be looked at.
	WarnLevel

	// Critical events that require immediate attention.
	ErrorLevel
)

// ParseLevel takes a string level and returns the defined log level constant.
// If the level is not defined, it returns false as the second parameter.
func ParseLevel(level string) (Level, bool) {
	switch strings.ToLower(level) {
	case "trace":
		return TraceLevel, true

	case "debug":
		return DebugLevel, true

	case "info":
		return InfoLevel, true

	case "warn", "warning":
		return WarnLevel, true

	case "error":
		return ErrorLevel, true
	}

	return Level(999), false
}

// String converts a Level to string.
func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "trace"

	case DebugLevel:
		return "debug"

	case InfoLevel:
		return "info"

	case WarnLevel:
		return "warn"

	case ErrorLevel:
		return "error"
	}

	return "unknown"
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (l *Level) UnmarshalText(text []byte) error {
	level, ok := ParseLevel(string(text))
	if !ok {
		return fmt.Errorf("undefined level: %q", string(text))
	}

	*l = level

	return nil
}
