package grpc_service

import (
	"context"
	"errors"
	"fmt"
	"gateway/pkg/config/service_address_config"
	"gateway/pkg/model/music_model"
	"gateway/pkg/proto/music_service_proto"
	"gateway/pkg/service/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func AddSong(songData *music_model.MusicAddModel) (*music_service_proto.GetResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details: %v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.AddReq{
		AuthorId:      songData.AuthorId,
		Collaborators: songData.Collaborators,
		SongName:      songData.SongName,
		Genre:         songData.Genre,
		Exclusive:     strconv.FormatBool(songData.Exclusive),
	}
	resp, err := musicService.AddSong(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add song. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add song. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}

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

func UpdateSong(newSongData *music_model.MusicUpdateModel) (*music_service_proto.GetResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details: %v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.UpdateReq{
		AuthorId:      newSongData.AuthorId,
		SongId:        newSongData.SongId,
		SongName:      newSongData.SongName,
		Collaborators: newSongData.Collaborators,
	}
	resp, err := musicService.UpdateSong(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update song. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update song. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}

func DeleteSong(deleteReq *music_model.MusicDeleteModel) (*music_service_proto.DeleteResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details: %v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.DeleteReq{
		AuthorId: deleteReq.UserId,
		SongId:   deleteReq.SongId,
	}
	resp, err := musicService.DeleteSong(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete song. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete song. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}

func GetByAuthor(authorId string) (*music_service_proto.GetSongsResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details: %v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.GetByAuthorReq{AuthorId: authorId}
	resp, err := musicService.GetByAuthor(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get songs by author. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get songs by author. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}

func GetByCollaborator(collaboratorId string) (*music_service_proto.GetSongsResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details: %v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.GetByCollaboratorsReq{CollaboratorId: collaboratorId}
	resp, err := musicService.GetByCollaborators(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get songs by collaborator. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get songs by collaborator. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}

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

func LikeSong(userId string, songId string) (*music_service_proto.GetResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service"))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.LikeSongReq{
		UserId: userId,
		SongId: songId,
	}
	resp, err := musicService.LikeSong(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't like song"))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return resp, nil
}

func DislikeSong(userId string, songId string) (*music_service_proto.GetResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service"))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.DislikeSongReq{
		UserId: userId,
		SongId: songId,
	}
	resp, err := musicService.DislikeSong(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't dislike song"))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return resp, nil
}

func GetUserLikedSongs(userId string) (*music_service_proto.GetSongsResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service"))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.GetLikedSongsReq{
		UserId: userId,
	}
	resp, err := musicService.GetUserLiked(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get liked songs of user %s", userId))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return resp, nil
}

func GetUserDislikedSongs(userId string) (*music_service_proto.GetSongsResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service"))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.GetDislikedSongsReq{
		UserId: userId,
	}
	resp, err := musicService.GetUserDisliked(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get disliked songs of user %s", userId))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return resp, nil
}

func AddUserToPlayedQuantity(userId string, songId string) (*music_service_proto.GetResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details %v\n", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.IncreasePlayedQuantityReq{
		UserId: userId,
		SongId: songId,
	}
	song, err := musicService.IncreasePlayedQuantity(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add user to song played quantity. Details: %v\n", err))
		return nil, err
	}
	return song, nil
}

func GetTrends() (*music_service_proto.GetSongsResp, error) {
	conn, err := grpc.Dial(*service_address_config.MusicServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details: %v", err))
		return nil, err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.GetTrendsReq{}
	resp, err := musicService.GetTrends(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't trends. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get trends. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}
