package global_vars_config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBClient     *mongo.Client
	DBContext    context.Context
	DBCollection *mongo.Collection
)
