package album_handler

import (
	"fmt"
	"gateway/pkg/config/constants"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/model/playlist_model"
	"gateway/pkg/service/album_service"
	"gateway/pkg/service/jwt_service"
	"gateway/pkg/service/logger"
	"gateway/pkg/service/playlist_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	global_vars_config.Router.Group("albums")
	{
		global_vars_config.Router.POST("/api/album", addAlbum)
		global_vars_config.Router.GET("/api/album/:id", getAlbum)
		global_vars_config.Router.GET("/api/album/author/:authorId", getAlbumsByAuthor)
		global_vars_config.Router.DELETE("/api/album/:playlist_id", deleteAlbum)
		global_vars_config.Router.POST("/api/album/like/:album_id", likeAlbum)
		global_vars_config.Router.POST("/api/album/dislike/:album_id", dislikeAlbum)
		global_vars_config.Router.GET("/api/albums/liked", getUserLikedAlbums)
		global_vars_config.Router.GET("/api/albums/disliked", getUserDislikedAlbums)
	}
}

func addAlbum(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	albumReq := &playlist_model.PlaylistAddModel{}
	if err := albumReq.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't validate request. Details: %v", err))
		return
	}
	albumReq.AuthorId = claims.Id
	albumReq.Type = constants.AlbumType
	addedPlaylist, err := playlist_service.AddPlaylist(albumReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add album. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, addedPlaylist)
}

func getAlbum(ctx *gin.Context) {
	albumId := ctx.Param("id")
	album, err := album_service.GetAlbum(albumId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("Error: couldn't get album by id. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, album)
}

func getAlbumsByAuthor(ctx *gin.Context) {
	authorId := ctx.Param("authorId")
	albums, err := album_service.GetAlbumsByAuthor(authorId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("Couldn't find albums. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, albums)
}

func deleteAlbum(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	albumId := ctx.Param("playlist_id")
	playlistData := &playlist_model.PlaylistDeleteModel{
		AuthorId:     claims.Id,
		PlaylistId:   albumId,
		PlaylistType: constants.AlbumType,
	}
	if err := playlist_service.Delete(playlistData); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't delete album. Delete: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, "album was deleted")
}

func likeAlbum(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	albumId := ctx.Param("album_id")
	likeDislikeModel := playlist_model.LikeDislikePlaylist{
		PlaylistId:   albumId,
		UserId:       claims.Id,
		PlaylistType: constants.AlbumType,
	}
	album, err := playlist_service.LikePlaylist(likeDislikeModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't like album %s by user %s. Details: %v", albumId, claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, album)
}

func dislikeAlbum(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	albumId := ctx.Param("album_id")
	likeDislikeModel := playlist_model.LikeDislikePlaylist{
		PlaylistId:   albumId,
		UserId:       claims.Id,
		PlaylistType: constants.AlbumType,
	}
	album, err := playlist_service.DislikePlaylist(likeDislikeModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't like album %s by user %s. Details: %v", albumId, claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, album)
}

func getUserLikedAlbums(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	album, err := playlist_service.GetUserLiked(claims.Id, constants.AlbumType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get liked albums of user %s. Details: %v", claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, album)
	logger.InfoLog("response data was sent")
}

func getUserDislikedAlbums(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	album, err := playlist_service.GetUserDisliked(claims.Id, constants.AlbumType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get liked albums of user %s. Details: %v", claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, album)
	logger.InfoLog("response data was sent")
}
