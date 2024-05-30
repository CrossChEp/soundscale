package models

type SongModel struct {
	SongId    string `json:"song_id"`
	ReadedAt  int64  `json:"readed_at"`
	ReadedSum int64  `json:"readed_sum"`
}
