package post

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"photo_service/pkg/config/path_config"
	"photo_service/pkg/logger"
	"photo_service/pkg/model"
	"photo_service/pkg/proto/photo_service_proto"
	"photo_service/pkg/repository"
	"photo_service/pkg/service/grpc_funcs"
)

func UploadPostPhotos(r *photo_service_proto.UploadPostPhotosReq) ([]string, error) {
	post, err := grpc_funcs.GetPost(r.PostId)
	if err != nil {
		return nil, err
	}
	if post.AuthorId != r.UserId {
		logger.ErrorLog(fmt.Sprintf("Error: user %s is not author of post. Author: %s", r.UserId, post.AuthorId))
		return nil, errors.New(fmt.Sprintf("Error: user %s is not author of post. Author: %s", r.UserId, post.AuthorId))
	}
	return uploadPhotos(r.PostId, r.Photos)
}

func uploadPhotos(postId string, photos []string) ([]string, error) {
	path, err := filepath.Abs(fmt.Sprintf("%s/%s", path_config.PostPath, postId))
	if err := os.Mkdir(path, 0750); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't create directory. Details: %v", err))
		return nil, err
	}
	var added []string
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't create file path of %s post photo. Details: %v", postId, err))
		return nil, err
	}
	for index, photo := range photos {
		err := upload(model.UploadPhotoModel{
			Id:    postId,
			Photo: photo,
			Path:  path,
			Pos:   index,
		})
		if err != nil {
			logger.ErrorLog(fmt.Sprintf("Error: couldn't upload photo"))
			continue
		}
		added = append(added, photo)
	}
	return added, nil
}

func upload(uploadData model.UploadPhotoModel) error {
	file, err := base64.StdEncoding.DecodeString(uploadData.Photo)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode file. Details: %v", err))
		return err
	}
	uploadData.Path = fmt.Sprintf("%s/%d.png", uploadData.Path, uploadData.Pos)
	if err := repository.Upload(uploadData.Path, file); err != nil {
		return err
	}
	return nil
}

func DownloadPhotos(postId string) ([]string, error) {
	filePath, err := filepath.Abs(fmt.Sprintf("%s/%s", path_config.PostPath, postId))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't create file path of %s post", postId))
		return nil, errors.New(fmt.Sprintf("Error: couldn't crete file path of %s post", postId))
	}
	photos, err := os.ReadDir(filePath)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get files of %s post", postId))
		return nil, errors.New(fmt.Sprintf("Error: couldn't get files of %s post", postId))
	}
	var encodedPhotos []string
	for i := 0; i < len(photos); i++ {
		photo, err := download(i, filePath)
		if err != nil {
			logger.ErrorLog(fmt.Sprintf("Error: couldn't download photo %d of post %s", i, postId))
			continue
		}
		encodedPhotos = append(encodedPhotos, photo)
	}
	return encodedPhotos, nil
}

func download(photoId int, postPath string) (string, error) {
	filePath, err := filepath.Abs(fmt.Sprintf("%s/%d.png", postPath, photoId))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't download photo %d.png of post %s. Details: %v", photoId, postPath, err))
		return "", err
	}
	file, err := repository.Download(filePath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(file), nil
}
