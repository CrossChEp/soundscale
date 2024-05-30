package setup

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net"
	"os"
	"player_service/config"
	"player_service/funcs/logger"
	"player_service/service"
)

func ServiceSetup() {
	setLoggers()
	setPublicKey()
	go setService()
}

func setLoggers() {
	file, err := os.OpenFile(config.LogsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	config.ErrorFileLogger = log.New(file, "ERROR\t", log.LstdFlags)
	config.InfoFileLogger = log.New(file, "INFO\t", log.LstdFlags)
}

func setPublicKey() {
	bytes, err := os.ReadFile(config.PublicKeyPath)
	if err != nil {
		panic("something wrong with private key")
	}
	key, err := jwt.ParseECPublicKeyFromPEM(bytes)
	config.PublicKey = key
}

func setService() {
	ln, err := net.Listen("tcp", *config.TCPAddress)
	if err != nil {
		panic("couldn't start tcp server")
	}
	logger.InfoLog(fmt.Sprintf("TCP server started on %s", *config.TCPAddress))
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.ErrorLog(fmt.Sprintf("Error: couldn't accept connection. Details: %v", err))
			return
		}
		service.DownloadSong(conn)
	}
}
