package main

import (
	"collection_service/pkg/service/setup"
	"flag"
)

func main() {
	flag.Parse()
	setup.ServiceSetup()
}
