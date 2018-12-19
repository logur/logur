package logur

// WithFields returns a new contextual logger instance with context added to it.
func WithFields(logger Logger, fields map[string]interface{}) Logger {
	if len(fields) == 0 {
		return logger
	}

	// Do not add a new layer
	// Create a new logger instead with the parent fields
	if ctxlogger, ok := logger.(*ContextualLogger); ok && len(ctxlogger.fields) > 0 {
		_fields := make(map[string]interface{}, len(ctxlogger.fields)+len(fields))

		for key, value := range ctxlogger.fields {
			_fields[key] = value
		}

		for key, value := range fields {
			_fields[key] = value
		}

		fields = _fields
		logger = ctxlogger.logger
	}

	ctxlogger := &ContextualLogger{logger: logger, fields: fields}

	if levelEnabler, ok := logger.(LevelEnabler); ok {
		ctxlogger.levelEnabler = levelEnabler
	}

	return ctxlogger
}

type ContextualLogger struct {
	logger       Logger
	fields       map[string]interface{}
	levelEnabler LevelEnabler
}

func (l *ContextualLogger) Trace(msg string, fields map[string]interface{}) {
	if !l.levelEnabled(Trace) {
		return
	}

	l.logger.Trace(msg, l.mergeFields(fields))
}

func (l *ContextualLogger) Debug(msg string, fields map[string]interface{}) {
	if !l.levelEnabled(Debug) {
		return
	}

	l.logger.Debug(msg, l.mergeFields(fields))
}

func (l *ContextualLogger) Info(msg string, fields map[string]interface{}) {
	if !l.levelEnabled(Info) {
		return
	}

	l.logger.Info(msg, l.mergeFields(fields))
}

func (l *ContextualLogger) Warn(msg string, fields map[string]interface{}) {
	if !l.levelEnabled(Warn) {
		return
	}

	l.logger.Warn(msg, l.mergeFields(fields))
}

func (l *ContextualLogger) Error(msg string, fields map[string]interface{}) {
	if !l.levelEnabled(Error) {
		return
	}

	l.logger.Error(msg, l.mergeFields(fields))
}

func (l *ContextualLogger) mergeFields(fields map[string]interface{}) map[string]interface{} {
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

func (l *ContextualLogger) levelEnabled(level Level) bool {
	if l.levelEnabler != nil {
		return l.levelEnabler.LevelEnabled(level)
	}

	return true
}
