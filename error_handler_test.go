package logur_test

import (
	"testing"

	. "github.com/goph/logur"
	logtesting "github.com/goph/logur/testing"
)

type errorStub struct {
	parent  error
	msg     string
	context []interface{}
}

func (e *errorStub) Error() string {
	return e.msg
}

func (e *errorStub) Cause() error {
	return e.parent
}

func (e *errorStub) Context() []interface{} {
	return e.context
}

type errorsStub struct {
	errors []error
}

func (e *errorsStub) Error() string {
	return "multiple error happened"
}

func (e *errorsStub) Errors() []error {
	return e.errors
}

func TestErrorHandler_Handle(t *testing.T) {
	tests := map[error][]LogEvent{
		&errorStub{
			msg: "error",
		}: {
			{
				Line:  "error",
				Level: Error,
			},
		},
		&errorsStub{
			errors: []error{
				&errorStub{msg: "error 1"},
				&errorStub{msg: "error 2"},
			},
		}: {
			{
				Line:   "error 1",
				Level:  Error,
				Fields: map[string]interface{}{"parent": "multiple error happened"},
			},
			{
				Line:   "error 2",
				Level:  Error,
				Fields: map[string]interface{}{"parent": "multiple error happened"},
			},
		},
		&errorStub{
			msg:     "error",
			context: []interface{}{"key", "value"},
		}: {
			{
				Line:   "error",
				Level:  Error,
				Fields: map[string]interface{}{"key": "value"},
			},
		},
		&errorStub{
			msg:     "error 1",
			context: []interface{}{"key", "value 1"},
			parent: &errorStub{
				msg:     "error 1",
				context: []interface{}{"key", "value 2"},
			},
		}: {
			{
				Line:   "error 1",
				Level:  Error,
				Fields: map[string]interface{}{"key": "value 1"},
			},
		},
		&errorStub{
			msg:     "error 1",
			context: []interface{}{"key 1", "value 1"},
			parent: &errorStub{
				msg:     "error 2",
				context: []interface{}{"key 2", "value 2"},
			},
		}: {
			{
				Line:   "error 1",
				Level:  Error,
				Fields: map[string]interface{}{"key 1": "value 1", "key 2": "value 2"},
			},
		},
		&errorStub{
			msg:     "error",
			context: []interface{}{"key"},
		}: {
			{
				Line:   "error",
				Level:  Error,
				Fields: map[string]interface{}{"key": nil},
			},
		},
		&errorStub{
			msg:     "error 1",
			context: []interface{}{"key 1", "value 1"},
			parent: &errorStub{
				msg:     "error 2",
				context: []interface{}{"key 2", "value 2"},
				parent: &errorStub{
					msg:     "error 3",
					context: []interface{}{"key 3"},
					parent: &errorStub{
						msg: "error 4",
						context: []interface{}{
							"key 1", "value 1",
							"key 2", "value 2",
							"key 3", "value 3",
							"key 4", "value 4",
						},
					},
				},
			},
		}: {
			{
				Line:   "error 1",
				Level:  Error,
				Fields: map[string]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": nil, "key 4": "value 4"},
			},
		},
	}

	for err, expectedEvents := range tests {
		err, expectedEvents := err, expectedEvents

		t.Run("", func(t *testing.T) {
			logger := NewTestLogger()
			handler := NewErrorHandler(logger)

			handler.Handle(err)

			if got, want := logger.Count(), len(expectedEvents); got != want {
				t.Fatalf("recorded %d events, but expected %d", got, want)
			}

			events := logger.Events()

			for key, expectedEvent := range expectedEvents {
				logtesting.AssertLogEventsEqual(t, expectedEvent, events[key])
			}
		})
	}
}
