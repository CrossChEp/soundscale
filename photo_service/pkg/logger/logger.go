package logger

import (
	"photo_service/pkg/config/logger_config"
)

func ErrorLog(message string) {
	logger_config.ErrorConsoleLogger.Println(message)
	logger_config.ErrorFileLogger.Println(message)
}

func InfoLog(message string) {
	logger_config.InfoConsoleLogger.Println(message)
	logger_config.InfoFileLogger.Println(message)
}
