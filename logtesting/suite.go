package logtesting

import (
	"testing"

	"logur.dev/logur"
	"logur.dev/logur/conformance"
)

// LoggerTestSuite implements a minimal set of tests that every logur compatible logger implementation must satisfy.
type LoggerTestSuite struct {
	LoggerFactory func(level logur.Level) (logur.Logger, func() []logur.LogEvent)

	TraceFallbackToDebug bool
}

func wrapLoggerFactory(
	fn func(logur.Level) (logur.Logger, func() []logur.LogEvent),
) func(logur.Level) (logur.Logger, conformance.TestLogger) {
	return func(level logur.Level) (logger logur.Logger, testLogger conformance.TestLogger) {
		logger, getEvents := fn(level)

		return logger, conformance.TestLoggerFunc(getEvents)
	}
}

// Execute executes the complete test suite.
//
// Deprecated: use logur.dev/conformance.TestSuite.Run.
func (s *LoggerTestSuite) Execute(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	conformance.TestSuite{
		LoggerFactory: wrapLoggerFactory(s.LoggerFactory),
		NoTraceLevel:  s.TraceFallbackToDebug,
	}.Run(t)
}

// TestLevels tests leveled logging capabilities.
//
// Deprecated: use logur.dev/conformance.TestSuite.RunLevelTest.
func (s *LoggerTestSuite) TestLevels(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	conformance.TestSuite{
		LoggerFactory: wrapLoggerFactory(s.LoggerFactory),
		NoTraceLevel:  s.TraceFallbackToDebug,
	}.RunLevelTest(t)
}

// TestLevelsContext tests leveled logging capabilities of a LoggerContext instance.
//
// Deprecated: use logur.dev/conformance.TestSuite.RunLevelContextTest.
func (s *LoggerTestSuite) TestLevelsContext(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	conformance.TestSuite{
		LoggerFactory: wrapLoggerFactory(s.LoggerFactory),
		NoTraceLevel:  s.TraceFallbackToDebug,
	}.RunLevelContextTest(t)
}

// TestLevelEnabler tests enabled levels.
// Note: this is not mandatory, incompatible loggers will be skipped.
//
// Deprecated: use logur.dev/conformance.TestSuite.RunLevelEnablerTest.
func (s *LoggerTestSuite) TestLevelEnabler(t *testing.T) {
	if s.LoggerFactory == nil {
		t.Fatal("logger factory is not configured")
	}

	conformance.TestSuite{
		LoggerFactory: wrapLoggerFactory(s.LoggerFactory),
		NoTraceLevel:  s.TraceFallbackToDebug,
	}.RunLevelEnablerTest(t)
}

// TestLevelEnablerUnknownReturnsTrue tests unknown enabled levels.
// Note: this is not mandatory, incompatible loggers will be skipped.
//
// Deprecated: use logur.dev/conformance.TestSuite.RunLevelEnablerTest.
func (s *LoggerTestSuite) TestLevelEnablerUnknownReturnsTrue(t *testing.T) {
	t.Skip("this test case is deprecated and is already covered by TestLevelEnabler")
}
