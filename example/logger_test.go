package example

import (
	"fmt"
	"testing"

	"github.com/goph/logur"
	"github.com/goph/logur/testing"
)

func TestAdapter_Levels(t *testing.T) {
	tests := map[string]struct {
		logFunc func(logger *Adapter, msg string, fields map[string]interface{})
	}{
		"trace": {
			logFunc: (*Adapter).Trace,
		},
		"debug": {
			logFunc: (*Adapter).Debug,
		},
		"info": {
			logFunc: (*Adapter).Info,
		},
		"warn": {
			logFunc: (*Adapter).Warn,
		},
		"error": {
			logFunc: (*Adapter).Error,
		},
	}

	for name, test := range tests {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			testLogger := logur.NewTestLogger()
			logger := NewAdapter(testLogger)

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

func TestAdapter_WithFields(t *testing.T) {
	testLogger := logur.NewTestLogger()

	var logger Logger = NewAdapter(testLogger)

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
