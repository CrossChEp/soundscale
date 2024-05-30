package base_functions

import (
	"playlist_serivce/pkg/config/constants"
	"playlist_serivce/pkg/service/checkers"
)

func DeleteFromArr(arr []string, item string) []string {
	var newArr []string
	flag := false
	for _, element := range arr {
		if item == element && !flag {
			flag = true
			continue
		}
		newArr = append(newArr, item)
	}
	return newArr
}

func RemoveUnexistingSong(songs []string) []string {
	var newSongs []string
	for _, song := range songs {
		if checkers.IsSongExists(song) {
			newSongs = append(newSongs, song)
		}
	}
	return newSongs
}

func IsPlaylistTypeExists(checkingPlaylistType string) bool {
	for _, playlistType := range constants.Playlisttypes {
		if checkingPlaylistType == playlistType {
			return true
		}
	}
	return false
}
