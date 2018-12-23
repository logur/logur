/*
Package mysqllog provides a MySQL database driver logger.

The MySQL driver logs critical errors through a logger that implements the
github.com/go-sql-driver/mysql.Logger interface.

With logur you can easily wire the logging library of your choice into the MySQL driver:

	package main

	import (
		"github.com/goph/logur"
		"github.com/goph/logur/integrations/mysqllog"
		"github.com/go-sql-driver/mysql"
	)

	func main() {
		logger := logur.NewNoop() // choose an actual implementation
		mysql.SetLogger(mysqllog.New(logger))
	}
*/
package mysqllog

import (
	"github.com/goph/logur"
)

// New returns a new MySQL database driver logger.
func New(logger logur.Logger) *logur.PrintLogger {
	return logur.NewPrintErrorLogger(logger)
}
