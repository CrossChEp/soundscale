package redis_repo

import (
	"context"
	"encoding/json"
	"fmt"
	"playlist_serivce/pkg/config/global_vars_config"
	models2 "playlist_serivce/pkg/models"
	"playlist_serivce/pkg/service/checkers"
	"playlist_serivce/pkg/service/logger"
)

func Save(playlist models2.PlaylistModel) error {
	existingPlaylist, err := global_vars_config.Redis.Get(context.Background(), playlist.UserId).Result()
	if err == nil {
		if err := addPlaylistDataToExistingPlaylist(existingPlaylist, playlist); err != nil {
			return err
		}
		return nil
	}
	if err := save(playlist); err != nil {
		return err
	}
	return nil
}

func addPlaylistDataToExistingPlaylist(existingPlaylist string, newPlaylist models2.PlaylistModel) error {
	var songs []models2.SongModel
	err := json.Unmarshal([]byte(existingPlaylist), &songs)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(SavePlaylist) Couldn't marshal playlist_service. Details: %v", err))
		return err
	}
	err = addSongToExistingPlaylist(songs, newPlaylist)
	if err != nil {
		return err
	}
	return nil
}

func addSongToExistingPlaylist(songs []models2.SongModel, newPlaylist models2.PlaylistModel) error {
	for _, song := range newPlaylist.Songs {
		if !checkers.IsElementInSlice(song.SongId, songs) {
			songs = append(songs, song)
		}
	}
	songsStr, err := json.Marshal(songs)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(SavePlaylist) Couldn't marshal playlist_service. Details: %v", err))
		return err
	}
	err = global_vars_config.Redis.Set(context.Background(), newPlaylist.UserId, songsStr, 0).Err()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(SavePlaylist) Couldn't save playlist_service. Details: %v", err))
		return err
	}
	return nil
}

func save(playlist models2.PlaylistModel) error {
	songsStr, err := json.Marshal(playlist.Songs)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(SavePlaylist) Couldn't marshal playlist_service. Details: %v", err))
		return err
	}
	err = global_vars_config.Redis.Set(context.Background(), playlist.UserId, songsStr, 0).Err()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(SavePlaylist) Couldn't save playlist_service. Details: %v", err))
		return err
	}
	return nil
}
