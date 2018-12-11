package zerologadapter

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
	"github.com/rs/zerolog"
)

func newTestSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
		LogEventAssertionFlags: 0 | loggertesting.SkipRawLine | loggertesting.AllowNoNewLine,
		TraceFallbackToDebug:   true,
		LoggerFactory: func() (Logger, func() []LogEvent) {
			var buf bytes.Buffer
			logger := zerolog.New(&buf)

			return New(logger), func() []LogEvent {
				lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

				events := make([]LogEvent, len(lines))

				for key, line := range lines {
					var event map[string]interface{}

					err := json.Unmarshal([]byte(line), &event)
					if err != nil {
						panic(err)
					}

					level, _ := ParseLevel(strings.ToLower(event["level"].(string)))
					msg := event["message"].(string)

					delete(event, "level")
					delete(event, "message")

					fields := Fields(event)

					events[key] = LogEvent{
						Line:   msg,
						Level:  level,
						Fields: fields,
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
