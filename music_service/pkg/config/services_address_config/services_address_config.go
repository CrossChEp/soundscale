package services_address_config

import (
	"flag"
	"fmt"
)

var (
	PORT                            = flag.Int("port", defaultPort, "server port of music transport")
	defaultUserServiceAddress       = fmt.Sprintf("%s:%d", "localhost", defaultUserServicePort)
	UserServiceAddress              = flag.String("usaddr", defaultUserServiceAddress, "address of user transport")
	DBName                          = flag.String("dbn", defaultDbName, "name of music database")
	DBAddress                       = flag.String("dbaddr", defaultDbAddress, "address of music database")
	defaultCollectionServiceAddress = fmt.Sprintf("%s:%d", "localhost", defaultCollectionServicePort)
	CollectionServiceAddress        = flag.String("colladdr", defaultCollectionServiceAddress, "address of collection service")
	RedisPort                       = flag.Int("rport", defaultRedisPort, "port of redis")
	RedisPassword                   = flag.String("rpassword", "", "password for redis")
	redisHost                       = flag.String("rhost", "localhost", "host of redis")
	RedisAddress                    = fmt.Sprintf("%s:%d", *redisHost, *RedisPort)
)

const (
	defaultPort                  = 7070
	defaultUserServicePort       = 8080
	defaultCollectionServicePort = 9876
	defaultDbAddress             = "mongodb://localhost:27017"
	defaultDbName                = "postogram_songs"
	defaultRedisPort             = 6379
)
