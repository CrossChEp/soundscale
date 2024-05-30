package service_address_config

import (
	"flag"
	"fmt"
)

var (
	Port                            = flag.Int("port", defaultPort, "port of collection service")
	defaultAddress                  = fmt.Sprintf("%s:%d", "localhost", *Port)
	Address                         = flag.String("addr", defaultAddress, "address of collection service")
	defaultUserServiceAddress       = fmt.Sprintf("%s:%d", "localhost", defaultUserServicePort)
	UserServiceAddress              = flag.String("usaddr", defaultUserServiceAddress, "address of user service")
	defaultMusicServiceAddress      = fmt.Sprintf("%s:%d", "localhost", defaultMusicServicePort)
	MusicServiceAddress             = flag.String("msaddr", defaultMusicServiceAddress, "address of music service")
	defaultPostsAddress             = fmt.Sprintf("%s:%d", "localhost", defaultPostServicePort)
	PostServiceAddress              = flag.String("postaddr", defaultPostsAddress, "post service address")
	defaultCollectionServiceAddress = fmt.Sprintf("%s:%d", "localhost", defaultCollectionServicePort)
	CollectionServiceAddress        = flag.String("csaddr", defaultCollectionServiceAddress, "address of collection service")
)

const (
	defaultPort                  = 4561
	defaultUserServicePort       = 8080
	defaultMusicServicePort      = 7070
	defaultPostServicePort       = 5555
	defaultCollectionServicePort = 9876
)
