package checkers

import "gateway/pkg/config/constants"

func IsPlaylistTypeExists(playlistType string) bool {
	for _, playlist := range constants.PlaylistTypeList {
		if playlist == playlistType {
			return true
		}
	}
	return false
}
