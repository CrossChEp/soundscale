package logger_config

import (
	"github.com/charmbracelet/log"
)

var (
	InfoConsoleLogger  = log.New(log.WithLevel(log.InfoLevel))
	InfoFileLogger     = log.New(log.WithLevel(log.InfoLevel))
	ErrorConsoleLogger = log.New(log.WithLevel(log.ErrorLevel))
	ErrorFileLogger    = log.New(log.WithLevel(log.ErrorLevel))
	DebugConsoleLogger = log.New(log.WithLevel(log.DebugLevel))
	DebugFileLogger    = log.New(log.WithLevel(log.DebugLevel))
)
