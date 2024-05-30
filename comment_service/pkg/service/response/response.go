package response

import (
	model "comment_service/pkg/model/comment"
	"comment_service/pkg/proto/comment_service_proto"
)

func CreateManyCommentsResponse(comments []model.CommentGetModel) *comment_service_proto.GetManyCommentsResp {
	return &comment_service_proto.GetManyCommentsResp{
		Comments: createRespArr(comments),
	}
}

func createRespArr(comments []model.CommentGetModel) []*comment_service_proto.GetCommentResp {
	var commentResps []*comment_service_proto.GetCommentResp
	for _, comment := range comments {
		commentResps = append(commentResps, comment.ConvertToResponse())
	}
	return commentResps
}
