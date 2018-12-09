/*
Package logur is an opinionated collection of logging best practices.

Features:
	- Unified logger interface
	- Test logger for testing log event recording
	- Noop logger for discarding log events
	- `io.Writer` support
	- Standard library logger support

An opinionated library should come with some best practices for usage and so does this one.


Create a custom interface

Interfaces should be defined by the consumer, so the `Logger` interface in this library should not be used directly.
A custom interface should be defined instead:

	type MyLogger interface {
		Debug(args ...interface{})
		Info(args ...interface{})
		Warn(args ...interface{})
		Error(args ...interface{})
	}

In a lucky scenario all logur loggers are compatible with the above interface, so you can just use them in your code:

	func main() {
		logger := logur.NewNoop()

		myFunc(logger)
	}

	func myFunc(logger MyLogger) {
		logger.Debug("myFunc ran")
	}

In case you need structured logging, the interface becomes a bit more complicated:

	type MyLogger interface {
		Debug(args ...interface{})
		// ...
		WithFields(fields map[string]interface{}) MyLogger
	}

As you can see `MyLogger` holds a reference to itself, which makes it incompatible with the logur interface.
The solution in this case is implementing a custom adapter:

	type myLogger struct {
		logger logur.Logger
	}

	func (l *myLogger) Debug(args ...interface{}) { l.logger.Debug(args...) }
	// ...
	func (l *myLogger) WithFields(fields map[string]interface{}) MyLogger {
		return myLogger{l.logger.WithFields(logur.Fields(fields))}
	}

Now you can easily use logur provided loggers inside your code:

	func main() {
		logger := &myLogger{logur.NewNoop()}

		myFunc(logger)
	}

	func myFunc(logger MyLogger) {
		logger.WithFields(map[string]interface{}{"key": "value"}).Debug("myFunc ran")
	}


Wrap helper functions with custom ones

In many cases it is unavoidable to maintain a simple integration layer between third-party libraries and your
application. Logur is no exception. In the previous section you saw how the main interface works with adapters,
but that's not all logur provides. It comes with a set of other tools (eg. creating a standard library logger)
to make logging easier. It might be tempting to just use them in your application, but writing an integration
layer is recommended, even around functions.

The following example creates a simple standard library logger for using as an HTTP server error log:

	func newStandardErrorLogger() *log.Logger {
		return logur.NewStandardLogger(logur.NewNoop(), logur.ErrorLevel, "", 0)
	}

	func main() {
		server := &http.Server{
			Handler: nil,
			ErrorLog: newStandardErrorLogger(),
		}
	}
*/
package logur
