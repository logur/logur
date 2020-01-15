package example

import (
	"github.com/go-sql-driver/mysql"

	"logur.dev/logur"
)

func Example_mysqlDriver() {
	logger := logur.NoopLogger{} // choose an actual implementation

	_ = mysql.SetLogger(logur.NewErrorPrintLogger(logger))

	// Output:
}
