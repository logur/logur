package logur_test

import (
	"strings"
	"testing"

	. "logur.dev/logur"
	"logur.dev/logur/logtesting"
)

// nolint: gochecknoglobals
var printLoggerTestMap = map[string]*struct {
	logger func(logger Logger) *PrintLogger
	level  Level
}{
	"info": {
		logger: func(logger Logger) *PrintLogger {
			return NewPrintLogger(LevelFunc(logger, Info))
		},
		level: Info,
	},
	"error": {
		logger: NewErrorPrintLogger,
		level:  Error,
	},
}

func TestPrintLogger_Print(t *testing.T) {
	for name, test := range printLoggerTestMap {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			logger := &TestLogger{}
			printLogger := test.logger(logger)

			printLogger.Print("message", 1, "message", 2)

			event := LogEvent{
				Level: test.level,
				Line:  "message1message2",
			}

			logtesting.AssertLogEventsEqual(t, event, *(logger.LastEvent()))
		})
	}
}

func TestPrintLogger_Println(t *testing.T) {
	for name, test := range printLoggerTestMap {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			logger := &TestLogger{}
			printLogger := test.logger(logger)

			printLogger.Println("message", 1, "message", 2)

			event := LogEvent{
				Level: test.level,
				Line:  "message 1 message 2",
			}

			logtesting.AssertLogEventsEqual(t, event, *(logger.LastEvent()))
		})
	}
}

func TestPrintLogger_Printf(t *testing.T) {
	for name, test := range printLoggerTestMap {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			logger := &TestLogger{}
			printLogger := test.logger(logger)

			printLogger.Printf("this is my %s", "message")

			event := LogEvent{
				Level: test.level,
				Line:  "this is my message",
			}

			logtesting.AssertLogEventsEqual(t, event, *(logger.LastEvent()))
		})
	}
}

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
			testLogger := &TestLogger{}
			logger := NewMessageLogger(testLogger)

			test.logFunc(logger, "message")

			event := LogEvent{
				Line:  "message",
				Level: level,
			}

			logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
		})
	}
}
