// Package db package for connecting to the database
package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"playlist_serivce/pkg/config/global_vars_config"
	"playlist_serivce/pkg/config/services_address_config"
	"playlist_serivce/pkg/service/logger"
	"time"
)

func ConnectDB() {
	setConfigVars()
	connect()
	ping()
}

func setConfigVars() {
	global_vars_config.DBClient, _ = mongo.NewClient(options.Client().ApplyURI(*services_address_config.DBAddress))
	global_vars_config.DBContext, _ = context.WithTimeout(context.Background(), 10*time.Hour)
	global_vars_config.DBCollection = global_vars_config.DBClient.Database(*services_address_config.DBName).Collection("playlists")
}

func connect() {
	err := global_vars_config.DBClient.Connect(global_vars_config.DBContext)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Panic error: couldn't disconnect from database. Details: %v", err))
		panic(err)
	}
}

func ping() {
	err := global_vars_config.DBClient.Ping(global_vars_config.DBContext, nil)
	if err != nil {
		panic("panic error: couldn't connect to database")
	}
}

func DisconnectDB() {
	err := global_vars_config.DBClient.Disconnect(global_vars_config.DBContext)
	if err != nil {
		panic("panic error: couldn't disconnect from database")
	}
}
