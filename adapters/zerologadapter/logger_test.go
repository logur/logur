package zerologadapter

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/rs/zerolog"

	"github.com/goph/logur"
	"github.com/goph/logur/logtesting"
)

func newTestSuite() *logtesting.LoggerTestSuite {
	return &logtesting.LoggerTestSuite{
		TraceFallbackToDebug: true,
		LoggerFactory: func(level logur.Level) (logur.Logger, func() []logur.LogEvent) {
			var buf bytes.Buffer
			logger := zerolog.New(&buf).Level(zerolog.Level(level))

			return New(logger), func() []logur.LogEvent {
				lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

				events := make([]logur.LogEvent, len(lines))

				for key, line := range lines {
					var event map[string]interface{}

					err := json.Unmarshal([]byte(line), &event)
					if err != nil {
						panic(err)
					}

					level, _ := logur.ParseLevel(strings.ToLower(event["level"].(string)))
					msg := event["message"].(string)

					delete(event, "level")
					delete(event, "message")

					events[key] = logur.LogEvent{
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

func TestLoggerSuite(t *testing.T) {
	newTestSuite().Execute(t)
}
