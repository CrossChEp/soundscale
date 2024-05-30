package service

import (
	"time"
	"token_service/config"
	"token_service/protos/user_service_proto"

	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// Generates a token from user's name and id
//
// Returns token if it was generated successfully or errror if not
func generateToken(user *user_service_proto.GetPrivateResponse) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)

	claims := TokenClaims{
		Id:       user.Id,
		Username: user.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.EXPIRATION_TIME).Unix(),
		},
	}
	token.Claims = claims
	ecdsaKey := config.PRIVATE_KEY
	token_str, err := token.SignedString(ecdsaKey)
	if err != nil {
		return "", err
	}
	return token_str, err
}
