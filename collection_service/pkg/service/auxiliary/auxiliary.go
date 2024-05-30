package auxiliary

import (
	"collection_service/pkg/config/constants"
	"collection_service/pkg/service/checkers"
)

func RemoveInvalidSongs(songs []string) []string {
	var newSongs []string
	for _, song := range songs {
		if checkers.IsSongExists(song) {
			newSongs = append(newSongs, song)
		}
	}
	return newSongs
}

func RemoveInvalidPlaylists(userId string, playlists []string) []string {
	var updatedPlaylists []string
	for _, playlist := range playlists {
		res, err := checkers.IsPlaylistInCollection(userId, playlist)
		if err != nil {
			continue
		}
		if checkers.IsPlaylistExists(playlist, constants.PlaylistType) && !res {
			updatedPlaylists = append(updatedPlaylists, playlist)
		}
	}
	return updatedPlaylists
}

func RemoveInvalidAlbums(userId string, albums []string) []string {
	var updatedAlbums []string
	for _, album := range albums {
		res, err := checkers.IsAlbumInCollection(userId, album)
		if err != nil {
			continue
		}
		if checkers.IsPlaylistExists(album, constants.AlbumType) && !res {
			updatedAlbums = append(updatedAlbums, album)
		}
	}
	return updatedAlbums
}

func RemoveInvalidAlbumsForeRemoving(userId string, albums []string) []string {
	var updatedAlbums []string
	for _, album := range albums {
		res, err := checkers.IsAlbumInCollection(userId, album)
		if err != nil {
			continue
		}
		if checkers.IsPlaylistExists(album, constants.AlbumType) && res {
			updatedAlbums = append(updatedAlbums, album)
		}
	}
	return updatedAlbums
}

func RemoveInvalidPlaylistsForRemoving(userId string, playlists []string) []string {
	var updatedPlaylists []string
	for _, playlist := range playlists {
		res, err := checkers.IsPlaylistInCollection(userId, playlist)
		if err != nil {
			continue
		}
		if checkers.IsPlaylistExists(playlist, constants.PlaylistType) && res {
			updatedPlaylists = append(updatedPlaylists, playlist)
		}
	}
	return updatedPlaylists
}

func RemoveFromArr(element string, arr []string) []string {
	var newArr []string
	for _, el := range arr {
		if el != element {
			newArr = append(newArr, el)
		}
	}
	return newArr
}
