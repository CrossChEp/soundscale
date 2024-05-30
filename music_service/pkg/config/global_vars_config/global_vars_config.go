package global_vars_config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"music_service/pkg/config/services_address_config"
)

var (
	DBCollection *mongo.Collection
	DBClient     *mongo.Client
	DBContext    context.Context
	Redis        = redis.NewClient(&redis.Options{
		Addr:     services_address_config.RedisAddress,
		Password: *services_address_config.RedisPassword,
		DB:       0,
	})
)
