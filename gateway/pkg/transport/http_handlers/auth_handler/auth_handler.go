package auth_handler

import (
	"fmt"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/model/auth_model"
	"gateway/pkg/service/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	global_vars_config.Router.Group("auth")
	{
		global_vars_config.Router.POST("/api/auth", authenticate)
		global_vars_config.Router.PUT("/api/token/refresh/:token", refreshToken)
	}
}

func authenticate(ctx *gin.Context) {
	credentials := &auth_model.CredentialModel{}
	credentials.ToModel(ctx)
	token, err := auth_service.Authenticate(credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	ctx.JSON(http.StatusOK, token)
}

func refreshToken(ctx *gin.Context) {
	token := ctx.Param("token")
	newToken, err := auth_service.RefreshToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	ctx.JSON(http.StatusOK, newToken)
}
