package music_handler

import (
	"fmt"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/model/music_model"
	"gateway/pkg/service/jwt_service"
	"gateway/pkg/service/logger"
	"gateway/pkg/service/music_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	global_vars_config.Router.Group("music")
	{
		global_vars_config.Router.POST("/api/song", addSong)
		global_vars_config.Router.GET("/api/song/:songId", getSong)
		global_vars_config.Router.PUT("/api/song", updateSong)
		global_vars_config.Router.DELETE("/api/song/:songId", deleteSong)
		global_vars_config.Router.GET("/api/songs/author/:authorId", getByAuthor)
		global_vars_config.Router.GET("/api/songs/collaborator/:collaboratorId", getByCollaborator)
		global_vars_config.Router.GET("/api/songs/singer/:singerId", getBySinger)
		global_vars_config.Router.POST("/api/song/like/:songId", likeSong)
		global_vars_config.Router.POST("/api/song/dislike/:songId", dislikeSong)
		global_vars_config.Router.GET("/api/songs/liked", getLikedSongs)
		global_vars_config.Router.GET("/api/songs/disliked", getDislikedSongs)
		global_vars_config.Router.POST("/api/song/played/:songId", addUserToPlayed)
		global_vars_config.Router.GET("/api/songs/trends", getTrends)
	}
}

func addSong(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songModel := &music_model.MusicAddModel{}
	if err := songModel.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't validate request. Details: %v", err))
		return
	}
	songModel.AuthorId = claims.Id
	song, err := music_service.AddSong(songModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add song. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, song)
}

func getSong(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songId := ctx.Param("songId")
	song, err := music_service.GetSongById(claims.Id, songId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get song by id. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, song)
}

func updateSong(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	newSongData := &music_model.MusicUpdateModel{}
	if err := newSongData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't validate request. Details: %v", err))
		return
	}
	newSongData.AuthorId = claims.Id
	song, err := music_service.UpdateSong(newSongData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't update song. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, song)
}

func deleteSong(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	deleteReq := &music_model.MusicDeleteModel{
		UserId: claims.Id,
		SongId: ctx.Param("songId"),
	}
	if err := music_service.DeleteSong(deleteReq); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't delete song. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("Song was deleted"))
}

func getByAuthor(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songs, err := music_service.GetByAuthor(claims.Id, ctx.Param("authorId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get author's songs. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, songs)
}

func getByCollaborator(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songs, err := music_service.GetByCollaborator(claims.Id, ctx.Param("collaboratorId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get collaborator's songs. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, songs)
}

func getBySinger(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songs, err := music_service.GetBySinger(claims.Id, ctx.Param("singerId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get singer's songs. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, songs)
}

func likeSong(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	userId, songId := claims.Id, ctx.Param("songId")
	song, err := music_service.LikeSong(userId, songId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't like song %s by user %s. Details: %v", songId, userId, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, song)
	logger.InfoLog("response data was sent")
}

func dislikeSong(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	userId, songId := claims.Id, ctx.Param("songId")
	song, err := music_service.DislikeSong(userId, songId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't dislike song %s by user %s. Details: %v", songId, userId, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, song)
	logger.InfoLog("response data was sent")
}

func getLikedSongs(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songs, err := music_service.GetLikedSongs(claims.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get liked songs of user %s. Details: %v", claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, songs)
	logger.InfoLog("response data was sent")
}

func getDislikedSongs(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songs, err := music_service.GetDislikedSongs(claims.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get disliked songs of user %s. Details: %v", claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, songs)
	logger.InfoLog("response data was sent")
}

func addUserToPlayed(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songId := ctx.Param("songId")
	song, err := music_service.AddUserToPlayed(claims.Id, songId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add user %s to played users of song %s. Details: %v", claims.Id, songId, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, song)
	logger.InfoLog("response data was sent")
}

func getTrends(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songs, err := music_service.GetTrends(claims.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get songs. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, songs)
}
