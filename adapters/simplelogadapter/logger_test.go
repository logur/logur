package simplelogadapter

import (
	"reflect"
	"testing"

	"github.com/goph/logur"
)

type testLogger struct {
	*logur.TestLogger
}

func newTestLogger() *testLogger {
	return &testLogger{logur.NewTestLogger()}
}

func (l *testLogger) WithFields(fields Fields) Logger {
	return &testLogger{l.TestLogger.WithFields(logur.Fields(fields)).(*logur.TestLogger)}
}

func TestAdapter(t *testing.T) {
	tests := map[string]struct {
		level     logur.Level
		loglnFunc func(logger logur.Logger, args ...interface{})
		logfFunc  func(logger logur.Logger, format string, args ...interface{})
	}{
		"Trace": {
			level:     logur.TraceLevel,
			loglnFunc: logur.Logger.Traceln,
			logfFunc:  logur.Logger.Tracef,
		},
		"Debug": {
			level:     logur.DebugLevel,
			loglnFunc: logur.Logger.Debugln,
			logfFunc:  logur.Logger.Debugf,
		},
		"Info": {
			level:     logur.InfoLevel,
			loglnFunc: logur.Logger.Infoln,
			logfFunc:  logur.Logger.Infof,
		},
		"Warn": {
			level:     logur.WarnLevel,
			loglnFunc: logur.Logger.Warnln,
			logfFunc:  logur.Logger.Warnf,
		},
		"Error": {
			level:     logur.ErrorLevel,
			loglnFunc: logur.Logger.Errorln,
			logfFunc:  logur.Logger.Errorf,
		},
	}

	for level, test := range tests {
		level, test := level, test

		t.Run(level, func(t *testing.T) {
			simpleLogger := newTestLogger()

			fields := logur.Fields{"key": "value"}
			logger := New(simpleLogger).WithFields(fields)

			args := []interface{}{"message", 1, "message", 2}
			format := "formatted msg: %s %d %s %d"

			test.loglnFunc(logger, args...)
			test.logfFunc(logger, format, args...)

			if got, want := simpleLogger.Count(), 2; got != want {
				t.Fatalf("expected %d log events, got %d", want, got)
			}

			events := simpleLogger.Events()

			loglnEvent := logur.LogEvent{
				Line:    "message 1 message 2",
				RawLine: []interface{}{"message 1 message 2"},
				Level:   test.level,
				Fields:  fields,
			}

			if !reflect.DeepEqual(events[0], loglnEvent) {
				t.Errorf("expected log events to be equal\ngot:  %v\nwant: %v", events[0], loglnEvent)
			}

			logfEvent := logur.LogEvent{
				Line:    "formatted msg: message 1 message 2",
				RawLine: []interface{}{"formatted msg: message 1 message 2"},
				Level:   test.level,
				Fields:  fields,
			}

			if !reflect.DeepEqual(events[1], logfEvent) {
				t.Errorf("expected log events to be equal\ngot:  %v\nwant: %v", events[1], logfEvent)
			}

			lastEvent := simpleLogger.LastEvent()

			if !reflect.DeepEqual(*lastEvent, logfEvent) {
				t.Errorf("expected log events to be equal\ngot:  %v\nwant: %v", *lastEvent, logfEvent)
			}
		})
	}
}
