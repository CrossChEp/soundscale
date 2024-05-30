package playlist_model

type PlaylistDeleteModel struct {
	AuthorId     string
	PlaylistId   string `json:"playlist_id"`
	PlaylistType string
}
