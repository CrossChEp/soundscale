package redis_models

type PlaylistModel struct {
	UserId string
	Songs  []SongsModel
}
