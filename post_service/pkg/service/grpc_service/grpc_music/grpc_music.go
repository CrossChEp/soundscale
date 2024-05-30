package grpc_music

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"post_service/pkg/config/service_address_config"
	"post_service/pkg/proto/music_service_proto"
	"post_service/pkg/service/logger"
)

func GetSongById(songId string) (*music_service_proto.GetResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details: %v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.GetReq{Id: songId}
	resp, err := musicService.GetSong(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get song. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get song. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}
