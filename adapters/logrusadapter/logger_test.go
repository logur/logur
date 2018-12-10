package logrusadapter

import (
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
	"github.com/sirupsen/logrus"
	logrustest "github.com/sirupsen/logrus/hooks/test"
)

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
