package collection_model

type CollectionAddModel struct {
	UserId        string   `bson:"user_id"`
	Songs         []string `bson:"songs"`
	Playlists     []string `bson:"playlists"`
	Albums        []string `bson:"albums"`
	Genres        []string `bson:"genres"`
	CreatedGenres []string `bson:"created_genres"`
	Followed      []string `bson:"followed"`
	Subscribed    []string `bson:"subscribed"`
}
