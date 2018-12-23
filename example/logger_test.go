package example

import (
	"fmt"
	"testing"

	"github.com/goph/logur"
	"github.com/goph/logur/testing"
)

func TestLoggerAdapter_Levels(t *testing.T) {
	tests := map[string]struct {
		logFunc func(logger *LoggerAdapter, msg string, fields map[string]interface{})
	}{
		"trace": {
			logFunc: (*LoggerAdapter).Trace,
		},
		"debug": {
			logFunc: (*LoggerAdapter).Debug,
		},
		"info": {
			logFunc: (*LoggerAdapter).Info,
		},
		"warn": {
			logFunc: (*LoggerAdapter).Warn,
		},
		"error": {
			logFunc: (*LoggerAdapter).Error,
		},
	}

	for name, test := range tests {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			testLogger := logur.NewTestLogger()
			logger := NewLoggerAdapter(testLogger)

			test.logFunc(logger, fmt.Sprintf("message: %s", name), nil)

			level, _ := logur.ParseLevel(name)

			event := logur.LogEvent{
				Level: level,
				Line:  "message: " + name,
			}

			logtesting.AssertLogEvents(t, event, *(testLogger.LastEvent()))
		})
	}
}

func TestLoggerAdapter_WithFields(t *testing.T) {
	testLogger := logur.NewTestLogger()

	var logger Logger = NewLoggerAdapter(testLogger)

	fields := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	logger = logger.WithFields(fields)

	logger.Debug("message", nil)

	event := logur.LogEvent{
		Level:  logur.Debug,
		Line:   "message",
		Fields: fields,
	}

	logtesting.AssertLogEvents(t, event, *(testLogger.LastEvent()))
}
