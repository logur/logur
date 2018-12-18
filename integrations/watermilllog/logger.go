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
	l.logger.WithFields(logur.Fields(fields)).WithFields(logur.Fields{"err": err}).Error(msg, nil)
}

func (l *logger) Info(msg string, fields watermill.LogFields) {
	l.logger.WithFields(logur.Fields(fields)).Info(msg, nil)
}

func (l *logger) Debug(msg string, fields watermill.LogFields) {
	l.logger.WithFields(logur.Fields(fields)).Debug(msg, nil)
}

func (l *logger) Trace(msg string, fields watermill.LogFields) {
	l.logger.WithFields(logur.Fields(fields)).Trace(msg, nil)
}
