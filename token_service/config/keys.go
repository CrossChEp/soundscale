package config

import "crypto/ecdsa"

var (
	PRIVATE_KEY_PATH = "store/ecdsa_jwt_private_key.pem"
	PUBLIC_KEY_PATH  = "store/ecdsa_jwt_public_key.pem"
	PRIVATE_KEY      *ecdsa.PrivateKey
	PUBLIC_KEY       *ecdsa.PublicKey
)
