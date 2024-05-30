package grpc_playlist

import (
	"collection_service/pkg/config/services_address_config"
	"collection_service/pkg/proto/playlist_service_proto"
	"collection_service/pkg/service/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetPlaylist(playlistId string, playlistType string) (*playlist_service_proto.GetPlaylistResp, error) {
	conn, err := grpc.Dial(*services_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to playlist service. Details: %v", err))
		return nil, err
	}
	req := &playlist_service_proto.GetPlaylistReq{
		Id:           playlistId,
		PlaylistType: playlistType,
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	resp, err := playlistService.GetPlaylist(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get playlist. Details: %v", err))
		return nil, err
	}
	return resp, nil
}
