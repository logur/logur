package kitlogadapter

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
)

func newTestSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
		LogEventAssertionFlags: 0 | loggertesting.AllowNoNewLine,
		TraceFallbackToDebug:   true,
		LoggerFactory: func() (Logger, func() []LogEvent) {
			var buf bytes.Buffer
			logger := log.NewJSONLogger(&buf)

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
					msg := event["msg"].(string)

					delete(event, "level")
					delete(event, "msg")

					events[key] = LogEvent{
						Line:   msg,
						Level:  level,
						Fields: event,
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
