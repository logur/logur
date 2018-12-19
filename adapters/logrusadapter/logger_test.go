package logrusadapter

import (
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
	logrustest "github.com/sirupsen/logrus/hooks/test"
)

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
