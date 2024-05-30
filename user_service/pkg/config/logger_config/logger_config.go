package logger_config

import (
	"log"
	"os"
)

var (
	InfoConsoleLogger  = log.New(os.Stdout, "INFO\t", log.LstdFlags)
	InfoFileLogger     *log.Logger
	ErrorConsoleLogger = log.New(os.Stdout, "ERROR\t", log.LstdFlags)
	ErrorFileLogger    *log.Logger
)
