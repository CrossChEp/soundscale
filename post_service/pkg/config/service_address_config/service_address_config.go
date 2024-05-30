package service_address_config

import (
	"flag"
	"fmt"
)

var (
	Port                          = flag.Int("port", defaultPort, "port of post service")
	defaultAddress                = fmt.Sprintf("%s:%d", "localhost", *Port)
	ServiceAddress                = flag.String("addr", defaultAddress, "address of post service")
	DBAddress                     = flag.String("dbaddr", defaultDBAddress, "address of posts database")
	DBName                        = flag.String("dbname", defaultDBName, "name of posts database")
	MusicServicePort              = flag.Int("msport", defaultMusicServicePort, "port of music service")
	defaultMusicServiceAddress    = fmt.Sprintf("%s:%d", "localhost", *MusicServicePort)
	MusicServiceAddress           = flag.String("msaddr", defaultMusicServiceAddress, "address of music service")
	PlaylistServicePort           = flag.Int("plsport", defaultPlaylistServicePort, "port of playlist service")
	defaultPlaylistServiceAddress = fmt.Sprintf("%s:%d", "localhost", *PlaylistServicePort)
	PlaylistServiceAddress        = flag.String("pladdr", defaultPlaylistServiceAddress, "address of playlist service")
)

const (
	defaultPort                = 5555
	defaultDBAddress           = "mongodb://127.0.0.1:27017"
	defaultDBName              = "postogram_posts"
	defaultMusicServicePort    = 7070
	defaultPlaylistServicePort = 6061
)
