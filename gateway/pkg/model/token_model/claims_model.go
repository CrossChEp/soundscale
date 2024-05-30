package token_model

import "github.com/golang-jwt/jwt"

type ClaimsModel struct {
	jwt.StandardClaims
	Id       string
	Nickname string
}
