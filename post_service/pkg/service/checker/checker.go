package checker

import (
	"post_service/pkg/constants"
	"post_service/pkg/proto/post_service_proto"
	"post_service/pkg/service/grpc_service/grpc_music"
	"post_service/pkg/service/grpc_service/grpc_playlist"
)

func IsSongExists(songId string) bool {
	if _, err := grpc_music.GetSongById(songId); err != nil {
		return false
	}
	return true
}

func IsPlaylistExists(playlistId string) bool {
	if _, err := grpc_playlist.GetPlaylist(playlistId, constants.PlaylistType); err != nil {
		return false
	}
	return true
}

func IsIdInArr(id string, ids []string) bool {
	for _, arrId := range ids {
		if arrId == id {
			return true
		}
	}
	return false
}

func IsAlbumExists(albumId string) bool {
	if _, err := grpc_playlist.GetPlaylist(albumId, constants.AlbumType); err != nil {
		return false
	}
	return true
}

func IsPostValid(postReq *post_service_proto.AddPostReq) bool {
	if len(postReq.Songs) > 10 || len(postReq.Albums) > 3 || len(postReq.Playlists) > 3 {
		return false
	}
	return true
}

func IsUpdatePostValid(postReq *post_service_proto.UpdatePostReq) bool {
	if len(postReq.Songs) > 10 || len(postReq.Albums) > 3 || len(postReq.Playlists) > 3 {
		return false
	}
	return true
}

func IsUserInArr(userId string, arr []string) bool {
	for _, user := range arr {
		if user == userId {
			return true
		}
	}
	return false
}
