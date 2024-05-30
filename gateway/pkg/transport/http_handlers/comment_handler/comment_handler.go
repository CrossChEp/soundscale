package comment_handler

import (
	"fmt"
	"gateway/pkg/config/constants"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/model/comment_model"
	"gateway/pkg/service/comment_service"
	"gateway/pkg/service/jwt_service"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Init() {
	global_vars_config.Router.Group("comment")
	{
		global_vars_config.Router.POST("/api/comment/post", addPostComment)
		global_vars_config.Router.POST("/api/comment/song", addSongComment)
		global_vars_config.Router.POST("/api/comment/playlist", addPlaylistComment)
		global_vars_config.Router.POST("/api/comment/album", addAlbumComment)
		global_vars_config.Router.GET("/api/comments/post/:id", getPostComments)
		global_vars_config.Router.GET("/api/comments/song/:id", getSongComments)
		global_vars_config.Router.GET("/api/comments/playlist/:id", getPlaylistsComments)
		global_vars_config.Router.GET("/api/comments/album/:id", getAlbumComments)
		global_vars_config.Router.PUT("/api/comment/:id", updateComment)
		global_vars_config.Router.DELETE("/api/comment/:id", deleteComment)
		global_vars_config.Router.GET("/api/comments/user", getUserComments)
	}
}

func addPostComment(ctx *gin.Context) {
	addComment(ctx, constants.PostType)
}

func addSongComment(ctx *gin.Context) {
	addComment(ctx, constants.SongType)
}

func addPlaylistComment(ctx *gin.Context) {
	addComment(ctx, constants.PlaylistType)
}

func addAlbumComment(ctx *gin.Context) {
	addComment(ctx, constants.AlbumType)
}

func addComment(ctx *gin.Context, entityType string) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	commentModel := &comment_model.CommentAddModel{}
	dir, _ := os.Getwd()
	if err := commentModel.ToModel(ctx); err != nil {
		logger.ErrorWithDebugLog("Error: couldn't add comment to post", err, dir)
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add comment to post. Details: %v", err))
	}
	commentModel.UserId = claims.Id
	commentModel.EntityType = entityType
	comment, err := comment_service.AddComment(commentModel)
	if err != nil {
		logger.ErrorWithDebugLog("Error: couldn't add comment to post", err, dir)
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't add comment to post. Details: %v", err))
	}
	ctx.JSON(http.StatusOK, comment)
}

func getPostComments(ctx *gin.Context) {
	getEntityComments(ctx, constants.PostType)
}

func getSongComments(ctx *gin.Context) {
	getEntityComments(ctx, constants.SongType)
}

func getPlaylistsComments(ctx *gin.Context) {
	getEntityComments(ctx, constants.PlaylistType)
}

func getAlbumComments(ctx *gin.Context) {
	getEntityComments(ctx, constants.AlbumType)
}

func getEntityComments(ctx *gin.Context, entityType string) {
	dir, _ := os.Getwd()
	entityId := ctx.Param("id")
	comments, err := comment_service.GetEntityComments(entityId, entityType)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't get %s %s comments", entityType, entityId), err, dir)
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get %s %s comments. Details: %v", entityType, entityId, err))
	}
	ctx.JSON(http.StatusOK, comments)
}

func updateComment(ctx *gin.Context) {
	dir, _ := os.Getwd()
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	commentId := ctx.Param("id")
	commentModel := &comment_model.CommentUpdateModel{}
	if err := commentModel.ToModel(ctx); err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't update comment %s", commentId), err, dir)
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't update comment %s. Details: %v", commentId, err))
	}
	commentModel.UserId, commentModel.CommentId = claims.Id, commentId
	comment, err := comment_service.UpdateComment(commentModel)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't update comment %s", commentId), err, dir)
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't update comment %s. Details: %v", commentId, err))
	}
	ctx.JSON(http.StatusOK, comment)
}

func deleteComment(ctx *gin.Context) {
	dir, _ := os.Getwd()
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	commentId := ctx.Param("id")
	resp, err := comment_service.DeleteComment(commentId, claims.Id)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't delete comment %s", commentId), err, dir)
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't delete comment %s. Details: %v", commentId, err))
	}
	ctx.JSON(http.StatusOK, resp)
}

func getUserComments(ctx *gin.Context) {
	dir, _ := os.Getwd()
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	comments, err := comment_service.GetUserComments(claims.Id)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't get user %s comments", claims.Id), err, dir)
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get user %s comments. Details: %v", claims.Id, err))
	}
	ctx.JSON(http.StatusOK, comments)
}
