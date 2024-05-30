package global_vars_config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"playlist_serivce/pkg/config/services_address_config"
)

var (
	Redis = redis.NewClient(&redis.Options{
		Addr:     services_address_config.DefaultRedisAddress,
		Password: *services_address_config.RedisPassword,
		DB:       0,
	})
	DBClient     *mongo.Client
	DBContext    context.Context
	DBCollection *mongo.Collection
)
