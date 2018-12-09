package loggertesting

import (
	"reflect"
	"strings"
	"testing"

	"github.com/goph/logur"
)

var testLevelMap = map[logur.Level]struct {
	logFunc   func(logger logur.Logger, args ...interface{})
	loglnFunc func(logger logur.Logger, args ...interface{})
	logfFunc  func(logger logur.Logger, format string, args ...interface{})
}{
	logur.TraceLevel: {
		logFunc:   logur.Logger.Trace,
		loglnFunc: logur.Logger.Traceln,
		logfFunc:  logur.Logger.Tracef,
	},
	logur.DebugLevel: {
		logFunc:   logur.Logger.Debug,
		loglnFunc: logur.Logger.Debugln,
		logfFunc:  logur.Logger.Debugf,
	},
	logur.InfoLevel: {
		logFunc:   logur.Logger.Info,
		loglnFunc: logur.Logger.Infoln,
		logfFunc:  logur.Logger.Infof,
	},
	logur.WarnLevel: {
		logFunc:   logur.Logger.Warn,
		loglnFunc: logur.Logger.Warnln,
		logfFunc:  logur.Logger.Warnf,
	},
	logur.ErrorLevel: {
		logFunc:   logur.Logger.Error,
		loglnFunc: logur.Logger.Errorln,
		logfFunc:  logur.Logger.Errorf,
	},
}

func assertLogEvents(t *testing.T, expected logur.LogEvent, actual logur.LogEvent, compareRawLine bool) {
	t.Helper()

	if expected.Level != actual.Level {
		t.Errorf("expected log levels to be equal\ngot:  %s\nwant: %s", expected.Level, actual.Level)
	}

	if expected.Line != actual.Line {
		t.Errorf("expected log lines to be equal\ngot:  %s\nwant: %s", expected.Line, actual.Line)
	}

	if compareRawLine {
		if !reflect.DeepEqual(expected.RawLine, actual.RawLine) {
			t.Errorf("expected raw log lines to be equal\ngot:  %v\nwant: %v", expected.RawLine, actual.RawLine)
		}
	}

	if !reflect.DeepEqual(expected.Fields, actual.Fields) {
		t.Errorf("expected log fields to be equal\ngot:  %v\nwant: %v", expected.Fields, actual.Fields)
	}
}

type LoggerTestSuite struct {
	LoggerFactory   func() (logur.Logger, func() []logur.LogEvent)
	CompareRawLines bool
}

func (s *LoggerTestSuite) TestLevels(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	for level, test := range testLevelMap {
		level, test := level, test

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
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

			assertLogEvents(t, logEvent, logEvents[0], s.CompareRawLines)
		})
	}
}

func (s *LoggerTestSuite) TestLevelsln(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	for level, test := range testLevelMap {
		level, test := level, test

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
			fields := logur.Fields{"key": "value"}

			logger, getLogEvents := s.LoggerFactory()

			logger = logger.WithFields(fields)

			args := []interface{}{"message", 1, "message", 2}

			test.loglnFunc(logger, args...)

			logEvents := getLogEvents()

			if got, want := len(logEvents), 1; got != want {
				t.Fatalf("expected %d log events, got %d", want, got)
			}

			logEvent := logur.LogEvent{
				Line:    "message 1 message 2\n",
				RawLine: args,
				Level:   level,
				Fields:  fields,
			}

			assertLogEvents(t, logEvent, logEvents[0], s.CompareRawLines)
		})
	}
}

func (s *LoggerTestSuite) TestLevelsf(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	for level, test := range testLevelMap {
		level, test := level, test

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
			fields := logur.Fields{"key": "value"}

			logger, getLogEvents := s.LoggerFactory()

			logger = logger.WithFields(fields)

			args := []interface{}{"message", 1, "message", 2}
			format := "formatted msg: %s %d %s %d"

			test.logfFunc(logger, format, args...)

			logEvents := getLogEvents()

			if got, want := len(logEvents), 1; got != want {
				t.Fatalf("expected %d log events, got %d", want, got)
			}

			logEvent := logur.LogEvent{
				Line:    "formatted msg: message 1 message 2",
				RawLine: args,
				Level:   level,
				Fields:  fields,
			}

			assertLogEvents(t, logEvent, logEvents[0], s.CompareRawLines)
		})
	}
}
