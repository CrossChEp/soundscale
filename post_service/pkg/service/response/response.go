package response

import (
	"post_service/pkg/model"
	"post_service/pkg/proto/post_service_proto"
)

func CreatePostsResponse(posts []model.PostModel) *post_service_proto.GetPostsResp {
	var postResps []*post_service_proto.PostResp
	for _, post := range posts {
		postResps = append(postResps, CreatePostResponse(&post))
	}
	return &post_service_proto.GetPostsResp{Posts: postResps}
}

func CreatePostResponse(post *model.PostModel) *post_service_proto.PostResp {
	return &post_service_proto.PostResp{
		Id:           post.Id,
		AuthorId:     post.AuthorId,
		Content:      post.Content,
		Songs:        post.Songs,
		Playlists:    post.Playlists,
		Albums:       post.Albums,
		Liked:        post.Liked,
		Disliked:     post.Disliked,
		CreationDate: post.CreationDate,
	}
}
