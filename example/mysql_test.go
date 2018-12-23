package example

import (
	"github.com/go-sql-driver/mysql"
	"github.com/goph/logur"
)

func ExampleMysqlDriver() {
	logger := logur.NewNoopLogger() // choose an actual implementation

	_ = mysql.SetLogger(logur.NewPrintErrorLogger(logger))

	// Output:
}
