package global_vars_config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DbClient     *mongo.Client
	DbContext    context.Context
	DbCollection *mongo.Collection
)
