package grpc_comment

import (
	"context"
	"fmt"
	"gateway/pkg/config/service_address_config"
	"gateway/pkg/model/comment_model"
	"gateway/pkg/proto/comment_service_proto"
	"gateway/pkg/service/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func AddComment(commentModel *comment_model.CommentAddModel) (*comment_service_proto.GetCommentResp, error) {
	dir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.CommentServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorWithDebugLog("Error: couldn't connect to comment service.", err, dir)
		return nil, err
	}
	commentService := comment_service_proto.NewCommentServiceClient(conn)
	req := &comment_service_proto.AddCommentReq{
		UserId:     commentModel.UserId,
		EntityId:   commentModel.EntityId,
		EntityType: commentModel.EntityType,
		Content:    commentModel.Content,
	}
	response, err := commentService.AddComment(context.TODO(), req)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't add comment by user %s.", commentModel.UserId), err, dir)
		return nil, err
	}
	return response, nil
}

func GetEntityComments(entityId string, entityType string) (*comment_service_proto.GetManyCommentsResp, error) {
	dir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.CommentServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorWithDebugLog("Error: couldn't connect to comment service.", err, dir)
		return nil, err
	}
	commentService := comment_service_proto.NewCommentServiceClient(conn)
	req := &comment_service_proto.GetEntityCommentsReq{
		EntityType: entityType,
		EntityId:   entityId,
	}
	response, err := commentService.GetEntityComments(context.TODO(), req)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't get entity %s comments by user", entityId), err, dir)
		return nil, err
	}
	return response, nil
}

func UpdateComment(commentModel *comment_model.CommentUpdateModel) (*comment_service_proto.GetCommentResp, error) {
	dir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.CommentServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorWithDebugLog("Error: couldn't connect to comment service.", err, dir)
		return nil, err
	}
	commentService := comment_service_proto.NewCommentServiceClient(conn)
	req := &comment_service_proto.UpdateCommentReq{
		UserId:     commentModel.UserId,
		CommentId:  commentModel.CommentId,
		NewContent: commentModel.NewContent,
	}
	response, err := commentService.UpdateComment(context.TODO(), req)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf(
			"Error: couldn't update comment %s by user %s",
			commentModel.CommentId,
			commentModel.UserId), err, dir)
		return nil, err
	}
	return response, nil
}

func DeleteComment(commentId string, userId string) (*comment_service_proto.Message, error) {
	dir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.CommentServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorWithDebugLog("Error: couldn't connect to comment service.", err, dir)
		return nil, err
	}
	commentService := comment_service_proto.NewCommentServiceClient(conn)
	req := &comment_service_proto.DeleteCommentReq{
		UserId:    userId,
		CommentId: commentId,
	}
	response, err := commentService.DeleteComment(context.TODO(), req)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf(
			"Error: couldn't delete comment %s by user %s",
			commentId,
			userId), err, dir)
		return nil, err
	}
	return response, nil
}

func GetUserComments(userId string) (*comment_service_proto.GetManyCommentsResp, error) {
	dir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.CommentServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorWithDebugLog("Error: couldn't connect to comment service.", err, dir)
		return nil, err
	}
	commentService := comment_service_proto.NewCommentServiceClient(conn)
	req := &comment_service_proto.GetUserCommentsReq{UserId: userId}
	response, err := commentService.GetUserComments(context.TODO(), req)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't get user %s comments", userId), err, dir)
		return nil, err
	}
	return response, nil
}
