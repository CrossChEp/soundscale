package post_model

import (
	"gateway/pkg/config/constants"
	"gateway/pkg/model/comment_model"
	"gateway/pkg/model/music_model"
	"gateway/pkg/model/playlist_model"
	"gateway/pkg/proto/post_service_proto"
	"gateway/pkg/service/comment_service"
	grpc_service "gateway/pkg/service/grpc_service/music"
	playlist_grpc_service "gateway/pkg/service/grpc_service/playlist"
)

type PostGetModel struct {
	Id           string                            `json:"id"`
	AuthorId     string                            `json:"author_id"`
	Content      string                            `json:"content"`
	Songs        []music_model.SongGetModel        `json:"songs"`
	Playlists    []playlist_model.PlaylistGetModel `json:"playlists"`
	Albums       []playlist_model.PlaylistGetModel `json:"albums"`
	Liked        int                               `json:"liked"`
	Disliked     int                               `json:"disliked"`
	CreationDate string                            `json:"creation_date"`
	Comments     []comment_model.CommentGetModel   `json:"comments"`
}

func (model *PostGetModel) ToModel(resp *post_service_proto.PostResp) {
	model.Id = resp.Id
	model.AuthorId = resp.AuthorId
	model.Content = resp.Content
	model.Songs = getSongsModels(resp.Songs)
	model.Playlists = getPlaylistResps(resp.Playlists, constants.PlaylistType)
	model.Albums = getPlaylistResps(resp.Albums, constants.AlbumType)
	model.Liked = len(resp.Liked)
	model.Disliked = len(resp.Disliked)
	model.CreationDate = resp.CreationDate
	model.Comments = getComments(resp.Id)
}

func getSongsModels(songIds []string) []music_model.SongGetModel {
	var songs []music_model.SongGetModel
	for _, songId := range songIds {
		song, err := grpc_service.GetSongById(songId)
		if err != nil {
			continue
		}
		songModel := music_model.SongGetModel{}
		songModel.ToModel(song)
		songs = append(songs, songModel)
	}
	return songs
}

func getPlaylistResps(playlistIds []string, playlistType string) []playlist_model.PlaylistGetModel {
	var playlists []playlist_model.PlaylistGetModel
	for _, playlistId := range playlistIds {
		playlistResp, err := playlist_grpc_service.GetPlaylist(playlistId, playlistType)
		if err != nil {
			continue
		}
		playlist := playlist_model.PlaylistGetModel{}
		playlist.ToModel(playlistResp)
		playlists = append(playlists, playlist)
	}
	return playlists
}

func getComments(postId string) []comment_model.CommentGetModel {
	comments, err := comment_service.GetEntityComments(postId, constants.PostType)
	if err != nil {
		return []comment_model.CommentGetModel{}
	}
	return comments
}
