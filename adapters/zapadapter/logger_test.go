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

// nolint: gochecknoglobals
var levelMap = map[Level]zapcore.Level{
	Trace: zap.DebugLevel,
	Debug: zap.DebugLevel,
	Info:  zap.InfoLevel,
	Warn:  zap.WarnLevel,
	Error: zap.ErrorLevel,
}

func newTestSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
		TraceFallbackToDebug: true,
		LoggerFactory: func(level Level) (Logger, func() []LogEvent) {
			var buf bytes.Buffer

			logger := zap.New(
				zapcore.NewCore(
					zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
					zapcore.AddSync(&buf),
					levelMap[level],
				),
			)

			return New(logger), func() []LogEvent {
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

func TestLoggerSuite(t *testing.T) {
	newTestSuite().Execute(t)
}
