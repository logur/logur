package hclogadapter

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/goph/logur"
	"github.com/goph/logur/testing"
	"github.com/hashicorp/go-hclog"
)

// nolint: gochecknoglobals
var logLineRegex = regexp.MustCompile(`.* \[(.*)\] {1,2}(.*): (.*)`)

func newTestSuite() *logtesting.LoggerTestSuite {
	return &logtesting.LoggerTestSuite{
		LoggerFactory: func(level logur.Level) (logur.Logger, func() []logur.LogEvent) {
			var buf bytes.Buffer
			logger := hclog.New(&hclog.LoggerOptions{
				Level:  hclog.Level(level + 1),
				Output: &buf,
			})

			return New(logger), func() []logur.LogEvent {
				lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

				events := make([]logur.LogEvent, len(lines))

				for key, line := range lines {
					match := logLineRegex.FindStringSubmatch(line)

					level, _ := logur.ParseLevel(strings.ToLower(match[1]))

					rawFields := strings.Fields(match[3])
					fields := make(map[string]interface{})

					for _, rawField := range rawFields {
						field := strings.SplitN(rawField, "=", 2)

						fields[field[0]] = field[1]
					}

					events[key] = logur.LogEvent{
						Line:   match[2],
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
