package setup

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"music_service/pkg/config/logger_config"
	"music_service/pkg/config/path_config"
	"music_service/pkg/config/services_address_config"
	"music_service/pkg/proto/music_service_proto"
	"music_service/pkg/service/db"
	"music_service/pkg/service/logger"
	"music_service/pkg/transport/grpc"
	"net"
	"os"
)

func ServiceSetup() {
	setLoggers()
	db.ConnectBD()
	setService()
}

func setLoggers() {
	file, err := os.OpenFile(path_config.LogsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	logger_config.ErrorFileLogger = log.New(file, "ERROR\t", log.LstdFlags)
	logger_config.InfoFileLogger = log.New(file, "INFO\t", log.LstdFlags)
}

func setService() {
	logger.InfoLog("Trying to start music transport")
	address := fmt.Sprintf(":%d", *services_address_config.PORT)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("panic error: couldn't start transport on port %d. Details: %v",
			*services_address_config.PORT, err))
		panic(err)
	}
	logger.InfoLog("Music transport was started")
	server := grpc.NewServer()
	music_service_proto.RegisterMusicServiceServer(server, &transport.MusicService{})
	logger.InfoLog(fmt.Sprintf("Started on %s", address))
	if err := server.Serve(listener); err != nil {
		logger.ErrorLog(fmt.Sprintf("panic error. Details: %v", err))
		panic(err)
	}
}
