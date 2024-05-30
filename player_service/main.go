package main

import (
	"flag"
	"player_service/config"
	"player_service/funcs/setup"
	"player_service/funcs/uploader"

	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()
	config.SongsPath = *config.StaticDir + "songs/"
	r := gin.Default()
	setup.ServiceSetup()
	r.Static("/static", *config.StaticDir + "songs/")
	r.POST("/upload", uploader.Upload)
	if err := r.Run(":8081"); err != nil {
		panic("couldn't start service")
	}
}
