package setup

import (
	"comment_service/pkg/config/logger_config"
	"comment_service/pkg/config/path_config"
	"comment_service/pkg/config/services_address_config"
	"comment_service/pkg/proto/comment_service_proto"
	"comment_service/pkg/service/db"
	"comment_service/pkg/service/logger"
	transport "comment_service/pkg/transport/grpc"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"google.golang.org/grpc"
	"net"
	"os"
)

func ServiceSetup() {
	setupLoggers()
	db.ConnectToDb()
	setupService()
	defer db.DisconnectDB()
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
		SetString("[ERROR] [COMMENT]").
		Foreground(lipgloss.Color("#ED1E26")).
		Bold(true)
	log.InfoLevelStyle = lipgloss.NewStyle().
		SetString("[INFO] [COMMENT]").
		Foreground(lipgloss.Color("#18B894")).
		Bold(true)
	log.DebugLevelStyle = lipgloss.NewStyle().
		SetString("[DEBUG] [COMMENT]").
		Foreground(lipgloss.Color("#FFC60A")).
		Bold(true)
}

func setupService() {
	curDir, _ := os.Getwd()
	logger.InfoLog(fmt.Sprintf("Trying to start comment service"))
	address := fmt.Sprintf(":%d", *services_address_config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Couldn't start comment service on address %s.", address))
		logger.DebugLog(fmt.Sprintf("%v Details: %v", curDir, err))
		panic(err)
	}
	logger.InfoLog(fmt.Sprintf("Comment service was started on %s", address))
	server := grpc.NewServer()
	comment_service_proto.RegisterCommentServiceServer(server, &transport.CommentService{})
	if err := server.Serve(listener); err != nil {
		logger.ErrorLog(fmt.Sprintf("Couldn't start comment service on %s.", address))
		logger.DebugLog(fmt.Sprintf("%v Details: %v", curDir, err))
		panic(err)
	}
}
