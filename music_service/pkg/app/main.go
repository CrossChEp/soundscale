package main

import (
	"flag"
	"music_service/pkg/service/db"
	"music_service/pkg/service/setup"
)

func main() {
	flag.Parse()
	setup.ServiceSetup()
	defer db.DisconnectDB()
}
