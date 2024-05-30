package state

import (
	"comment_service/pkg/config/constants"
	"comment_service/pkg/service/grpc_service/grpc_playlist"
)

type PlaylistType struct {
	State
}

func (s PlaylistType) Get(entityId string) error {
	_, err := grpc_playlist.GetById(entityId, constants.TypePlaylist)
	return err
}
