package logur

// Logger is a unified interface for various logging use cases and practices, including:
//		- leveled logging
// 		- leveled formatters
// 		- structured logging
type Logger interface {
	// Trace logs a Trace event.
	//
	// Even more fine-grained information than Debug events.
	// Loggers not supporting this level should fall back to Debug.
	Trace(msg string)

	// Debug logs a Debug event.
	//
	// A verbose series of information events.
	// They are useful when debugging the system.
	Debug(msg string)

	// Info logs an Info event.
	//
	// General information about what's happening inside the system.
	Info(msg string)

	// Warn logs a Warn(ing) event.
	//
	// Non-critical events that should be looked at.
	Warn(msg string)

	// Error logs an Error event.
	//
	// Critical events that require immediate attention.
	// Loggers commonly provide Fatal and Panic levels above Error level,
	// but exiting and panicing is out of scope for a logging library.
	Error(msg string)

	// WithFields appends structured fields to a new (child) logger instance.
	WithFields(fields map[string]interface{}) Logger
}

// Fields is used to define structured fields which are appended to log events.
// It can be used as a shorthand for map[string]interface{}.
type Fields map[string]interface{}

// LogFunc is a function recording a log event.
type LogFunc func(msg string)
