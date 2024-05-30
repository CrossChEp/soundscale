package repository

import (
	"fmt"
	"player_service/config"
	commandline "player_service/funcs/command_line"
	"player_service/funcs/logger"
	"player_service/models"

	"github.com/gin-gonic/gin"
)

func Upload(uploadData *models.UploadModel, ctx *gin.Context) error {
	err := ctx.SaveUploadedFile(uploadData.MusicFile, config.SongsPath+uploadData.SongId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't save file. Details: %v", err))
		return err
	}
	err = convertToFlac(uploadData.SongId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert file to flac. Details: %v", err))
		return err
	}
	return nil
}

func convertToFlac(songId string) error {
	execResult := commandline.ShellOut(fmt.Sprintf("ffmpeg -i %ssongs/%s -acodec flac %ssongs/%s.flac", *config.StaticDir, songId, *config.StaticDir, songId))
	if execResult.Error != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't execute command. Details: %v\n", execResult.Stderr))
		err := deleteTempFile(songId)
		if err != nil {
			execResult.Error = err
		}
		return execResult.Error
	}
	if err := deleteTempFile(songId); err != nil {
		return err
	}
	return nil
}

func deleteTempFile(songId string) error {
	execResult := commandline.ShellOut(fmt.Sprintf("rm -r %ssongs/%s", *config.StaticDir, songId))
	if execResult.Error != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't execute command. Details: %v\n", execResult.Stderr))
		return execResult.Error
	}
	return nil
}
