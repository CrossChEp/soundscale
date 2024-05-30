package model

import "comment_service/pkg/proto/comment_service_proto"

type CommentGetModel struct {
	Id           string `bson:"_id"`
	AuthorId     string `bson:"user_id"`
	EntityType   string `bson:"entity_type"`
	EntityId     string `bson:"entity_id"`
	Content      string `bson:"content"`
	CreationDate string `bson:"creation_date"`
}

func (m *CommentGetModel) ConvertToResponse() *comment_service_proto.GetCommentResp {
	return &comment_service_proto.GetCommentResp{
		CommentId:    m.Id,
		AuthorId:     m.AuthorId,
		EntityType:   m.EntityType,
		EntityId:     m.EntityId,
		Content:      m.Content,
		CreationDate: m.CreationDate,
	}
}
