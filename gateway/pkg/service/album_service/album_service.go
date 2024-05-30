package album_service

import (
	"gateway/pkg/config/constants"
	"gateway/pkg/model/playlist_model"
	grpc_service "gateway/pkg/service/grpc_service/playlist"
)

func GetAlbumsByAuthor(authorId string) (*playlist_model.PlaylistsGetModel, error) {
	playlistResp, err := grpc_service.GetPlaylistByAuthor(authorId, constants.AlbumType)
	if err != nil {
		return nil, err
	}
	playlists := playlist_model.PlaylistsGetModel{}
	playlists.ToModel(playlistResp)
	return &playlists, nil
}

func GetAlbum(albumId string) (*playlist_model.PlaylistGetModel, error) {
	albumResp, err := grpc_service.GetPlaylist(albumId, constants.AlbumType)
	if err != nil {
		return nil, err
	}
	resAlbum := &playlist_model.PlaylistGetModel{}
	resAlbum.ToModel(albumResp)
	return resAlbum, nil
}
