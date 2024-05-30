package grpc_collection

import (
	"context"
	"fmt"
	"gateway/pkg/config/service_address_config"
	"gateway/pkg/model/collection_model"
	"gateway/pkg/proto/collection_service_proto"
	"gateway/pkg/service/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitCollection(userId string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	initReq := &collection_service_proto.InitReq{UserId: userId}
	response, err := collectionService.InitCollection(context.TODO(), initReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't initialize user %s grpc_collection. Details: %d", userId, err))
		return nil, err
	}
	return response, nil
}

func GetCollection(userId string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	getReq := &collection_service_proto.GetCollectionReq{UserId: userId}
	response, err := collectionService.GetCollection(context.TODO(), getReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user %s grpc_collection. Details: %d", userId, err))
		return nil, err
	}
	return response, nil
}

func AddSongs(userId string, songs []string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	addSongsReq := &collection_service_proto.SongsReq{
		UserId:  userId,
		SongIds: songs,
	}
	response, err := collectionService.AddSongs(context.TODO(), addSongsReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add songs to user %s grpc_collection. Details: %d", userId, err))
		return nil, err
	}
	return response, nil
}

func RemoveSongs(userId string, songs []string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	addSongsReq := &collection_service_proto.SongsReq{
		UserId:  userId,
		SongIds: songs,
	}
	response, err := collectionService.RemoveSongs(context.TODO(), addSongsReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't remove songs to user %s grpc_collection. Details: %d", userId, err))
		return nil, err
	}
	return response, nil
}

func AddPlaylists(userId string, playlists []string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	playlistsReq := &collection_service_proto.PlaylistsReq{
		UserId:      userId,
		PlaylistIds: playlists,
	}
	response, err := collectionService.AddPlaylists(context.TODO(), playlistsReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add playlists to user %s grpc_collection. Details: %d", userId, err))
		return nil, err
	}
	return response, nil
}

func RemovePlaylists(userId string, playlists []string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	playlistsReq := &collection_service_proto.PlaylistsReq{
		UserId:      userId,
		PlaylistIds: playlists,
	}
	response, err := collectionService.RemovePlaylists(context.TODO(), playlistsReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add playlists to user %s grpc_collection. Details: %d", userId, err))
		return nil, err
	}
	return response, nil
}

func AddAlbums(userId string, albums []string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	albumsReq := &collection_service_proto.AlbumsReq{
		UserId:   userId,
		AlbumIds: albums,
	}
	response, err := collectionService.AddAlbums(context.TODO(), albumsReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add albums to user %s grpc_collection. Details: %d", userId, err))
		return nil, err
	}
	return response, nil
}

func RemoveAlbums(userId string, albums []string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	albumsReq := &collection_service_proto.AlbumsReq{
		UserId:   userId,
		AlbumIds: albums,
	}
	response, err := collectionService.RemoveAlbums(context.TODO(), albumsReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't remove albums to user %s grpc_collection. Details: %d", userId, err))
		return nil, err
	}
	return response, nil
}

func AddGenres(genresModel *collection_model.GenresModel) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	genresReq := &collection_service_proto.AddGenresReq{
		UserId: genresModel.UserId,
		Genres: genresModel.Genres,
	}
	response, err := collectionService.AddGenres(context.TODO(), genresReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add genres to user %s collection. Details: %d", genresModel.UserId, err))
		return nil, err
	}
	return response, nil
}

func Follow(userId string, musicianId string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	req := &collection_service_proto.FollowReq{
		UserId:     userId,
		MusicianId: musicianId,
	}
	response, err := collectionService.Follow(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("couldn't follow by user %s on musician %s. Details: %v", userId, musicianId, err))
		return nil, err
	}
	return response, nil
}

func Unfollow(userId string, musicianId string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	req := &collection_service_proto.UnfollowReq{
		UserId:     userId,
		MusicianId: musicianId,
	}
	response, err := collectionService.Unfollow(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("couldn't unfollow by user %s on musician %s. Details: %v", userId, musicianId, err))
		return nil, err
	}
	return response, nil
}

func Subscribe(userId string, musicianId string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	req := &collection_service_proto.SubscribeReq{
		UserId:     userId,
		MusicianId: musicianId,
	}
	response, err := collectionService.Subscribe(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("couldn't subscribe by user %s on musician %s. Details: %v", userId, musicianId, err))
		return nil, err
	}
	return response, nil
}

func Unsubscribe(userId string, musicianId string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	req := &collection_service_proto.UnsubscribeReq{
		UserId:     userId,
		MusicianId: musicianId,
	}
	response, err := collectionService.Unsubscribe(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("couldn't unsubscribe by user %s on musician %s. Details: %v", userId, musicianId, err))
		return nil, err
	}
	return response, nil
}
