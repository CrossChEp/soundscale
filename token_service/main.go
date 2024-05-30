package main

import (
	"fmt"
	"log"
	"net"

	"token_service/config"
	"token_service/protos/token_service_proto"
	"token_service/service"

	"google.golang.org/grpc"
)

func serve() {
	address := fmt.Sprintf(":%d", config.CONFIG.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	token_service_proto.RegisterTokenServiceServer(server, &service.TokenService{})
	config.Logger.Infof("Server has started at %v", address)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
		panic("Failed to serve")
	}
}

func main() {
	config.Init()
	serve()
}
