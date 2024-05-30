package playlist_models

type PlaylistGetModel struct {
	Id          string   `bson:"_id"`
	Name        string   `bson:"name"`
	Author      string   `bson:"author"`
	Songs       []string `bson:"songs"`
	Listened    int64    `bson:"listened"`
	ReleaseDate string   `bson:"release_date"`
	LastUpdate  string   `bson:"last_update"`
	Type        string   `bson:"type"`
	Liked       []string `bson:"liked"`
	Disliked    []string `json:"disliked"`
}
