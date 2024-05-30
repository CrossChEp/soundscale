package post_service

import (
	"gateway/pkg/model/post_model"
	"gateway/pkg/service/grpc_service/grpc_post"
)

func AddPost(postData *post_model.PostAddModel) (*post_model.PostGetModel, error) {
	resp, err := grpc_post.AddPost(postData)
	if err != nil {
		return nil, err
	}
	post := &post_model.PostGetModel{}
	post.ToModel(resp)
	return post, nil
}

func GetPost(postId string) (*post_model.PostGetModel, error) {
	resp, err := grpc_post.GetPost(postId)
	if err != nil {
		return nil, err
	}
	post := &post_model.PostGetModel{}
	post.ToModel(resp)
	return post, nil
}

func GetUserPosts(userId string) (*post_model.PostsGetModel, error) {
	resp, err := grpc_post.GetUserPosts(userId)
	if err != nil {
		return nil, err
	}
	posts := &post_model.PostsGetModel{}
	posts.ToModel(resp)
	posts.AuthorId = userId
	return posts, nil
}

func UpdatePost(postData *post_model.PostUpdateModel) (*post_model.PostGetModel, error) {
	resp, err := grpc_post.UpdatePost(postData)
	if err != nil {
		return nil, err
	}
	post := &post_model.PostGetModel{}
	post.ToModel(resp)
	return post, nil
}

func DeletePost(postId string, userId string) error {
	if _, err := grpc_post.DeletePost(postId, userId); err != nil {
		return err
	}
	return nil
}

func LikePost(userId string, postId string) (*post_model.PostGetModel, error) {
	resp, err := grpc_post.LikePost(userId, postId)
	if err != nil {
		return nil, err
	}
	post := &post_model.PostGetModel{}
	post.ToModel(resp)
	return post, nil
}

func DislikePost(userId string, postId string) (*post_model.PostGetModel, error) {
	resp, err := grpc_post.DislikePost(userId, postId)
	if err != nil {
		return nil, err
	}
	post := &post_model.PostGetModel{}
	post.ToModel(resp)
	return post, nil
}

func GetUserLikedPosts(userId string) (*post_model.PostsGetModel, error) {
	resp, err := grpc_post.GetUserLiked(userId)
	if err != nil {
		return nil, err
	}
	posts := &post_model.PostsGetModel{}
	posts.ToModel(resp)
	posts.AuthorId = userId
	return posts, nil
}

func GetUserDislikedPosts(userId string) (*post_model.PostsGetModel, error) {
	resp, err := grpc_post.GetUserDisliked(userId)
	if err != nil {
		return nil, err
	}
	posts := &post_model.PostsGetModel{}
	posts.ToModel(resp)
	posts.AuthorId = userId
	return posts, nil
}
