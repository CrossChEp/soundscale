package response

import (
	"playlist_serivce/pkg/models/playlist_models"
	"playlist_serivce/pkg/proto/playlist_service_proto"
)

func CreateGetPlaylistsResponse(playlists []playlist_models.PlaylistGetModel) *playlist_service_proto.GetPlaylistsResp {
	var playlistsProtos []*playlist_service_proto.GetPlaylistResp
	for _, playlist := range playlists {
		playlistsProtos = append(playlistsProtos, CreateGetPlaylistResponse(&playlist))
	}
	return &playlist_service_proto.GetPlaylistsResp{
		Playlists: playlistsProtos,
	}
}

func CreateGetPlaylistResponse(playlist *playlist_models.PlaylistGetModel) *playlist_service_proto.GetPlaylistResp {
	return &playlist_service_proto.GetPlaylistResp{
		Id:          playlist.Id,
		Name:        playlist.Name,
		AuthorId:    playlist.Author,
		Songs:       playlist.Songs,
		CreatedAt:   playlist.ReleaseDate,
		LastUpdated: playlist.LastUpdate,
		Listened:    playlist.Listened,
		Type:        playlist.Type,
		Liked:       playlist.Liked,
		Disliked:    playlist.Disliked,
	}
}
