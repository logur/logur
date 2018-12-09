package logur

import (
	"testing"
)

func TestLevel_String(t *testing.T) {
	tests := map[string]Level{
		"trace": TraceLevel,
		"debug": DebugLevel,
		"info":  InfoLevel,
		"warn":  WarnLevel,
		"error": ErrorLevel,
	}

	for levelName, level := range tests {
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
		"trace":   TraceLevel,
		"debug":   DebugLevel,
		"info":    InfoLevel,
		"warn":    WarnLevel,
		"warning": WarnLevel,
		"error":   ErrorLevel,
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
