package logger

import "player_service/config"

func InfoLog(message string) {
	config.InfoFileLogger.Println(message)
	config.InfoConsoleLogger.Println(message)
}

func ErrorLog(message string) {
	config.ErrorFileLogger.Println(message)
	config.ErrorConsoleLogger.Println(message)
}
