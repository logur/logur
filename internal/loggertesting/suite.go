package loggertesting

import (
	"strings"
	"testing"

	"github.com/goph/logur"
)

// nolint: gochecknoglobals
var testLevelMap = map[logur.Level]struct {
	logFunc func(logger logur.Logger, args ...interface{})
}{
	logur.Trace: {
		logFunc: logur.Logger.Trace,
	},
	logur.Debug: {
		logFunc: logur.Logger.Debug,
	},
	logur.Info: {
		logFunc: logur.Logger.Info,
	},
	logur.Warn: {
		logFunc: logur.Logger.Warn,
	},
	logur.Error: {
		logFunc: logur.Logger.Error,
	},
}

type LoggerTestSuite struct {
	LoggerFactory          func() (logur.Logger, func() []logur.LogEvent)
	LogEventAssertionFlags uint8
	TraceFallbackToDebug   bool
}

func (s *LoggerTestSuite) TestLevels(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	for level, test := range testLevelMap {
		level, test := level, test

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
			if level == logur.Trace && s.TraceFallbackToDebug {
				level = logur.Debug
			}

			fields := logur.Fields{"key": "value"}

			logger, getLogEvents := s.LoggerFactory()

			logger = logger.WithFields(fields)

			args := []interface{}{"message", 1, "message", 2}

			test.logFunc(logger, args...)

			logEvents := getLogEvents()

			if got, want := len(logEvents), 1; got != want {
				t.Fatalf("expected %d log events, got %d", want, got)
			}

			logEvent := logur.LogEvent{
				Line:    "message1message2",
				RawLine: args,
				Level:   level,
				Fields:  fields,
			}

			AssertLogEvents(t, logEvent, logEvents[0], s.LogEventAssertionFlags)
		})
	}
}
