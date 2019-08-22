# Logur Examples

[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/logur/example)

This package contains various usage and implementation examples for Logur:

- [MySQL driver](https://github.com/go-sql-driver/mysql) integration [example](mysql_test.go)
- [gRPC log](https://godoc.org/google.golang.org/grpc/grpclog) integration [example](grpc_test.go)
- [Bugsnag](https://bugsnag.com) integration [example](bugsnag_test.go)
- [Rollbar](https://rollbar.com) integration [example](rollbar_test.go)

It also contains a custom example logger you might want to use in your application.
Feel free to copy and customize it to your needs.

For a real application example, check out [sagikazarmark/modern-go-application](https://github.com/sagikazarmark/modern-go-application):

- [Interface](https://github.com/sagikazarmark/modern-go-application/blob/65edb2b/internal/common/logger.go#L7-L29)
- [Implementation](https://github.com/sagikazarmark/modern-go-application/blob/65edb2b/internal/common/commonadapter/logger.go)
