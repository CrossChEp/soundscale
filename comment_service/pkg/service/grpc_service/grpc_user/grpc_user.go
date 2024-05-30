package grpc_user

import (
	"comment_service/pkg/config/services_address_config"
	"comment_service/pkg/proto/user_service_proto"
	"comment_service/pkg/service/logger"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetUserById(userId string) (*user_service_proto.GetResponse, error) {
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*services_address_config.UserServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog("Error: couldn't connect to user service.")
		logger.DebugLog(fmt.Sprintf("%v Details: %v", curDir, err))
		return nil, err
	}
	userService := user_service_proto.NewUserServiceClient(conn)
	req := &user_service_proto.GetByIdRequest{Id: userId}
	resp, err := userService.GetById(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get response from user service. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}
