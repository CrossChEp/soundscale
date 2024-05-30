package playlist_handler

import (
	"fmt"
	"gateway/pkg/config/constants"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/model/playlist_model"
	"gateway/pkg/service/jwt_service"
	"gateway/pkg/service/logger"
	"gateway/pkg/service/playlist_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	global_vars_config.Router.Group("playlist")
	{
		global_vars_config.Router.POST("/api/playlist", AddPlaylist)
		global_vars_config.Router.GET("/api/playlist/:id", GetPlaylist)
		global_vars_config.Router.GET("/api/playlists/author/:authorId", GetPlaylistByAuthor)
		global_vars_config.Router.POST("/api/playlist/songs", AddSongToPlaylist)
		global_vars_config.Router.DELETE("/api/playlist/songs", DeleteSongFromPlaylist)
		global_vars_config.Router.POST("/api/playlist/current", addCurrentPlaylist)
		global_vars_config.Router.DELETE("/api/playlist/:playlist_id", DeletePlaylist)
		global_vars_config.Router.POST("/api/playlist/like/:playlist_id", likePlaylist)
		global_vars_config.Router.POST("/api/playlist/dislike/:playlist_id", dislikePlaylist)
		global_vars_config.Router.GET("/api/playlists/liked", getUserLikedPlaylists)
		global_vars_config.Router.GET("/api/playlists/disliked", getUserDislikedPlaylists)
	}
}

func AddPlaylist(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlistReq := &playlist_model.PlaylistAddModel{}
	if err := playlistReq.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't validate request. Details: %v", err))
		return
	}
	playlistReq.AuthorId = claims.Id
	addedPlaylist, err := playlist_service.AddPlaylist(playlistReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add playlist. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, addedPlaylist)
}

func GetPlaylist(ctx *gin.Context) {
	playlistId := ctx.Param("id")
	playlist, err := playlist_service.GetPlaylist(playlistId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("Error: couldn't get playlist by id. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, playlist)
}

func GetPlaylistByAuthor(ctx *gin.Context) {
	authorId := ctx.Param("authorId")
	playlists, err := playlist_service.GetPlaylistsByAuthor(authorId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("Couldn't find playlists. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, playlists)
}

func AddSongToPlaylist(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlistData := &playlist_model.AddSongsToPlaylistModel{}
	if err := playlistData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return
	}
	playlistData.AuthorId = claims.Id
	resPlaylist, err := playlist_service.AddSongToPlaylist(playlistData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add song to playlist. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, resPlaylist)
}

func DeleteSongFromPlaylist(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlistData := &playlist_model.DeleteSongFromPlaylistModel{}
	if err := playlistData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return
	}
	playlistData.AuthorId = claims.Id
	resPlaylist, err := playlist_service.DeleteSongFromPlaylist(playlistData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't delete song from playlist. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, resPlaylist)
}

func DeletePlaylist(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlistId := ctx.Param("playlist_id")
	playlistData := &playlist_model.PlaylistDeleteModel{
		AuthorId:     claims.Id,
		PlaylistId:   playlistId,
		PlaylistType: constants.PlaylistType,
	}
	if err := playlist_service.Delete(playlistData); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't delete playlist. Delete: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, "playlist was deleted")
}

func addCurrentPlaylist(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlistData := &playlist_model.AddCurrentPlaylistModel{}
	if err := playlistData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't validate request. Details: %v", err))
		return
	}
	playlistData.UserId = claims.Id
	if err := playlist_service.AddCurrentPlaylist(playlistData); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add playlist. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("Playlist was added"))
}

func likePlaylist(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlistId := ctx.Param("playlist_id")
	likeDislikeModel := playlist_model.LikeDislikePlaylist{
		PlaylistId:   playlistId,
		UserId:       claims.Id,
		PlaylistType: constants.PlaylistType,
	}
	playlist, err := playlist_service.LikePlaylist(likeDislikeModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't like playlist %s by user %s. Details: %v", playlistId, claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, playlist)
	logger.InfoLog("response data was sent")
}

func dislikePlaylist(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlistId := ctx.Param("playlist_id")
	likeDislikeModel := playlist_model.LikeDislikePlaylist{
		PlaylistId:   playlistId,
		UserId:       claims.Id,
		PlaylistType: constants.PlaylistType,
	}
	playlist, err := playlist_service.DislikePlaylist(likeDislikeModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't dislike playlist %s by user %s. Details: %v", playlistId, claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, playlist)
	logger.InfoLog("response data was sent")
}

func getUserLikedPlaylists(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlist, err := playlist_service.GetUserLiked(claims.Id, constants.PlaylistType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get liked playlists of user %s. Details: %v", claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, playlist)
	logger.InfoLog("response data was sent")
}

func getUserDislikedPlaylists(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlist, err := playlist_service.GetUserDisliked(claims.Id, constants.PlaylistType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get liked playlists of user %s. Details: %v", claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, playlist)
	logger.InfoLog("response data was sent")
}
