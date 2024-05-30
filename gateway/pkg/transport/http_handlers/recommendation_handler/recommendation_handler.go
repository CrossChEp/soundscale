package recommendation_handler

import (
	"fmt"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/service/jwt_service"
	"gateway/pkg/service/logger"
	"gateway/pkg/service/recommendation_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Init() {
	global_vars_config.Router.Group("recommendation")
	{
		global_vars_config.Router.GET("/api/recommendation", getRecommendation)
	}
}

func getRecommendation(ctx *gin.Context) {
	dir, _ := os.Getwd()
	claims, err := jwt_service.GetUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("Error: couldn't authorize user. Details: %v", err))
		return
	}
	recommendation, err := recommendation_service.GetRecommendation(claims.Id)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't get recommendation for user %s.", claims.Id), err, dir)
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error: couldn't get recommendation for user %s. Details: %v", claims.Id, err))
	}
	ctx.JSON(http.StatusOK, recommendation)
}
