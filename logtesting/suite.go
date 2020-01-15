package logtesting

import (
	"context"
	"strings"
	"testing"

	"logur.dev/logur"
)

// nolint: gochecknoglobals
var testLevelMap = map[logur.Level]struct {
	logFunc    func(logger logur.Logger, msg string, fields ...map[string]interface{})
	logCtxFunc func(logger logur.LoggerContext, ctx context.Context, msg string, fields ...map[string]interface{})
}{
	logur.Trace: {
		logFunc:    logur.Logger.Trace,
		logCtxFunc: logur.LoggerContext.TraceContext,
	},
	logur.Debug: {
		logFunc:    logur.Logger.Debug,
		logCtxFunc: logur.LoggerContext.DebugContext,
	},
	logur.Info: {
		logFunc:    logur.Logger.Info,
		logCtxFunc: logur.LoggerContext.InfoContext,
	},
	logur.Warn: {
		logFunc:    logur.Logger.Warn,
		logCtxFunc: logur.LoggerContext.WarnContext,
	},
	logur.Error: {
		logFunc:    logur.Logger.Error,
		logCtxFunc: logur.LoggerContext.ErrorContext,
	},
}

// LoggerTestSuite implements a minimal set of tests that every logur compatible logger implementation must satisfy.
type LoggerTestSuite struct {
	LoggerFactory        func(level logur.Level) (logur.Logger, func() []logur.LogEvent)
	LoggerContextFactory func(level logur.Level) (logur.LoggerContext, func() []logur.LogEvent)
	TraceFallbackToDebug bool
}

// Execute executes the complete test suite.
func (s *LoggerTestSuite) Execute(t *testing.T) {
	t.Parallel()

	t.Run("Levels", s.TestLevels)

	if s.LoggerContextFactory != nil {
		t.Run("Levels", s.TestLevelsContext)
	}

	t.Run("LevelEnabler", s.TestLevelEnabler)
	t.Run("LevelEnabler_UnknownReturnsTrue", s.TestLevelEnablerUnknownReturnsTrue)
}

// TestLevels tests leveled logging capabilities.
func (s *LoggerTestSuite) TestLevels(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	for level, test := range testLevelMap {
		level, test := level, test

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
			if level == logur.Trace && s.TraceFallbackToDebug {
				level = logur.Debug
			}

			fields := map[string]interface{}{"key": "value"}

			logger, getLogEvents := s.LoggerFactory(logur.Trace)

			test.logFunc(logger, "message1message2", fields)

			logEvents := getLogEvents()

			if got, want := len(logEvents), 1; got != want {
				t.Fatalf("expected %d log events, got %d", want, got)
			}

			logEvent := logur.LogEvent{
				Line:   "message1message2",
				Level:  level,
				Fields: fields,
			}

			AssertLogEventsEqual(t, logEvent, logEvents[0])
		})
	}
}

// TestLevelsContext tests leveled logging capabilities of a LoggerContext instance.
func (s *LoggerTestSuite) TestLevelsContext(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	for level, test := range testLevelMap {
		level, test := level, test

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
			if level == logur.Trace && s.TraceFallbackToDebug {
				level = logur.Debug
			}

			fields := map[string]interface{}{"key": "value"}

			logger, getLogEvents := s.LoggerContextFactory(logur.Trace)

			test.logCtxFunc(logger, context.Background(), "message1message2", fields)

			logEvents := getLogEvents()

			if got, want := len(logEvents), 1; got != want {
				t.Fatalf("expected %d log events, got %d", want, got)
			}

			logEvent := logur.LogEvent{
				Line:   "message1message2",
				Level:  level,
				Fields: fields,
			}

			AssertLogEventsEqual(t, logEvent, logEvents[0])
		})
	}
}

// nolint: gochecknoglobals
var allLevels = []logur.Level{logur.Trace, logur.Debug, logur.Info, logur.Warn, logur.Error}

// TestLevelEnabler tests enabled levels.
// Note: this is not mandatory, incompatible loggers will be skipped.
func (s *LoggerTestSuite) TestLevelEnabler(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	for _, level := range allLevels {
		level := level

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
			if level == logur.Trace && s.TraceFallbackToDebug {
				return
			}

			logger, _ := s.LoggerFactory(level)

			enabler, ok := logger.(logur.LevelEnabler)
			if !ok {
				t.Skip("logger does not implement logur.LevelEnabler interface")
			}

			for _, l := range allLevels {
				if l == logur.Trace && s.TraceFallbackToDebug {
					continue
				}

				enabled := enabler.LevelEnabled(l)

				if l >= level && !enabled {
					t.Errorf("expected level %q to be enabled when the minimum level is %q", l, level)
				} else if l < level && enabled {
					t.Errorf("expected level %q to be disabled when the minimum level is %q", l, level)
				}
			}
		})
	}
}

// TestLevelEnablerUnknownReturnsTrue tests unknown enabled levels.
// Note: this is not mandatory, incompatible loggers will be skipped.
func (s *LoggerTestSuite) TestLevelEnablerUnknownReturnsTrue(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	logger, _ := s.LoggerFactory(logur.Trace)

	enabler, ok := logger.(logur.LevelEnabler)
	if !ok {
		t.Skip("logger does not implement logur.LevelEnabler interface")
	}

	enabled := enabler.LevelEnabled(logur.Level(999))

	if !enabled {
		t.Error("logur.LevelEnabler implementation should return true when it cannot detect a level")
	}
}
