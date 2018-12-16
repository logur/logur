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

	// Traceln is the Println counterpart of Trace.
	//
	// Loggers not supporting this level should fall back to Debug.
	Traceln(args ...interface{})

	// Debugln is the Println counterpart of Debug.
	Debugln(args ...interface{})

	// Infoln is the Println counterpart of Info.
	Infoln(args ...interface{})

	// Warnln is the Println counterpart of Warn.
	Warnln(args ...interface{})

	// Errorln is the Println counterpart of Error.
	Errorln(args ...interface{})

	// Tracef is the formatter counterpart of Trace.
	//
	// Loggers not supporting this level should fall back to Debug.
	Tracef(format string, args ...interface{})

	// Debugf is the formatter counterpart of Debug.
	Debugf(format string, args ...interface{})

	// Infof is the formatter counterpart of Info.
	Infof(format string, args ...interface{})

	// Warnf is the formatter counterpart of Warn.
	Warnf(format string, args ...interface{})

	// Errorf is the formatter counterpart of Error.
	Errorf(format string, args ...interface{})

	// WithFields appends structured fields to a new (child) logger instance.
	WithFields(fields map[string]interface{}) Logger
}

// Fields is used to define structured fields which are appended to log events.
// It can be used as a shorthand for map[string]interface{}.
type Fields map[string]interface{}

// LogFunc is a function recording a log event.
type LogFunc func(args ...interface{})
