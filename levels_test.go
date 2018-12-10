package logur

import (
	"testing"
)

// nolint: gochecknoglobals
var levelMap = map[string]Level{
	"trace": Trace,
	"debug": Debug,
	"info":  Info,
	"warn":  Warn,
	"error": Error,
}

func TestLevel_String(t *testing.T) {
	for levelName, level := range levelMap {
		levelName, level := levelName, level

		t.Run(levelName, func(t *testing.T) {
			if levelString := level.String(); levelString != levelName {
				t.Errorf("level %q does not match the expected %q", levelString, levelName)
			}
		})
	}
}

func TestLevel_String_Unknown(t *testing.T) {
	level := Level(999)

	if levelString := level.String(); levelString != "unknown" {
		t.Errorf("level %q does not match the expected \"unknown\"", levelString)
	}
}

func TestParseAndUnmarshalLevel(t *testing.T) {
	tests := map[string]Level{
		"trace":   Trace,
		"debug":   Debug,
		"info":    Info,
		"warn":    Warn,
		"warning": Warn,
		"error":   Error,
	}

	for levelName, level := range tests {
		levelName, level := levelName, level

		t.Run("parse:"+levelName, func(t *testing.T) {
			parsedLevel, ok := ParseLevel(levelName)

			if !ok {
				t.Fatalf("parsing level failed: %q", levelName)
			}

			if parsedLevel != level {
				t.Errorf("parsed level %q does not match the expected %q", parsedLevel, level)
			}
		})

		t.Run("unmarshal:"+levelName, func(t *testing.T) {
			var l Level

			err := l.UnmarshalText([]byte(levelName))
			if err != nil {
				t.Fatal("unmarshaling level failed:", err.Error())
			}

			if l != level {
				t.Errorf("unmarshaled level %q does not match the expected %q", l, level)
			}
		})
	}
}

func TestParseLevel_Unknown(t *testing.T) {
	_, ok := ParseLevel("unknown")

	if ok {
		t.Error("parsing unknown levels should fail")
	}
}

func TestLevelFunc(t *testing.T) {
	for levelName, level := range levelMap {
		levelName, level := levelName, level

		t.Run(levelName, func(t *testing.T) {
			logger := NewTestLogger()

			logFunc := LevelFunc(logger, level)
			const msg = "message"

			logFunc(msg)

			if logger.Count() < 1 {
				t.Fatal("logger did not record any events")
			}

			event := logger.LastEvent()

			if event.Level != level {
				t.Errorf("expected level %q instead of %q", level.String(), event.Level.String())
			}

			if got, want := event.Line, msg; got != want {
				t.Errorf("expected message %q instead of %q", want, got)
			}
		})
	}
}
