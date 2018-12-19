package logrusadapter

import (
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
	"github.com/sirupsen/logrus"
	logrustest "github.com/sirupsen/logrus/hooks/test"
)

// nolint: gochecknoglobals
var levelMap = map[Level]logrus.Level{
	Trace: logrus.TraceLevel,
	Debug: logrus.DebugLevel,
	Info:  logrus.InfoLevel,
	Warn:  logrus.WarnLevel,
	Error: logrus.ErrorLevel,
}

func newTestSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
		LoggerFactory: func(level Level) (Logger, func() []LogEvent) {
			logrusLogger, hook := logrustest.NewNullLogger()
			logrusLogger.SetLevel(levelMap[level])

			return New(logrusLogger), func() []LogEvent {
				entries := hook.AllEntries()

				events := make([]LogEvent, len(entries))

				for key, entry := range entries {
					level, _ := ParseLevel(entry.Level.String())

					events[key] = LogEvent{
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
