package conformance

import (
	"context"
	"strings"
	"testing"

	"logur.dev/logur"
)

// TestLogger acts as a recorder for the Logger under test.
type TestLogger interface {
	// Events returns all recorded events in the logger.
	Events() []logur.LogEvent
}

// TestLoggerFunc converts an event function to a TestLogger if it's definition compatible with the interface.
type TestLoggerFunc func() []logur.LogEvent

func (fn TestLoggerFunc) Events() []logur.LogEvent {
	return fn()
}

// TestSuite runs a set of conformance tests against a logger.
type TestSuite struct {
	LoggerFactory func(level logur.Level) (logur.Logger, TestLogger)

	NoTraceLevel bool
}

// Run executes the complete test suite.
func (s TestSuite) Run(t *testing.T) {
	t.Parallel()

	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	t.Run("Level", s.RunLevelTest)
	t.Run("LevelContext", s.RunLevelContextTest)
	t.Run("LevelsEnabler", s.RunLevelEnablerTest)
}

// RunLevelTest tests leveled logging capabilities of a Logger.
func (s TestSuite) RunLevelTest(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	for level, test := range testLevelMap {
		level, test := level, test

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
			if level == logur.Trace && s.NoTraceLevel {
				level = logur.Debug
			}

			const message = "message"
			fields := logur.Fields{"key": "value"}

			logger, testLogger := s.LoggerFactory(logur.Trace)

			test.logFunc(logger, message, fields)

			logEvents := testLogger.Events()

			if want, have := 1, len(logEvents); want != have {
				t.Errorf("unexpexted log event count\nexpected: %v\nactual:   %v", want, have)
			}

			logEvent := logur.LogEvent{
				Line:   message,
				Level:  level,
				Fields: fields,
			}

			if err := logEvents[0].AssertEquals(logEvent); err != nil {
				t.Errorf("%+v", err)
			}
		})
	}
}

// RunLevelContextTest tests leveled logging capabilities of a LoggerContext.
// Note: this is not mandatory, incompatible loggers will be skipped.
func (s TestSuite) RunLevelContextTest(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	logger, _ := s.LoggerFactory(logur.Trace)

	if _, ok := logger.(logur.LoggerContext); !ok {
		t.Skip("logger does not implement logur.LoggerContext interface")
	}

	for level, test := range testLevelMap {
		level, test := level, test

		t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
			if level == logur.Trace && s.NoTraceLevel {
				level = logur.Debug
			}

			const message = "message"
			fields := logur.Fields{"key": "value"}

			logger, testLogger := s.LoggerFactory(logur.Trace)

			loggerCtx, ok := logger.(logur.LoggerContext)
			if !ok {
				t.Skip("logger does not implement logur.LoggerContext interface")
			}

			test.logCtxFunc(loggerCtx, context.Background(), message, fields)

			logEvents := testLogger.Events()

			if want, have := 1, len(logEvents); want != have {
				t.Errorf("unexpexted log event count\nexpected: %v\nactual:   %v", want, have)
			}

			logEvent := logur.LogEvent{
				Line:   message,
				Level:  level,
				Fields: fields,
			}

			if err := logEvents[0].AssertEquals(logEvent); err != nil {
				t.Errorf("%+v", err)
			}
		})
	}
}

// RunLevelEnablerTest tests enabled levels.
// Note: this is not mandatory, incompatible loggers will be skipped.
// nolint: gocognit
func (s TestSuite) RunLevelEnablerTest(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	logger, _ := s.LoggerFactory(logur.Trace)

	if _, ok := logger.(logur.LevelEnabler); !ok {
		t.Skip("logger does not implement logur.LevelEnabler interface")
	}

	t.Run("Levels", func(t *testing.T) {
		for _, level := range allLevels {
			level := level

			t.Run(strings.ToTitle(level.String()), func(t *testing.T) {
				if level == logur.Trace && s.NoTraceLevel {
					return
				}

				logger, _ := s.LoggerFactory(level)

				enabler, ok := logger.(logur.LevelEnabler)
				if !ok {
					t.Skip("logger does not implement logur.LevelEnabler interface")
				}

				for _, l := range allLevels {
					if l == logur.Trace && s.NoTraceLevel {
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
	})

	t.Run("UnknownLevel", func(t *testing.T) {
		logger, _ := s.LoggerFactory(logur.Trace)

		enabler, ok := logger.(logur.LevelEnabler)
		if !ok {
			t.Skip("logger does not implement logur.LevelEnabler interface")
		}

		enabled := enabler.LevelEnabled(logur.Level(999))

		if !enabled {
			t.Error("LevelEnabler should return true when the level is not supported")
		}
	})
}
