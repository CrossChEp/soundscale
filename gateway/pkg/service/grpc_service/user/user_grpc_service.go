package grpc_service

import (
	"context"
	"errors"
	"fmt"
	"gateway/pkg/config/service_address_config"
	"gateway/pkg/model/user_models"
	"gateway/pkg/proto/user_service_proto"
	"gateway/pkg/service/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func GetUserByNickname(nickname string) (*user_service_proto.GetResponse, error) {
	conn, err := grpc.Dial(*service_address_config.UserServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to user service. Details: %v", err))
		return nil, err
	}
	userService := user_service_proto.NewUserServiceClient(conn)
	req := &user_service_proto.GetByNicknameRequest{
		Nickname: nickname,
	}
	resp, err := userService.GetByNickname(context.TODO(), req)
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

func AddUser(userData *user_models.UserRegisterModel) (*user_service_proto.GetPrivateResponse, error) {
	conn, err := grpc.Dial(*service_address_config.UserServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to user service. Details: %v", err))
		return nil, err
	}
	userService := user_service_proto.NewUserServiceClient(conn)
	req := &user_service_proto.AddRequest{
		Nickname:    userData.Nickname,
		Email:       userData.Email,
		PhoneNumber: userData.PhoneNumber,
		Password:    userData.Password,
	}
	resp, err := userService.Add(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: can't register user. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: can't register user. Details: %v", resp.Error))
		return nil, err
	}
	return resp, nil
}

func UpdateUser(newUserData *user_models.UserUpdateModel) (*user_service_proto.GetResponse, error) {
	conn, err := grpc.Dial(*service_address_config.UserServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to user service. Details: %v", err))
		return nil, err
	}
	userService := user_service_proto.NewUserServiceClient(conn)
	req := &user_service_proto.UpdateRequest{
		UserId:      newUserData.UserId,
		Nickname:    newUserData.Nickname,
		Login:       newUserData.Login,
		Email:       newUserData.Email,
		PhoneNumber: newUserData.PhoneNumber,
		Password:    newUserData.Password,
	}
	resp, err := userService.Update(context.Background(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update user: Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update user: Details: %v", err))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}

func DeleteUser(userId string) (*user_service_proto.GetResponse, error) {
	conn, err := grpc.Dial(*service_address_config.UserServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to user service. Details: %v", err))
		return nil, err
	}
	userService := user_service_proto.NewUserServiceClient(conn)
	req := &user_service_proto.DeleteRequest{UserId: userId}
	resp, err := userService.Delete(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete user. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete user. Details: %v", err))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}
