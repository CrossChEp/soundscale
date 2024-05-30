package models

import "github.com/dgrijalva/jwt-go"

type ClaimsModel struct {
	jwt.StandardClaims
	Id       string
	Username string
}
