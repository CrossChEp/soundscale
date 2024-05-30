package comment_model

import "gateway/pkg/proto/comment_service_proto"

type CommentGetModel struct {
	Id           string `json:"id"`
	Author       string `json:"author"`
	EntityType   string `json:"entity_type"`
	EntityId     string `json:"entity_id"`
	CreationDate string `json:"creation_date"`
	Content      string `json:"content"`
}

func (m *CommentGetModel) ToModel(resp *comment_service_proto.GetCommentResp) {
	m.Id = resp.CommentId
	m.Author = resp.AuthorId
	m.EntityType = resp.EntityType
	m.EntityId = resp.EntityId
	m.CreationDate = resp.CreationDate
	m.Content = resp.Content
}
