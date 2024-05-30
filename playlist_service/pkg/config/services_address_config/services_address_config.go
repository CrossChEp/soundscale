package services_address_config

import (
	"flag"
	"fmt"
)

var (
	PORT                       = flag.Int("port", defaultPort, "port for playlist_service transport")
	MusicServicePort           = flag.Int("msport", defaultMusicServicePort, "port for music transport")
	defaultMusicServiceAddress = fmt.Sprintf("%s:%d", "localhost", *MusicServicePort)
	MusicServiceAddress        = flag.String("msaddress", defaultMusicServiceAddress, "address of music transport")
	UserServicePort            = flag.Int("usport", defaultUserServicePort, "port for user transport")
	defaultUserServiceAddress  = fmt.Sprintf("%s:%d", "localhost", *UserServicePort)
	UserServiceAddress         = flag.String("usaddress", defaultUserServiceAddress, "address of user transport")
	RedisPort                  = flag.Int("rport", defaultRedisPort, "port of redis")
	DefaultRedisAddress        = fmt.Sprintf("%s:%d", *RedisHost, *RedisPort)
	RedisPassword              = flag.String("rpassword", "", "password for redis")
	RedisHost                  = flag.String("rhost", "localhost", "host of redis")
	defaultDBAddress           = "mongodb://127.0.0.1:27017"
	defaultDBName              = "postogram_playlists"
	DBAddress                  = flag.String("dbaddr", defaultDBAddress, "address for database")
	DBName                     = flag.String("dbname", defaultDBName, "name of user database")
)

const (
	defaultMusicServicePort = 7070
	defaultPort             = 6061
	defaultUserServicePort  = 8080
	defaultRedisPort        = 6379
)
