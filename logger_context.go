package logur

// WithFields returns a new logger instance that attaches the given fields to every subsequent log call.
func WithFields(logger Logger, fields map[string]interface{}) Logger {
	if len(fields) == 0 {
		return logger
	}

	// Do not add a new layer
	// Create a new logger instead with the parent fields
	if l, ok := logger.(*fieldLogger); ok && len(l.fields) > 0 {
		_fields := make(map[string]interface{}, len(l.fields)+len(fields))

		for key, value := range l.fields {
			_fields[key] = value
		}

		for key, value := range fields {
			_fields[key] = value
		}

		fields = _fields
		logger = l.logger
	}

	l := &fieldLogger{logger: logger, fields: fields}

	if levelEnabler, ok := logger.(LevelEnabler); ok {
		l.levelEnabler = levelEnabler
	}

	return l
}

// fieldLogger holds a context and passes it to the underlying logger when a log event is recorded.
type fieldLogger struct {
	logger       Logger
	fields       map[string]interface{}
	levelEnabler LevelEnabler
}

// Trace implements the logur.Logger interface.
func (l *fieldLogger) Trace(msg string, fields ...map[string]interface{}) {
	l.log(Trace, l.logger.Trace, msg, fields)
}

// Debug implements the logur.Logger interface.
func (l *fieldLogger) Debug(msg string, fields ...map[string]interface{}) {
	l.log(Debug, l.logger.Debug, msg, fields)
}

// Info implements the logur.Logger interface.
func (l *fieldLogger) Info(msg string, fields ...map[string]interface{}) {
	l.log(Info, l.logger.Info, msg, fields)
}

// Warn implements the logur.Logger interface.
func (l *fieldLogger) Warn(msg string, fields ...map[string]interface{}) {
	l.log(Warn, l.logger.Warn, msg, fields)
}

// Error implements the logur.Logger interface.
func (l *fieldLogger) Error(msg string, fields ...map[string]interface{}) {
	l.log(Error, l.logger.Error, msg, fields)
}

// log deduplicates some field logger code.
func (l *fieldLogger) log(level Level, logFunc LogFunc, msg string, fields []map[string]interface{}) {
	if !l.levelEnabled(level) {
		return
	}

	var f = l.fields
	if len(fields) > 0 {
		f = l.mergeFields(fields[0])
	}

	logFunc(msg, f)
}

func (l *fieldLogger) mergeFields(fields map[string]interface{}) map[string]interface{} {
	if len(fields) == 0 { // Not having any fields passed to the log function has a higher chance
		return l.fields
	}

	if len(l.fields) == 0 { // This is possible too, but has a much lower probability
		return fields
	}

	f := make(map[string]interface{}, len(fields)+len(l.fields))

	for key, value := range l.fields {
		f[key] = value
	}

	for key, value := range fields {
		f[key] = value
	}

	return f
}

func (l *fieldLogger) levelEnabled(level Level) bool {
	if l.levelEnabler != nil {
		return l.levelEnabler.LevelEnabled(level)
	}

	return true
}
