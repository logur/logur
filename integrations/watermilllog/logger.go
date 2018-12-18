package watermilllog

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/goph/logur"
)

type logger struct {
	logger logur.Logger
}

// New returns a github.com/ThreeDotsLabs/watermill.LoggerAdapter compatible logger.
func New(l logur.Logger) watermill.LoggerAdapter {
	return &logger{l}
}

func (l *logger) Error(msg string, err error, fields watermill.LogFields) {
	fields["err"] = err

	l.logger.Error(msg, fields)
}

func (l *logger) Info(msg string, fields watermill.LogFields) {
	l.logger.Info(msg, fields)
}

func (l *logger) Debug(msg string, fields watermill.LogFields) {
	l.logger.Debug(msg, fields)
}

func (l *logger) Trace(msg string, fields watermill.LogFields) {
	l.logger.Trace(msg, fields)
}
