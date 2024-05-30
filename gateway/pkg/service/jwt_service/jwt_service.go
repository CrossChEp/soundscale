package jwt_service

import (
	"errors"
	"fmt"
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/model/token_model"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
)

func GetUserClaims(ctx *gin.Context) (*token_model.ClaimsModel, error) {
	token, err := getTokenFromRequest(ctx)
	if err != nil {
		return nil, err
	}
	claims, err := DecodeJwt(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func getTokenFromRequest(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		logger.ErrorLog("Couldn't get token as authorization header is empty")
		return "", errors.New("can't get token as auth header is empty")
	}
	return strings.Split(header, " ")[1], nil
}

func DecodeJwt(token string) (*token_model.ClaimsModel, error) {
	dir, _ := os.Getwd()
	claims := token_model.ClaimsModel{}
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) { return global_vars_config.PublicKey, nil })
	/*if err != nil && !strings.Contains(err.Error(), "token is expired") {
		logger_config.ErrorConsoleLogger.Error("token has expired")
		return nil, err
	}*/
	if err != nil {
		logger.ErrorWithDebugLog("Error: couldn't decode token.", err, dir)
		return nil, errors.New(fmt.Sprintf("Error: couldn't decode token. Details: %v", err))
	}
	return &claims, nil
}
