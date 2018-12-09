package logur

import (
	"reflect"
	"testing"
)

func TestTestLogger(t *testing.T) {
	tests := map[string]struct {
		level     Level
		logFunc   func(logger *TestLogger, args ...interface{})
		loglnFunc func(logger *TestLogger, args ...interface{})
		logfFunc  func(logger *TestLogger, format string, args ...interface{})
	}{
		"Trace": {
			level:     TraceLevel,
			logFunc:   (*TestLogger).Trace,
			loglnFunc: (*TestLogger).Traceln,
			logfFunc:  (*TestLogger).Tracef,
		},
		"Debug": {
			level:     DebugLevel,
			logFunc:   (*TestLogger).Debug,
			loglnFunc: (*TestLogger).Debugln,
			logfFunc:  (*TestLogger).Debugf,
		},
		"Info": {
			level:     InfoLevel,
			logFunc:   (*TestLogger).Info,
			loglnFunc: (*TestLogger).Infoln,
			logfFunc:  (*TestLogger).Infof,
		},
		"Warn": {
			level:     WarnLevel,
			logFunc:   (*TestLogger).Warn,
			loglnFunc: (*TestLogger).Warnln,
			logfFunc:  (*TestLogger).Warnf,
		},
		"Error": {
			level:     ErrorLevel,
			logFunc:   (*TestLogger).Error,
			loglnFunc: (*TestLogger).Errorln,
			logfFunc:  (*TestLogger).Errorf,
		},
	}

	for level, test := range tests {
		level, test := level, test

		t.Run(level, func(t *testing.T) {
			fields := Fields{"key": "value"}
			logger := NewTestLogger().WithFields(fields).(*TestLogger)

			args := []interface{}{"message", 1, "message", 2}
			format := "formatted msg: %s %d %s %d"

			test.logFunc(logger, args...)
			test.loglnFunc(logger, args...)
			test.logfFunc(logger, format, args...)

			if got, want := logger.Count(), 3; got != want {
				t.Fatalf("expected %d log events, got %d", want, got)
			}

			events := logger.Events()

			logEvent := LogEvent{
				Line:    "message1message2",
				RawLine: args,
				Level:   test.level,
				Fields:  fields,
			}

			if !reflect.DeepEqual(events[0], logEvent) {
				t.Errorf("expected log events to be equal\ngot:  %v\nwant: %v", events[0], logEvent)
			}

			loglnEvent := LogEvent{
				Line:    "message 1 message 2\n",
				RawLine: args,
				Level:   test.level,
				Fields:  fields,
			}

			if !reflect.DeepEqual(events[1], loglnEvent) {
				t.Errorf("expected log events to be equal\ngot:  %v\nwant: %v", events[1], loglnEvent)
			}

			logfEvent := LogEvent{
				Line:    "formatted msg: message 1 message 2",
				RawLine: append([]interface{}{format}, args...),
				Level:   test.level,
				Fields:  fields,
			}

			if !reflect.DeepEqual(events[2], logfEvent) {
				t.Errorf("expected log events to be equal\ngot:  %v\nwant: %v", events[2], logfEvent)
			}

			lastEvent := logger.LastEvent()

			if !reflect.DeepEqual(*lastEvent, logfEvent) {
				t.Errorf("expected log events to be equal\ngot:  %v\nwant: %v", *lastEvent, logfEvent)
			}
		})
	}
}
