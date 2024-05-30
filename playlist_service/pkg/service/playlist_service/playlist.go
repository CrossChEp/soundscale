package playlist_service

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"playlist_serivce/pkg/config/constants"
	"playlist_serivce/pkg/models"
	"playlist_serivce/pkg/models/playlist_models"
	"playlist_serivce/pkg/proto/playlist_service_proto"
	"playlist_serivce/pkg/reposiory/mongo_repo/playlist_repo"
	"playlist_serivce/pkg/reposiory/redis_repo"
	"playlist_serivce/pkg/service/base_functions"
	"playlist_serivce/pkg/service/checkers"
	"playlist_serivce/pkg/service/logger"
)

func AddPlaylist(r *playlist_service_proto.AddPlaylistReq) (*playlist_models.PlaylistGetModel, error) {
	if !checkers.IsUserExists(r.AuthorId) {
		logger.ErrorLog("Error: user with such id doesn't exist")
		return nil, errors.New("user with such id doesn't exist")
	}
	if err := checkSongs(r.SongIds); err != nil {
		return nil, err
	}
	playlistId, err := playlist_repo.Save(r)
	if err != nil {
		return nil, err
	}
	playlist, err := playlist_repo.Get(playlistId.(primitive.ObjectID), r.Type)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func GetPlaylist(r *playlist_service_proto.GetPlaylistReq) (*playlist_models.PlaylistGetModel, error) {
	objectId, err := primitive.ObjectIDFromHex(r.Id)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert string to object id. Details: %v", err))
		return nil, err
	}
	playlist, err := playlist_repo.Get(objectId, r.PlaylistType)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func GetPlaylistsByAuthor(r *playlist_service_proto.GetPlaylistsByAuthorReq) ([]playlist_models.PlaylistGetModel, error) {
	playlists, err := playlist_repo.GetByAuthor(r.Id, r.PlaylistType)
	if err != nil {
		return nil, err
	}
	return playlists, nil
}

func AddSongsToPlaylist(r *playlist_service_proto.AddSongsToPlaylistReq) (*playlist_models.PlaylistGetModel, error) {
	r.SongIds = base_functions.RemoveUnexistingSong(r.SongIds)
	playlistId, err := primitive.ObjectIDFromHex(r.PlaylistId)
	if err != nil {
		return nil, err
	}
	playlist, err := playlist_repo.Get(playlistId, r.PlaylistType)
	if playlist == nil {
		logger.ErrorLog(fmt.Sprintf("Error: playlist with id %s does not exist", playlistId))
		return nil, errors.New(fmt.Sprintf("Error: playlist with id %s does not exist", playlistId))
	}
	if r.AuthorId != playlist.Author {
		logger.ErrorLog(fmt.Sprintf("Error user with id %s is not author of playlist with id %s", r.AuthorId, playlist.Author))
		return nil, errors.New(fmt.Sprintf("Error user with id %s is not author of playlist with id %s", r.AuthorId, playlist.Author))
	}
	err = playlist_repo.AddToPlaylist(playlist_models.AddRemoveSongsModel{
		PlaylistId:   playlistId,
		SongsIds:     r.SongIds,
		PlaylistType: r.PlaylistType,
	})
	if err != nil {
		return nil, err
	}
	oid, err := primitive.ObjectIDFromHex(r.PlaylistId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert playlist id to object id. Details: %v", err))
		return nil, err
	}
	playlist, err = playlist_repo.Get(oid, r.PlaylistType)
	return playlist, nil
}

func DeleteSongsFromPlaylist(r *playlist_service_proto.DeleteSongsFromPlaylistReq) (*playlist_models.PlaylistGetModel, error) {
	r.SongsIds = base_functions.RemoveUnexistingSong(r.SongsIds)
	objectId, err := primitive.ObjectIDFromHex(r.PlaylistId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert id to object id. Details: %v", err))
		return nil, err
	}
	playlist, err := playlist_repo.Get(objectId, r.PlaylistType)
	if r.AuthorId != playlist.Author {
		logger.ErrorLog(fmt.Sprintf("Error user with id %s is not author of playlist with id %s", r.AuthorId, playlist.Author))
		return nil, errors.New(fmt.Sprintf("Error user with id %s is not author of playlist with id %s", r.AuthorId, playlist.Author))
	}
	err = playlist_repo.DeleteFromPlaylist(playlist_models.AddRemoveSongsModel{
		PlaylistId:   objectId,
		SongsIds:     r.SongsIds,
		PlaylistType: r.PlaylistType,
	})
	if err != nil {
		return nil, err
	}
	playlist, err = playlist_repo.Get(objectId, r.PlaylistType)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func AddCurrentPlaylist(r *playlist_service_proto.AddCurrentPlaylistReq) error {
	if !checkers.IsUserExists(r.UserId) {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user with id %s", r.UserId))
		return errors.New(fmt.Sprintf("couldn't get user with id %s", r.UserId))
	}
	if err := checkSongs(r.Songs); err != nil {
		return err
	}
	err := redis_repo.Save(models.PlaylistModel{
		Songs:  generateSongsModel(r.Songs),
		UserId: r.UserId,
	})
	if err != nil {
		return err
	}
	return nil
}

func DeletePlaylist(r *playlist_service_proto.DeletePlaylistReq) error {
	objectId, err := primitive.ObjectIDFromHex(r.Id)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert id to object id. Details: %v", err))
		return err
	}
	playlist, err := playlist_repo.Get(objectId, r.PlaylistType)
	if r.AuthorId != playlist.Author {
		logger.ErrorLog(fmt.Sprintf("Error user with id %s is not author of playlist with id %s", r.AuthorId, playlist.Author))
		return errors.New(fmt.Sprintf("Error user with id %s is not author of playlist with id %s", r.AuthorId, playlist.Author))
	}
	if playlist.Type == constants.TypeFavourite {
		logger.ErrorLog(fmt.Sprintf("Error: can't delete favourite playlist"))
		return errors.New("error: can't delete favourite playlist")
	}
	if err := playlist_repo.DeletePlaylist(objectId); err != nil {
		return err
	}
	return nil
}

func GetFavouriteSongs(userId string) (*playlist_models.PlaylistGetModel, error) {
	if !checkers.IsUserExists(userId) {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user with id %s.", userId))
		return nil, errors.New(fmt.Sprintf("couldn't get user with id %s.", userId))
	}
	playlist, err := playlist_repo.GetUserFavourite(userId)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func generateSongsModel(songs []string) []models.SongModel {
	var songModels []models.SongModel
	for _, songId := range songs {
		model := models.SongModel{
			SongId:    songId,
			ReadedAt:  0,
			ReadedSum: 0,
		}
		songModels = append(songModels, model)
	}
	return songModels
}

func checkSongs(songs []string) error {
	for _, songId := range songs {
		if !checkers.IsSongExists(songId) {
			logger.ErrorLog(fmt.Sprintf("Error: song with id %s doesn't exist", songId))
			return errors.New(fmt.Sprintf("song with id %s doesn't exist", songId))
		}
	}
	return nil
}

func LikePlaylist(r *playlist_service_proto.LikePlaylistReq) (*playlist_models.PlaylistGetModel, error) {
	playlistId, err := primitive.ObjectIDFromHex(r.PlaylistId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert id %s to object id. Details: %v", playlistId, err))
		return nil, err
	}
	playlist, err := playlist_repo.Get(playlistId, r.PlaylistType)
	if err != nil {
		return nil, err
	}
	if checkers.IsUserInArr(r.UserId, playlist.Liked) {
		logger.InfoLog(fmt.Sprintf("user %s already liked playlist %s", r.UserId, r.PlaylistId))
		return playlist, nil
	}
	if checkers.IsUserInArr(r.UserId, playlist.Disliked) {
		playlist.Disliked = base_functions.DeleteFromArr(playlist.Disliked, r.UserId)
	}
	likeModel := &playlist_models.LikeDislikeModel{
		Playlist:     playlist,
		UserId:       r.UserId,
		PlaylistType: r.PlaylistType,
	}
	if err := playlist_repo.LikePlaylist(likeModel); err != nil {
		return nil, err
	}
	return playlist_repo.Get(playlistId, r.PlaylistType)
}

func DislikePlaylist(r *playlist_service_proto.DislikePlaylistReq) (*playlist_models.PlaylistGetModel, error) {
	playlistId, err := primitive.ObjectIDFromHex(r.PlaylistId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert id %s to object id. Details: %v", playlistId, err))
		return nil, err
	}
	playlist, err := playlist_repo.Get(playlistId, r.PlaylistType)
	if err != nil {
		return nil, err
	}
	if checkers.IsUserInArr(r.UserId, playlist.Disliked) {
		logger.InfoLog(fmt.Sprintf("user %s already disliked playlist %s", r.UserId, r.PlaylistId))
		return playlist, nil
	}
	if checkers.IsUserInArr(r.UserId, playlist.Liked) {
		playlist.Liked = base_functions.DeleteFromArr(playlist.Liked, r.UserId)
	}
	likeModel := &playlist_models.LikeDislikeModel{
		Playlist:     playlist,
		UserId:       r.UserId,
		PlaylistType: r.PlaylistType,
	}
	if err := playlist_repo.DislikedPlaylist(likeModel); err != nil {
		return nil, err
	}
	return playlist_repo.Get(playlistId, r.PlaylistType)
}

func GetLikedPlaylists(userId string, playlistType string) ([]playlist_models.PlaylistGetModel, error) {
	if !checkers.IsUserExists(userId) {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user with id %s.", userId))
		return nil, errors.New(fmt.Sprintf("couldn't get user with id %s.", userId))
	}
	playlist, err := playlist_repo.GetUserLiked(userId, playlistType)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func GetDislikedPlaylists(userId string, playlistType string) ([]playlist_models.PlaylistGetModel, error) {
	if !checkers.IsUserExists(userId) {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user with id %s.", userId))
		return nil, errors.New(fmt.Sprintf("couldn't get user with id %s.", userId))
	}
	playlist, err := playlist_repo.GetUserDisliked(userId, playlistType)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}
