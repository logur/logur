# Logur

[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-orange.svg?style=flat-square)](https://github.com/goph/logur)
[![Go Report Card](https://goreportcard.com/badge/github.com/goph/logur?style=flat-square)](https://goreportcard.com/report/github.com/goph/logur)
[![GolangCI](https://golangci.com/badges/github.com/goph/logur.svg)](https://golangci.com/r/github.com/goph/logur)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/logur)

[![Build Status](https://img.shields.io/travis/com/goph/logur.svg?style=flat-square)](https://travis-ci.com/goph/logur)

**Logur is an opinionated collection of logging best practices.**

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [FAQ](#faq)
    - [Why not just X logger?](#why-not-just-x-logger)
    - [Why not go-kit logger?](#why-not-go-kit-logger)
    - [Why not `logger.With(keyvals ...interface{})`?](#why-not-loggerwithkeyvals-interface)
    - [Why no `*f` (format) functions?](#why-no-f-format-functions)
    - [Why no `*ln` functions?](#why-no-ln-functions)
- [Inspiration](#inspiration)


## Features

- Unified logger interface
- Test logger for testing log event recording
- Noop logger for discarding log events
- `io.Writer` support
- Standard library logger support
- [github.com/goph/emperror](https://github.com/goph/emperror) compatible error handler
- Adapters for well-known logging libraries:
    * [hclog](https://github.com/hashicorp/go-hclog)
    * [go-kit log](https://github.com/go-kit/kit)
    * [logrus](https://github.com/sirupsen/logrus)
    * [zap](https://github.com/uber-go/zap)
    * [zerolog](https://github.com/rs/zerolog)
    * Next one contributed by You!


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

**TL;DR:** See example usage and best practices in [github.com/sagikazarmark/modern-go-application](https://github.com/sagikazarmark/modern-go-application).

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


### Why not `logger.With(keyvals ...interface{})`?

There is an increasing tendency of logging libraries implementing the following interface:

```go
type Logger interface {
	// ...
	With(keyvals ...interface{})
}
```

The arguments behind this interface are being simple and convenient, not as verbose as a map of fields.
There are also usual arguments *against* the alternative solutions, (eg. a `map[string]interface{}` endorsed by this library)
like not being able to order fields or being forced to import a concrete type, like `logur.Fields`.
Ultimately, this is not completely independent from personal taste, so one will always prefer one or the other.
(You can read more in the above linked [go-kit proposal](https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU)).

Let's take a look at these arguments one by one:

**1. Simplicity and verbosity**

This is something that's hard to argue with, but also a little bit subjective.
Here is a comparison of a single line context logging:

```go
logger = log.With(logger, "key", "value", "key2", "value")
logger = logger.WithFields(log.Fields{"key": "value", "key2": "value"})
```

Obviously the second one is more verbose, takes a bit more efforts to write, but this is rather a question of habits.

Let's take a look at a multiline example as well:

```go
logger = log.With(
	logger,
	"key", "value",
	"key2", "value",
)

logger = logger.WithFields(log.Fields{
	"key": "value",
	"key2": "value",
})
```

The difference is less visible in this case and harder to argue that one is better than the other.


**2. Ordering fields**

This is one of the less known arguments against maps in the context of logging, you can read about it in the
[go-kit proposal](https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU).

Since maps are unordered in Go, fields added to a log line might not always look like the same on the output.
Variadic (slice) arguments do not suffer from this issue. However, most implementations convert slices internally to maps,
so if ordering matters, most logging libraries won't work anyway.

Also, this library is not a one size fits all proposal and doesn't try to solve [100% of the problems](https://dev.to/nickjj/optimize-your-programming-decisions-for-the-95-not-the-5-2n42)
(unlike the official logger proposal), but rather aim for the most common use cases which doesn't include ordering of fields.


**3. Importing a concrete type**

As discussed earlier, one should not import the interface in this library directly to the application's code,
but instead use adapters, so this problem is basically nonexistent.


Comparing the slice and the map solution, there are also some arguments against using a variadic slice:

**1. Odd number of arguments**

The variadic slice interface implementation has to deal with the case when an odd number of arguments are passed to
the function. While the [go-kit proposal](https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU)
argues that this is extremely hard mistake to make, the risk is still there that the logs will lack some information.

**2. Converting the slice to key value pairs**

In order to display the context as key-value pairs the logging implementations has to convert the key parameters
to string in most of the cases (while the value parameter can be handled by the marshaling protocol).
This adds an extra step to outputting the logs (an extra loop going through all the parameters).
While I don't have scientific evidence proving one to be slower than the other (yet), it seems to be an unnecessary
complication at first.


### Why no `*f` (format) functions?

A previous version of this interface contained a set of functions that allowed messages to be formatted with arguments:

```go
type Logger interface {
	// ...
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
```

The reason why they were originally included in the interface is that most logging libraries implement these methods,
but experience showed that they are not used frequently. Also, nowadays structured logging is a better practice than
formatting log messages with structured data, thus these methods were removed from the core interface.


### Why no `*ln` functions?

Another common group of logging functions originally included in the interface is `*ln` function group:

```go
type Logger interface {
	// ...
	Traceln(args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
}
```

Usually separate log events are represented on separate lines anyway, so the added value is not newlines in this case,
but the different semantics between `fmt.Print` and `fmt.Println`.
See [this](https://play.golang.org/p/32GnCpXttbH) example illustrating the difference.

Common logging libraries include these functions, but experience showed they are not used frequently,
so they got removed.



## Inspiration

This package is heavily inspired by a set of logging libraries:

- [github.com/InVisionApp/go-logger](https://github.com/InVisionApp/go-logger)
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
- [github.com/go-kit/kit](https://github.com/go-kit/kit)


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
