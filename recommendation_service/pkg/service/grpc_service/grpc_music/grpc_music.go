package grpc_music

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"recommendation_service/pkg/config/service_address_config"
	"recommendation_service/pkg/proto/music_service_proto"
	"recommendation_service/pkg/service/logger"
)

func GetBySinger(singerId string) (*music_service_proto.GetSongsResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details: %v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.GetBySingerReq{SingerId: singerId}
	resp, err := musicService.GetBySinger(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get songs by singer. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get songs by singer. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}
