package collection_handler

import (
	"fmt"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/model/collection_model"
	"gateway/pkg/service/collection_service"
	"gateway/pkg/service/jwt_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	global_vars_config.Router.Group("collection")
	{
		global_vars_config.Router.POST("/api/collection/songs", addSongsToCollection)
		global_vars_config.Router.DELETE("/api/collection/songs", deleteSongsFromCollection)
		global_vars_config.Router.GET("/api/collection", getCollection)
		global_vars_config.Router.POST("/api/collection/playlists", addPlaylistToCollection)
		global_vars_config.Router.DELETE("/api/collection/playlists", deletePlaylistsFromCollection)
		global_vars_config.Router.POST("/api/collection/albums", addAlbumToCollection)
		global_vars_config.Router.DELETE("/api/collection/albums", deleteAlbumsFromCollection)
		global_vars_config.Router.POST("/api/collection/genres", addGenresToUserCollection)
	}
}

func addSongsToCollection(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songsData := &collection_model.SongsModel{}
	if err := songsData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return
	}
	songsData.UserId = claims.Id
	resPlaylist, err := collection_service.AddSongs(songsData.UserId, songsData.SongsIds)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add songs to user collection. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, resPlaylist)
}

func deleteSongsFromCollection(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	songs := &collection_model.SongsModel{}
	if err := songs.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return
	}
	songs.UserId = claims.Id
	resPlaylist, err := collection_service.RemoveSongs(songs.UserId, songs.SongsIds)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't delete songs from user collection. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, resPlaylist)
}

func addPlaylistToCollection(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlistData := &collection_model.PlaylistModel{}
	if err := playlistData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return
	}
	playlistData.UserId = claims.Id
	resPlaylist, err := collection_service.AddPlaylists(playlistData.UserId, playlistData.PlaylistsIds)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add playlists to user collection. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, resPlaylist)
}

func addGenresToUserCollection(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	genres := &collection_model.GenresModel{}
	if err := genres.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return
	}
	genres.UserId = claims.Id
	resPlaylist, err := collection_service.AddGenres(genres)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add genres to user %s collection. Details: %v", claims.Id, err))
		return
	}
	ctx.JSON(http.StatusOK, resPlaylist)
}

func deletePlaylistsFromCollection(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	playlistData := &collection_model.PlaylistModel{}
	if err := playlistData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return
	}
	playlistData.UserId = claims.Id
	collection, err := collection_service.RemovePlaylists(playlistData.UserId, playlistData.PlaylistsIds)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't delete playlists from collection. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, collection)
}

func addAlbumToCollection(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	albumData := &collection_model.AlbumModel{}
	if err := albumData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return
	}
	albumData.UserId = claims.Id
	resPlaylist, err := collection_service.AddAlbums(albumData.UserId, albumData.AlbumsIds)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add playlists to user collection. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, resPlaylist)
}

func deleteAlbumsFromCollection(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	albumData := &collection_model.AlbumModel{}
	if err := albumData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return
	}
	albumData.UserId = claims.Id
	collection, err := collection_service.RemoveAlbums(albumData.UserId, albumData.AlbumsIds)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't delete albums from collection. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, collection)
}

func getCollection(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	favourites, err := collection_service.GetCollection(claims.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get favourites of user with id %s. Details: %v", claims.Id, err))
		return
	}
	ctx.JSON(http.StatusOK, favourites)
}
