package album_cover

import (
	"encoding/base64"
	"errors"
	"fmt"
	"path/filepath"
	"photo_service/pkg/config/constants"
	"photo_service/pkg/config/path_config"
	"photo_service/pkg/logger"
	"photo_service/pkg/proto/photo_service_proto"
	"photo_service/pkg/proto/playlist_service_proto"
	"photo_service/pkg/repository"
	"photo_service/pkg/service/grpc_funcs"
)

func Upload(r *photo_service_proto.UploadAlbumCoverReq) error {
	album, err := grpc_funcs.GetPlaylist(r.AlbumId, constants.AlbumType)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get album with id %s. Details: %v", r.AlbumId, err))
		return err
	}
	if err := checkAlbumData(album, r.UserId); err != nil {
		return err
	}
	if err := upload(album, r.File); err != nil {
		return err
	}
	return nil
}

func checkAlbumData(album *playlist_service_proto.GetPlaylistResp, userId string) error {
	if album.Type != constants.AlbumType {
		logger.ErrorLog(fmt.Sprintf("Error playlist with id %s is not an album", album.Id))
		return errors.New(fmt.Sprintf("error playlist with id %s is not an album", album.Id))
	}
	if album.AuthorId != userId {
		logger.ErrorLog(fmt.Sprintf("Error: user with id %s is not an author of album with id %s", userId, album.Id))
		return errors.New(fmt.Sprintf("Error: user with id %s is not an author of album with id %s", userId, album.Id))
	}
	return nil
}

func upload(album *playlist_service_proto.GetPlaylistResp, file string) error {
	fileBytes, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode request base64 string to file. Details: %v", err))
		return err
	}
	filePath, err := filepath.Abs(fmt.Sprintf("%s/%s.png", path_config.AlbumCoverPath, album.Id))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't create file path of album with id %s. Details: %v", album.Id, err))
		return err
	}
	if err := repository.Upload(filePath, fileBytes); err != nil {
		return err
	}
	return nil
}

func Download(r *photo_service_proto.DownloadAlbumCoverReq) (string, error) {
	_, err := grpc_funcs.GetPlaylist(r.AlbumId, constants.AlbumType)
	if err != nil {
		return "", err
	}
	filePath, err := filepath.Abs(fmt.Sprintf("%s/%s.png", path_config.AlbumCoverPath, r.AlbumId))
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
