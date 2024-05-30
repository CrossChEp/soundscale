package token_models

import "github.com/golang-jwt/jwt"

type ClaimsModel struct {
	jwt.StandardClaims
	Id       string
	Username string
}
