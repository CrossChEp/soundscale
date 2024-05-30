package grpc_playlist

import (
	"comment_service/pkg/config/services_address_config"
	"comment_service/pkg/proto/playlist_service_proto"
	"comment_service/pkg/service/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetById(playlistId string, playlistType string) (*playlist_service_proto.GetPlaylistResp, error) {
	curDir, _ := os.Getwd()
	conn, err := grpc.Dial(*services_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog("Error: couldn't connect to playlist service.")
		logger.DebugLog(fmt.Sprintf("%v Details: %v", curDir, err))
		return nil, err
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	req := &playlist_service_proto.GetPlaylistReq{
		Id:           playlistId,
		PlaylistType: playlistType,
	}
	resp, err := playlistService.GetPlaylist(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get response from playlist service. Details: %v", err))
		return nil, err
	}
	return resp, nil
}
