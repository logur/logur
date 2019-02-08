package watermilllog

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/goph/logur"
)

type ErrorLogger struct {
	logger logur.Logger
}

// New returns a github.com/ThreeDotsLabs/watermill.LoggerAdapter compatible logger.
func NewErrorLogger(l logur.Logger) *ErrorLogger {
	return &ErrorLogger{l}
}

func (l *ErrorLogger) Error(msg string, err error, fields watermill.LogFields) {
	fields["err"] = err

	l.logger.Error(msg, fields)
}

func (l *ErrorLogger) Info(msg string, fields watermill.LogFields) {
	l.logger.Info(msg, fields)
}

func (l *ErrorLogger) Debug(msg string, fields watermill.LogFields) {
	l.logger.Debug(msg, fields)
}

func (l *ErrorLogger) Trace(msg string, fields watermill.LogFields) {
	l.logger.Trace(msg, fields)
}

func (l *ErrorLogger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	if len(fields) == 0 {
		return l
	}

	return New(logur.WithFields(l.logger, fields))
}
