package main

import (
	"flag"
	"playlist_serivce/pkg/service/setup"
)

func main() {
	flag.Parse()
	setup.ServiceSetup()
}
