package music_model

import (
	"gateway/pkg/config/constants"
	"gateway/pkg/model/comment_model"
	"gateway/pkg/proto/music_service_proto"
	"gateway/pkg/service/comment_service"
	"strconv"
)

type SongGetModel struct {
	SongId        string                          `json:"song_id"`
	SongName      string                          `json:"song_name"`
	AuthorId      string                          `json:"author_id"`
	Collaborators []string                        `json:"collaborators"`
	Genre         string                          `json:"genre"`
	ReleaseData   string                          `json:"release_data"`
	Played        int                             `json:"played"`
	Listened      int                             `json:"listened"`
	Liked         []string                        `json:"liked"`
	Disliked      []string                        `json:"disliked"`
	Likes         int                             `json:"likes"`
	Dislikes      int                             `json:"dislikes"`
	Comments      []comment_model.CommentGetModel `json:"comments"`
	Exclusive     bool                            `json:"exclusive"`
}

func (model *SongGetModel) ToModel(resp *music_service_proto.GetResp) {
	ex, _ := strconv.ParseBool(resp.Exclusive)
	model.SongId = resp.Id
	model.SongName = resp.SongName
	model.AuthorId = resp.AuthorId
	model.Collaborators = resp.Collaborators
	model.Genre = resp.Genre
	model.ReleaseData = resp.ReleaseDate.AsTime().String()
	model.Played = len(resp.Played)
	model.Listened = len(resp.Listened)
	model.Liked = resp.Liked
	model.Disliked = resp.Disliked
	model.Likes = len(resp.Liked)
	model.Dislikes = len(resp.Disliked)
	model.Comments = getComments(model.SongId)
	model.Exclusive = ex
}

func getComments(songId string) []comment_model.CommentGetModel {
	comments, err := comment_service.GetEntityComments(songId, constants.SongType)
	if err != nil {
		return []comment_model.CommentGetModel{}
	}
	return comments
}
