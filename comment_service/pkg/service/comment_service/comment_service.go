package comment_service

import (
	"comment_service/pkg/config/constants"
	model "comment_service/pkg/model/comment"
	"comment_service/pkg/proto/comment_service_proto"
	"comment_service/pkg/repository"
	"comment_service/pkg/service/checker"
	"comment_service/pkg/service/grpc_service/grpc_user"
	"comment_service/pkg/service/logger"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)

func AddComment(r *comment_service_proto.AddCommentReq) (*model.CommentGetModel, error) {
	_, err := grpc_user.GetUserById(r.UserId)
	if err != nil {
		return nil, err
	}
	if r.EntityType == "" || !checker.IsConstantInArr(r.EntityType, constants.EntityTypes) {
		logger.ErrorLog(fmt.Sprintf("Error: invalid entity type"))
		return nil, errors.New("invalid entity type")
	}
	if !checker.IsEntityExists(r.EntityType, r.EntityId) {
		logger.ErrorLog(fmt.Sprintf("Error: entity doesn't exist"))
		return nil, errors.New("entity doesn't exist")
	}
	res, err := repository.Save(r)
	if err != nil {
		return nil, err
	}
	return repository.GetByObjectId(res.InsertedID.(primitive.ObjectID))
}

func Get(r *comment_service_proto.GetCommentReq) (*model.CommentGetModel, error) {
	return repository.GetWithType(r.CommentId, r.EntityType)
}

func Update(r *comment_service_proto.UpdateCommentReq) (*model.CommentGetModel, error) {
	dir, _ := os.Getwd()
	comment, err := repository.Get(r.CommentId)
	if err != nil {
		return nil, err
	}
	if r.UserId != comment.AuthorId {
		err = errors.New(fmt.Sprintf("user %s is nott an author of comment %s", r.UserId, r.CommentId))
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: user %s is not an authpr of comment %s", r.UserId, r.CommentId), err, dir)
		return nil, err
	}
	if err := repository.Update(r.CommentId, r.NewContent); err != nil {
		return nil, err
	}
	return repository.Get(r.CommentId)
}

func Delete(r *comment_service_proto.DeleteCommentReq) error {
	dir, _ := os.Getwd()
	comment, err := repository.Get(r.CommentId)
	if err != nil {
		return err
	}
	if comment.AuthorId != r.UserId {
		err = errors.New(fmt.Sprintf("user %s is nott an author of comment %s", r.UserId, r.CommentId))
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: user %s is not an authpr of comment %s", r.UserId, r.CommentId), err, dir)
		return err
	}
	return repository.Delete(r.CommentId)
}

func GetEntityComments(r *comment_service_proto.GetEntityCommentsReq) ([]model.CommentGetModel, error) {
	if !checker.IsEntityExists(r.EntityType, r.EntityId) {
		logger.ErrorLog(fmt.Sprintf("Error: entity doesn't exist"))
		return nil, errors.New("entity doesn't exist")
	}
	return repository.GetEntityComments(r.EntityId, r.GetEntityType())
}

func GetUserComments(r *comment_service_proto.GetUserCommentsReq) ([]model.CommentGetModel, error) {
	return repository.GetUserComments(r.UserId)
}
