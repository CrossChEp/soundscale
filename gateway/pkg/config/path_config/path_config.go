package path_config

import (
	"flag"
)

var (
	KeyPath = flag.String("kp", defaultKeyPath, "path of public keys")
)

const (
	LogsPath       = "./logs/logs.log"
	defaultKeyPath = "./pkg/keys/ecdsa_jwt_public_key.pem"
)
