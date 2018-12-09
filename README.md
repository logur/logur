# Logur

[![Go Report Card](https://goreportcard.com/badge/github.com/goph/logur?style=flat-square)](https://goreportcard.com/report/github.com/goph/logur)
[![GolangCI](https://golangci.com/badges/github.com/goph/logur.svg)](https://golangci.com/r/github.com/goph/logur)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/logur)

[![Build Status](https://img.shields.io/travis/com/goph/logur.svg?style=flat-square)](https://travis-ci.com/goph/logur)

**Logur is an opinionated collection of logging best practices.**


## Features

- Unified logger interface
- Test logger for testing log event recording
- Noop logger for discarding log events
- `io.Writer` support
- Standard library logger support


## Installation

Logur uses [Go Modules](https://github.com/golang/go/wiki/Modules) introduced in Go 1.11, so the recommended way is
using `go get`:

```bash
$ go get github.com/goph/logur
```

Alternatively, you can install it via [Dep](https://golang.github.io/dep/):

```bash
$ dep ensure -add github.com/goph/logur
```


## Usage

An opinionated library should come with some best practices for usage and so does this one.

### Create a custom interface

Interfaces should be defined by the consumer, so the `Logger` interface in this library should not be used directly.
A custom interface should be defined instead:

```go
type MyLogger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
```

In a lucky scenario all logur loggers are compatible with the above interface, so you can just use them in your code:

```go
func main() {
    logger := logur.NewNoop()
    
    myFunc(logger)
}

func myFunc(logger MyLogger) {
	logger.Debug("myFunc ran")
}
```

In case you need structured logging, the interface becomes a bit more complicated:

```go
type MyLogger interface {
	Debug(args ...interface{})
	// ...
	WithFields(fields map[string]interface{}) MyLogger
}
```

As you can see `MyLogger` holds a reference to itself, which makes it incompatible with the logur interface.
The solution in this case is implementing a custom adapter:

```go
type myLogger struct {
	logger logur.Logger
}

func (l *myLogger) Debug(args ...interface{}) { l.logger.Debug(args...) }
// ...
func (l *myLogger) WithFields(fields map[string]interface{}) MyLogger { 
	return myLogger{l.logger.WithFields(logur.Fields(fields))}
}
```

Now you can easily use logur provided loggers inside your code:

```go
func main() {
    logger := &myLogger{logur.NewNoop()}
    
    myFunc(logger)
}

func myFunc(logger MyLogger) {
	logger.WithFields(map[string]interface{}{"key": "value"}).Debug("myFunc ran")
}
```

### Wrap helper functions with custom ones

In many cases it is unavoidable to maintain a simple integration layer between third-party libraries and your
application. Logur is no exception. In the previous section you saw how the main interface works with adapters,
but that's not all logur provides. It comes with a set of other tools (eg. creating a standard library logger)
to make logging easier. It might be tempting to just use them in your application, but writing an integration
layer is recommended, even around functions.

The following example creates a simple standard library logger for using as an HTTP server error log:

```go
func newStandardErrorLogger() *log.Logger {
	return logur.NewStandardLogger(logur.NewNoop(), logur.ErrorLevel, "", 0)
}

func main() {
	server := &http.Server{
		Handler: nil,
		ErrorLog: newStandardErrorLogger(),
	}
}
```


## FAQ

### Why not just X logger?

To be honest: mostly because I don't care. Loggers proliferated in the Go ecosystem in the past few years.
Each tries to convince you it's the most performant or the easiest to use.
But the fact is your application doesn't care which you use.
In fact, it's happier if it doesn't know anything about it at all.
Logger libraries (just like every third-party library) are external dependencies.
If you wire them into your application, it will be tied to the chosen libraries for ever.
That's why using a custom interface is a highly recommended practice.

Let's consider the following logger interface:

```go
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
```

You can easily create an interface like this and implement an adapter for the logging library of your choice.
In many cases you might not even have to do that as most logging libraries already implement these four methods,
so you can just use it without wiring it into your application which makes the actual library a less important detail.

### Why not go-kit logger?

Go-kit deserves it's own FAQ entry because for quite some time I was really fond of the its logger interface
and it was the closest thing to become an [official Go logging solution](https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU).
I still think it is great, because the interface is very simple, yet it's incredibly powerful.
But this simplicity is why I ultimately stopped using it as my primary logger.

Just a short recap of the interface itself:

```go
type Logger interface {
	Log(keyvals ...interface{}) error
}
```

It's really simple and easy to use in any application. Following Go's guidelines of using interfaces,
one can easily copy this interface and just use it to decouple the code from go-kit itself.

The problem with this interface appears when you try to do "advanced stuff"
(like structured logging or adding a level to a log event):

```go
import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// ...

logger := log.With(logger, "key", "value")
level.Info(logger).Log("msg", "message")
```

As you can see doing any kind of structured or leveled logging requires to import go-kit packages after all,
which takes us back to [Why not just X logger?](#why-not-just-x-logger).

In short: Using go-kit directly - no matter how awesome its interface is - suffers from the same problem as
using any other logging library.

One could implement all those functions for a custom interface based on go-kit,
but it probably isn't worth the hassle. Defining a more verbose, custom interface is a lot more easier to work with.
That being said, go-kit logger can very well serve as a perfect base for an implementation of that interface.

The [proposal](https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU) linked above contains many examples
why the authors ended up with an interface like this. Go check it out, you might just have the same use cases
which could make the go-kit interface a better fit than the one in this library.


## Inspiration

This package is heavily inspired by a set of logging libraries:

- [github.com/InVisionApp/go-logger](https://github.com/InVisionApp/go-logger)
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
- [github.com/go-kit/kit](https://github.com/go-kit/kit)


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
