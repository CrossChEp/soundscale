package logger

import (
	"fmt"
	"recommendation_service/pkg/config/logger_config"
	"time"
)

func InfoLog(message string) {
	prefix := getPrefix()
	logger_config.InfoConsoleLogger.Info(fmt.Sprintf("%s  %s", prefix, message))
	logger_config.InfoFileLogger.Info(fmt.Sprintf("%s  %s", prefix, message))
}

func ErrorLog(message string) {
	prefix := getPrefix()
	logger_config.ErrorConsoleLogger.Error(fmt.Sprintf("%s  %s", prefix, message))
	logger_config.ErrorFileLogger.Error(fmt.Sprintf("%s  %s", prefix, message))
}

func DebugLog(message string) {
	prefix := getPrefix()
	logger_config.DebugConsoleLogger.Debug(fmt.Sprintf("%s  %s", prefix, message))
	logger_config.DebugFileLogger.Debug(fmt.Sprintf("%s  %s", prefix, message))
}

func ErrorWithDebugLog(msg string, err error, directory string) {
	ErrorLog(msg)
	DebugLog(fmt.Sprintf("%s Details: %v", directory, err))
}

func getPrefix() string {
	curTime := time.Now().UTC()
	timeFormat := "15:04:05 01/02/2006"
	return curTime.Format(timeFormat)
}
