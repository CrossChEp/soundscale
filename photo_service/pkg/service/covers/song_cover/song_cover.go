package song_cover

import (
	"encoding/base64"
	"errors"
	"fmt"
	"path/filepath"
	"photo_service/pkg/config/path_config"
	"photo_service/pkg/logger"
	"photo_service/pkg/proto/music_service_proto"
	"photo_service/pkg/proto/photo_service_proto"
	"photo_service/pkg/repository"
	"photo_service/pkg/service/grpc_funcs"
)

func UploadSongCover(r *photo_service_proto.UploadCoverReq) error {
	song, err := grpc_funcs.GetSong(r.SongId)
	if err != nil {
		return err
	}
	if song == nil {
		return errors.New("couldn't get song with such id")
	}
	if song.AuthorId != r.UserId {
		logger.ErrorLog("Error: user is not an author of this songs!")
		return errors.New("user is not an author of this songs")
	}
	if err := upload(r, song); err != nil {
		return err
	}
	return nil
}

func upload(r *photo_service_proto.UploadCoverReq, song *music_service_proto.GetResp) error {
	fileBytes, err := base64.StdEncoding.DecodeString(r.File)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode base64 string. Details: %v", err))
		return err
	}
	filePath, err := filepath.Abs(fmt.Sprintf("%s%s.png", path_config.SongCoversPath, song.Id))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't create path. Details: %v", err))
		return err
	}
	if err := repository.Upload(filePath, fileBytes); err != nil {
		return err
	}

	return nil
}

func DownloadCover(r *photo_service_proto.DownloadSongCoverReq) (string, error) {
	_, err := grpc_funcs.GetSong(r.SongId)
	if err != nil {
		return "", err
	}
	filePath, err := filepath.Abs(fmt.Sprintf("%s%s.png", path_config.SongCoversPath, r.SongId))
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
