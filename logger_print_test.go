package logur_test

import (
	"testing"

	. "github.com/goph/logur"
	logtesting "github.com/goph/logur/testing"
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
			logger := NewTestLogger()
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
			logger := NewTestLogger()
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
			logger := NewTestLogger()
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
