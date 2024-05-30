package playlist_model

import (
	"gateway/pkg/model/comment_model"
	"gateway/pkg/proto/playlist_service_proto"
	"gateway/pkg/service/comment_service"
)

type PlaylistGetModel struct {
	Id         string                          `json:"id"`
	Name       string                          `json:"name"`
	Author     string                          `json:"author"`
	Songs      []string                        `json:"songs"`
	CreatedAt  string                          `json:"created_at"`
	LastUpdate string                          `json:"updated_at"`
	Listened   int64                           `json:"listened"`
	Type       string                          `json:"type"`
	Liked      []string                        `json:"liked"`
	Disliked   []string                        `json:"disliked"`
	Likes      int                             `json:"likes"`
	Dislikes   int                             `json:"dislikes"`
	Comments   []comment_model.CommentGetModel `json:"comments"`
}

func (model *PlaylistGetModel) ToModel(resp *playlist_service_proto.GetPlaylistResp) {
	model.Id = resp.Id
	model.Name = resp.Name
	model.Author = resp.AuthorId
	model.Songs = resp.Songs
	model.CreatedAt = resp.CreatedAt
	model.LastUpdate = resp.LastUpdated
	model.Listened = resp.Listened
	model.Type = resp.Type
	model.Liked = resp.Liked
	model.Disliked = resp.Disliked
	model.Likes = len(resp.Liked)
	model.Dislikes = len(resp.Disliked)
	model.Comments = getComments(resp.Id, resp.Type)
}

func getComments(playlistId string, playlistType string) []comment_model.CommentGetModel {
	comments, err := comment_service.GetEntityComments(playlistId, playlistType)
	if err != nil {
		return []comment_model.CommentGetModel{}
	}
	return comments
}
