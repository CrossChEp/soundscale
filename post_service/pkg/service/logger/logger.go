package logger

import (
	"fmt"
	"post_service/pkg/config/logger_config"
	"time"
)

func InfoLog(msg string) {
	prefix := getPrefix()
	logger_config.InfoConsoleLogger.Info(fmt.Sprintf("%s  %s", prefix, msg))
	logger_config.InfoFileLogger.Info(fmt.Sprintf("%s  %s", prefix, msg))
}

func ErrorLog(msg string) {
	prefix := getPrefix()
	logger_config.ErrorConsoleLogger.Error(fmt.Sprintf("%s  %s", prefix, msg))
	logger_config.ErrorFileLogger.Error(fmt.Sprintf("%s  %s", prefix, msg))
}

func DebugLog(msg string) {
	prefix := getPrefix()
	logger_config.DebugConsoleLogger.Debug(fmt.Sprintf("%s  %s", prefix, msg))
	logger_config.DebugFileLogger.Debug(fmt.Sprintf("%s  %s", prefix, msg))
}

func getPrefix() string {
	curTime := time.Now().UTC()
	timeFormat := "15:04:05 01/02/2006"
	return curTime.Format(timeFormat)
}
