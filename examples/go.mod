module github.com/goph/logur/examples

go 1.12

require (
	github.com/bugsnag/bugsnag-go v1.5.3
	github.com/go-sql-driver/mysql v1.4.1
	github.com/goph/logur v0.11.2
	github.com/rollbar/rollbar-go v1.1.0
	google.golang.org/grpc v1.23.0
)

replace github.com/goph/logur => ../
