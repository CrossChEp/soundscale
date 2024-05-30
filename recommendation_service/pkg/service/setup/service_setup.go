package setup

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"recommendation_service/pkg/config/logger_config"
	"recommendation_service/pkg/config/path_config"
	"recommendation_service/pkg/config/service_address_config"
	"recommendation_service/pkg/proto/recommendation_service_proto"
	"recommendation_service/pkg/service/logger"
	transport "recommendation_service/pkg/transport/grpc"
)

func ServiceSetup() {
	setupLoggers()
	setupService()
}

func setupLoggers() {
	file, err := os.OpenFile(path_config.LoggersPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	logger_config.InfoFileLogger.SetOutput(file)
	logger_config.ErrorFileLogger.SetOutput(file)
	logger_config.DebugFileLogger.SetOutput(file)
	setLoggersStyle()
}

func setLoggersStyle() {
	log.ErrorLevelStyle = lipgloss.NewStyle().
		SetString("[ERROR] [RECOMMENDATION]").
		Foreground(lipgloss.Color("#ED1E26")).
		Bold(true)
	log.InfoLevelStyle = lipgloss.NewStyle().
		SetString("[INFO] [RECOMMENDATION]").
		Foreground(lipgloss.Color("#18B894")).
		Bold(true)
	log.DebugLevelStyle = lipgloss.NewStyle().
		SetString("[DEBUG] [RECOMMENDATION]").
		Foreground(lipgloss.Color("#FFC60A")).
		Bold(true)
}

func setupService() {
	curDir, _ := os.Getwd()
	logger.InfoLog(fmt.Sprintf("Trying to start recommendation service"))
	address := fmt.Sprintf(":%d", *service_address_config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Couldn't start recommendation service on address %s.", address))
		logger.DebugLog(fmt.Sprintf("%v Details: %v", curDir, err))
		panic(err)
	}
	logger.InfoLog(fmt.Sprintf("Recommendation service was started on %s", address))
	server := grpc.NewServer()
	recommendation_service_proto.RegisterRecommendationServiceServer(server, &transport.RecommendationService{})
	if err := server.Serve(listener); err != nil {
		logger.ErrorLog(fmt.Sprintf("Couldn't start recommendation service on %s.", address))
		logger.DebugLog(fmt.Sprintf("%v Details: %v", curDir, err))
		panic(err)
	}
}
