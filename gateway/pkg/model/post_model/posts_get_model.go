package post_model

import "gateway/pkg/proto/post_service_proto"

type PostsGetModel struct {
	AuthorId string         `json:"author_id"`
	Posts    []PostGetModel `json:"posts"`
}

func (model *PostsGetModel) ToModel(resp *post_service_proto.GetPostsResp) {
	model.Posts = createPostsModel(resp)
}

func createPostsModel(postsResps *post_service_proto.GetPostsResp) []PostGetModel {
	var posts []PostGetModel
	for _, postResp := range postsResps.Posts {
		post := PostGetModel{}
		post.ToModel(postResp)
		posts = append(posts, post)
	}
	return posts
}
