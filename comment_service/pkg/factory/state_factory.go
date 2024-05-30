package factory

import (
	"comment_service/pkg/config/constants"
	"comment_service/pkg/state"
)

func GetState(entityType string) state.State {
	switch entityType {
	case constants.TypePost:
		return state.PostState{}
	case constants.TypeSong:
		return state.SongState{}
	case constants.TypeAlbum:
		return state.AlbumState{}
	case constants.TypePlaylist:
		return state.PlaylistType{}
	}
	return nil
}
