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
	Trace Level = iota

	// A verbose series of information events.
	// They are useful when debugging the system.
	Debug

	// General information about what's happening inside the system.
	Info

	// Non-critical events that should be looked at.
	Warn

	// Critical events that require immediate attention.
	Error
)

// Levels returns a list of available Level values.
func Levels() []Level {
	return []Level{Trace, Debug, Info, Warn, Error}
}

// ParseLevel takes a string level and returns the defined log level constant.
// If the level is not defined, it returns false as the second parameter.
func ParseLevel(level string) (Level, bool) {
	switch strings.ToLower(level) {
	case "trace":
		return Trace, true

	case "debug":
		return Debug, true

	case "info":
		return Info, true

	case "warn", "warning":
		return Warn, true

	case "error":
		return Error, true
	}

	return Level(999), false
}

// String converts a Level to string.
func (l Level) String() string {
	switch l {
	case Trace:
		return "trace"

	case Debug:
		return "debug"

	case Info:
		return "info"

	case Warn:
		return "warn"

	case Error:
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

// LevelFunc returns a logger function for a level.
// If the level is invalid it falls back to Info level.
func LevelFunc(logger Logger, level Level) LogFunc {
	switch level {
	case Trace:
		return logger.Trace

	case Debug:
		return logger.Debug

	case Info:
		return logger.Info

	case Warn:
		return logger.Warn

	case Error:
		return logger.Error

	default:
		return logger.Info
	}
}

// LevelEnabler checks if a level is enabled in a logger.
// If the logger cannot reliably decide the correct level this method MUST return true.
type LevelEnabler interface {
	// LevelEnabled checks if a level is enabled in a logger.
	LevelEnabled(level Level) bool
}
