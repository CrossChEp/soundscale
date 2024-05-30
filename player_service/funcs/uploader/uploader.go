package uploader

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"player_service/funcs/grpc_funcs"
	"player_service/funcs/jwt_funcs"
	"player_service/funcs/logger"
	"player_service/models"
	"player_service/repository"
)

func Upload(ctx *gin.Context) {
	logger.InfoLog("Upload function was called")
	var upload = &models.UploadModel{}
	if err := upload.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	claims, err := jwt_funcs.DecodeJwt(upload.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	song, err := grpc_funcs.GetSong(upload.SongId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if song.AuthorId != claims.Id {
		ctx.JSON(http.StatusForbidden, "user is not author of this song")
		return
	}
	if err := repository.Upload(upload, ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	logger.InfoLog("song was uploaded")
	ctx.JSON(http.StatusOK, "song was uploaded")
}
