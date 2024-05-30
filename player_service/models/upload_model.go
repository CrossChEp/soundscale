package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"player_service/funcs/logger"
)

type UploadModel struct {
	Token     string                `json:"token"`
	SongId    string                `json:"songId"`
	MusicFile *multipart.FileHeader `json:"songFile"`
}

func (model *UploadModel) ToModel(ctx *gin.Context) error {
	model.Token = ctx.PostForm("token")
	model.SongId = ctx.PostForm("songId")
	var err error
	model.MusicFile, err = ctx.FormFile("songFile")
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request to model. Details: %v\n", err))
		return err
	}
	return nil
}
