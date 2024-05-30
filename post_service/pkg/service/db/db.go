package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"post_service/pkg/config/global_vars_config"
	"post_service/pkg/config/service_address_config"
	"post_service/pkg/service/logger"
	"time"
)

func Connect() {
	logger.InfoLog(fmt.Sprintf("Trying to connect to database"))
	connectToDB()
	logger.InfoLog(fmt.Sprintf("Info: connected to database"))
}

func connectToDB() {
	setupVars()
	connect()
	ping()
}

func setupVars() {
	global_vars_config.DBClient, _ =
		mongo.NewClient(options.Client().ApplyURI(*service_address_config.DBAddress))
	global_vars_config.DBContext, _ =
		context.WithTimeout(context.Background(), 10*time.Hour)
	global_vars_config.DBCollection =
		global_vars_config.DBClient.Database(*service_address_config.DBName).Collection("posts")
}

func connect() {
	if err := global_vars_config.DBClient.Connect(global_vars_config.DBContext); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to database. Details: %v", err))
		panic(err)
	}
}

func ping() {
	if err := global_vars_config.DBClient.Ping(global_vars_config.DBContext, nil); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't ping datbase. Details: %v", err))
		panic(err)
	}
}

func Disconnect() {
	if err := global_vars_config.DBClient.Disconnect(global_vars_config.DBContext); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't disconnect from database. Details: %v", err))
		panic(err)
	}
}
