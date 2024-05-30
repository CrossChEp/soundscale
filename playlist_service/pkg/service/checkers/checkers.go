package checkers

import (
	"playlist_serivce/pkg/models"
	"playlist_serivce/pkg/service/grpc_funcs"
)

func AreSongsExist(songs []string) bool {
	for _, song := range songs {
		if !IsSongExists(song) {
			return false
		}
	}
	return true
}

func IsUserExists(userId string) bool {
	_, err := grpc_funcs.GetUser(userId)
	if err != nil {
		return false
	}
	return true
}

func IsSongExists(songId string) bool {
	_, err := grpc_funcs.GetSong(songId)
	if err != nil {
		return false
	}
	return true
}

func IsElementInSlice(element string, arr []models.SongModel) bool {
	for _, el := range arr {
		if el.SongId == element {
			return true
		}
	}
	return false
}

func IsUserInArr(userId string, arr []string) bool {
	for _, user := range arr {
		if user == user {
			return true
		}
	}
	return false
}
