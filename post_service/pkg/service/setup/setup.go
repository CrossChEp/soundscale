package setup

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"post_service/pkg/config/logger_config"
	"post_service/pkg/config/path_config"
	"post_service/pkg/config/service_address_config"
	"post_service/pkg/proto/post_service_proto"
	"post_service/pkg/service/db"
	"post_service/pkg/service/logger"
	transport "post_service/pkg/transport/grpc"
)

func ServiceSetup() {
	setLoggers()
	db.Connect()
	setService()
}

func setLoggers() {
	file, err := os.OpenFile(path_config.LogsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(fmt.Sprintf("couldn't open file for logs. Details: %v", err))
	}
	logger_config.InfoFileLogger.SetOutput(file)
	logger_config.ErrorFileLogger.SetOutput(file)
	logger_config.DebugFileLogger.SetOutput(file)
	setLoggersStyle()
}

func setLoggersStyle() {
	log.ErrorLevelStyle = lipgloss.NewStyle().
		SetString("[ERROR] [POST SERVICE]").
		Foreground(lipgloss.Color("#ED1E26")).
		Bold(true)
	log.InfoLevelStyle = lipgloss.NewStyle().
		SetString("[INFO] [POST SERVICE]").
		Foreground(lipgloss.Color("#18B894")).
		Bold(true)
	log.DebugLevelStyle = lipgloss.NewStyle().
		SetString("[DEBUG] [POST SERVICE]").
		Foreground(lipgloss.Color("#FFC60A")).
		Bold(true)
}

func setService() {
	dir, _ := os.Getwd()
	address := fmt.Sprintf(":%d", *service_address_config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("panic error: couldn't start post service on port %d Details: %v\n", *service_address_config.Port, err))
		logger.DebugLog(fmt.Sprintf("%v(setService): %v", dir, err))
		panic(err)
	}
	logger.InfoLog(fmt.Sprintf("Post service was started at %s\n", *service_address_config.ServiceAddress))
	server := grpc.NewServer()
	post_service_proto.RegisterPostServiceServer(server, &transport.PostService{})
	if err := server.Serve(listener); err != nil {
		logger.ErrorLog(fmt.Sprintf("Panic error: couldn't start service on: %s", *service_address_config.ServiceAddress))
		logger.DebugLog(fmt.Sprintf("%v(setService): %v", dir, err))
		panic(err)
	}
}
