// Package transport Пакет, который содержит хэндлеры, связанные crud песен
package transport

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"music_service/pkg/proto/music_service_proto"
	"music_service/pkg/repository"
	"music_service/pkg/service/music"
	"music_service/pkg/service/response"
)

// MusicService Структура, содержащая методы для работы с песнями
type MusicService struct {
	music_service_proto.MusicServiceServer
}

// AddSong Хендлер, который добавляет данные песни в БД
//
// # ---Принимает аргумент типа AddReq:
//
//	type AddReq struct {
//		Token string		// jwt токен пользователя
//		Collaborators []string 	// Массив, который содержит id колабораторов
//		SongName      string	// Название песни
//		Genre         string
//	}
//
// # ---Возвращает объект типа GetResp:
//
//	type GetResp struct {
//		Id            string	//  	Id песни
//		SongName      string	// 	Название песни
//		AuthorId      string	//	Id автора
//		Collaborators []string	// 	Массив, который содержит id коллабораторов
//		Genre         string	//	Жанр песни
//		ReleaseDate   *timestamppb.Timestamp	//	Дата релиза
//		Error         string	// Поле, которое содержит информацию об ошибке, если она произошла
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *MusicService) AddSong(_ context.Context, r *music_service_proto.AddReq) (*music_service_proto.GetResp, error) {
	fmt.Println(r.Exclusive)
	song, err := music.Add(r)
	if err != nil {
		return response.CreateErrorResp("couldn't add song"), nil
	}
	return response.CreateGetResp(song), nil
}

// GetSong Хендлер, который получает данные песни, используя id песни
//
// # ---Принимает аргумент типа GetReq:
//
//	type GetReq struct {
//		Id string	// Id песни
//	}
//
// # ---Возвращает объект типа GetResp:
//
//	type GetResp struct {
//		Id            string	//  	Id песни
//		SongName      string	// 	Название песни
//		AuthorId      string	//	Id автора
//		Collaborators []string	// 	Массив, который содержит id коллабораторов
//		Genre         string	//	Жанр песни
//		ReleaseDate   *timestamppb.Timestamp	//	Дата релиза
//		Error         string	// Поле, которое содержит информацию об ошибке, если она произошла
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *MusicService) GetSong(_ context.Context, r *music_service_proto.GetReq) (*music_service_proto.GetResp, error) {
	oid, err := primitive.ObjectIDFromHex(r.Id)
	if err != nil {
		return response.CreateErrorResp("id is invalid"), nil
	}
	song, err := repository.GetByObjectId(oid)
	if err != nil {
		return response.CreateErrorResp("couldn't get user"), nil
	}
	return response.CreateGetResp(song), nil
}

// GetByAuthor Хендлер, который получает данные всех песен, в которых присутствовал автор
//
// # ---Принимает аргумент типа GetByAuthorReq:
//
//	type GetByAuthorReq struct {
//		AuthorId string	// Id автора
//	}
//
// # ---Возвращает объект типа GetSongsResp:
//
//	type GetSongsResp struct {
//		Songs 	[]*GetResp	// 	Поле, которое содержит массив из GetResp. Это все песни, в которых учавствовал автор
//		Error 	string      	//	Поле, которое содержит детали ошибки, если она возникла
//	}
//
// # ---Объект типа GetResp:
//
//	type GetResp struct {
//		Id            string	//  	Id песни
//		SongName      string	// 	Название песни
//		AuthorId      string	//	Id автора
//		Collaborators []string	// 	Массив, который содержит id коллабораторов
//		Genre         string	//	Жанр песни
//		ReleaseDate   *timestamppb.Timestamp	//	Дата релиза
//		Error         string	// Поле, которое содержит информацию об ошибке, если она произошла
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *MusicService) GetByAuthor(_ context.Context, r *music_service_proto.GetByAuthorReq) (*music_service_proto.GetSongsResp, error) {
	songs, err := repository.GetByAuthor(r.AuthorId)
	if err != nil {
		return response.CreateGetErrorSongsResp("couldn't get songs"), nil
	}
	return response.CreateGetSongsResp(songs), nil
}

// GetByCollaborators Хендлер, который получает данные всех песен, где пользователь является коллаборатором
//
// # ---Принимает аргумент типа GetByCollaboratorsReq:
//
//	type GetByCollaboratorsReq struct {
//		CollaboratorId string	// Id коллаборатора
//	}
//
// # ---Возвращает объект типа GetSongsResp:
//
//	type GetSongsResp struct {
//		Songs 	[]*GetResp	// 	Поле, которое содержит массив из GetResp. Это все песни, в которых учавствовал пользователь
//		Error 	string      	//	Поле, которое содержит детали ошибки, если она возникла
//	}
//
// # ---Объект типа GetResp:
//
//	type GetResp struct {
//		Id            string	//  	Id песни
//		SongName      string	// 	Название песни
//		AuthorId      string	//	Id автора
//		Collaborators []string	// 	Массив, который содержит id коллабораторов
//		Genre         string	//	Жанр песни
//		ReleaseDate   *timestamppb.Timestamp	//	Дата релиза
//		Error         string	// Поле, которое содержит информацию об ошибке, если она произошла
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *MusicService) GetByCollaborators(_ context.Context, r *music_service_proto.GetByCollaboratorsReq) (*music_service_proto.GetSongsResp, error) {
	songs, err := repository.GetByAuthorFeaturing(r.CollaboratorId)
	if err != nil {
		return response.CreateGetErrorSongsResp("couldn't get songs"), nil
	}
	return response.CreateGetSongsResp(songs), nil
}

// GetBySinger Хендлер, который получает данные всех песен, где присутствует пользователь(певец). Этот хендлер обращается к поиску песен
// по автору и по коллаборатору
//
// # ---Принимает аргумент типа GetBySingerReq:
//
//	type GetBySingerReq struct {
//		SingerId string	 // Id певца
//	}
//
// # ---Возвращает объект типа GetSongsResp:
//
//	type GetSongsResp struct {
//		Songs 	[]*GetResp	// 	Поле, которое содержит массив из GetResp. Это все песни, в которых учавствовал автор
//		Error 	string      	//	Поле, которое содержит детали ошибки, если она возникла
//	}
//
// # ---Объект типа GetResp:
//
//	type GetResp struct {
//		Id            string	//  	Id песни
//		SongName      string	// 	Название песни
//		AuthorId      string	//	Id автора
//		Collaborators []string	// 	Массив, который содержит id коллабораторов
//		Genre         string	//	Жанр песни
//		ReleaseDate   *timestamppb.Timestamp	//	Дата релиза
//		Error         string	// Поле, которое содержит информацию об ошибке, если она произошла
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *MusicService) GetBySinger(_ context.Context, r *music_service_proto.GetBySingerReq) (*music_service_proto.GetSongsResp, error) {
	songs, err := music.GetBySinger(r)
	if err != nil {
		return response.CreateGetErrorSongsResp("couldn't get song"), nil
	}
	return response.CreateGetSongsResp(songs), nil
}

// UpdateSong Хендлер, который обноляет данные песни и возвращает обновленные данные песни
//
// # ---Принимает аргумент типа UpdateReq:
//
//	type UpdateReq struct {
//		Token         string	// jwt токен пользователяя
//		SongId        string	// id песни, которая должна быть обновлена
//		SongName      string	// новое название песни
//		Collaborators []string	// изменения коллабораторов
//	}
//
// # ---Возвращает объект типа GetResp:
//
//	type GetResp struct {
//		Id            string	//  	Id песни
//		SongName      string	// 	Название песни
//		AuthorId      string	//	Id автора
//		Collaborators []string	// 	Массив, который содержит id коллабораторов
//		Genre         string	//	Жанр песни
//		ReleaseDate   *timestamppb.Timestamp	//	Дата релиза
//		Error         string	// Поле, которое содержит информацию об ошибке, если она произошла
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *MusicService) UpdateSong(_ context.Context, req *music_service_proto.UpdateReq) (*music_service_proto.GetResp, error) {
	updatedSong, err := music.Update(req)
	if err != nil {
		return response.CreateErrorResp("couldn't update song"), nil
	}
	return response.CreateGetResp(updatedSong), nil
}

// DeleteSong Хендлер, который удаляет данные песни из БД
//
// # ---Принимает аргумент типа DeleteReq:
//
//	type DeleteReq struct {
//		Token  string	// jwt токен пользователя
//		SongId string	// id песни, которая будет удалена
//	}
//
// # ---Возвращает объект типа DeleteResp:
//
//	type DeleteResp struct {
//		Message string	// Поле, которое содержит сообщение при успешном выполнении хендлера
//		Error   string	// Поле, которое содержит об ошибке, если она имеется
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *MusicService) DeleteSong(_ context.Context, r *music_service_proto.DeleteReq) (*music_service_proto.DeleteResp, error) {
	err := music.Delete(r)
	if err != nil {
		return &music_service_proto.DeleteResp{Error: "couldn't delete user"}, nil
	}
	return &music_service_proto.DeleteResp{Message: fmt.Sprintf("song with id %s was deleted", r.SongId)}, nil
}

func (s *MusicService) IncreasePlayedQuantity(_ context.Context,
	r *music_service_proto.IncreasePlayedQuantityReq) (*music_service_proto.GetResp, error) {
	song, err := music.AddUserToPlayed(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: couldn't add user to played. Details: %v", err))
	}
	return response.CreateGetResp(song), nil
}

func (s *MusicService) IncreaseListenedQuantity(_ context.Context,
	r *music_service_proto.IncreaseListenedQuantityReq) (*music_service_proto.GetResp, error) {
	song, err := music.AddUserToListeners(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't add user to listened. Details: %v", err))
	}
	return response.CreateGetResp(song), nil
}

func (s *MusicService) LikeSong(_ context.Context,
	r *music_service_proto.LikeSongReq) (*music_service_proto.GetResp, error) {
	song, err := music.AddUserToLiked(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't like song %s. Details: %v", song.Id, err))
	}
	return response.CreateGetResp(song), nil
}

func (s *MusicService) DislikeSong(_ context.Context,
	r *music_service_proto.DislikeSongReq) (*music_service_proto.GetResp, error) {
	song, err := music.AddUserToDisliked(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't dislike song %s. Details: %v", song.Id, err))
	}
	return response.CreateGetResp(song), nil
}

func (s *MusicService) GetUserLiked(_ context.Context,
	r *music_service_proto.GetLikedSongsReq) (*music_service_proto.GetSongsResp, error) {
	songs, err := music.GetUserLiked(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get liked songs of user %s. Details: %v", r.UserId, err))
	}
	return response.CreateGetSongsResp(songs), nil
}

func (s *MusicService) GetUserDisliked(_ context.Context,
	r *music_service_proto.GetDislikedSongsReq) (*music_service_proto.GetSongsResp, error) {
	songs, err := music.GetUserDisliked(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get disliked songs of user %s. Details: %v", r.UserId, err))
	}
	return response.CreateGetSongsResp(songs), nil
}

func (s *MusicService) GetTrends(_ context.Context,
	_ *music_service_proto.GetTrendsReq) (*music_service_proto.GetSongsResp, error) {
	trends, err := music.GetTrends()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get trends. Details: %v", err))
	}
	return response.CreateSimpleGetSongsResp(trends), nil
}
