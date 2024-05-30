package grpc_service

import (
	"context"
	"errors"
	"fmt"
	"gateway/pkg/config/constants"
	"gateway/pkg/config/service_address_config"
	"gateway/pkg/model/playlist_model"
	"gateway/pkg/proto/playlist_service_proto"
	"gateway/pkg/service/checkers"
	"gateway/pkg/service/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AddPlaylist(playlist *playlist_model.PlaylistAddModel) (*playlist_service_proto.GetPlaylistResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to playlist service. Details: %v", err))
		return nil, err
	}
	if !checkers.IsPlaylistTypeExists(playlist.Type) {
		playlist.Type = constants.PlaylistType
	}
	req := &playlist_service_proto.AddPlaylistReq{
		Name:     playlist.Name,
		AuthorId: playlist.AuthorId,
		SongIds:  playlist.Songs,
		Type:     playlist.Type,
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	resp, err := playlistService.AddPlaylist(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add playlist. Details: %v", err))
		return nil, err
	}
	return resp, nil
}

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

func GetPlaylistByAuthor(authorId string, playlistType string) (*playlist_service_proto.GetPlaylistsResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to playlist service. Details: %v", err))
		return nil, err
	}
	req := &playlist_service_proto.GetPlaylistsByAuthorReq{
		Id:           authorId,
		PlaylistType: playlistType,
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	resp, err := playlistService.GetPlaylistsByAuthor(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get playlists. Details: %v", err))
		return nil, err
	}
	return resp, nil
}

func AddSongToPlaylist(addSongData *playlist_model.AddSongsToPlaylistModel) (*playlist_service_proto.GetPlaylistResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to playlist service. Details: %v", err))
		return nil, err
	}
	req := &playlist_service_proto.AddSongsToPlaylistReq{
		AuthorId:     addSongData.AuthorId,
		PlaylistId:   addSongData.PlaylistId,
		SongIds:      addSongData.SongIds,
		PlaylistType: addSongData.PlaylistType,
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	resp, err := playlistService.AddSongToPlaylist(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add song to playlist. Details: %v", err))
		return nil, err
	}
	return resp, nil
}

func DeleteSongFromPlaylist(songData *playlist_model.DeleteSongFromPlaylistModel) (*playlist_service_proto.GetPlaylistResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to playlist service. Details: %v", err))
		return nil, err
	}
	req := &playlist_service_proto.DeleteSongsFromPlaylistReq{
		AuthorId:     songData.AuthorId,
		PlaylistId:   songData.PlaylistId,
		SongsIds:     songData.SongIds,
		PlaylistType: songData.PlaylistType,
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	resp, err := playlistService.DeleteSongFromPlaylist(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete song from playlist. Details: %v", err))
		return nil, err
	}
	return resp, nil
}

func Delete(deleteData *playlist_model.PlaylistDeleteModel) (*playlist_service_proto.MessageResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to playlist service. Details: %v", err))
		return nil, err
	}
	req := &playlist_service_proto.DeletePlaylistReq{
		AuthorId:     deleteData.AuthorId,
		Id:           deleteData.PlaylistId,
		PlaylistType: deleteData.PlaylistType,
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	resp, err := playlistService.DeletePlaylist(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete playlist. Details: %v", err))
		return nil, err
	}
	return resp, nil
}

func AddCurrentPlaylist(playlistData *playlist_model.AddCurrentPlaylistModel) (*playlist_service_proto.AddCurrentPlaylistResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details: %v", err))
		return nil, err
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	req := &playlist_service_proto.AddCurrentPlaylistReq{
		UserId: playlistData.UserId,
		Songs:  playlistData.Songs,
	}
	resp, err := playlistService.AddCurrentPlaylist(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add playlist. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add playlist. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}

func LikePlaylist(likeModel playlist_model.LikeDislikePlaylist) (*playlist_service_proto.GetPlaylistResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog("Error: couldn't connect to playlist service")
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	req := &playlist_service_proto.LikePlaylistReq{
		UserId:       likeModel.UserId,
		PlaylistId:   likeModel.PlaylistId,
		PlaylistType: likeModel.PlaylistType,
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	resp, err := playlistService.LikePlaylist(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't like playlist %s by user %s.", likeModel.PlaylistId, likeModel.UserId))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return resp, nil
}

func DislikePlaylist(likeModel playlist_model.LikeDislikePlaylist) (*playlist_service_proto.GetPlaylistResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog("Error: couldn't connect to playlist service")
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	req := &playlist_service_proto.DislikePlaylistReq{
		UserId:       likeModel.UserId,
		PlaylistId:   likeModel.PlaylistId,
		PlaylistType: likeModel.PlaylistType,
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	resp, err := playlistService.DislikePlaylist(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't dislike playlist %s by user %s.", likeModel.PlaylistId, likeModel.UserId))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return resp, nil
}

func GetUserLikedPlaylists(userId string, playlistType string) (*playlist_service_proto.GetPlaylistsResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog("Error: couldn't connect to playlist service")
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	req := &playlist_service_proto.GetUserLikedReq{
		UserId:       userId,
		PlaylistType: playlistType,
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	resp, err := playlistService.GetUserLiked(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get liked playlists of user %s", userId))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return resp, nil
}

func GetUserDislikedPlaylists(userId string, playlistType string) (*playlist_service_proto.GetPlaylistsResp, error) {
	conn, err := grpc.Dial(*service_address_config.PlaylistServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog("Error: couldn't connect to playlist service")
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	req := &playlist_service_proto.GetUserDislikedReq{
		UserId:       userId,
		PlaylistType: playlistType,
	}
	playlistService := playlist_service_proto.NewPlaylistServiceClient(conn)
	resp, err := playlistService.GetUserDisliked(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get disliked playlists of user %s", userId))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return resp, nil
}
