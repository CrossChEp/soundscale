// Package transport Пакет, который содержит хэндлеры, связанные с управлением плейлистами
package transport

import (
	"context"
	"errors"
	"fmt"
	"playlist_serivce/pkg/proto/playlist_service_proto"
	"playlist_serivce/pkg/service/playlist_service"
	"playlist_serivce/pkg/service/response"
)

// PlaylistService Структура, содержащая методы для работы с плейлистами
type PlaylistService struct {
	playlist_service_proto.PlaylistServiceServer
}

// AddCurrentPlaylist AddCurrentPlaylist Хендлер, который добавляет песни в текущий проигрываемый плейлист
//
// # ---Принимает аргумент типа AddPlaylistReq:
//
//	type AddPlaylistReq struct {
//		Token string		// jwt токен пользователя
//		Songs []string 		// Массив, который содержит id песен
//	}
//
// # ---Возвращает объект типа AddPlaylistResp:
//
//	type AddPlaylistResp struct {
//		Message string	// Поле, которое содержит сообщение при успешном выполнении хендлера
//		Error   string	// Поле, которое содержит сообщение об ошибке, если хендлер выполнился провально
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *PlaylistService) AddCurrentPlaylist(_ context.Context, r *playlist_service_proto.AddCurrentPlaylistReq) (*playlist_service_proto.AddCurrentPlaylistResp, error) {
	err := playlist_service.AddCurrentPlaylist(r)
	if err != nil {
		return &playlist_service_proto.AddCurrentPlaylistResp{Error: "err"}, nil
	}
	return &playlist_service_proto.AddCurrentPlaylistResp{Message: "msg"}, nil
}

func (s *PlaylistService) AddPlaylist(_ context.Context, r *playlist_service_proto.AddPlaylistReq) (*playlist_service_proto.GetPlaylistResp, error) {
	playlist, err := playlist_service.AddPlaylist(r)
	if err != nil {
		return nil, err
	}
	return response.CreateGetPlaylistResponse(playlist), nil
}

func (s *PlaylistService) GetPlaylist(_ context.Context, r *playlist_service_proto.GetPlaylistReq) (*playlist_service_proto.GetPlaylistResp, error) {
	playlist, err := playlist_service.GetPlaylist(r)
	if err != nil {
		return nil, err
	}
	return response.CreateGetPlaylistResponse(playlist), nil
}

func (s *PlaylistService) GetPlaylistsByAuthor(_ context.Context,
	r *playlist_service_proto.GetPlaylistsByAuthorReq) (*playlist_service_proto.GetPlaylistsResp, error) {
	playlists, err := playlist_service.GetPlaylistsByAuthor(r)
	if err != nil {
		return nil, err
	}
	return response.CreateGetPlaylistsResponse(playlists), nil
}

func (s *PlaylistService) AddSongToPlaylist(_ context.Context,
	r *playlist_service_proto.AddSongsToPlaylistReq) (*playlist_service_proto.GetPlaylistResp, error) {
	playlist, err := playlist_service.AddSongsToPlaylist(r)
	if err != nil {
		return nil, err
	}
	return response.CreateGetPlaylistResponse(playlist), nil
}

func (s *PlaylistService) DeleteSongFromPlaylist(_ context.Context,
	r *playlist_service_proto.DeleteSongsFromPlaylistReq) (*playlist_service_proto.GetPlaylistResp, error) {
	playlist, err := playlist_service.DeleteSongsFromPlaylist(r)
	if err != nil {
		return nil, err
	}
	return response.CreateGetPlaylistResponse(playlist), nil
}

func (s *PlaylistService) DeletePlaylist(_ context.Context,
	r *playlist_service_proto.DeletePlaylistReq) (*playlist_service_proto.MessageResp, error) {
	if err := playlist_service.DeletePlaylist(r); err != nil {
		return nil, err
	}
	return &playlist_service_proto.MessageResp{Content: "playlist was deleted"}, nil
}

func (s *PlaylistService) GetFavouriteSongs(_ context.Context,
	r *playlist_service_proto.GetUserFavouriteReq) (*playlist_service_proto.GetPlaylistResp, error) {
	favourites, err := playlist_service.GetFavouriteSongs(r.UserId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't get favourites of user with id %s. Details: %v", r.UserId, err))
	}
	return response.CreateGetPlaylistResponse(favourites), nil
}

func (s *PlaylistService) LikePlaylist(_ context.Context,
	r *playlist_service_proto.LikePlaylistReq) (*playlist_service_proto.GetPlaylistResp, error) {
	playlist, err := playlist_service.LikePlaylist(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't like playlist %s. Details: %v", r.PlaylistId, err))
	}
	return response.CreateGetPlaylistResponse(playlist), nil
}

func (s *PlaylistService) DislikePlaylist(_ context.Context,
	r *playlist_service_proto.DislikePlaylistReq) (*playlist_service_proto.GetPlaylistResp, error) {
	playlist, err := playlist_service.DislikePlaylist(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't dislike playlist %s. Details: %v", r.PlaylistId, err))
	}
	return response.CreateGetPlaylistResponse(playlist), nil
}

func (s *PlaylistService) GetUserLiked(_ context.Context,
	r *playlist_service_proto.GetUserLikedReq) (*playlist_service_proto.GetPlaylistsResp, error) {
	playlists, err := playlist_service.GetLikedPlaylists(r.UserId, r.PlaylistType)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't get user %s liked playlists. Details: %v", r.UserId, err))
	}
	return response.CreateGetPlaylistsResponse(playlists), nil
}

func (s *PlaylistService) GetUserDisliked(_ context.Context,
	r *playlist_service_proto.GetUserDislikedReq) (*playlist_service_proto.GetPlaylistsResp, error) {
	playlists, err := playlist_service.GetDislikedPlaylists(r.UserId, r.PlaylistType)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't get user %s disliked playlists. Details: %v", r.UserId, err))
	}
	return response.CreateGetPlaylistsResponse(playlists), nil
}
