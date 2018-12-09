package hclogadapter

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	. "github.com/goph/logur"
	"github.com/hashicorp/go-hclog"
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
			var buf bytes.Buffer
			hclogLogger := hclog.New(&hclog.LoggerOptions{
				Level:  hclog.Trace,
				Output: &buf,
			})

			fields := Fields{"key": "value"}
			logger := New(hclogLogger).WithFields(fields)

			args := []interface{}{"message", 1, "message", 2}
			format := "formatted msg: %s %d %s %d"

			test.logFunc(logger, args...)
			test.loglnFunc(logger, args...)
			test.logfFunc(logger, format, args...)

			lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

			if got, want := len(lines), 3; got != want {
				t.Fatalf("expected %d log events, got %d", want, got)
			}

			levelString := fmt.Sprintf(
				"[%s]%s",
				strings.ToUpper(test.level.String()),
				strings.Repeat(" ", 6-len(test.level.String())),
			)

			line := strings.SplitN(lines[0], " ", 2)[1]
			if got, want := line, fmt.Sprintf("%smessage1message2: key=value", levelString); got != want {
				t.Errorf("expected log messages to be equal\ngot:  %s\nwant: %s", got, want)
			}

			lineln := strings.SplitN(lines[1], " ", 2)[1]
			if got, want := lineln, fmt.Sprintf("%smessage 1 message 2: key=value", levelString); got != want {
				t.Errorf("expected log messages to be equal\ngot:  %s\nwant: %s", got, want)
			}

			linef := strings.SplitN(lines[2], " ", 2)[1]
			if got, want := linef, fmt.Sprintf("%sformatted msg: message 1 message 2: key=value", levelString); got != want {
				t.Errorf("expected log messages to be equal\ngot:  %s\nwant: %s", got, want)
			}
		})
	}
}
