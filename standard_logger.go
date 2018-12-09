package logur

import "log"

// NewStandardLogger returns a new standard library logger.
func NewStandardLogger(logger Logger, level Level, prefix string, flag int) *log.Logger {
	return log.New(NewLevelWriter(logger, level), prefix, flag)
}
