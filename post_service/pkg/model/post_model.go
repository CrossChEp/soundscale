package model

type PostModel struct {
	Id           string   `bson:"_id"`
	AuthorId     string   `bson:"author_id"`
	Content      string   `bson:"content"`
	Songs        []string `bson:"songs"`
	Playlists    []string `bson:"playlists"`
	Albums       []string `bson:"albums"`
	CreationDate string   `bson:"creation_date"`
	Liked        []string `bson:"liked"`
	Disliked     []string `bson:"disliked"`
}
