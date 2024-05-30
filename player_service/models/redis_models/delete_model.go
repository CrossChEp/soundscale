package redis_models

type DeleteFromPlaylistModel struct {
	UserId   string
	Position int
	SongId   string
}
