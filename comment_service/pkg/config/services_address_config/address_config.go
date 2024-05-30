package services_address_config

import (
	"flag"
	"fmt"
)

var (
	Port                          = flag.Int("port", defaultPort, "port of collection service")
	defaultAddress                = fmt.Sprintf("%s:%d", "localhost", *Port)
	Address                       = flag.String("addr", defaultAddress, "address of collection service")
	UserServicePort               = flag.Int("usport", defaultUserServicePort, "port of user service")
	defaultUserServiceAddress     = fmt.Sprintf("%s:%d", "localhost", *UserServicePort)
	UserServiceAddress            = flag.String("usaddr", defaultUserServiceAddress, "address of user service")
	PlaylistServicePort           = flag.Int("plport", defaultPlaylistServicePort, "port of playlist service")
	defaultPlaylistServiceAddress = fmt.Sprintf("%s:%d", "localhost", *PlaylistServicePort)
	PlaylistServiceAddress        = flag.String("pladdr", defaultPlaylistServiceAddress, "address of playlist service")
	MusicServicePort              = flag.Int("msport", defaultMusicServicePort, "port of music service")
	defaultMusicServiceAddress    = fmt.Sprintf("%s:%d", "localhost", *MusicServicePort)
	MusicServiceAddress           = flag.String("msaddr", defaultMusicServiceAddress, "address of music service")
	defaultPostsAddress           = fmt.Sprintf("%s:%d", "localhost", defaultPostServicePort)
	PostServiceAddress            = flag.String("postaddr", defaultPostsAddress, "post service address")
	DbName                        = flag.String("dbname", defaultDbName, "name of collections database")
	DbAddress                     = flag.String("dbaddrr", DefaultDbAddress, "address of collections database")
)

const (
	defaultPort                = 1000
	defaultUserServicePort     = 8080
	defaultPlaylistServicePort = 6061
	defaultMusicServicePort    = 7070
	defaultPostServicePort     = 5555
	defaultDbName              = "postogram_comments"
	DefaultDbAddress           = "mongodb://127.0.0.1:27017"
)
