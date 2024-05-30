package service_addresses_config

import (
	"flag"
)

var (
	PORT          = flag.Int("port", 8080, "port for user transport")
	PublicKeyPath = flag.String("skp", "pkg/keys/ecdsa_jwt_public_key.pem", "path of secret keys for jwt token")
	DBAddress     = flag.String("dbaddr", defaultDBAddress, "address for database")
	DBName        = flag.String("dbname", defaultDBName, "name of user database")
)

const (
	defaultDBAddress = "mongodb://127.0.0.1:27017"
	defaultDBName    = "postogram_users"
)
