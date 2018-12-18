package keyvals

import "testing"

func TestFromMap(t *testing.T) {
	m := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	kvs := FromMap(m)

	for i := 0; i < len(m)*2; i += 2 {
		if got, want := kvs[i+1], m[kvs[i].(string)]; got != want {
			t.Errorf("key %q is expected to be followed by value %q, got %q", kvs[i], want, got)
		}
	}
}

func TestFromMap_Empty(t *testing.T) {
	kvs := FromMap(nil)

	if len(kvs) != 0 {
		t.Error("key-value pairs is expected to be empty")
	}

	if kvs == nil {
		t.Error("key-value pairs should not be nil (even if the map was)")
	}
}
