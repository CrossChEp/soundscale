package logger

import "collection_service/pkg/config/logger_config"

func InfoLog(msg string) {
	logger_config.InfoConsoleLogger.Println(msg)
	logger_config.InfoFileLogger.Println(msg)
}

func ErrorLog(msg string) {
	logger_config.ErrorConsoleLogger.Println(msg)
	logger_config.ErrorFileLogger.Println(msg)
}
