package zapadapter

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newTestSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
		LogEventAssertionFlags: 0 | loggertesting.AllowNoNewLine,
		TraceFallbackToDebug:   true,
		LoggerFactory: func() (Logger, func() []LogEvent) {
			var buf bytes.Buffer

			logger := zap.New(
				zapcore.NewCore(
					zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
					zapcore.AddSync(&buf),
					zap.DebugLevel,
				),
			)

			return New(logger.Sugar()), func() []LogEvent {
				lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

				events := make([]LogEvent, len(lines))

				for key, line := range lines {
					log := strings.SplitN(line, "\t", 4)

					level, _ := ParseLevel(strings.ToLower(log[1]))

					var fields map[string]interface{}

					if len(log) > 3 {
						err := json.Unmarshal([]byte(log[3]), &fields)
						if err != nil {
							panic(err)
						}
					}

					events[key] = LogEvent{
						Line:   log[2],
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
