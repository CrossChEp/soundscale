package music_model

type MusicDeleteModel struct {
	UserId string
	SongId string `json:"song_id"`
}
