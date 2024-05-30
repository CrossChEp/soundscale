package redis_repo

import (
	"context"
	"encoding/json"
	"fmt"
	"music_service/pkg/config/global_vars_config"
	"music_service/pkg/model"
	"music_service/pkg/service/logger"
	"time"
)

func SetTrends(songs []model.Song) error {
	trendsStr, err := json.Marshal(songs)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert songs to string. Details: %v", err))
		return err
	}
	err = global_vars_config.Redis.Set(context.Background(), "trends", trendsStr, 10*time.Minute).Err()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't set trends. Details: %v", err))
		return err
	}
	return nil
}

func GetTrends() ([]model.Song, error) {
	trendsStr, err := global_vars_config.Redis.Get(context.Background(), "trends").Result()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get trends. Details: %v", err))
		return nil, err
	}
	var trends []model.Song
	if err := json.Unmarshal([]byte(trendsStr), &trends); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't unmarshall redis trends record. Details: %v", err))
		return nil, err
	}
	return trends, nil
}
