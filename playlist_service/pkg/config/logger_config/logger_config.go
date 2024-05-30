package logger_config

import (
	"log"
	"os"
)

var (
	ErrorConsoleLogger = log.New(os.Stdout, "ERROR\t", log.LstdFlags)
	ErrorFileLogger    *log.Logger
	InfoConsoleLogger  = log.New(os.Stdout, "INFO\t", log.LstdFlags)
	InfoFileLogger     *log.Logger
)
