package filter

import "post_service/pkg/service/checker"

func RemoveNonExistingSongs(songs []string) []string {
	var newSongs []string
	for _, song := range songs {
		if checker.IsSongExists(song) && !checker.IsIdInArr(song, newSongs) {
			newSongs = append(newSongs, song)
		}
	}
	return newSongs
}

func RemoveNonExistingPlaylists(playlists []string) []string {
	var newPlaylists []string
	for _, playlist := range playlists {
		if checker.IsPlaylistExists(playlist) && !checker.IsIdInArr(playlist, newPlaylists) {
			newPlaylists = append(newPlaylists, playlist)
		}
	}
	return newPlaylists
}

func RemoveNonExistingAlbums(albums []string) []string {
	var newAlbums []string
	for _, album := range albums {
		if checker.IsAlbumExists(album) && !checker.IsIdInArr(album, newAlbums) {
			newAlbums = append(newAlbums, album)
		}
	}
	return newAlbums
}
