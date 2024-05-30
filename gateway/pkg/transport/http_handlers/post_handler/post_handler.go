package post_handler

import (
	"fmt"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/model/post_model"
	"gateway/pkg/service/jwt_service"
	"gateway/pkg/service/logger"
	"gateway/pkg/service/post_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	global_vars_config.Router.Group("Posts")
	{
		global_vars_config.Router.POST("/api/post", addPost)
		global_vars_config.Router.GET("/api/post/:postId", getPost)
		global_vars_config.Router.GET("/api/posts/user/:userId", getUserPosts)
		global_vars_config.Router.PUT("/api/post/:postId", updatePost)
		global_vars_config.Router.DELETE("/api/post/:postId", deletePost)
		global_vars_config.Router.POST("/api/post/like/:post_id", likePost)
		global_vars_config.Router.POST("/api/post/dislike/:post_id", dislikePost)
		global_vars_config.Router.GET("/api/posts/liked", getUserLikedPosts)
		global_vars_config.Router.GET("/api/posts/disliked", getUserDislikedPosts)
	}
}

func addPost(ctx *gin.Context) {
	logger.InfoLog(fmt.Sprintf("%v handler was called. Request IP: %v", ctx.Request.URL, ctx.ClientIP()))
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("%v Error: couldn't authorize user. Details: %v", ctx.Request.URL, err))
		logger.InfoLog(fmt.Sprintf("%v handler sent error data to IP: %v", ctx.Request.URL, ctx.ClientIP()))
		return
	}
	postModel := &post_model.PostAddModel{}
	if err := postModel.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%v Error: couldn't parse request data. Details: %v", ctx.Request.URL, err))
		logger.InfoLog(fmt.Sprintf("%v handler sent error data to IP: %v", ctx.Request.URL, ctx.ClientIP()))
		return
	}
	postModel.AuthorId = claims.Id
	post, err := post_service.AddPost(postModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%v Error: couldn't add post to user %s profile. Details: %v", ctx.Request.URL, claims.Id, err))
		logger.InfoLog(fmt.Sprintf("%v handler sent error data to IP: %v", ctx.Request.URL, ctx.ClientIP()))
		return
	}
	ctx.JSON(http.StatusOK, post)
	logger.InfoLog(fmt.Sprintf("%v handler sent response to IP %v", ctx.Request.URL, ctx.ClientIP()))
}

func getPost(ctx *gin.Context) {
	logger.InfoLog(fmt.Sprintf("%v handler was called. Request IP: %v", ctx.Request.URL, ctx.ClientIP()))
	postId := ctx.Param("postId")
	post, err := post_service.GetPost(postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%v Error: couldn't get post %s. Details: %v", ctx.Request.URL, postId, err))
		logger.InfoLog(fmt.Sprintf("%v handler sent error data to IP: %v", ctx.Request.URL, ctx.ClientIP()))
		return
	}
	ctx.JSON(http.StatusOK, post)
	logger.InfoLog(fmt.Sprintf("%v handler sent response to IP %v", ctx.Request.URL, ctx.ClientIP()))
}

func getUserPosts(ctx *gin.Context) {
	userId := ctx.Param("userId")
	posts, err := post_service.GetUserPosts(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%v Error: couldn't get post of user %s. Details: %v", ctx.Request.URL, userId, err))
		logger.InfoLog(fmt.Sprintf("%v handler sent error data to IP: %v", ctx.Request.URL, ctx.ClientIP()))
		return
	}
	ctx.JSON(http.StatusOK, posts)
	logger.InfoLog(fmt.Sprintf("%v handler sent response to IP %v", ctx.Request.URL, ctx.ClientIP()))
}

func updatePost(ctx *gin.Context) {
	logger.InfoLog(fmt.Sprintf("%v handler was called. Request IP: %v", ctx.Request.URL, ctx.ClientIP()))
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("%v Error: couldn't authorize user. Details: %v", ctx.Request.URL, err))
		logger.InfoLog(fmt.Sprintf("%v handler sent error data to IP: %v", ctx.Request.URL, ctx.ClientIP()))
		return
	}
	postId := ctx.Param("postId")
	updateData := &post_model.PostUpdateModel{}
	if err := updateData.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%v Error: couldn't parse request data. Details: %v", ctx.Request.URL, err))
		logger.InfoLog(fmt.Sprintf("%v handler sent error data to IP: %v", ctx.Request.URL, ctx.ClientIP()))
		return
	}
	updateData.AuthorId = claims.Id
	updateData.PostId = postId
	post, err := post_service.UpdatePost(updateData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%v Error: couldn't update post %s. Details: %v", ctx.Request.URL, postId, err))
		logger.InfoLog(fmt.Sprintf("%v handler sent error data to IP: %v", ctx.Request.URL, ctx.ClientIP()))
		return
	}
	ctx.JSON(http.StatusOK, post)
	logger.InfoLog(fmt.Sprintf("%v handler sent response to IP %v", ctx.Request.URL, ctx.ClientIP()))
}

func deletePost(ctx *gin.Context) {
	logger.InfoLog(fmt.Sprintf("%v handler was called. Request IP: %v", ctx.Request.URL, ctx.ClientIP()))
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("%v Error: couldn't authorize user. Details: %v", ctx.Request.URL, err))
		logger.InfoLog(fmt.Sprintf("%v handler sent error data to IP: %v", ctx.Request.URL, ctx.ClientIP()))
		return
	}
	postId := ctx.Param("postId")
	if err := post_service.DeletePost(postId, claims.Id); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%v Error: couldn't delete post %s. Details: %v", ctx.Request.URL, postId, err))
		logger.InfoLog(fmt.Sprintf("%v handler sent error data to IP: %v", ctx.Request.URL, ctx.ClientIP()))
		return
	}
	ctx.JSON(http.StatusOK, "post was deleted")
	logger.InfoLog(fmt.Sprintf("%v handler sent response to IP %v", ctx.Request.URL, ctx.ClientIP()))
}

func likePost(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	postId := ctx.Param("post_id")
	post, err := post_service.LikePost(claims.Id, postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't like post %s by user %s. Details: %v", postId, claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, post)
	logger.InfoLog("response data was sent")
}

func dislikePost(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	postId := ctx.Param("post_id")
	post, err := post_service.DislikePost(claims.Id, postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't dislike post %s by user %s. Details: %v", postId, claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, post)
	logger.InfoLog("response data was sent")
}

func getUserLikedPosts(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	posts, err := post_service.GetUserLikedPosts(claims.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get liked posts of user %s. Details: %v", claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, posts)
	logger.InfoLog("response data was sent")
}

func getUserDislikedPosts(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	posts, err := post_service.GetUserDislikedPosts(claims.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get disliked posts of user %s. Details: %v", claims.Id, err))
		logger.InfoLog("error response was sent")
		return
	}
	ctx.JSON(http.StatusOK, posts)
	logger.InfoLog("response data was sent")
}
