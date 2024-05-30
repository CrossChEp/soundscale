package user_handler

import (
	"fmt"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/model/user_models"
	"gateway/pkg/service/collection_service"
	"gateway/pkg/service/jwt_service"
	"gateway/pkg/service/logger"
	"gateway/pkg/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	global_vars_config.Router.Group("user")
	{
		global_vars_config.Router.POST("/api/user/register", register)
		global_vars_config.Router.GET("/api/user/:userId", getUser)
		global_vars_config.Router.GET("/api/user/nickname/:nickname", getUserByNickname)
		global_vars_config.Router.PUT("/api/user", update)
		global_vars_config.Router.DELETE("/api/user", deleteUser)
		global_vars_config.Router.POST("/api/user/follow/:musicianId", follow)
		global_vars_config.Router.POST("/api/user/unfollow/:musicianId", unfollow)
		global_vars_config.Router.POST("/api/user/subscribe/:musicianId", subscribe)
		global_vars_config.Router.POST("/api/user/unsubscribe/:musicianId", unsubscribe)
		global_vars_config.Router.GET("/api/user/me/:token", authUser)
	}
}

func getUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, "invalid request")
		return
	}
	user, err := user_service.GetById(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("%v", err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func getUserByNickname(ctx *gin.Context) {
	userId := ctx.Param("nickname")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, "invalid request")
		return
	}
	user, err := user_service.GetByNickname(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("%v", err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func register(ctx *gin.Context) {
	var userData = &user_models.UserRegisterModel{}
	userData.ToModel(ctx)
	user, err := user_service.Register(userData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func update(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Invalid token. Details: %v", err))
		return
	}
	updateModel := &user_models.UserUpdateModel{}
	if err := updateModel.ToModel(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't convert request to model. Details: %v", err))
		return
	}
	updateModel.UserId = claims.Id
	user, err := user_service.Update(updateModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Couldn't update token. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func deleteUser(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Invalid token. Details: %v", err))
		return
	}
	user, err := user_service.Delete(claims.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid token. Details: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func follow(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Invalid token. Details: %v", err))
		return
	}
	musicianId := ctx.Param("musicianId")
	collection, err := collection_service.Follow(claims.Id, musicianId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't follow by %s on %s. Details: %v", claims.Id, musicianId, err))
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't follow by %s on %s. Details: %v", claims.Id, musicianId, err))
	}
	ctx.JSON(http.StatusOK, collection)
}

func unfollow(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Invalid token. Details: %v", err))
		return
	}
	musicianId := ctx.Param("musicianId")
	collection, err := collection_service.Unfollow(claims.Id, musicianId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't unfollow by %s on %s. Details: %v", claims.Id, musicianId, err))
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't unfollow by %s on %s. Details: %v", claims.Id, musicianId, err))
	}
	ctx.JSON(http.StatusOK, collection)
}

func subscribe(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Invalid token. Details: %v", err))
		return
	}
	musicianId := ctx.Param("musicianId")
	collection, err := collection_service.Subscribe(claims.Id, musicianId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't subscribe by %s on %s. Details: %v", claims.Id, musicianId, err))
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't subscribe by %s on %s. Details: %v", claims.Id, musicianId, err))
	}
	ctx.JSON(http.StatusOK, collection)
}

func unsubscribe(ctx *gin.Context) {
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Invalid token. Details: %v", err))
		return
	}
	musicianId := ctx.Param("musicianId")
	collection, err := collection_service.Unsubscribe(claims.Id, musicianId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't unsubscribe by %s on %s. Details: %v", claims.Id, musicianId, err))
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't unsubscribe by %s on %s. Details: %v", claims.Id, musicianId, err))
	}
	ctx.JSON(http.StatusOK, collection)
}

func authUser(ctx *gin.Context) {
	token := ctx.Param("token")
	claims, err := jwt_service.DecodeJwt(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Invalid token. Details: %v", err))
		return
	}
	user, err := user_service.GetById(claims.Id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("%v", err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
