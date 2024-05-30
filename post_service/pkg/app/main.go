package main

import (
	"flag"
	"post_service/pkg/service/db"
	"post_service/pkg/service/setup"
)

func main() {
	flag.Parse()
	setup.ServiceSetup()
	defer db.Disconnect()
}
