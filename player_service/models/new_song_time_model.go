package models

import (
	"player_service/models/redis_models"
)

type SongTimeModel struct {
	Song     redis_models.SongsModel
	Duration int
	UserId   string
}
