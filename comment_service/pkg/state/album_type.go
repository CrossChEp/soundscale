package state

import (
	"comment_service/pkg/config/constants"
	"comment_service/pkg/service/grpc_service/grpc_playlist"
)

type AlbumState struct {
	State
}

func (s AlbumState) Get(entityId string) error {
	_, err := grpc_playlist.GetById(entityId, constants.TypeAlbum)
	return err
}
