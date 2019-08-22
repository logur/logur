package kit

import (
	"fmt"
	"strings"

	"logur.dev/logur"
	"logur.dev/logur/internal/keyvals"
)

// Logger is a go-kit logger.
type Logger struct {
	logFuncs       map[string]logur.LogFunc
	defaultLogFunc logur.LogFunc
}

// New returns a new go-kit logger.
func New(logger logur.Logger) *Logger {
	l := &Logger{
		logFuncs: map[string]logur.LogFunc{
			"trace":   logger.Trace,
			"debug":   logger.Debug,
			"info":    logger.Info,
			"warn":    logger.Warn,
			"warning": logger.Warn,
			"error":   logger.Error,
		},
		defaultLogFunc: logger.Info,
	}

	return l
}

func (l *Logger) Log(kvs ...interface{}) error {
	if len(kvs)%2 == 1 {
		kvs = append(kvs, "(MISSING)")
	}

	fields := keyvals.ToMap(kvs)

	logFunc := l.defaultLogFunc

	if lf, ok := l.logFuncs[strings.ToLower(fmt.Sprintf("%s", fields["level"]))]; ok {
		delete(fields, "level")

		logFunc = lf
	}

	var msg string
	if m, ok := fields["msg"]; ok {
		delete(fields, "msg")
		msg = fmt.Sprintf("%s", m)
	}

	logFunc(msg, fields)

	return nil
}
