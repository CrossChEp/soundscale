package redis_repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"player_service/config"
	"player_service/funcs/logger"
	"player_service/models/redis_models"

	"github.com/hcl/audioduration"
)

func Get(userId string) ([]redis_models.SongsModel, error) {
	songsStr, err := config.Redis.Get(context.Background(), userId).Result()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get record with user id %s. Details: %d", userId, err))
		return nil, err
	}
	var songs []redis_models.SongsModel
	if err := json.Unmarshal([]byte(songsStr), &songs); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't unmarshall redis record. Details: %v", err))
		return nil, err
	}
	return songs, nil
}

func DeleteFromPlaylist(deleteModel redis_models.DeleteFromPlaylistModel) error {
	playlist, err := Get(deleteModel.UserId)
	if err != nil {
		return err
	}
	playlist, err = remove(playlist, deleteModel.Position)
	if err != nil {
		return err
	}
	playlistStr, err := json.Marshal(playlist)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete song from current playlist. Details: %v", err))
		return err
	}
	err = config.Redis.Set(context.Background(), deleteModel.UserId, playlistStr, 0).Err()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update current playlist. Details: %v", err))
		return err
	}
	return nil
}

func remove(arr []redis_models.SongsModel, pos int) ([]redis_models.SongsModel, error) {
	if pos > len(arr) {
		return nil, errors.New("index is out of bounds")
	}
	return append(arr[:pos], arr[pos+1:]...), nil
}

func Delete(userId string) error {
	err := config.Redis.Del(context.Background(), userId).Err()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete playlist. Details: %v", err))
		return err
	}
	return nil
}

func UpdateDuration(userId string, duration int) error {
	songs, err := Get(userId)
	if err != nil {
		return err
	}
	songs[0].ReadedAt = songs[0].ReadedAt + duration
	songsStr, err := json.Marshal(songs)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't marshal songs. Details: %v", err))
		return err
	}
	if err := config.Redis.Set(context.Background(), userId, songsStr, 0).Err(); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update playlist. Details: %v", err))
		return err
	}
	return nil
}

func UpdateReadedSum(userId string, duration int) error {
	songs, err := Get(userId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get songs of user with id %s. Details: %v", userId, err))
		return err
	}
	songs[0].ReadedSum += duration
	songsStr, err := json.Marshal(songs)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert songs to string. Details: %v", err))
		return err
	}
	if err := config.Redis.Set(context.Background(), userId, songsStr, 0).Err(); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't add readed sum. Details: %v", err))
		return err
	}
	return nil
}

func GetSongDuration(userId string, songPos int) (float64, error) {
	songs, err := Get(userId)
	if err != nil {
		return 0, err
	}
	if songPos > len(songs) || songPos < 0 {
		logger.ErrorLog("Error: song position can't out of bounds.")
		return 0, errors.New("song position is out of bounds")
	}
	filePath, err := filepath.Abs(fmt.Sprintf("%s%s.flac", config.SongsPath, songs[songPos].SongId))
	f, _ := os.Open(filePath)
	defer f.Close()
	d, err := audioduration.FLAC(f)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get duration of song. Details: %v", err))
		return 0, err
	}
	return d, nil
}
