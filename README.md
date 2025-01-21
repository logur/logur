> [!WARNING]
> This project is archived and no longer maintained. Consider using [`log/slog`](https://pkg.go.dev/log/slog) instead.
>
> Read more about why it was archived in [this post](https://sagikazarmark.com/blog/posts/less-is-more-archive-projects-for-a-better-open-source-ecosystem/).

![Logur](.github/logo.png?raw=true)

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go#logging)

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/logur/logur/CI?style=flat-square)](https://github.com/logur/logur/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/logur/logur?style=flat-square)](https://codecov.io/gh/logur/logur)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/logur?style=flat-square)](https://goreportcard.com/report/logur.dev/logur)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.12-61CFDD.svg?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/mod/logur.dev/logur)](https://pkg.go.dev/mod/logur.dev/logur)


**Logur is an opinionated collection of logging best practices.**

## Table of Contents

- [Preface](#preface)
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


## Preface

Logur is an opinionated logging library targeted at producing (application) logs.
It does not try to solve every problem around logging,
only a few considered important by the developers, thus it's highly opinionated.

The main focus of the library:

- provide a unified interface that does not require developers to import external dependencies
- encourage leveled and structured logging
- provide tools for easy integration of other logging libraries and components

Logur does not care about log output, you can use whatever library/formatting/forwarder you want
(ie. use an existing logging library with one of the [adapters](https://github.com/search?q=topic%3Aadapter+org%3Alogur)).

Despite the opinionated nature, Logur encourages you to create custom logging interfaces for your needs
and only use Logur as an integration layer/tool. Use the features you need/want and just omit the rest.

As mentioned above, Logur aims to cover only [95% of the use cases](https://dev.to/nickjj/optimize-your-programming-decisions-for-the-95-not-the-5-2n42),
so Logur might not be for you. Read on for more details.


## Features

- Unified logger interface
- Test logger for testing log event recording
- Noop logger for discarding log events
- `io.Writer` support
- Standard library logger support
- [Integrations](https://github.com/search?q=topic%3Aintegration+org%3Alogur) with well-known libraries:
    - [gRPC log](https://godoc.org/google.golang.org/grpc/grpclog)
    - [MySQL driver](https://github.com/go-sql-driver/mysql)
    - [Watermill](https://watermill.io/)
    - [InvisionApp logger](https://github.com/InVisionApp/go-logger) interface
    - [Bugsnag](https://bugsnag.com) [SDK](https://godoc.org/github.com/bugsnag/bugsnag-go) (logger, for error handling see [github.com/emperror/emperror](https://github.com/emperror/emperror))
    - [Rollbar](https://rollbar.com) [SDK](https://godoc.org/github.com/rollbar/rollbar-go) (logger, for error handling see [github.com/emperror/emperror](https://github.com/emperror/emperror))
    - [logr](https://github.com/go-logr/logr) logger interface
    - [go-kit](https://github.com/go-kit/kit) logger
    - [zap](https://github.com/uber-go/zap) logger
- [Adapters](https://github.com/search?q=topic%3Aadapter+org%3Alogur) for well-known logging libraries:
    - [hclog](https://github.com/hashicorp/go-hclog)
    - [go-kit log](https://github.com/go-kit/kit)
    - [logrus](https://github.com/sirupsen/logrus)
    - [zap](https://github.com/uber-go/zap)
    - [zerolog](https://github.com/rs/zerolog)
    - Next one contributed by You!


## Installation

Logur uses [Go Modules](https://github.com/golang/go/wiki/Modules) introduced in Go 1.11, so the recommended way is
using `go get`:

```bash
$ go get logur.dev/logur
```


## Usage

An opinionated library should come with some best practices for usage and so does this one.

**TL;DR:** See example usage and best practices in [github.com/sagikazarmark/modern-go-application](https://github.com/sagikazarmark/modern-go-application).
Also, check out the [examples](examples) package in this repository.

### Create a custom interface

Interfaces should be defined by the consumer, so the `Logger` interface in this library should not be used directly.
A custom interface should be defined instead:

```go
type MyLogger interface {
	Trace(msg string, fields ...map[string]interface{})
	Debug(msg string, fields ...map[string]interface{})
	Info(msg string, fields ...map[string]interface{})
	Warn(msg string, fields ...map[string]interface{})
	Error(msg string, fields ...map[string]interface{})
}
```

In a lucky scenario all Logur loggers are compatible with the above interface, so you can just use them in your code:

```go
func main() {
    logger := logur.NewNoopLogger()

    myFunc(logger)
}

func myFunc(logger MyLogger) {
	logger.Debug("myFunc ran")
	// OR
	logger.Debug("myFunc ran", map[string]interface{}{"key": "value"})
}
```

In case you need to populate the logger with some common context, the interface becomes a bit more complicated:

```go
type MyLogger interface {
	Trace(msg string, fields ...map[string]interface{})
	Debug(msg string, fields ...map[string]interface{})
	// ...
	WithFields(fields map[string]interface{}) MyLogger
}
```

As you can see `MyLogger` holds a reference to itself, which makes it incompatible with the Logur implementations.
The solution in this case is implementing a custom adapter:

```go
type myLogger struct {
	logger logur.Logger
}

func (l *myLogger) Debug(msg string, fields ...map[string]interface{}) { l.logger.Debug(msg, fields...) }
// ...
func (l *myLogger) WithFields(fields map[string]interface{}) MyLogger {
	return myLogger{logur.WithFields(l.logger, fields)}
}
```

Now you can easily use Logur provided loggers inside your code:

```go
func main() {
    logger := &myLogger{logur.NewNoopLogger()}

    myFunc(logger)
}

func myFunc(logger MyLogger) {
	logger.WithFields(map[string]interface{}{"key": "value"}).Debug("myFunc ran", nil)
}
```

### Wrap helper functions with custom ones

In many cases it is unavoidable to maintain a simple integration layer between third-party libraries and your
application. Logur is no exception. In the previous section you saw how the main interface works with adapters,
but that's not all Logur provides. It comes with a set of other tools (eg. a standard library logger compatible `io.Writer`)
to make logging easier. It might be tempting to just use them in your application, but writing an integration
layer is recommended, even around functions.

The following example creates a simple standard library logger for using as an HTTP server error log:

```go
func newStandardErrorLogger() *log.Logger {
	return logur.NewStandardLogger(logur.NewNoopLogger(), logur.ErrorLevel, "", 0)
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

To be honest: mostly because I don't care. Logging libraries proliferated in the Go ecosystem in the past few years.
Each tries to convince you it's the most performant or the easiest to use.
But the fact is your application doesn't care which you use.
In fact, it's happier if it doesn't know anything about it at all.
Logging libraries (just like every third-party library) are external dependencies.
If you wire them into your application, it will be tied to the chosen libraries forever.
That's why using a custom interface is a highly recommended practice.

Let's consider the following logger interface:

```go
type Logger interface {
	Trace(msg string, fields ...map[string]interface{})
	Debug(msg string, fields ...map[string]interface{})
	Info(msg string, fields ...map[string]interface{})
	Warn(msg string, fields ...map[string]interface{})
	Error(msg string, fields ...map[string]interface{})
}
```

You can easily create an interface like this and implement an adapter for the logging library of your choice
without wiring it into your application which makes the actual library a less important detail.


### Why not Go kit logger?

Go-kit deserves it's own entry because for quite some time I was really fond of the its logger interface
and it was the closest thing to become an [official Go logging solution](https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU).
I still think it is great, because the interface is very simple, yet it's incredibly powerful.
But this simplicity is why I ultimately stopped using it as my primary logger
(or I should say: stopped knowing that I actually use it).

Just a short recap of the interface itself:

```go
type Logger interface {
	Log(keyvals ...interface{}) error
}
```

It's really simple and easy to use in any application. Following Go's guidelines of using interfaces,
one can easily copy this interface and just use it to decouple the code from Go kit itself.

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

As you can see doing any kind of structured or leveled logging requires to import Go kit packages after all,
which takes us back to [Why not just X logger?](#why-not-just-x-logger).

In short: Using Go kit directly - no matter how awesome its interface is - suffers from the same problem as
using any other logging library.

One could implement all those functions for a custom interface based on Go kit,
but it probably isn't worth the hassle. Defining a more verbose, custom interface is a lot easier to work with.
That being said, Go kit logger can very well serve as a perfect base for an implementation of that interface.

The [proposal](https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU) linked above contains many examples
why the authors ended up with an interface like this. Go check it out, you might just have the same use cases
which could make the Go kit interface a better fit than the one in this library.


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
like not being able to order fields, being forced to import a concrete type, like `logur.Fields`
or (quite often) performance questions.

Ultimately, this is not completely independent from personal taste, so one will always prefer one or the other.
(You can read more in the above linked [Go kit proposal](https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU)).

Let's take a look at these arguments one by one:

**1. Simplicity and verbosity**

This is something that's hard to argue with, but also a little bit subjective.
Here is a comparison of a single line context logging:

```go
logger = log.With(logger, "key", "value", "key2", "value")
logger = logur.WithFields(logger, map[string]interface{}{"key": "value", "key2": "value"})
```

Obviously the second one is more verbose, takes a bit more efforts to write, but it is rather a question of habits.

Let's take a look at a multiline example as well:

```go
logger = log.With(logger,
	"key", "value",
	"key2", "value",
)

logger = logur.WithFields(logger, map[string]interface{}{
	"key": "value",
	"key2": "value",
})
```

The difference is less visible in this case and harder to argue that one is better than the other.

Also, defining a custom type is relatively easy which makes the difference even smaller:

```go
logger = log.With(logger,
	"key", "value",
	"key2", "value",
)

type LogFields map[string]interface{}

logger = logur.WithFields(logger, LogFields{
	"key": "value",
	"key2": "value",
})
```


**2. Ordering fields**

This is one of the less known arguments against maps in the context of logging, you can read about it in the
[Go kit proposal](https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU).

Since maps are unordered in Go, fields added to a log line might not always look like the same on the output.
Variadic (slice) arguments do not suffer from this issue. However, most implementations convert slices internally to maps,
so if ordering matters, most logging libraries won't work anyway.

Also, this library is not a one size fits all proposal and doesn't try to solve [100% of the problems](https://dev.to/nickjj/optimize-your-programming-decisions-for-the-95-not-the-5-2n42)
(unlike the official logger proposal), but rather aim for the most common use cases which doesn't include ordering of fields.


**3. Performance**

Comparing the performance of different solutions is not easy and depends heavily on the interface.

For example, Uber's [Zap](https://github.com/uber-go/zap) comes with a rich interface
(thus requires you to couple your code to Zap), but also promises zero-allocation in certain scenarios,
making it an extremely fast logging library. If being *very fast* is a requirement for you,
even at the expense tight coupling, then Zap is a great choice.

More generic interfaces often use ... well `interface{}` for structured context,
but [interface allocations](https://commaok.xyz/post/interface-allocs/) are much cheaper now.

The performance debate these days is often between two approaches:

- variadic slices (`...interface{}`)
- maps (`map[string]interface{}`)

Specifically, how the provided context can be *merged* with an existing, internal context of the logger.
Admittedly, `append`ing to a slice is much cheaper than merging a map, so if performance is crucial,
using an interface with variadic slices will always be **slightly** faster.
But that difference in performance is negligible in most of the cases, so you won't even notice it,
unless you start logging with hundreds of structured context fields (which will have other problems anyway).

There is a problem with Logur adapters though: since the interface uses maps for structured context,
libraries like Zap that use variadic slices does not perform too well because of the map -> slice conversion.
In most of the cases this should still be acceptable, but you should be aware of this fact when choosing an adapter.

Partly because of this, I plan to add a `KVLogger` interface that uses variadic slices for structured context.
Obviously, map based libraries will suffer from the same performance penalty, but then the choice of the interface
will be up to the developer.


Comparing the slice and the map solution, there are also some arguments against using a variadic slice:

**1. Odd number of arguments**

The variadic slice interface implementation has to deal with the case when an odd number of arguments are passed to
the function. While the [Go kit proposal](https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU)
argues that this is extremely hard mistake to make, the risk is still there that the logs will lack some information.

**2. Converting the slice to key value pairs**

In order to display the context as key-value pairs the logging implementations has to convert the key parameters
to string in most of the cases (while the value parameter can be handled by the marshaling protocol).
This adds an extra step to outputting the logs (an extra loop going through all the parameters).
While there is no scientific evidence proving one to be slower than the other (yet), it seems to be an unnecessary
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


## Development

Contributions are welcome! :)

1. Clone the repository
1. Make changes on a new branch
1. Run the test suite:
    ```bash
    ./pleasew build
    ./pleasew test
    ./pleasew gotest
    ./pleasew lint
    ```
1. Commit, push and open a PR


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
