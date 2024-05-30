package grpc_user

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"recommendation_service/pkg/config/service_address_config"
	"recommendation_service/pkg/proto/user_service_proto"
	"recommendation_service/pkg/service/logger"
)

func GetUserById(userId string) (*user_service_proto.GetResponse, error) {
	conn, err := grpc.Dial(*service_address_config.UserServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to user service. Details: %v", err))
		return nil, err
	}
	userService := user_service_proto.NewUserServiceClient(conn)
	req := &user_service_proto.GetByIdRequest{Id: userId}
	resp, err := userService.GetById(context.TODO(), req)
	if err != nil {
		logger.ErrorLog("Error: couldn't get response from user service")
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user"))
		logger.DebugLog(resp.Error)
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}
