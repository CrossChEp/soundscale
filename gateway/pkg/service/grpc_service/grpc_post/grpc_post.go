package grpc_post

import (
	"context"
	"fmt"
	"gateway/pkg/config/service_address_config"
	"gateway/pkg/model/post_model"
	"gateway/pkg/proto/post_service_proto"
	"gateway/pkg/service/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func AddPost(postModel *post_model.PostAddModel) (*post_service_proto.PostResp, error) {
	logger.InfoLog("(AddPost) Connecting to post service...")
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to collection service"))
		logger.DebugLog(fmt.Sprintf("(AddPost) %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(AddPost) Connected to post service")
	postService := post_service_proto.NewPostServiceClient(conn)
	postReq := &post_service_proto.AddPostReq{
		UserId:    postModel.AuthorId,
		Content:   postModel.Content,
		Songs:     postModel.Songs,
		Playlists: postModel.Playlists,
		Albums:    postModel.Albums,
	}
	logger.InfoLog("(AddPost) Sending add post request to post service")
	response, err := postService.AddPost(context.TODO(), postReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add post of user %s", postModel.AuthorId))
		logger.DebugLog(fmt.Sprintf("(AddPost) %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(AddPost) Request sent successfully and data was got")
	return response, nil
}

func GetPost(postId string) (*post_service_proto.PostResp, error) {
	logger.InfoLog("(GetPost) Connecting to post service...")
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to collection service"))
		logger.DebugLog(fmt.Sprintf("GetPost %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(GetPost) Connected to post service")
	postService := post_service_proto.NewPostServiceClient(conn)
	postReq := &post_service_proto.GetPostReq{PostId: postId}
	logger.InfoLog("(GetPost) Sending get post request to post service")
	response, err := postService.GetPost(context.TODO(), postReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get post %s", postId))
		logger.DebugLog(fmt.Sprintf("(GetPost) %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(GetPost) Request sent successfully and data was got")
	return response, nil
}

func GetUserPosts(userId string) (*post_service_proto.GetPostsResp, error) {
	logger.InfoLog("(GetUserPosts) Connecting to post service...")
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to collection service"))
		logger.DebugLog(fmt.Sprintf("GetUserPosts %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(GetUserPosts) Connected to post service")
	postService := post_service_proto.NewPostServiceClient(conn)
	postReq := &post_service_proto.GetUserPostsReq{UserId: userId}
	logger.InfoLog("(GetUserPosts) Sending get user posts request to post service")
	response, err := postService.GetUserPosts(context.TODO(), postReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get post of user %s", userId))
		logger.DebugLog(fmt.Sprintf("(GetUserPosts) %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(GetUserPosts) Request sent successfully and data was got")
	return response, nil
}

func UpdatePost(updateData *post_model.PostUpdateModel) (*post_service_proto.PostResp, error) {
	logger.InfoLog("(UpdatePost) Connecting to post service...")
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to collection service"))
		logger.DebugLog(fmt.Sprintf("UpdatePost %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(UpdatePost) Connected to post service")
	postService := post_service_proto.NewPostServiceClient(conn)
	postReq := &post_service_proto.UpdatePostReq{
		PostId:    updateData.PostId,
		Content:   updateData.Content,
		Songs:     updateData.Songs,
		Playlists: updateData.Playlists,
		Albums:    updateData.Albums,
		UserId:    updateData.AuthorId,
	}
	logger.InfoLog("(UpdatePost) Sending update post request to post service")
	response, err := postService.UpdatePost(context.TODO(), postReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update post %s", updateData.PostId))
		logger.DebugLog(fmt.Sprintf("(UpdatePost) %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(UpdatePost) Request sent successfully and data was got")
	return response, nil
}

func DeletePost(postId string, userId string) (*post_service_proto.Resp, error) {
	logger.InfoLog("(DeletePost) Connecting to post service...")
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to collection service"))
		logger.DebugLog(fmt.Sprintf("DeletePost %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(DeletePost) Connected to post service")
	postService := post_service_proto.NewPostServiceClient(conn)
	postReq := &post_service_proto.DeletePostReq{
		PostId: postId,
		UserId: userId,
	}
	logger.InfoLog("(DeletePost) Sending delete post request to post service")
	response, err := postService.DeletePost(context.TODO(), postReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete post %s", postId))
		logger.DebugLog(fmt.Sprintf("(DeletePost) %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(DeletePost) Request sent successfully and post was deleted")
	return response, nil
}

func LikePost(userId string, postId string) (*post_service_proto.PostResp, error) {
	logger.InfoLog("(LikePost) Connecting to post service...")
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to collection service"))
		logger.DebugLog(fmt.Sprintf("LikePost %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(LikePost) Connected to post service")
	postService := post_service_proto.NewPostServiceClient(conn)
	postReq := &post_service_proto.LikePostReq{
		PostId: postId,
		UserId: userId,
	}
	logger.InfoLog("(LikePost) Sending like post request to post service")
	response, err := postService.LikePost(context.TODO(), postReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't like post %s", postId))
		logger.DebugLog(fmt.Sprintf("(LikePost) %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(LikePost) Request sent successfully and post was liked")
	return response, nil
}

func DislikePost(userId string, postId string) (*post_service_proto.PostResp, error) {
	logger.InfoLog("(DislikePost) Connecting to post service...")
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to collection service"))
		logger.DebugLog(fmt.Sprintf("DislikePost %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(DislikePost) Connected to post service")
	postService := post_service_proto.NewPostServiceClient(conn)
	postReq := &post_service_proto.DislikePostReq{
		PostId: postId,
		UserId: userId,
	}
	logger.InfoLog("(DislikePost) Sending dislike post request to post service")
	response, err := postService.DislikePost(context.TODO(), postReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't dislike post %s", postId))
		logger.DebugLog(fmt.Sprintf("(DislikePost) %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(LikePost) Request sent successfully and post was liked")
	return response, nil
}

func GetUserLiked(userId string) (*post_service_proto.GetPostsResp, error) {
	logger.InfoLog("(GetUserLiked) Connecting to post service...")
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to collection service"))
		logger.DebugLog(fmt.Sprintf("GetUserLiked %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(GetUserLiked) Connected to post service")
	postService := post_service_proto.NewPostServiceClient(conn)
	postReq := &post_service_proto.GetUserLikedReq{
		UserId: userId,
	}
	logger.InfoLog("(GetUserLiked) Sending get user liked post request to post service")
	response, err := postService.GetUserLiked(context.TODO(), postReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user %s liked post", userId))
		logger.DebugLog(fmt.Sprintf("(GetUserLiked) %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(GetUserLiked) Request sent successfully and user liked post were got")
	return response, nil
}

func GetUserDisliked(userId string) (*post_service_proto.GetPostsResp, error) {
	logger.InfoLog("(GetUserDisliked) Connecting to post service...")
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*service_address_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to collection service"))
		logger.DebugLog(fmt.Sprintf("GetUserDisliked %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(GetUserDisliked) Connected to post service")
	postService := post_service_proto.NewPostServiceClient(conn)
	postReq := &post_service_proto.GetUserDislikedReq{
		UserId: userId,
	}
	logger.InfoLog("(GetUserDisliked) Sending get user disliked post request to post service")
	response, err := postService.GetUserDisliked(context.TODO(), postReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user %s disliked post", userId))
		logger.DebugLog(fmt.Sprintf("(GetUserDisliked) %v : %v", curDir, err))
		return nil, err
	}
	logger.InfoLog("(GetUserDisliked) Request sent successfully and user disliked post were got")
	return response, nil
}
