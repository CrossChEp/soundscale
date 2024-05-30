package models

import (
	"net"
	"player_service/models/redis_models"
)

type StartPlaylistModel struct {
	UserId string
	Songs  []redis_models.SongsModel
	Conn   net.Conn
}
