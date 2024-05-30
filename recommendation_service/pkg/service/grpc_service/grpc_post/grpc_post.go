package grpc_post

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"recommendation_service/pkg/config/service_address_config"
	"recommendation_service/pkg/proto/post_service_proto"
	"recommendation_service/pkg/service/logger"
)

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
