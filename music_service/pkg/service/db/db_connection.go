// Package db package for connection to database
package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"music_service/pkg/config/global_vars_config"
	"music_service/pkg/config/services_address_config"
	"music_service/pkg/service/logger"
	"time"
)

func ConnectBD() {
	logger.InfoLog("Trying to connect to database")
	connectToDb()
	logger.InfoLog("Connected to database")
}

func connectToDb() {
	setupConfigVariables()
	connect()
	ping()
}

func setupConfigVariables() {
	global_vars_config.DBClient, _ = mongo.NewClient(options.Client().ApplyURI(*services_address_config.DBAddress))
	global_vars_config.DBContext, _ = context.WithTimeout(context.Background(), 10*time.Hour)
	global_vars_config.DBCollection = global_vars_config.DBClient.Database(*services_address_config.DBName).Collection("songs")
}

func connect() {
	err := global_vars_config.DBClient.Connect(global_vars_config.DBContext)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("panic error: couldn't disconnect to database. Details: %v", err))
		panic(err)
	}
}

func ping() {
	err := global_vars_config.DBClient.Ping(global_vars_config.DBContext, nil)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Panic error: couldn't ping database. Details: %v", err))
		panic("panic error: couldn't connect to database")
	}
}

func DisconnectDB() {
	err := global_vars_config.DBClient.Disconnect(global_vars_config.DBContext)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Panic error: couldn't disconnect from db. Details: %v", err))
		panic("panic error: couldn't disconnect from database")
	}
}
