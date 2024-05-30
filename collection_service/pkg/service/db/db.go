package db

import (
	"collection_service/pkg/config/global_vars_config"
	"collection_service/pkg/config/services_address_config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectToDb() {
	setConfigVars()
	connect()
	ping()
}

func setConfigVars() {
	global_vars_config.DbClient, _ = mongo.NewClient(options.Client().ApplyURI(*services_address_config.DbAddress))
	global_vars_config.DbContext, _ = context.WithTimeout(context.Background(), 10*time.Hour)
	global_vars_config.DbCollection = global_vars_config.DbClient.Database(*services_address_config.DbName).Collection("collections")
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
