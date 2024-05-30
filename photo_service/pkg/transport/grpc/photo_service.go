// Package transport Пакет, который содержит хэндлеры, связанные с фотографиями
package transport

import (
	"context"
	"errors"
	"fmt"
	"photo_service/pkg/logger"
	"photo_service/pkg/proto/photo_service_proto"
	"photo_service/pkg/service/covers/album_cover"
	"photo_service/pkg/service/covers/playlist_cover"
	"photo_service/pkg/service/covers/post"
	"photo_service/pkg/service/covers/song_cover"
	"photo_service/pkg/service/pfp"
)

// PhotoService Структура, содержащая методы для работы с фотографиями
type PhotoService struct {
	photo_service_proto.PhotoServiceServer
}

// UploadPFP Хендлер, который загружает аватарку пользователю
//
// # Принимает аргумент типа PFPRequest:
//
//	type PFPRequest struct {
//		Token string	// jwt токен пользователя
//		Photo []byte 	// Массив, который содержит байты фотографии
//	}
//
// # Возвращает объект типа Response:
//
//	type Response struct {
//		Content string 	// Поле, которое содержит ответ сервера(обычно строка)
//		Error   string 	// При ошибке поле содержит строку о деталях ошибки
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *PhotoService) UploadPFP(_ context.Context, r *photo_service_proto.PFPRequest) (*photo_service_proto.Response, error) {
	if err := pfp.Upload(r); err != nil {
		return &photo_service_proto.Response{Error: "couldn't upload pfp"}, nil
	}
	return &photo_service_proto.Response{Content: "pfp was uploaded"}, nil
}

// DownloadPFP Хендлер, который скачивает аватарку пользователя
//
// # Принимает аргумент типа PFPRequest:
//
//	type GetPFPReq struct {
//		Token string  // jwt токен пользователя
//	}
//
// # Возвращает объект типа GetPhotoResp:
//
//	type GetPhotoResp struct {
//		File  []byte 	// Поле массив байтов фотографии пользователя
//		Error   string 	// При ошибке поле содержит строку о деталях ошибки
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *PhotoService) DownloadPFP(_ context.Context, r *photo_service_proto.GetPFPReq) (*photo_service_proto.Response, error) {
	photo, err := pfp.Download(r)
	if err != nil {
		return &photo_service_proto.Response{Error: "couldn't download file"}, nil
	}
	return &photo_service_proto.Response{Content: photo}, nil
}

// UploadCover Хендлер, который загружает обложку для песни
//
// # Принимает аргумент типа UploadCoverReq:
//
//	type UploadCoverReq struct {
//		Token string 	// jwt токен пользователя
//		SongId string 	// Id песни, для которой загружается обложка
//		File   string 	// Массив байтов обложки, который закодирован в base64
//	}
//
// # Возвращает объект типа Response:
//
//	type Response struct {
//		Content string 	// Поле содержит ответ сервера(обычно строка)
//		Error   string 	// При ошибке поле содержит строку о деталях ошибки
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *PhotoService) UploadCover(_ context.Context, r *photo_service_proto.UploadCoverReq) (*photo_service_proto.Response, error) {
	if err := song_cover.UploadSongCover(r); err != nil {
		return &photo_service_proto.Response{Error: fmt.Sprintf("Couldn't upload songs cover. Details: %v", err)}, nil
	}
	return &photo_service_proto.Response{Content: "Music cover was uploaded"}, nil
}

// DownloadSongCover Хендлер, который скачивает обложку песни
//
// # Принимает аргумент типа DownloadSongCoverReq:
//
//	type DownloadSongCoverReq struct {
//		SongId string  // id песни, обложка которой будет скачана
//	}
//
// # Возвращает объект типа Response:
//
//	type Response struct {
//		Content string 	// Поле содержит ответ сервера(обычно строка)
//		Error   string 	// При ошибке поле содержит строку о деталях ошибки
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *PhotoService) DownloadSongCover(_ context.Context, r *photo_service_proto.DownloadSongCoverReq) (*photo_service_proto.Response, error) {
	songStr, err := song_cover.DownloadCover(r)
	if err != nil {
		return &photo_service_proto.Response{Error: fmt.Sprintf("Couldn't get cover. Details: %v", err)}, nil
	}
	return &photo_service_proto.Response{Content: songStr}, nil
}

func (s *PhotoService) UploadPlaylistCover(_ context.Context,
	r *photo_service_proto.UploadPlaylistCoverReq) (*photo_service_proto.Response, error) {
	if err := playlist_cover.UploadCover(r); err != nil {
		return nil, errors.New(fmt.Sprintf("error: couldn't upload cover. Details: %v", err))
	}
	return &photo_service_proto.Response{Content: "Cover was uploaded"}, nil
}

func (s *PhotoService) DownloadPlaylistCover(_ context.Context,
	r *photo_service_proto.DownloadPlaylistCoverReq) (*photo_service_proto.Response, error) {
	cover, err := playlist_cover.DownloadCover(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: couldn't download cover. Details: %v", err))
	}
	return &photo_service_proto.Response{Content: cover}, nil
}

func (s *PhotoService) UploadAlbumCover(_ context.Context,
	r *photo_service_proto.UploadAlbumCoverReq) (*photo_service_proto.Response, error) {
	if err := album_cover.Upload(r); err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't upload cover for album with id %s. Detaisl: %v", r.AlbumId, err))
	}
	return &photo_service_proto.Response{
		Content: fmt.Sprintf("Cover for album %s was successfully updated", r.AlbumId),
	}, nil
}

func (s *PhotoService) DownloadAlbumCover(_ context.Context,
	r *photo_service_proto.DownloadAlbumCoverReq) (*photo_service_proto.Response, error) {
	cover, err := album_cover.Download(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: couldn't download cover. Details: %v", err))
	}
	return &photo_service_proto.Response{Content: cover}, nil
}

func (s *PhotoService) UploadPostPhotos(_ context.Context,
	r *photo_service_proto.UploadPostPhotosReq) (*photo_service_proto.Response, error) {
	logger.InfoLog(fmt.Sprintf("UploadPostPhotos function was called"))
	_, err := post.UploadPostPhotos(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't upload photos. Details: %v", err))
	}
	return &photo_service_proto.Response{Content: "photos were uploaded"}, nil
}

func (s *PhotoService) DownloadPostPhotos(_ context.Context,
	r *photo_service_proto.DownloadPostPhotosReq) (*photo_service_proto.PhotosResp, error) {
	photos, err := post.DownloadPhotos(r.PostId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't download photos of post %s. Details: %v", r.PostId, err))
	}
	return &photo_service_proto.PhotosResp{Photos: photos}, nil
}
