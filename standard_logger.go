package logur

import "log"

// NewStandardLogger returns a new standard library logger.
func NewStandardLogger(logger Logger, level Level, prefix string, flag int) *log.Logger {
	return log.New(NewLevelWriter(logger, level), prefix, flag)
}

// NewErrorStandardLogger returns a new standard library logger for error level logging (eg. for HTTP servers).
func NewErrorStandardLogger(logger Logger, prefix string, flag int) *log.Logger {
	return NewStandardLogger(logger, Error, prefix, flag)
}
