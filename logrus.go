// Package logrus allows the creation of Loggers based on logrus.StandardLogger().
package logrus

import (
	"github.com/go-log/log"
	"github.com/go-log/log/print"
	"github.com/sirupsen/logrus"
)

// WithFields creates a new logger based on logrus.WithFields().
func WithFields(f logrus.Fields) log.Logger {
	return print.New(logrus.WithFields(f))
}

// WithFields creates a new logger based on logrus.StandardLogger().
func New() log.Logger {
	return print.New(logrus.StandardLogger())
}
