package watermilllog

import (
	"fmt"
	"io"
)

// errorWithMessage annotates err with a new message.
// Copied from https://github.com/emperror/errors/blob/acc47e5/errors.go#L169-L231
func errorWithMessage(err error, message string) error {
	if err == nil {
		return nil
	}

	return &withMessage{
		error: err,
		msg:   message,
	}
}

type withMessage struct {
	error error
	msg   string
}

func (w *withMessage) Error() string { return w.msg + ": " + w.error.Error() }
func (w *withMessage) Cause() error  { return w.error }
func (w *withMessage) Unwrap() error { return w.error }

// nolint: errcheck
func (w *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.error)
			io.WriteString(s, w.msg)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, w.Error())
	}
}
