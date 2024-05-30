package state

import (
	"comment_service/pkg/service/grpc_service/grpc_song"
)

type SongState struct {
	State
}

func (s SongState) Get(entityId string) error {
	_, err := grpc_song.GetById(entityId)
	return err
}
