package logur

// mergeFields merges some current fields with incoming log fields.
func mergeFields(currentFields Fields, fields []map[string]interface{}) Fields {
	if len(fields) == 0 {
		return currentFields
	}

	if len(currentFields) == 0 {
		return fields[0]
	}

	// the maximum length of the map is the sum of the two map's length
	f := make(map[string]interface{}, len(fields)+len(fields[0]))

	for key, value := range currentFields {
		f[key] = value
	}

	for key, value := range fields[0] {
		f[key] = value
	}

	return f
}
