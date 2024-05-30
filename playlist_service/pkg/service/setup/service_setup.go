package setup

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"playlist_serivce/pkg/config/logger_config"
	"playlist_serivce/pkg/config/path_config"
	"playlist_serivce/pkg/config/services_address_config"
	"playlist_serivce/pkg/proto/playlist_service_proto"
	"playlist_serivce/pkg/service/logger"
	"playlist_serivce/pkg/service/setup/db"
	transport "playlist_serivce/pkg/transport/grpc"
)

func ServiceSetup() {
	setLoggers()
	db.ConnectDB()
	setupService()
	defer db.DisconnectDB()
}

func setLoggers() {
	file, err := os.OpenFile(path_config.LogsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	logger_config.ErrorFileLogger = log.New(file, "ERROR\t", log.LstdFlags)
	logger_config.InfoFileLogger = log.New(file, "INFO\t", log.LstdFlags)
}

func setupService() {
	logger.InfoLog("Trying to start playlist_service service")
	address := fmt.Sprintf(":%d", *services_address_config.PORT)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("panic error: couldn't start service on port %d. Details: %v", *services_address_config.PORT, err))
		panic(err)
	}
	logger.InfoLog("Playlist service was started")
	server := grpc.NewServer()
	playlist_service_proto.RegisterPlaylistServiceServer(server, &transport.PlaylistService{})
	if err := server.Serve(listener); err != nil {
		logger.ErrorLog("panic error: ")
		panic(err)
	}
}
