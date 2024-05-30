package setup

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"user_service/pkg/config/logger_config"
	"user_service/pkg/config/path_config"
	"user_service/pkg/config/service_addresses_config"
	"user_service/pkg/proto/user_service_proto"
	"user_service/pkg/service/db"
	"user_service/pkg/service/logger"
	transport "user_service/pkg/transport/grpc"
)

func ServiceSetup() {
	setLoggers()
	db.ConnectBD()
	logger.InfoLog("Starting user service...")
	address := fmt.Sprintf(":%d", *service_addresses_config.PORT)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("panic error: couldn't start user service on port %d Details: %v\n", *service_addresses_config.PORT, err))
		panic(err)
	}
	logger.InfoLog(fmt.Sprintf("User service was started at %d\n", *service_addresses_config.PORT))
	server := grpc.NewServer()
	user_service_proto.RegisterUserServiceServer(server, &transport.UserService{})
	if err := server.Serve(listener); err != nil {
		logger.ErrorLog(fmt.Sprintf("Painc error: %v", err))
		panic(err)
	}
}

func setLoggers() {
	file, err := os.OpenFile(path_config.LogsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	logger_config.ErrorFileLogger = log.New(file, "ERROR\t", log.LstdFlags)
	logger_config.InfoFileLogger = log.New(file, "INFO\t", log.LstdFlags)
}
