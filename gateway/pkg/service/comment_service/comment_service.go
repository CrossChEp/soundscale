package comment_service

import (
	"gateway/pkg/model/comment_model"
	"gateway/pkg/service/grpc_service/grpc_comment"
)

func AddComment(commentModel *comment_model.CommentAddModel) (*comment_model.CommentGetModel, error) {
	commentResp, err := grpc_comment.AddComment(commentModel)
	if err != nil {
		return nil, err
	}
	comment := &comment_model.CommentGetModel{}
	comment.ToModel(commentResp)
	return comment, nil
}

func GetEntityComments(entityId string, entityType string) ([]comment_model.CommentGetModel, error) {
	var comments []comment_model.CommentGetModel
	commentResps, err := grpc_comment.GetEntityComments(entityId, entityType)
	if err != nil {
		return nil, err
	}
	for _, commentResp := range commentResps.Comments {
		comment := comment_model.CommentGetModel{}
		comment.ToModel(commentResp)
		comments = append(comments, comment)
	}
	return comments, nil
}

func GetUserComments(userId string) ([]comment_model.CommentGetModel, error) {
	var comments []comment_model.CommentGetModel
	commentResps, err := grpc_comment.GetUserComments(userId)
	if err != nil {
		return nil, err
	}
	for _, commentResp := range commentResps.Comments {
		comment := comment_model.CommentGetModel{}
		comment.ToModel(commentResp)
		comments = append(comments, comment)
	}
	return comments, nil
}

func UpdateComment(commentData *comment_model.CommentUpdateModel) (*comment_model.CommentGetModel, error) {
	resp, err := grpc_comment.UpdateComment(commentData)
	if err != nil {
		return nil, err
	}
	comment := &comment_model.CommentGetModel{}
	comment.ToModel(resp)
	return comment, nil
}

func DeleteComment(commentId string, userId string) (string, error) {
	resp, err := grpc_comment.DeleteComment(commentId, userId)
	if err != nil {
		return "", err
	}
	return resp.Content, nil
}
