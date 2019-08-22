module logur.dev/logur/benchmarks

go 1.12

require (
	github.com/go-kit/kit v0.9.0
	github.com/hashicorp/go-hclog v0.9.2
	github.com/rs/zerolog v1.15.0
	github.com/sirupsen/logrus v1.4.2
	go.uber.org/zap v1.10.0
	logur.dev/adapter/hclog v0.2.0
	logur.dev/adapter/kit v0.2.0
	logur.dev/adapter/logrus v0.2.0
	logur.dev/adapter/zap v0.2.0
	logur.dev/adapter/zerolog v0.2.0
	logur.dev/logur v0.15.0
)

replace logur.dev/logur => ../
