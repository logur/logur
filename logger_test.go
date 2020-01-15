package logur_test

import (
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	. "logur.dev/logur"
	"logur.dev/logur/logtesting"
)

func newFieldLoggerTestSuite() *logtesting.LoggerTestSuite {
	return &logtesting.LoggerTestSuite{
		LoggerFactory: func(_ Level) (Logger, func() []LogEvent) {
			logger := &TestLogger{}

			return WithFields(logger, map[string]interface{}{"key": "value"}), logger.Events
		},
	}
}

func TestFieldLogger_Levels(t *testing.T) {
	newFieldLoggerTestSuite().TestLevels(t)
}

func TestWithFields(t *testing.T) {
	logger := &TestLogger{}
	l := WithFields(
		WithFields(
			WithFields(logger, map[string]interface{}{"key": "value"}),
			map[string]interface{}{"key": "value2"},
		),
		map[string]interface{}{"key": "value3"},
	)

	l.Info("message", map[string]interface{}{"key2": "value"})

	logEvent := LogEvent{
		Line:   "message",
		Level:  Info,
		Fields: map[string]interface{}{"key": "value3", "key2": "value"},
	}

	logtesting.AssertLogEventsEqual(t, logEvent, logEvent)
}

func TestWithField(t *testing.T) {
	logger := &TestLogger{}
	l := WithField(
		WithField(
			WithField(logger, "key", "value"),
			"key", "value2",
		),
		"key", "value3",
	)

	l.Info("message", map[string]interface{}{"key2": "value"})

	logEvent := LogEvent{
		Line:   "message",
		Level:  Info,
		Fields: map[string]interface{}{"key": "value3", "key2": "value"},
	}

	logtesting.AssertLogEventsEqual(t, logEvent, *logger.LastEvent())
}

func TestNewLoggerContext(t *testing.T) {
	t.Run("no_fields", func(t *testing.T) {
		testLogger := &TestLoggerFacade{}

		logger := NewLoggerContext(testLogger, func(ctx context.Context) map[string]interface{} {
			return nil
		})

		logger.InfoContext(context.Background(), "message")

		logEvent := LogEvent{
			Line:  "message",
			Level: Info,
		}

		logtesting.AssertLogEventsEqual(t, logEvent, *testLogger.LastEvent())
	})

	t.Run("ctx_fields", func(t *testing.T) {
		testLogger := &TestLoggerFacade{}

		logger := NewLoggerContext(testLogger, func(ctx context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key": "value",
			}
		})

		logger.InfoContext(context.Background(), "message")

		logEvent := LogEvent{
			Line:   "message",
			Level:  Info,
			Fields: map[string]interface{}{"key": "value"},
		}

		logtesting.AssertLogEventsEqual(t, logEvent, *testLogger.LastEvent())
	})

	t.Run("fields", func(t *testing.T) {
		testLogger := &TestLoggerFacade{}

		logger := NewLoggerContext(testLogger, func(ctx context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}
		})

		logger.InfoContext(context.Background(), "message", map[string]interface{}{
			"key":  "another value",
			"key3": "value3",
		})

		logEvent := LogEvent{
			Line:  "message",
			Level: Info,
			Fields: map[string]interface{}{
				"key":  "another value",
				"key2": "value2",
				"key3": "value3",
			},
		}

		logtesting.AssertLogEventsEqual(t, logEvent, *testLogger.LastEvent())
	})
}

func TestContextExtractors(t *testing.T) {
	extractor := ContextExtractors(
		func(_ context.Context) map[string]interface{} {
			return nil
		},
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}
		},
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key":  "another_value",
				"key3": "value3",
			}
		},
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key4": time.Minute,
			}
		},
	)

	expected := map[string]interface{}{
		"key":  "another_value",
		"key2": "value2",
		"key3": "value3",
		"key4": time.Minute,
	}

	if want, have := expected, extractor(context.Background()); !reflect.DeepEqual(want, have) {
		t.Errorf("unexpexted details\nexpected: %v\nactual:   %v", want, have)
	}
}

// nolint: gochecknoglobals
var printLoggerTestMap = map[string]*struct {
	logger func(logger Logger) *PrintLogger
	level  Level
}{
	"info": {
		logger: func(logger Logger) *PrintLogger {
			return NewPrintLogger(LevelFunc(logger, Info))
		},
		level: Info,
	},
	"error": {
		logger: NewErrorPrintLogger,
		level:  Error,
	},
}

func TestPrintLogger_Print(t *testing.T) {
	for name, test := range printLoggerTestMap {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			logger := &TestLogger{}
			printLogger := test.logger(logger)

			printLogger.Print("message", 1, "message", 2)

			event := LogEvent{
				Level: test.level,
				Line:  "message1message2",
			}

			logtesting.AssertLogEventsEqual(t, event, *(logger.LastEvent()))
		})
	}
}

func TestPrintLogger_Println(t *testing.T) {
	for name, test := range printLoggerTestMap {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			logger := &TestLogger{}
			printLogger := test.logger(logger)

			printLogger.Println("message", 1, "message", 2)

			event := LogEvent{
				Level: test.level,
				Line:  "message 1 message 2",
			}

			logtesting.AssertLogEventsEqual(t, event, *(logger.LastEvent()))
		})
	}
}

func TestPrintLogger_Printf(t *testing.T) {
	for name, test := range printLoggerTestMap {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			logger := &TestLogger{}
			printLogger := test.logger(logger)

			printLogger.Printf("this is my %s", "message")

			event := LogEvent{
				Level: test.level,
				Line:  "this is my message",
			}

			logtesting.AssertLogEventsEqual(t, event, *(logger.LastEvent()))
		})
	}
}

// TestLevels tests leveled logging capabilities.
func TestMessageLogger_Levels(t *testing.T) {
	tests := map[Level]struct {
		logFunc func(logger *MessageLogger, msg string)
	}{
		Trace: {
			logFunc: (*MessageLogger).Trace,
		},
		Debug: {
			logFunc: (*MessageLogger).Debug,
		},
		Info: {
			logFunc: (*MessageLogger).Info,
		},
		Warn: {
			logFunc: (*MessageLogger).Warn,
		},
		Error: {
			logFunc: (*MessageLogger).Error,
		},
	}

	for level, test := range tests {
		level, test := level, test

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
			testLogger := &TestLogger{}
			logger := NewMessageLogger(testLogger)

			test.logFunc(logger, "message")

			event := LogEvent{
				Line:  "message",
				Level: level,
			}

			logtesting.AssertLogEventsEqual(t, event, *(testLogger.LastEvent()))
		})
	}
}
