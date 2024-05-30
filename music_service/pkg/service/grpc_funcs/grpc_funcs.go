package grpc_funcs

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"music_service/pkg/config/services_address_config"
	"music_service/pkg/model"
	"music_service/pkg/proto/collection_service_proto"
	user_service_proto2 "music_service/pkg/proto/user_service_proto"
	"music_service/pkg/service/logger"
)

func GetUserById(id string) (*user_service_proto2.GetResponse, error) {
	conn, err := grpc.Dial(*services_address_config.UserServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to user transport. Details: %v", err))
		return nil, err
	}
	userService := user_service_proto2.NewUserServiceClient(conn)
	data := user_service_proto2.GetByIdRequest{
		Id: id,
	}
	user, err := userService.GetById(context.TODO(), &data)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user by this id. Details: %v", err))
	}
	if user.Error != "" {
		return &user_service_proto2.GetResponse{Error: "user with such id doesn't exist"}, nil
	}
	return user, nil
}

func UpdateUser(newUserData *model.UserUpdateModel) (*user_service_proto2.GetResponse, error) {
	conn, err := grpc.Dial(*services_address_config.UserServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to user service. Details: %v", err))
		return nil, err
	}
	userService := user_service_proto2.NewUserServiceClient(conn)
	req := &user_service_proto2.UpdateRequest{
		UserId:      newUserData.UserId,
		Nickname:    newUserData.Nickname,
		Login:       newUserData.Login,
		Email:       newUserData.Email,
		PhoneNumber: newUserData.PhoneNumber,
		Password:    newUserData.Password,
		UserState:   newUserData.State,
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

func AddCreatedGenres(userId string, genres []string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*services_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	genresReq := &collection_service_proto.AddCreatedGenresReq{
		UserId: userId,
		Genres: genres,
	}
	response, err := collectionService.AddCreatedGenres(context.TODO(), genresReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add created genres to user %s collection. Details: %d", userId, err))
		return nil, err
	}
	return response, nil
}
