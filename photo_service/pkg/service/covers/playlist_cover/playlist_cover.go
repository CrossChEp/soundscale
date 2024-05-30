package playlist_cover

import (
	"encoding/base64"
	"errors"
	"fmt"
	"path/filepath"
	"photo_service/pkg/config/constants"
	"photo_service/pkg/config/path_config"
	"photo_service/pkg/logger"
	"photo_service/pkg/proto/photo_service_proto"
	"photo_service/pkg/repository"
	"photo_service/pkg/service/grpc_funcs"
)

func UploadCover(r *photo_service_proto.UploadPlaylistCoverReq) error {
	playlist, err := grpc_funcs.GetPlaylist(r.PlaylistId, constants.PlaylistType)
	if err != nil {
		return err
	}
	if r.UserId != playlist.AuthorId {
		logger.ErrorLog(fmt.Sprintf("Error: user is not an author of playlist"))
		return errors.New("user is not an author of playlist")
	}
	if err := upload(r); err != nil {
		return err
	}
	return nil
}

func upload(r *photo_service_proto.UploadPlaylistCoverReq) error {
	fileBytes, err := base64.StdEncoding.DecodeString(r.File)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode base64 file. Details: %v", err))
		return err
	}
	path, err := filepath.Abs(fmt.Sprintf("%s/%s.png", path_config.PlaylistCoversPath, r.PlaylistId))
	if err := repository.Upload(path, fileBytes); err != nil {
		return err
	}
	return nil
}

func DownloadCover(r *photo_service_proto.DownloadPlaylistCoverReq) (string, error) {
	_, err := grpc_funcs.GetPlaylist(r.PlaylistId, constants.PlaylistType)
	if err != nil {
		return "", err
	}
	filePath, err := filepath.Abs(fmt.Sprintf("%s/%s.png", path_config.PlaylistCoversPath, r.PlaylistId))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't create file path. Details: %v", err))
		return "", err
	}
	coverBytes, err := repository.Download(filePath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(coverBytes), nil
}
