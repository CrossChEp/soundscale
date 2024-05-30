package main

import (
	"flag"
	"user_service/pkg/service/db"
	"user_service/pkg/service/setup"
)

func main() {
	flag.Parse()
	setup.ServiceSetup()
	defer db.DisconnectDB()
}
