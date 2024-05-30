package setup

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"photo_service/pkg/config/logger_config"
	"photo_service/pkg/config/path_config"
	"photo_service/pkg/config/services_adress_config"
	"photo_service/pkg/logger"
	"photo_service/pkg/proto/photo_service_proto"
	transport "photo_service/pkg/transport/grpc"
)

func ServiceSetup() {
	setupLoggers()
	setupService()
}

func setupLoggers() {
	file, err := os.OpenFile(path_config.LogsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	logger_config.ErrorFileLogger = log.New(file, "ERROR\t", log.LstdFlags)
	logger_config.InfoFileLogger = log.New(file, "INFO\t", log.LstdFlags)
}

func setupService() {
	logger.InfoLog("Trying to start photo service")
	address := fmt.Sprintf(":%d", *services_adress_config.PORT)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("panic error: couldn't start service on port %d. Details: %v",
			err, *services_adress_config.PORT))
		panic(err)
	}
	logger.InfoLog("Photo service was started")
	server := grpc.NewServer()
	photo_service_proto.RegisterPhotoServiceServer(server, &transport.PhotoService{})
	if err := server.Serve(listener); err != nil {
		logger.ErrorLog(fmt.Sprintf("panic error: Details: %v", err))
		panic(err)
	}
}
