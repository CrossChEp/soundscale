package logger

import (
	"playlist_serivce/pkg/config/logger_config"
)

func InfoLog(message string) {
	logger_config.InfoFileLogger.Println(message)
	logger_config.InfoConsoleLogger.Println(message)
}

func ErrorLog(message string) {
	logger_config.ErrorFileLogger.Println(message)
	logger_config.ErrorConsoleLogger.Println(message)
}
