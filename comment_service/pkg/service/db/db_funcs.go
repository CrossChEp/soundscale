package db

import (
	"comment_service/pkg/config/constants"
	"comment_service/pkg/config/global_vars_config"
	"comment_service/pkg/config/services_address_config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func ConnectToDb() {
	setDbAddress()
	setConfigVars()
	connect()
	ping()
}

func setDbAddress() {
	mongoAddr := os.Getenv(fmt.Sprintf("%s", constants.MongoAddrEnvName))
	user, password := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")
	if mongoAddr == "" {
		mongoAddr = services_address_config.DefaultDbAddress
	}
	if user != "" && password != "" {
		mongoAddr = fmt.Sprintf("%s:%s@%s", user, password, mongoAddr)
	}
	*services_address_config.DbAddress = mongoAddr
}

func setConfigVars() {
	global_vars_config.DbClient, _ = mongo.NewClient(options.Client().ApplyURI(*services_address_config.DbAddress))
	global_vars_config.DbContext, _ = context.WithTimeout(context.Background(), 10*time.Hour)
	global_vars_config.DbCollection = global_vars_config.DbClient.Database(*services_address_config.DbName).Collection("comments")
}

func ping() {
	if err := global_vars_config.DbClient.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
}

func connect() {
	if err := global_vars_config.DbClient.Connect(global_vars_config.DbContext); err != nil {
		panic(err)
	}
}

func DisconnectDB() {
	if err := global_vars_config.DbClient.Disconnect(global_vars_config.DbContext); err != nil {
		panic(err)
	}
}
