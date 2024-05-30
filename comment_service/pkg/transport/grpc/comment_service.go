package transport

import (
	"comment_service/pkg/proto/comment_service_proto"
	"comment_service/pkg/service/comment_service"
	"comment_service/pkg/service/logger"
	"comment_service/pkg/service/response"
	"context"
	"errors"
	"fmt"
)

type CommentService struct {
	comment_service_proto.CommentServiceServer
}

func (s *CommentService) AddComment(_ context.Context,
	req *comment_service_proto.AddCommentReq) (*comment_service_proto.GetCommentResp, error) {
	logger.InfoLog(fmt.Sprintf("Add comment function was called."))
	comment, err := comment_service.AddComment(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't add comment by user %s. Details: %v", req.UserId, err))
	}
	logger.InfoLog(fmt.Sprintf("Comment was added by user %s", req.UserId))
	return comment.ConvertToResponse(), nil
}

func (s *CommentService) GetComment(_ context.Context,
	req *comment_service_proto.GetCommentReq) (*comment_service_proto.GetCommentResp, error) {
	logger.InfoLog(fmt.Sprintf("Get comment function was called."))
	comment, err := comment_service.Get(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get comment %s. Details: %v", req.CommentId, err))
	}
	logger.InfoLog(fmt.Sprintf("Comment %s was getted", req.CommentId))
	return comment.ConvertToResponse(), nil
}

func (s *CommentService) UpdateComment(_ context.Context,
	req *comment_service_proto.UpdateCommentReq) (*comment_service_proto.GetCommentResp, error) {
	logger.InfoLog(fmt.Sprintf("Update comment function was called."))
	comment, err := comment_service.Update(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't update comment %s. Details: %v", req.CommentId, err))
	}
	logger.InfoLog(fmt.Sprintf("Commetn %s was updated by user %s", req.CommentId, req.UserId))
	return comment.ConvertToResponse(), nil
}

func (s *CommentService) DeleteComment(_ context.Context,
	req *comment_service_proto.DeleteCommentReq) (*comment_service_proto.Message, error) {
	logger.InfoLog(fmt.Sprintf("Delete comment function was called."))
	if err := comment_service.Delete(req); err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't delete comment %s. Details: %v", req.CommentId, err))
	}
	logger.InfoLog(fmt.Sprintf("Comment %s was deleted by user %s", req.CommentId, req.UserId))
	return &comment_service_proto.Message{Content: "comment was deleted"}, nil
}

func (s *CommentService) GetEntityComments(_ context.Context,
	req *comment_service_proto.GetEntityCommentsReq) (*comment_service_proto.GetManyCommentsResp, error) {
	logger.InfoLog("Get Entity comments function was called")
	comments, err := comment_service.GetEntityComments(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get entity comments. Details: %v", err))
	}
	logger.InfoLog(fmt.Sprintf("%s %s comments were gotten", req.EntityType, req.EntityId))
	return response.CreateManyCommentsResponse(comments), nil
}

func (s *CommentService) GetUserComments(_ context.Context,
	req *comment_service_proto.GetUserCommentsReq) (*comment_service_proto.GetManyCommentsResp, error) {
	logger.InfoLog("Get user comments function was called")
	comments, err := comment_service.GetUserComments(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get user comments. Details: %v", err))
	}
	logger.InfoLog(fmt.Sprintf("User %s comments were gotten", req.UserId))
	return response.CreateManyCommentsResponse(comments), nil
}
