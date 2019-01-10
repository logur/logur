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

// ContextualLogger holds a context and passes it to the underlying logger when a log event is recorded.
type ContextualLogger struct {
	logger       Logger
	fields       map[string]interface{}
	levelEnabler LevelEnabler
}

// Trace implements the logur.Logger interface.
func (l *ContextualLogger) Trace(msg string, fields ...map[string]interface{}) {
	if !l.levelEnabled(Trace) {
		return
	}

	var f = l.fields
	if len(fields) > 0 {
		f = l.mergeFields(fields[0])
	}

	l.logger.Trace(msg, f)
}

// Debug implements the logur.Logger interface.
func (l *ContextualLogger) Debug(msg string, fields ...map[string]interface{}) {
	if !l.levelEnabled(Debug) {
		return
	}

	var f = l.fields
	if len(fields) > 0 {
		f = l.mergeFields(fields[0])
	}

	l.logger.Debug(msg, f)
}

// Info implements the logur.Logger interface.
func (l *ContextualLogger) Info(msg string, fields ...map[string]interface{}) {
	if !l.levelEnabled(Info) {
		return
	}

	var f = l.fields
	if len(fields) > 0 {
		f = l.mergeFields(fields[0])
	}

	l.logger.Info(msg, f)
}

// Warn implements the logur.Logger interface.
func (l *ContextualLogger) Warn(msg string, fields ...map[string]interface{}) {
	if !l.levelEnabled(Warn) {
		return
	}

	var f = l.fields
	if len(fields) > 0 {
		f = l.mergeFields(fields[0])
	}

	l.logger.Warn(msg, f)
}

// Error implements the logur.Logger interface.
func (l *ContextualLogger) Error(msg string, fields ...map[string]interface{}) {
	if !l.levelEnabled(Error) {
		return
	}

	var f = l.fields
	if len(fields) > 0 {
		f = l.mergeFields(fields[0])
	}

	l.logger.Error(msg, f)
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
