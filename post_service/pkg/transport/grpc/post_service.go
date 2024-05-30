package transport

import (
	"context"
	"errors"
	"fmt"
	"post_service/pkg/proto/post_service_proto"
	"post_service/pkg/service/logger"
	"post_service/pkg/service/post_service"
	"post_service/pkg/service/response"
)

type PostService struct {
	post_service_proto.PostServiceServer
}

func (s *PostService) AddPost(_ context.Context,
	r *post_service_proto.AddPostReq) (*post_service_proto.PostResp, error) {
	logger.InfoLog(fmt.Sprintf("(AddPost) Request was recieved"))
	post, err := post_service.AddPost(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't add post of user %s. Details: %v", r.UserId, err))
	}
	return response.CreatePostResponse(post), nil
}

func (s *PostService) GetPost(_ context.Context,
	r *post_service_proto.GetPostReq) (*post_service_proto.PostResp, error) {
	logger.InfoLog(fmt.Sprintf("(GetPost) Request was recieved"))
	post, err := post_service.GetPost(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get post %s. Details: %v", r.PostId, err))
	}
	return response.CreatePostResponse(post), nil
}

func (s *PostService) UpdatePost(_ context.Context,
	r *post_service_proto.UpdatePostReq) (*post_service_proto.PostResp, error) {
	logger.InfoLog(fmt.Sprintf("(UpdatePost) Request was recieved"))
	post, err := post_service.UpdatePost(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't update post %s. Details: %v", r.PostId, err))
	}
	return response.CreatePostResponse(post), nil
}

func (s *PostService) DeletePost(_ context.Context,
	r *post_service_proto.DeletePostReq) (*post_service_proto.Resp, error) {
	logger.InfoLog(fmt.Sprintf("(DeletePost) Request was recieved"))
	if err := post_service.DeletePost(r); err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't delete post %s. Details: %v", r.PostId, err))
	}
	return &post_service_proto.Resp{Content: fmt.Sprintf("post %s was deleted", r.PostId)}, nil
}

func (s *PostService) GetUserPosts(_ context.Context,
	r *post_service_proto.GetUserPostsReq) (*post_service_proto.GetPostsResp, error) {
	logger.InfoLog(fmt.Sprintf("(GetUserPosts) Request was recieved"))
	posts, err := post_service.GetUserPosts(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get post of user %s. Details: %v", r.UserId, err))
	}
	return response.CreatePostsResponse(posts.Posts), nil
}

func (s *PostService) LikePost(_ context.Context,
	r *post_service_proto.LikePostReq) (*post_service_proto.PostResp, error) {
	post, err := post_service.LikePost(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't like post %s by user %s. Details: %v", r.PostId, r.UserId, err))
	}
	return response.CreatePostResponse(post), nil
}

func (s *PostService) DislikePost(_ context.Context,
	r *post_service_proto.DislikePostReq) (*post_service_proto.PostResp, error) {
	post, err := post_service.DislikePost(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't dislike post %s by user %s. Details: %v", r.PostId, r.UserId, err))
	}
	return response.CreatePostResponse(post), nil
}

func (s *PostService) GetUserLiked(_ context.Context,
	r *post_service_proto.GetUserLikedReq) (*post_service_proto.GetPostsResp, error) {
	posts, err := post_service.GetUserLiked(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get liked posts of user %s. Details: %v", r.UserId, err))
	}
	return response.CreatePostsResponse(posts.Posts), nil
}

func (s *PostService) GetUserDisliked(_ context.Context,
	r *post_service_proto.GetUserDislikedReq) (*post_service_proto.GetPostsResp, error) {
	posts, err := post_service.GetUserDisliked(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get disliked posts of user %s. Details: %v", r.UserId, err))
	}
	return response.CreatePostsResponse(posts.Posts), nil
}
