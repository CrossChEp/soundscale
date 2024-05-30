package photo_service

import (
	"gateway/pkg/model/photo_model"
	grpc_service "gateway/pkg/service/grpc_service/photo"
)

func UploadPFP(photoData photo_model.UploadPFPModel) error {
	if _, err := grpc_service.UploadPFP(photoData); err != nil {
		return err
	}
	return nil
}

func DownloadPFP(userId string) (string, error) {
	resp, err := grpc_service.DownloadPFP(userId)
	if err != nil {
		return "", err
	}
	return resp.Content, nil
}

func UploadSongCover(coverModel photo_model.UploadSongCoverModel) error {
	if _, err := grpc_service.UploadSongCover(coverModel); err != nil {
		return err
	}
	return nil
}

func DownloadSongCover(songId string) (string, error) {
	songCover, err := grpc_service.DownloadSongCover(songId)
	if err != nil {
		return "", err
	}
	return songCover.Content, nil
}

func UploadPlaylistCover(coverModel *photo_model.UploadPlaylistCoverModel) error {
	if _, err := grpc_service.UploadPlaylistCover(coverModel); err != nil {
		return err
	}
	return nil
}

func DownloadPlaylistCover(playlistId string) (string, error) {
	songCover, err := grpc_service.DownloadPlaylistCover(playlistId)
	if err != nil {
		return "", err
	}
	return songCover.Content, nil
}

func DownloadAlbumCover(albumId string) (string, error) {
	songCover, err := grpc_service.DownloadAlbumCover(albumId)
	if err != nil {
		return "", err
	}
	return songCover.Content, nil
}

func UploadAlbumCover(coverModel *photo_model.UploadPlaylistCoverModel) error {
	if _, err := grpc_service.UploadAlbumCover(coverModel); err != nil {
		return err
	}
	return nil
}
