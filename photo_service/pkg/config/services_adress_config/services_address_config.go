package services_adress_config

import (
	"flag"
	"fmt"
)

var (
	PORT                          = flag.Int("port", defaultPort, "server port for photo transport")
	defaultUserServiceAddress     = fmt.Sprintf("%s:%d", "localhost", defaultUserServicePort)
	UserServiceAddress            = flag.String("usaddr", defaultUserServiceAddress, "address of user sercvice")
	MusicServicePort              = flag.Int("msport", defaultMusicServicePort, "port of music transport")
	defaultMusicServiceAddress    = fmt.Sprintf("%s:%d", "localhost", *MusicServicePort)
	MusicServiceAddress           = flag.String("msaddr", defaultMusicServiceAddress, "address of music transport")
	PlaylistServicePort           = flag.Int("psport", defaultPlaylistServicePort, "port of playlist service")
	defaultPlaylistServiceAddress = fmt.Sprintf("%s:%d", "localhost", *PlaylistServicePort)
	PlaylistServiceAddress        = flag.String("psaddr", defaultPlaylistServiceAddress, "address of playlist service")
	PostServicePort               = flag.Int("posport", defaultPostServicePort, "port of grpc_post service")
	defaultPostServiceAddress     = fmt.Sprintf("%s:%d", "localhost", *PostServicePort)
	PostServiceAddress            = flag.String("posaddr", defaultPostServiceAddress, "address of grpc_post service")
)

const (
	defaultPort                = 7354
	defaultMusicServicePort    = 7000
	defaultUserServicePort     = 8080
	defaultPlaylistServicePort = 6061
	defaultPostServicePort     = 5555
)
