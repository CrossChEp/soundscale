package grpc_song

import (
	"comment_service/pkg/config/services_address_config"
	"comment_service/pkg/proto/music_service_proto"
	"comment_service/pkg/service/logger"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetById(songId string) (*music_service_proto.GetResp, error) {
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*services_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog("Error: couldn't connect to music service.")
		logger.DebugLog(fmt.Sprintf("%v Details: %v", curDir, err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.GetReq{Id: songId}
	resp, err := musicService.GetSong(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get response from music service. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get response from music service. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}
