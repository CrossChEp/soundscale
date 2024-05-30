package state

import (
	"comment_service/pkg/service/grpc_service/grpc_post"
)

type PostState struct {
	State
}

func (s PostState) Get(entityId string) error {
	_, err := grpc_post.GetById(entityId)
	return err
}
