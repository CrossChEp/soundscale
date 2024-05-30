package playlist_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddRemoveSongsModel struct {
	PlaylistId   primitive.ObjectID
	SongsIds     []string
	PlaylistType string
}
