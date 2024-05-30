package grpc_funcs

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"player_service/config"
	"player_service/funcs/logger"
	"player_service/proto/collection_service_proto"
	"player_service/proto/music_service_proto"
)

func GetSong(songId string) (*music_service_proto.GetResp, error) {
	conn, err := grpc.Dial(*config.MusicServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details %v\n", err))
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

func AddUserToPlayedQuantity(userId string, songId string) error {
	conn, err := grpc.Dial(*config.MusicServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details %v\n", err))
		return err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.IncreasePlayedQuantityReq{
		UserId: userId,
		SongId: songId,
	}
	_, err = musicService.IncreasePlayedQuantity(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add user to song played quantity. Details: %v\n", err))
		return err
	}
	return nil
}

func AddUserToListenedQuantity(userId string, songId string) error {
	conn, err := grpc.Dial(*config.MusicServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to music service. Details %v\n", err))
		return err
	}
	musicService := music_service_proto.NewMusicServiceClient(conn)
	req := &music_service_proto.IncreaseListenedQuantityReq{
		UserId: userId,
		SongId: songId,
	}
	_, err = musicService.IncreaseListenedQuantity(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add user to song listened quantity. Details: %v\n", err))
		return err
	}
	return nil
}

func AddSongGenreToUser(userId string, genre string) error {
	conn, err := grpc.Dial(*config.CollectionServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to collection service. Details %v\n", err))
		return err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	req := &collection_service_proto.AddGenresReq{
		UserId: userId,
		Genres: []string{genre},
	}
	_, err = collectionService.AddGenres(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add genres to user %s collection. Details: %v\n", userId, err))
		return err
	}
	return nil
}
