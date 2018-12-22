package logur

// MessageLogger simplifies the Logger interface by removing the second context parameter.
// Useful when there is no need for contextual logging.
type MessageLogger struct {
	logger Logger
}

// NewMessageLogger returns a new MessageLogger instance.
func NewMessageLogger(logger Logger) *MessageLogger {
	return &MessageLogger{
		logger: logger,
	}
}

// Trace logs a Trace level event.
func (l *MessageLogger) Trace(msg string) {
	l.logger.Trace(msg, nil)
}

// Debug logs a Debug level event.
func (l *MessageLogger) Debug(msg string) {
	l.logger.Debug(msg, nil)
}

// Info logs a Info level event.
func (l *MessageLogger) Info(msg string) {
	l.logger.Info(msg, nil)
}

// Warn logs a Warn level event.
func (l *MessageLogger) Warn(msg string) {
	l.logger.Warn(msg, nil)
}

// Error logs a Error level event.
func (l *MessageLogger) Error(msg string) {
	l.logger.Error(msg, nil)
}
