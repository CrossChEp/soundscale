// Package service Пакет, который содержит методы для управления подгрузчиком песен
package service

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"player_service/config"
	"player_service/funcs/grpc_funcs"
	"player_service/funcs/logger"
	"player_service/models"
	"player_service/models/redis_models"
	"player_service/repository/redis_repo"
	"strings"
	"time"
)

// DownloadSong Функция, которая принимает объект типа net.Conn. Эта функция обрабатывает все запросы на tcp
// сервис плеера.
//
// # ---Запрос на сервис посылается таким видом:
//
// "<jwt токен пользователя>\t<время, которое нужно прибавить к текущему проигрываему времени в форматe
// "час:минута:секунда">\n"
//
// # ---!Важно!---
//
// если вы не собираетесь перематывать песню, оставьте второй аргумент пустым
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func DownloadSong(conn net.Conn) {
	commandsStr, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't read commands. Details: %v", err))
		conn.Close()
		return
	}
	commands := strings.Split(commandsStr, "\t")
	if !isDataValid(commands) {
		conn.Close()
		return
	}
	execute(commands, conn)
}

func isDataValid(commands []string) bool {
	if commands[0] == "" {
		return false
	}
	return true
}

func execute(commands []string, conn net.Conn) {
	userId, duration := commands[0], ""
	if len(commands) == 2 {
		duration = commands[1]
	} else {
		userId = userId[:len(userId)-1]
	}
	songs, err := redis_repo.Get(userId)
	if err != nil {
		if err := conn.Close(); err != nil {
			logger.ErrorLog(fmt.Sprintf("Error: couldn't close tcp connection. Details: %v", err))
		}
		return
	}
	if duration != "" {
		durationTime, err := time.Parse("15:04:05", duration[:len(duration)-1])
		time := durationTime.Minute()*60 + durationTime.Second()
		if err != nil {
			logger.ErrorLog(fmt.Sprintf("Error: couldn't parse time. Details: %v", err))
			conn.Close()
			return
		}
		err = setNewFileTime(models.SongTimeModel{
			Song:     songs[0],
			Duration: time,
			UserId:   userId,
		})
		if err != nil {
			conn.Close()
			return
		}
	}
	startPlaylist(models.StartPlaylistModel{
		UserId: userId,
		Songs:  songs,
		Conn:   conn,
	})
	conn.Write([]byte("end\n"))
}

func setNewFileTime(newSongTime models.SongTimeModel) error {
	filePath, err := filepath.Abs(fmt.Sprintf("%s%s.flac", config.SongsPath, newSongTime.Song.SongId))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't create file path. Details: %v", err))
		return err
	}
	file, err := os.Stat(filePath)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get file. Details: %v", err))
		return err
	}
	size := file.Size()
	fileDuration, err := redis_repo.GetSongDuration(newSongTime.UserId, 0)
	if err != nil {
		return err
	}
	bytesPerSec, err := getBytesPerSec(size, fileDuration)
	if err != nil {
		return err
	}
	if err := redis_repo.UpdateDuration(newSongTime.UserId, int(bytesPerSec)*newSongTime.Duration); err != nil {
		return err
	}
	return nil
}

func getBytesPerSec(size int64, duration float64) (int64, error) {
	if duration == 0 || int(duration) > int(size) {
		logger.ErrorLog("Error: time duration is invalid")
		return 0, errors.New("time duration is invalid")
	}
	return size / int64(duration), nil
}

func startPlaylist(playlistData models.StartPlaylistModel) {
	for i := 0; i < len(playlistData.Songs); i++ {
		song, err := redis_repo.Get(playlistData.UserId)
		if err != nil {
			playlistData.Conn.Close()
			return
		}
		err = download(song[0], playlistData.Conn, playlistData.UserId)
		if err != nil {
			playlistData.Conn.Close()
			return
		}
		err = redis_repo.DeleteFromPlaylist(redis_models.DeleteFromPlaylistModel{
			UserId:   playlistData.UserId,
			Position: 0,
		})
		if err != nil {
			playlistData.Conn.Close()
			return
		}
	}
}

func download(song redis_models.SongsModel, conn net.Conn, userId string) error {
	filePath, err := filepath.Abs(fmt.Sprintf("./songs/%s.flac", song.SongId))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Errror: couldn't create filepath. Details: %v", err))
		return err
	}
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Errror: couldn't open file. Details: %v", err))
		return err
	}
	if err := grpc_funcs.AddUserToPlayedQuantity(userId, song.SongId); err != nil {
		return err
	}
	songModel, _ := grpc_funcs.GetSong(song.SongId)
	if err := grpc_funcs.AddSongGenreToUser(userId, songModel.Genre); err != nil {
		return nil
	}
	err = read(redis_models.ReaderModel{
		File:   file,
		SongId: song.SongId,
		Conn:   conn,
		UserId: userId,
	})
	fileStat, err := os.Stat(filePath)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get file stat. Details: %v", err))
		return err
	}
	songs, err := redis_repo.Get(userId)
	if err != nil {
		return err
	}
	if countListenedPercent(int64(songs[0].ReadedSum), fileStat.Size()) {
		if err := grpc_funcs.AddUserToListenedQuantity(userId, songs[0].SongId); err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func read(readerModel redis_models.ReaderModel) error {
	buf := make([]byte, 1024)
	for {
		songs, err := redis_repo.Get(readerModel.UserId)
		if err != nil {
			return err
		}
		readed, err := readerModel.File.ReadAt(buf, int64(songs[0].ReadedAt))
		if err != nil {
			logger.InfoLog(fmt.Sprintf("%v", err))
			if _, err := readerModel.Conn.Write([]byte("EOF\n")); err != nil {
				return err
			}
			return nil
		}
		if err := redis_repo.UpdateDuration(readerModel.UserId, readed); err != nil {
			return err
		}
		if err := redis_repo.UpdateReadedSum(readerModel.UserId, readed); err != nil {
			return err
		}
		r := bytes.NewReader(buf)
		_, err = io.Copy(readerModel.Conn, r)
	}
}
func countListenedPercent(listened int64, songSize int64) bool {
	return listened*100/songSize >= 99
}
