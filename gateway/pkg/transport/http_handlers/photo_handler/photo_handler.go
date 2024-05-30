package photo_handler

import (
	"fmt"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/model/photo_model"
	"gateway/pkg/service/jwt_service"
	"gateway/pkg/service/photo_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	global_vars_config.Router.Group("photo")
	{
		global_vars_config.Router.POST("/api/photo/pfp", uploadPFP)
		global_vars_config.Router.GET("/api/photo/pfp/:userId", downloadPFP)
		global_vars_config.Router.POST("/api/photo/song/cover", uploadSongCover)
		global_vars_config.Router.GET("/api/photo/song/cover/:songId", downloadSongCover)
		global_vars_config.Router.POST("/api/photo/playlist/cover", uploadPlaylistCover)
		global_vars_config.Router.GET("/api/photo/playlist/cover/:playlistId", downloadPlaylistCover)
		global_vars_config.Router.POST("/api/photo/album/cover", uploadAlbumCover)
		global_vars_config.Router.GET("/api/photo/album/cover/:albumId", downloadAlbumCover)
	}
}

func uploadPFP(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: invalid token. Details: %v", err))
		return
	}
	photoData := photo_model.UploadPFPModel{}
	if err := photoData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Couldn't convert request to model. Details: %v", err))
		return
	}
	photoData.UserId = claims.Id
	if err := photo_service.UploadPFP(photoData); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't upload profile picture. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, "Profile picture was uploaded")
}

func downloadPFP(ctx *gin.Context) {
	userId := ctx.Param("userId")
	pfp, err := photo_service.DownloadPFP(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't download pfp. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, pfp)
}

func uploadSongCover(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get token. Details: %v", err))
		return
	}
	coverModel := photo_model.UploadSongCoverModel{}
	if err := coverModel.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to model. Details: %v", err))
		return
	}
	coverModel.UserId = claims.Id
	if err := photo_service.UploadSongCover(coverModel); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't upload song cover. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, "Song cover was uploaded")
}

func downloadSongCover(ctx *gin.Context) {
	songId := ctx.Param("songId")
	songCover, err := photo_service.DownloadSongCover(songId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't download song cover. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("%s", songCover))
}

func uploadPlaylistCover(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	coverModel := &photo_model.UploadPlaylistCoverModel{}
	if err := coverModel.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't upload playlist cover. Details: %v", err))
		return
	}
	coverModel.UserId = claims.Id
	if err := photo_service.UploadPlaylistCover(coverModel); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't upload playlist cover. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, "Playlist cover was uploaded")
}

func downloadPlaylistCover(ctx *gin.Context) {
	playlistId := ctx.Param("playlistId")
	playlistCover, err := photo_service.DownloadPlaylistCover(playlistId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't download playlist cover. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("%s", playlistCover))
}

func uploadAlbumCover(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	coverModel := &photo_model.UploadPlaylistCoverModel{}
	if err := coverModel.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't upload album cover. Details: %v", err))
		return
	}
	coverModel.UserId = claims.Id
	if err := photo_service.UploadAlbumCover(coverModel); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't upload album cover. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, "Album cover was uploaded")
}

func downloadAlbumCover(ctx *gin.Context) {
	albumId := ctx.Param("albumId")
	albumCover, err := photo_service.DownloadAlbumCover(albumId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't download album cover. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("%s", albumCover))
}
