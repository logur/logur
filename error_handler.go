package logur

import (
	"fmt"
	"reflect"
)

// ErrorHandler is a github.com/goph/emperror compatible error handler for logging errors.
type ErrorHandler struct {
	logger Logger
}

// NewErrorHandler returns a new ErrorHandler.
func NewErrorHandler(logger Logger) *ErrorHandler {
	if logger == nil {
		logger = NewNoop()
	}

	return &ErrorHandler{
		logger: logger,
	}
}

// Handle records an error event and forwards it to the underlying logger.
func (h *ErrorHandler) Handle(err error) {
	if err == nil {
		return
	}

	fields := make(map[string]interface{})

	// Extract context from the error and attach it to the log
	if keyvals := errorContext(err); len(keyvals) > 0 {
		fields = errorContextToMap(keyvals)
	}

	type errorCollection interface {
		Errors() []error
	}

	if errs, ok := err.(errorCollection); ok {
		for _, e := range errs.Errors() {
			fields := fields
			fields["parent"] = err.Error()

			h.logger.Error(e.Error(), fields)
		}
	} else {
		h.logger.Error(err.Error(), fields)
	}
}

// errorContext extracts the context key-value pairs from an error (or error chain).
func errorContext(err error) []interface{} {
	type contextor interface {
		Context() []interface{}
	}

	var kvs []interface{}

	errorForEachCause(err, func(err error) bool {
		if cerr, ok := err.(contextor); ok {
			kv := cerr.Context()
			if len(kv)%2 == 1 {
				kv = append(kv, nil)
			}

			kvs = append(kv, kvs...)
		}

		return true
	})

	return kvs
}

// errorForEachCause loops through an error chain and calls a function for each of them,
// starting with the topmost one.
//
// The function can return false to break the loop before it ends.
func errorForEachCause(err error, fn func(err error) bool) {
	// causer is the interface defined in github.com/pkg/errors for specifying a parent error.
	type causer interface {
		Cause() error
	}

	for err != nil {
		continueLoop := fn(err)
		if !continueLoop {
			break
		}

		cause, ok := err.(causer)
		if !ok {
			break
		}

		err = cause.Cause()
	}
}

func errorContextToMap(keyvals []interface{}) map[string]interface{} {
	m := map[string]interface{}{}

	if len(keyvals) == 0 {
		return m
	}

	if len(keyvals)%2 == 1 {
		keyvals = append(keyvals, nil)
	}

	for i := 0; i < len(keyvals); i += 2 {
		mergeErrorContext(m, keyvals[i], keyvals[i+1])
	}

	return m
}

func mergeErrorContext(dst map[string]interface{}, k, v interface{}) {
	var key string

	switch x := k.(type) {
	case string:
		key = x
	case fmt.Stringer:
		key = safeString(x)
	default:
		key = fmt.Sprint(x)
	}

	switch x := v.(type) {
	case error:
		v = safeError(x)
	case fmt.Stringer:
		v = safeString(x)
	}

	dst[key] = v
}

func safeString(str fmt.Stringer) (s string) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(str); v.Kind() == reflect.Ptr && v.IsNil() {
				s = "NULL"
			} else {
				panic(panicVal)
			}
		}
	}()

	s = str.String()

	return
}

func safeError(err error) (s interface{}) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(err); v.Kind() == reflect.Ptr && v.IsNil() {
				s = nil
			} else {
				panic(panicVal)
			}
		}
	}()

	s = err.Error()

	return
}
