// Package mysqllog provides a MySQL database driver compatible logger.
package mysqllog

import (
	"fmt"

	"github.com/goph/logur"
)

// Logger is a MySQL database driver compatible logger.
type Logger struct {
	logger logur.Logger
}

// New returns a new MySQL database driver compatible logger.
func New(logger logur.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

// Print is used to log critical error messages.
func (l *Logger) Print(v ...interface{}) {
	l.logger.Error(fmt.Sprint(v...), nil)
}
