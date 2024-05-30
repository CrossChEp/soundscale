package checkers

import (
	"collection_service/pkg/config/logger_config"
	"collection_service/pkg/repository/collection_repo"
	"collection_service/pkg/service/grpc_service/grpc_music"
	"collection_service/pkg/service/grpc_service/grpc_playlist"
	"collection_service/pkg/service/grpc_service/grpc_user"
	"fmt"
)

func IsUserExists(userId string) bool {
	if _, err := grpc_user.GetUserById(userId); err != nil {
		logger_config.ErrorConsoleLogger.Println(fmt.Sprintf("(IsUserExists)Error: %v", err))
		return false
	}
	return true
}

func IsUserCollectionExists(userId string) bool {
	if _, err := collection_repo.GetUserCollection(userId); err != nil {
		return false
	}
	return true
}

func IsPlaylistExists(playlistId string, playlistType string) bool {
	if _, err := grpc_playlist.GetPlaylist(playlistId, playlistType); err != nil {
		return false
	}
	return true
}

func IsPlaylistInCollection(userId string, playlistId string) (bool, error) {
	collection, err := collection_repo.GetUserCollection(userId)
	if err != nil {
		return false, nil
	}
	for _, playlist := range collection.Playlists {
		if playlist == playlistId {
			return true, nil
		}
	}
	return false, nil
}

func IsAlbumInCollection(userId string, albumId string) (bool, error) {
	collection, err := collection_repo.GetUserCollection(userId)
	if err != nil {
		return false, nil
	}
	for _, album := range collection.Albums {
		if album == albumId {
			return true, nil
		}
	}
	return false, nil
}

func IsSongExists(songId string) bool {
	if _, err := grpc_music.GetSongById(songId); err != nil {
		return false
	}
	return true
}

func IsElementInArr(element string, arr []string) bool {
	for _, el := range arr {
		if el == element {
			return true
		}
	}
	return false
}
