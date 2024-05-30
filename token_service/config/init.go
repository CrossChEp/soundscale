package config

import (
	"crypto/ecdsa"
	"flag"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func ReadPrivateKey() (*ecdsa.PrivateKey, error) {
	pem_file, err := os.ReadFile(PRIVATE_KEY_PATH)
	if err != nil {
		Logger.Errorf("ReadPrivateKey: %v", err)
		return nil, err
	}
	key, err := jwt.ParseECPrivateKeyFromPEM(pem_file)
	return key, err
}

func ReadPublicKey() (*ecdsa.PublicKey, error) {
	pem_file, err := os.ReadFile(PUBLIC_KEY_PATH)
	if err != nil {
		Logger.Errorf("ReadPublicKey: %v", err)
		return nil, err
	}
	key, err := jwt.ParseECPublicKeyFromPEM(pem_file)
	return key, err
}

func initLoggers() {
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05.00 MST",
		PadLevelText:    true,
	})
}

func initKeys() {
	PUBLIC_KEY, _ = ReadPublicKey()
	PRIVATE_KEY, _ = ReadPrivateKey()
}

func initConfig() {
	Logger.Info("Reading configuration file...")
	config_path := flag.String("cfg", "./config/config_develop.yml", "path to config file")
	flag.Parse()
	conf := &Config{}
	file, err := os.ReadFile(*config_path)
	if err != nil {
		panic(err)
	}
	stat, _ := os.Stat(*config_path)
	if stat.IsDir() {
		panic("Provided path is a directory, not a file!")
	}

	yaml.Unmarshal(file, conf)
	CONFIG = *conf
	Logger.Info("Reading configurations is finished!")
}

func Init() {
	initLoggers()
	initConfig()
	initKeys()
}
