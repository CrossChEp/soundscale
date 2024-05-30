package grpc_post

import (
	"comment_service/pkg/config/services_address_config"
	"comment_service/pkg/proto/post_service_proto"
	"comment_service/pkg/service/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetById(postId string) (*post_service_proto.PostResp, error) {
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*services_address_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog("Error: couldn't connect to post service.")
		logger.DebugLog(fmt.Sprintf("%v Details: %v", curDir, err))
		return nil, err
	}
	postService := post_service_proto.NewPostServiceClient(conn)
	req := &post_service_proto.GetPostReq{PostId: postId}
	resp, err := postService.GetPost(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get response from post service. Details: %v", err))
		return nil, err
	}
	return resp, nil
}
