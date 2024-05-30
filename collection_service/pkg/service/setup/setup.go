package setup

import (
	"collection_service/pkg/config/logger_config"
	"collection_service/pkg/config/path_config"
	"collection_service/pkg/config/services_address_config"
	"collection_service/pkg/proto/collection_service_proto"
	"collection_service/pkg/service/db"
	"collection_service/pkg/service/logger"
	transport "collection_service/pkg/transport/grpc"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func ServiceSetup() {
	setLoggers()
	db.ConnectToDb()
	setupService()
	defer db.DisconnectDB()
}

func setLoggers() {
	file, err := os.OpenFile(path_config.LoggersPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	logger_config.InfoFileLogger = log.New(file, "INFO\t", log.LstdFlags)
	logger_config.ErrorFileLogger = log.New(file, "ERROR\t", log.LstdFlags)
}

func setupService() {
	logger.InfoLog(fmt.Sprintf("Trying to start collection service"))
	address := fmt.Sprintf(":%d", *services_address_config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Couldn't start collection service on address %s. Details: %v", address, err))
		panic(err)
	}
	logger.InfoLog(fmt.Sprintf("Collection service was started on %s", address))
	server := grpc.NewServer()
	collection_service_proto.RegisterCollectionServiceServer(server, &transport.CollectionService{})
	if err := server.Serve(listener); err != nil {
		logger.ErrorLog(fmt.Sprintf("Couldn't start collection service on %s. Details: %v", address, err))
		panic(err)
	}
}
