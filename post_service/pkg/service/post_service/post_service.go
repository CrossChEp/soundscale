package post_service

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post_service/pkg/model"
	"post_service/pkg/proto/post_service_proto"
	"post_service/pkg/repository"
	"post_service/pkg/service/auxilary"
	"post_service/pkg/service/checker"
	"post_service/pkg/service/filter"
	"post_service/pkg/service/logger"
)

func AddPost(r *post_service_proto.AddPostReq) (*model.PostModel, error) {
	if !checker.IsPostValid(r) {
		logger.ErrorLog(fmt.Sprintf("Error: post data is invalid"))
		return nil, errors.New(fmt.Sprintf("Error: post data is invalid"))
	}
	filterAddReq(r)
	postData, err := repository.Save(r)
	if err != nil {
		return nil, err
	}
	post, err := repository.GetByObjectId(postData.InsertedID.(primitive.ObjectID))
	if err != nil {
		return nil, err
	}
	return post, nil
}

func filterAddReq(r *post_service_proto.AddPostReq) {
	r.Songs = filter.RemoveNonExistingSongs(r.Songs)
	r.Playlists = filter.RemoveNonExistingPlaylists(r.Playlists)
	r.Albums = filter.RemoveNonExistingAlbums(r.Albums)
}

func GetPost(r *post_service_proto.GetPostReq) (*model.PostModel, error) {
	return repository.Get(r.PostId)
}

func UpdatePost(r *post_service_proto.UpdatePostReq) (*model.PostModel, error) {
	if !checker.IsUpdatePostValid(r) {
		logger.ErrorLog(fmt.Sprintf("Error: post data is invalid"))
		return nil, errors.New(fmt.Sprintf("Error: post data is invalid"))
	}
	filterUpdateReq(r)
	if err := repository.Update(r); err != nil {
		return nil, err
	}
	return repository.Get(r.PostId)
}

func filterUpdateReq(r *post_service_proto.UpdatePostReq) {
	r.Songs = filter.RemoveNonExistingSongs(r.Songs)
	r.Playlists = filter.RemoveNonExistingPlaylists(r.Playlists)
	r.Albums = filter.RemoveNonExistingAlbums(r.Albums)
}

func DeletePost(r *post_service_proto.DeletePostReq) error {
	post, err := repository.Get(r.PostId)
	if err != nil {
		return err
	}
	if post.AuthorId != r.UserId {
		logger.ErrorLog(fmt.Sprintf("Error: user %s is not an author of post %s. Post author: %s", r.UserId, r.PostId, post.AuthorId))
		return errors.New(fmt.Sprintf("Error: user %s is not an author of post %s. Post author: %s", r.UserId, r.PostId, post.AuthorId))
	}
	return repository.Delete(r.PostId)
}

func GetUserPosts(r *post_service_proto.GetUserPostsReq) (*model.PostsModel, error) {
	return repository.GetUserPosts(r.UserId)
}

func LikePost(r *post_service_proto.LikePostReq) (*model.PostModel, error) {
	post, err := repository.Get(r.PostId)
	if err != nil {
		return nil, err
	}
	if checker.IsUserInArr(r.UserId, post.Liked) {
		logger.InfoLog(fmt.Sprintf("User %s already liked post %s", r.UserId, r.PostId))
		return post, nil
	}
	post.Disliked = auxilary.RemoveUserIfExists(r.UserId, post.Disliked)
	if err := repository.Like(r.UserId, post); err != nil {
		return nil, err
	}
	return repository.Get(r.PostId)
}

func DislikePost(r *post_service_proto.DislikePostReq) (*model.PostModel, error) {
	post, err := repository.Get(r.PostId)
	if err != nil {
		return nil, err
	}
	if checker.IsUserInArr(r.UserId, post.Disliked) {
		logger.InfoLog(fmt.Sprintf("User %s already disliked post %s", r.UserId, r.PostId))
		return post, nil
	}
	post.Liked = auxilary.RemoveUserIfExists(r.UserId, post.Liked)
	if err := repository.Dislike(r.UserId, post); err != nil {
		return nil, err
	}
	return repository.Get(r.PostId)
}

func GetUserLiked(r *post_service_proto.GetUserLikedReq) (*model.PostsModel, error) {
	return repository.GetUserLiked(r.UserId)
}

func GetUserDisliked(r *post_service_proto.GetUserDislikedReq) (*model.PostsModel, error) {
	return repository.GetUserDisliked(r.UserId)
}
