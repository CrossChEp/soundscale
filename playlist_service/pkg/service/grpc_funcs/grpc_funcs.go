package grpc_funcs

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"playlist_serivce/pkg/config/services_address_config"
	"playlist_serivce/pkg/proto/music_service_proto"
	"playlist_serivce/pkg/proto/user_service_proto"
	"playlist_serivce/pkg/service/logger"
)

func GetSong(songId string) (*music_service_proto.GetResp, error) {
	conn, err := grpc.Dial(*services_address_config.MusicServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music transport. Details %v\n", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.GetReq{Id: songId}
	resp, err := musicService.GetSong(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get song. Details: %v\n", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get song. Details: %v\n", resp.Error))
		return nil, err
	}
	if resp == nil {
		logger.ErrorLog("Error: can't get song with such id")
		return nil, errors.New("can't get song with such id")
	}
	return resp, nil
}

func GetUser(userId string) (*user_service_proto.GetResponse, error) {
	conn, err := grpc.Dial(*services_address_config.UserServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to user transport. Details %v\n", err))
		return nil, err
	}
	userService := user_service_proto.NewUserServiceClient(conn)
	req := &user_service_proto.GetByIdRequest{Id: userId}
	resp, err := userService.GetById(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user. Details: %v\n", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user. Details: %v\n", resp.Error))
		return nil, err
	}
	if resp == nil {
		logger.ErrorLog("Error: can't get user with such id")
		return nil, errors.New("can't get user with such id")
	}
	return resp, nil
}
