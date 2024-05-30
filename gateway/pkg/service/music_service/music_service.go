package music_service

import (
	"errors"
	"fmt"
	"gateway/pkg/model/collection_model"
	"gateway/pkg/model/music_model"
	"gateway/pkg/service/auxilary"
	"gateway/pkg/service/collection_service"
	grpc_service "gateway/pkg/service/grpc_service/music"
	"gateway/pkg/service/logger"
	"os"
)

func AddSong(songData *music_model.MusicAddModel) (*music_model.SongGetModel, error) {
	dir, _ := os.Getwd()
	if songData.Genre == "" {
		err := errors.New(fmt.Sprintf("Error: this song has now genre. Genre: %s", songData.Genre))
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: this song has now genre. Genre: %s", songData.Genre), err, dir)
	}
	resp, err := grpc_service.AddSong(songData)
	if err != nil {
		return nil, err
	}
	getModel := &music_model.SongGetModel{}
	getModel.ToModel(resp)
	return getModel, nil
}

func GetSongById(userId string, songId string) (*music_model.SongGetModel, error) {
	resp, err := grpc_service.GetSongById(songId)
	if err != nil {
		return nil, err
	}
	getModel := &music_model.SongGetModel{}
	getModel.ToModel(resp)
	collection, err := collection_service.GetCollection(userId)
	if err != nil {
		return nil, err
	}
	if getModel.Exclusive && !auxilary.IsElementInArr(getModel.AuthorId, collection.Subscribed) &&
		userId != getModel.AuthorId {
		return nil, nil
	}
	return getModel, nil
}

func UpdateSong(newSongData *music_model.MusicUpdateModel) (*music_model.SongGetModel, error) {
	resp, err := grpc_service.UpdateSong(newSongData)
	if err != nil {
		return nil, err
	}
	getModel := &music_model.SongGetModel{}
	getModel.ToModel(resp)
	return getModel, nil
}

func DeleteSong(deleteReq *music_model.MusicDeleteModel) error {
	_, err := grpc_service.DeleteSong(deleteReq)
	if err != nil {
		return err
	}
	return nil
}

func GetByAuthor(userId string, authorId string) (*music_model.SongsGetModel, error) {
	resp, err := grpc_service.GetByAuthor(authorId)
	if err != nil {
		return nil, err
	}
	collection, err := collection_service.GetCollection(userId)
	songsModel := &music_model.SongsGetModel{}
	songsModel.ToModel(resp)
	songsModel.Songs = filter(collection, songsModel.Songs)
	return songsModel, nil
}

func filter(collection *collection_model.CollectionGetModel, songs []*music_model.SongGetModel) []*music_model.SongGetModel {
	var filtered []*music_model.SongGetModel
	for _, song := range songs {
		if collection.UserId != song.AuthorId &&
			song.Exclusive &&
			!auxilary.IsElementInArr(song.AuthorId, collection.Subscribed) {
			continue
		}
		filtered = append(filtered, song)
	}
	return filtered
}

func GetByCollaborator(userId string, collaboratorId string) (*music_model.SongsGetModel, error) {
	resp, err := grpc_service.GetByCollaborator(collaboratorId)
	if err != nil {
		return nil, err
	}
	songsModel := &music_model.SongsGetModel{}
	songsModel.ToModel(resp)
	collection, err := collection_service.GetCollection(userId)
	if err != nil {
		return nil, err
	}
	songsModel.Songs = filter(collection, songsModel.Songs)
	return songsModel, nil
}

func GetBySinger(userId string, singerId string) (*music_model.SongsGetModel, error) {
	resp, err := grpc_service.GetBySinger(singerId)
	if err != nil {
		return nil, err
	}
	songsModel := &music_model.SongsGetModel{}
	songsModel.ToModel(resp)
	collection, err := collection_service.GetCollection(userId)
	if err != nil {
		return nil, err
	}
	songsModel.Songs = filter(collection, songsModel.Songs)
	return songsModel, nil
}

func LikeSong(userId string, songId string) (*music_model.SongGetModel, error) {
	resp, err := grpc_service.LikeSong(userId, songId)
	if err != nil {
		return nil, err
	}
	songsModel := &music_model.SongGetModel{}
	songsModel.ToModel(resp)
	return songsModel, nil
}

func DislikeSong(userId string, songId string) (*music_model.SongGetModel, error) {
	resp, err := grpc_service.DislikeSong(userId, songId)
	if err != nil {
		return nil, err
	}
	songsModel := &music_model.SongGetModel{}
	songsModel.ToModel(resp)
	return songsModel, nil
}

func GetLikedSongs(userId string) (*music_model.SongsGetModel, error) {
	resp, err := grpc_service.GetUserLikedSongs(userId)
	if err != nil {
		return nil, err
	}
	songsModel := &music_model.SongsGetModel{}
	songsModel.ToModel(resp)
	return songsModel, nil
}

func GetDislikedSongs(userId string) (*music_model.SongsGetModel, error) {
	resp, err := grpc_service.GetUserDislikedSongs(userId)
	if err != nil {
		return nil, err
	}
	songsModel := &music_model.SongsGetModel{}
	songsModel.ToModel(resp)
	return songsModel, nil
}

func AddUserToPlayed(userId string, songId string) (*music_model.SongGetModel, error) {
	resp, err := grpc_service.AddUserToPlayedQuantity(userId, songId)
	if err != nil {
		return nil, err
	}
	songModel := &music_model.SongGetModel{}
	songModel.ToModel(resp)
	return songModel, nil
}

func GetTrends(userId string) (*music_model.SongsGetModel, error) {
	resp, err := grpc_service.GetTrends()
	if err != nil {
		return nil, err
	}
	songs := &music_model.SongsGetModel{}
	songs.ToModel(resp)
	collection, err := collection_service.GetCollection(userId)
	if err != nil {
		return nil, err
	}
	songs.Songs = filter(collection, songs.Songs)
	return songs, nil
}
