package grpc_service

import (
	"context"
	"errors"
	"fmt"
	"gateway/pkg/config/service_address_config"
	"gateway/pkg/model/photo_model"
	"gateway/pkg/proto/photo_service_proto"
	"gateway/pkg/service/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UploadPFP(photoData photo_model.UploadPFPModel) (*photo_service_proto.Response, error) {
	conn, err := grpc.Dial(*service_address_config.PhotoServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to photo service. Details: %v", err))
		return nil, err
	}
	photoService := photo_service_proto.NewPhotoServiceClient(conn)
	req := &photo_service_proto.PFPRequest{
		UserId: photoData.UserId,
		Photo:  photoData.Photo,
	}
	resp, err := photoService.UploadPFP(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't upload profile picture. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't upload profile picture. Details: %v", resp.Error))
		return nil, err
	}
	return resp, nil
}

func DownloadPFP(userId string) (*photo_service_proto.Response, error) {
	conn, err := grpc.Dial(*service_address_config.PhotoServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to photo service. Details: %v", err))
		return nil, err
	}
	photoService := photo_service_proto.NewPhotoServiceClient(conn)
	req := &photo_service_proto.GetPFPReq{UserId: userId}
	resp, err := photoService.DownloadPFP(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't download profile picture. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't download profile picture. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}

func UploadSongCover(songModel photo_model.UploadSongCoverModel) (*photo_service_proto.Response, error) {
	conn, err := grpc.Dial(*service_address_config.PhotoServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to photo service. Details: %v", err))
		return nil, err
	}
	photoService := photo_service_proto.NewPhotoServiceClient(conn)
	req := &photo_service_proto.UploadCoverReq{
		UserId: songModel.UserId,
		SongId: songModel.SongId,
		File:   songModel.File,
	}
	resp, err := photoService.UploadCover(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't upload song cover. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't upload song cover. Details: %v", err))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}

func DownloadSongCover(songId string) (*photo_service_proto.Response, error) {
	conn, err := grpc.Dial(*service_address_config.PhotoServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to photo service. Details: %v", err))
		return nil, err
	}
	photoService := photo_service_proto.NewPhotoServiceClient(conn)
	req := &photo_service_proto.DownloadSongCoverReq{SongId: songId}
	resp, err := photoService.DownloadSongCover(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't download song cover. Details: %v", err))
		return nil, err
	}
	if resp.Error != "" {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't download song cover. Details: %v", resp.Error))
		return nil, errors.New(resp.Error)
	}
	return resp, nil
}

func UploadPlaylistCover(playlistModel *photo_model.UploadPlaylistCoverModel) (*photo_service_proto.Response, error) {
	conn, err := grpc.Dial(*service_address_config.PhotoServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to photo service. Details: %v", err))
		return nil, err
	}
	photoService := photo_service_proto.NewPhotoServiceClient(conn)
	req := &photo_service_proto.UploadPlaylistCoverReq{
		UserId:     playlistModel.UserId,
		PlaylistId: playlistModel.Id,
		File:       playlistModel.File,
	}
	resp, err := photoService.UploadPlaylistCover(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't upload playlist cover. Details: %v", err))
		return nil, err
	}
	return resp, nil
}

func DownloadPlaylistCover(playlistId string) (*photo_service_proto.Response, error) {
	conn, err := grpc.Dial(*service_address_config.PhotoServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to photo service. Details: %v", err))
		return nil, err
	}
	photoService := photo_service_proto.NewPhotoServiceClient(conn)
	req := &photo_service_proto.DownloadPlaylistCoverReq{PlaylistId: playlistId}
	resp, err := photoService.DownloadPlaylistCover(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't download playlist cover. Details: %v", err))
		return nil, err
	}
	return resp, nil
}

func DownloadAlbumCover(playlistId string) (*photo_service_proto.Response, error) {
	conn, err := grpc.Dial(*service_address_config.PhotoServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to photo service. Details: %v", err))
		return nil, err
	}
	photoService := photo_service_proto.NewPhotoServiceClient(conn)
	req := &photo_service_proto.DownloadAlbumCoverReq{AlbumId: playlistId}
	resp, err := photoService.DownloadAlbumCover(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't download album cover. Details: %v", err))
		return nil, err
	}
	return resp, nil
}

func UploadAlbumCover(playlistModel *photo_model.UploadPlaylistCoverModel) (*photo_service_proto.Response, error) {
	conn, err := grpc.Dial(*service_address_config.PhotoServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to photo service. Details: %v", err))
		return nil, err
	}
	photoService := photo_service_proto.NewPhotoServiceClient(conn)
	req := &photo_service_proto.UploadAlbumCoverReq{
		UserId:  playlistModel.UserId,
		AlbumId: playlistModel.Id,
		File:    playlistModel.File,
	}
	resp, err := photoService.UploadAlbumCover(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't upload album cover. Details: %v", err))
		return nil, err
	}
	return resp, nil
}
