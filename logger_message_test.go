package logur_test

import (
	"strings"
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/testing"
)

// TestLevels tests leveled logging capabilities.
func TestMessageLogger_Levels(t *testing.T) {
	tests := map[Level]struct {
		logFunc func(logger *MessageLogger, msg string)
	}{
		Trace: {
			logFunc: (*MessageLogger).Trace,
		},
		Debug: {
			logFunc: (*MessageLogger).Debug,
		},
		Info: {
			logFunc: (*MessageLogger).Info,
		},
		Warn: {
			logFunc: (*MessageLogger).Warn,
		},
		Error: {
			logFunc: (*MessageLogger).Error,
		},
	}

	for level, test := range tests {
		level, test := level, test

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
			testLogger := NewTestLogger()
			logger := NewMessageLogger(testLogger)

			test.logFunc(logger, "message")

			event := LogEvent{
				Line:  "message",
				Level: level,
			}

			logtesting.AssertLogEvents(t, event, *(testLogger.LastEvent()))
		})
	}
}
