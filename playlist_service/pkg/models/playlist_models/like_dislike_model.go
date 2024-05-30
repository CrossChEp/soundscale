package playlist_models

type LikeDislikeModel struct {
	Playlist     *PlaylistGetModel
	UserId       string
	PlaylistType string
}
