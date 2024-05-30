package playlist_service

import (
	"gateway/pkg/config/constants"
	"gateway/pkg/model/playlist_model"
	grpc_service "gateway/pkg/service/grpc_service/playlist"
)

func AddPlaylist(playlistData *playlist_model.PlaylistAddModel) (*playlist_model.PlaylistGetModel, error) {
	playlistResp, err := grpc_service.AddPlaylist(playlistData)
	if err != nil {
		return nil, err
	}
	resPlaylist := &playlist_model.PlaylistGetModel{}
	resPlaylist.ToModel(playlistResp)
	return resPlaylist, nil
}

func GetPlaylist(playlistId string) (*playlist_model.PlaylistGetModel, error) {
	playlistResp, err := grpc_service.GetPlaylist(playlistId, constants.PlaylistType)
	if err != nil {
		return nil, err
	}
	resPlaylist := &playlist_model.PlaylistGetModel{}
	resPlaylist.ToModel(playlistResp)
	return resPlaylist, nil
}

func GetPlaylistsByAuthor(authorId string) (*playlist_model.PlaylistsGetModel, error) {
	playlistResp, err := grpc_service.GetPlaylistByAuthor(authorId, constants.PlaylistType)
	if err != nil {
		return nil, err
	}
	playlists := playlist_model.PlaylistsGetModel{}
	playlists.ToModel(playlistResp)
	return &playlists, nil
}

func AddSongToPlaylist(songAddModel *playlist_model.AddSongsToPlaylistModel) (*playlist_model.PlaylistGetModel, error) {
	songAddModel.PlaylistType = constants.PlaylistType
	playlistResp, err := grpc_service.AddSongToPlaylist(songAddModel)
	if err != nil {
		return nil, err
	}
	playlists := playlist_model.PlaylistGetModel{}
	playlists.ToModel(playlistResp)
	return &playlists, nil
}

func DeleteSongFromPlaylist(deleteSongModel *playlist_model.DeleteSongFromPlaylistModel) (*playlist_model.PlaylistGetModel, error) {
	deleteSongModel.PlaylistType = constants.PlaylistType
	playlistResp, err := grpc_service.DeleteSongFromPlaylist(deleteSongModel)
	if err != nil {
		return nil, err
	}
	playlists := playlist_model.PlaylistGetModel{}
	playlists.ToModel(playlistResp)
	return &playlists, nil
}

func Delete(deleteData *playlist_model.PlaylistDeleteModel) error {
	_, err := grpc_service.Delete(deleteData)
	if err != nil {
		return err
	}
	return nil
}

func AddCurrentPlaylist(playlistData *playlist_model.AddCurrentPlaylistModel) error {
	if _, err := grpc_service.AddCurrentPlaylist(playlistData); err != nil {
		return err
	}
	return nil
}

func LikePlaylist(likeDislikeModel playlist_model.LikeDislikePlaylist) (*playlist_model.PlaylistGetModel, error) {
	resp, err := grpc_service.LikePlaylist(likeDislikeModel)
	if err != nil {
		return nil, err
	}
	playlist := &playlist_model.PlaylistGetModel{}
	playlist.ToModel(resp)
	return playlist, nil
}

func DislikePlaylist(likeDislikeModel playlist_model.LikeDislikePlaylist) (*playlist_model.PlaylistGetModel, error) {
	resp, err := grpc_service.DislikePlaylist(likeDislikeModel)
	if err != nil {
		return nil, err
	}
	playlist := &playlist_model.PlaylistGetModel{}
	playlist.ToModel(resp)
	return playlist, nil
}

func GetUserLiked(userId string, playlistType string) (*playlist_model.PlaylistsGetModel, error) {
	resp, err := grpc_service.GetUserLikedPlaylists(userId, playlistType)
	if err != nil {
		return nil, err
	}
	playlist := &playlist_model.PlaylistsGetModel{}
	playlist.ToModel(resp)
	return playlist, nil
}

func GetUserDisliked(userId string, playlistType string) (*playlist_model.PlaylistsGetModel, error) {
	resp, err := grpc_service.GetUserDislikedPlaylists(userId, playlistType)
	if err != nil {
		return nil, err
	}
	playlist := &playlist_model.PlaylistsGetModel{}
	playlist.ToModel(resp)
	return playlist, nil
}
