package grpc_playlist

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"post_service/pkg/config/service_address_config"
	"post_service/pkg/proto/playlist_service_proto"
	"post_service/pkg/service/logger"
)

func GetPlaylist(playlistId string, playlistType string) (*playlist_service_proto.GetPlaylistResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress,
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
