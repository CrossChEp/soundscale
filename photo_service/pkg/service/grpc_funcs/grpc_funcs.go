package grpc_funcs

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"photo_service/pkg/config/services_adress_config"
	"photo_service/pkg/logger"
	"photo_service/pkg/proto/music_service_proto"
	"photo_service/pkg/proto/playlist_service_proto"
	"photo_service/pkg/proto/post_service_proto"
)

func GetSong(songId string) (*music_service_proto.GetResp, error) {
	conn, err := grpc.Dial(*services_adress_config.MusicServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music transport. Details: %v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := music_service_proto.GetReq{Id: songId}
	resp, err := musicService.GetSong(context.TODO(), &req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get songs. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get songs. Details: %s", resp.Error))
		return nil, err
	}
	return resp, nil
}

func GetPlaylist(playlistId string, playlistType string) (*playlist_service_proto.GetPlaylistResp, error) {
	conn, err := grpc.Dial(*services_adress_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to playlist service. Details: %v", err))
		return nil, err
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	req := &playlist_service_proto.GetPlaylistReq{
		Id:           playlistId,
		PlaylistType: playlistType,
	}
	resp, err := playlistService.GetPlaylist(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get playlist. Details: %v", err))
		return nil, err
	}
	return resp, nil
}

func GetPost(postId string) (*post_service_proto.PostResp, error) {
	conn, err := grpc.Dial(*services_adress_config.PostServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to playlist service. Details: %v", err))
		return nil, err
	}
	postService := post_service_proto.NewPostServiceClient(conn)
	req := &post_service_proto.GetPostReq{PostId: postId}
	resp, err := postService.GetPost(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get psot. Details: %v", err))
		return nil, err
	}
	return resp, nil
}
