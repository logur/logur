package hclogadapter

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	. "github.com/goph/logur"
	"github.com/goph/logur/internal/loggertesting"
	"github.com/hashicorp/go-hclog"
)

// nolint: gochecknoglobals
var logLineRegex = regexp.MustCompile(`.* \[(.*)\] {1,2}(.*): (.*)`)

func newTestSuite() *loggertesting.LoggerTestSuite {
	return &loggertesting.LoggerTestSuite{
		LogEventAssertionFlags: 0 | loggertesting.SkipRawLine | loggertesting.AllowNoNewLine,
		LoggerFactory: func() (Logger, func() []LogEvent) {
			var buf bytes.Buffer
			logger := hclog.New(&hclog.LoggerOptions{
				Level:  hclog.Trace,
				Output: &buf,
			})

			return New(logger), func() []LogEvent {
				lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

				events := make([]LogEvent, len(lines))

				for key, line := range lines {
					match := logLineRegex.FindStringSubmatch(line)

					level, _ := ParseLevel(strings.ToLower(match[1]))

					rawFields := strings.Fields(match[3])
					fields := make(Fields)

					for _, rawField := range rawFields {
						field := strings.SplitN(rawField, "=", 2)

						fields[field[0]] = field[1]
					}

					events[key] = LogEvent{
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

func TestLogger_Levels(t *testing.T) {
	newTestSuite().TestLevels(t)
}

func TestLogger_Levelsln(t *testing.T) {
	newTestSuite().TestLevelsln(t)
}

func TestLogger_Levelsf(t *testing.T) {
	newTestSuite().TestLevelsf(t)
}
