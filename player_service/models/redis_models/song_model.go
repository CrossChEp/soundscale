package redis_models

type SongsModel struct {
	SongId    string `json:"song_id"`
	ReadedAt  int    `json:"readed_at"`
	ReadedSum int    `json:"readed_sum"`
}
