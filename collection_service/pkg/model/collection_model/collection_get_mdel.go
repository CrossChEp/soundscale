package collection_model

type CollectionGetModel struct {
	Id            string   `bson:"_id"`
	UserId        string   `bson:"user_id"`
	Songs         []string `bson:"songs"`
	Playlists     []string `bson:"playlists"`
	Albums        []string `bson:"albums"`
	Played        []string `bson:"played"`
	Genres        []string `bson:"genres"`
	CreatedGenres []string `bson:"created_genres"`
	Followed      []string `bson:"followed"`
	Subscribed    []string `bson:"subscribed"`
}
