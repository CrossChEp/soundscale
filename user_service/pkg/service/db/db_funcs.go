package db

import (
	"context"
	"fmt"
	"time"
	"user_service/pkg/config/global_vars_config"
	"user_service/pkg/config/service_addresses_config"
	"user_service/pkg/service/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectBD() {
	connectToDb()
}

func connectToDb() {
	logger.InfoLog(fmt.Sprintf("Connecting to database at %s", *service_addresses_config.DBAddress))
	setupConfigVariables()
	connect()
	ping()
	logger.InfoLog("Connected to database")
}

func setupConfigVariables() {
	global_vars_config.DBClient, _ = mongo.NewClient(options.Client().ApplyURI(*service_addresses_config.DBAddress))
	global_vars_config.DBContext, _ = context.WithTimeout(context.Background(), 10*time.Hour)
	global_vars_config.DBCollection = global_vars_config.DBClient.Database(*service_addresses_config.DBName).Collection("users")
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
