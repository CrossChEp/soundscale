package config

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	// TODO: Rewrite logging!!!!
	ErrorConsoleLogger              = log.New(os.Stdout, "ERROR\t", log.LstdFlags)
	ErrorFileLogger                 *log.Logger
	InfoConsoleLogger               = log.New(os.Stdout, "INFO\t", log.LstdFlags)
	InfoFileLogger                  *log.Logger
	MusicServicePort                = flag.Int("msport", defaultMusicServicePort, "port for music service")
	defaultMusicServiceAddress      = fmt.Sprintf("%s:%d", "localhost", *MusicServicePort)
	MusicServiceAddress             = flag.String("msaddress", defaultMusicServiceAddress, "address of music service")
	PublicKey                       *ecdsa.PublicKey
	UserServicePort                 = flag.Int("usport", defaultUserServicePort, "port of user service")
	defaultUserServiceAddress       = fmt.Sprintf("%s:%d", "localhost", *UserServicePort)
	UserServiceAddress              = flag.String("usaddr", defaultUserServiceAddress, "address of user service")
	defaultCollectionServiceAddress = fmt.Sprintf("%s:%d", "localhost", defaultCollectionServicePort)
	CollectionServiceAddress        = flag.String("colladdr", defaultCollectionServiceAddress, "address of collection service")
	PORT                            = flag.Int("port", defaultPort, "port of player service")
	tcpDefaultAddress               = fmt.Sprintf("%s:%d", "127.0.0.1", *PORT)
	TCPAddress                      = flag.String("tcpaddr", tcpDefaultAddress, "address for tcp service")
	RedisPort                       = flag.Int("rport", defaultRedisPort, "port of redis")
	redisPassword                   = flag.String("rpassword", "", "password for redis")
	redisHost                       = flag.String("rhost", "localhost", "host of redis")
	RedisAddress                    = fmt.Sprintf("%s:%d", *redisHost, *RedisPort)
	Redis                           = redis.NewClient(&redis.Options{
		Addr:     RedisAddress,
		Password: *redisPassword,
		DB:       0,
	})
	StaticDir = flag.String("sdir", "./", "path to static directory")
	SongsPath = *StaticDir + "songs/"
)

const (
	defaultMusicServicePort      = 7070
	defaultRedisPort             = 6379
	defaultPort                  = 6060
	defaultUserServicePort       = 8080
	defaultCollectionServicePort = 9876
	PublicKeyPath                = "keys/ecdsa_jwt_public_key.pem"
	LogsPath                     = "logs/logs.log"
)
