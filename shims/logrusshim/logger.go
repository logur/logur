package logrusshim

import (
	"github.com/goph/logur"
	"github.com/sirupsen/logrus"
)

type shim struct {
	*logrus.Entry
}

// New returns a new logur compatible logger with Logrus as the logging library.
// If nil is passed as logger, the global Logrus instance is used as fallback.
func New(logger *logrus.Logger) logur.Logger {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	return &shim{logrus.NewEntry(logger)}
}

// WithFields returns a new logger based on the original logger with
// the additional supplied fields.
func (s *shim) WithFields(fields logur.Fields) logur.Logger {
	return &shim{
		s.Entry.WithFields(logrus.Fields(fields)),
	}
}
