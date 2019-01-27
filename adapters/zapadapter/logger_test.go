package zapadapter

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/goph/logur"
	logtesting "github.com/goph/logur/testing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// nolint: gochecknoglobals
var levelMap = map[logur.Level]zapcore.Level{
	logur.Trace: zap.DebugLevel,
	logur.Debug: zap.DebugLevel,
	logur.Info:  zap.InfoLevel,
	logur.Warn:  zap.WarnLevel,
	logur.Error: zap.ErrorLevel,
}

func newTestSuite() *logtesting.LoggerTestSuite {
	return &logtesting.LoggerTestSuite{
		TraceFallbackToDebug: true,
		LoggerFactory: func(level logur.Level) (logur.Logger, func() []logur.LogEvent) {
			var buf bytes.Buffer

			logger := zap.New(
				zapcore.NewCore(
					zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
					zapcore.AddSync(&buf),
					levelMap[level],
				),
			)

			return New(logger), func() []logur.LogEvent {
				lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

				events := make([]logur.LogEvent, len(lines))

				for key, line := range lines {
					log := strings.SplitN(line, "\t", 4)

					level, _ := logur.ParseLevel(strings.ToLower(log[1]))

					var fields map[string]interface{}

					if len(log) > 3 {
						err := json.Unmarshal([]byte(log[3]), &fields)
						if err != nil {
							panic(err)
						}
					}

					events[key] = logur.LogEvent{
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
