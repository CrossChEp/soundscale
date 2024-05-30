package logger

import "user_service/pkg/config/logger_config"

func ErrorLog(msg string) {
	logger_config.ErrorConsoleLogger.Println(msg)
	logger_config.ErrorFileLogger.Println(msg)
}

func InfoLog(msg string) {
	logger_config.InfoConsoleLogger.Println(msg)
	logger_config.InfoFileLogger.Println(msg)
}
