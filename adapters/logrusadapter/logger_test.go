package logrusadapter

import (
	"testing"

	"github.com/goph/logur"
	"github.com/goph/logur/testing"
	"github.com/sirupsen/logrus"
	logrustest "github.com/sirupsen/logrus/hooks/test"
)

// nolint: gochecknoglobals
var levelMap = map[logur.Level]logrus.Level{
	logur.Trace: logrus.TraceLevel,
	logur.Debug: logrus.DebugLevel,
	logur.Info:  logrus.InfoLevel,
	logur.Warn:  logrus.WarnLevel,
	logur.Error: logrus.ErrorLevel,
}

func newTestSuite() *logtesting.LoggerTestSuite {
	return &logtesting.LoggerTestSuite{
		LoggerFactory: func(level logur.Level) (logur.Logger, func() []logur.LogEvent) {
			logrusLogger, hook := logrustest.NewNullLogger()
			logrusLogger.SetLevel(levelMap[level])

			return New(logrusLogger), func() []logur.LogEvent {
				entries := hook.AllEntries()

				events := make([]logur.LogEvent, len(entries))

				for key, entry := range entries {
					level, _ := logur.ParseLevel(entry.Level.String())

					events[key] = logur.LogEvent{
						Line:   entry.Message,
						Level:  level,
						Fields: entry.Data,
					}
				}

				return events
			}
		},
	}
}

func TestLoggerSuite(t *testing.T) {
	newTestSuite().Execute(t)
}
