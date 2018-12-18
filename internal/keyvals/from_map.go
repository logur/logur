package keyvals

// FromMap converts a map of fields to a variadic key-value pair slice.
func FromMap(m map[string]interface{}) []interface{} {
	if len(m) == 0 {
		return make([]interface{}, 0)
	}

	keyvals := make([]interface{}, len(m)*2)
	i := 0

	for key, value := range m {
		keyvals[i] = key
		keyvals[i+1] = value

		i += 2
	}

	return keyvals
}
