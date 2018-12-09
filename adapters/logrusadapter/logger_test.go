package logrusadapter

import (
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
	"github.com/sirupsen/logrus"
	logrustest "github.com/sirupsen/logrus/hooks/test"
)

func TestAdapter(t *testing.T) {
	tests := map[string]struct {
		level     Level
		logFunc   func(logger Logger, args ...interface{})
		loglnFunc func(logger Logger, args ...interface{})
		logfFunc  func(logger Logger, format string, args ...interface{})
	}{
		"Trace": {
			level:     TraceLevel,
			logFunc:   Logger.Trace,
			loglnFunc: Logger.Traceln,
			logfFunc:  Logger.Tracef,
		},
		"Debug": {
			level:     DebugLevel,
			logFunc:   Logger.Debug,
			loglnFunc: Logger.Debugln,
			logfFunc:  Logger.Debugf,
		},
		"Info": {
			level:     InfoLevel,
			logFunc:   Logger.Info,
			loglnFunc: Logger.Infoln,
			logfFunc:  Logger.Infof,
		},
		"Warn": {
			level:     WarnLevel,
			logFunc:   Logger.Warn,
			loglnFunc: Logger.Warnln,
			logfFunc:  Logger.Warnf,
		},
		"Error": {
			level:     ErrorLevel,
			logFunc:   Logger.Error,
			loglnFunc: Logger.Errorln,
			logfFunc:  Logger.Errorf,
		},
	}

	for level, test := range tests {
		level, test := level, test

		t.Run(level, func(t *testing.T) {
			logrusLogger, hook := logrustest.NewNullLogger()
			logrusLogger.SetLevel(logrus.TraceLevel)

			fields := Fields{"key": "value"}
			logger := New(logrusLogger).WithFields(fields)

			args := []interface{}{"message", 1, "message", 2}
			format := "formatted msg: %s %d %s %d"

			test.logFunc(logger, args...)
			test.loglnFunc(logger, args...)
			test.logfFunc(logger, format, args...)

			entries := hook.AllEntries()

			if got, want := len(entries), 3; got != want {
				t.Fatalf("expected %d log events, got %d", want, got)
			}

			logrusLevel, err := logrus.ParseLevel(test.level.String())
			if err != nil {
				t.Fatal("invalid logrus level:", err.Error())
			}

			if entries[0].Level != logrusLevel {
				t.Errorf(
					"expected log event to be level %q but received %q instead",
					logrusLevel.String(),
					entries[0].Level.String(),
				)
			}

			if got, want := entries[0].Message, "message1message2"; got != want {
				t.Errorf("expected log messages to be equal\ngot:  %s\nwant: %s", got, want)
			}

			if entries[1].Level != logrusLevel {
				t.Errorf(
					"expected log event to be level %q but received %q instead",
					logrusLevel.String(),
					entries[1].Level.String(),
				)
			}

			if got, want := entries[1].Message, "message 1 message 2"; got != want {
				t.Errorf("expected log messages to be equal\ngot:  %s\nwant: %s", got, want)
			}

			if entries[2].Level != logrusLevel {
				t.Errorf(
					"expected log event to be level %q but received %q instead",
					logrusLevel.String(),
					entries[2].Level.String(),
				)
			}

			if got, want := entries[2].Message, "formatted msg: message 1 message 2"; got != want {
				t.Errorf("expected log messages to be equal\ngot:  %s\nwant: %s", got, want)
			}
		})
	}
}

func newTestSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
		LogEventAssertionFlags: 0 | loggertesting.SkipRawLine | loggertesting.AllowNoNewLine,
		LoggerFactory: func() (Logger, func() []LogEvent) {
			logrusLogger, hook := logrustest.NewNullLogger()
			logrusLogger.SetLevel(logrus.TraceLevel)

			return New(logrusLogger), func() []LogEvent {
				entries := hook.AllEntries()

				events := make([]LogEvent, len(entries))

				for key, entry := range entries {
					level, _ := ParseLevel(entry.Level.String())

					events[key] = LogEvent{
						Line:   entry.Message,
						Level:  level,
						Fields: Fields(entry.Data),
					}
				}

				return events
			}
		},
	}
}

func TestLogger_Levels(t *testing.T) {
	newTestSuite().TestLevels(t)
}

func TestLogger_Levelsln(t *testing.T) {
	newTestSuite().TestLevelsln(t)
}

func TestLogger_Levelsf(t *testing.T) {
	newTestSuite().TestLevelsf(t)
}
