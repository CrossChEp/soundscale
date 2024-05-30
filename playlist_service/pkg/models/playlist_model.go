package models

type PlaylistModel struct {
	Songs  []SongModel `json:"songs"`
	UserId string      `json:"user_id"`
}
