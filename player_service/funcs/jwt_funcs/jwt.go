package jwt_funcs

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"player_service/config"
	"player_service/funcs/logger"
	"player_service/models"
	"strings"
)

func DecodeJwt(token string) (*models.ClaimsModel, error) {
	claims := models.ClaimsModel{}
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) { return config.PublicKey, nil })
	if err != nil && strings.Contains(err.Error(), "token is expired") {
		logger.ErrorLog(fmt.Sprintf("Error: %v", err))
		return nil, err
	}
	return &claims, nil
}
