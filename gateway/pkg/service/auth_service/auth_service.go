package auth_service

import (
	"gateway/pkg/model/auth_model"
	grpc_service "gateway/pkg/service/grpc_service/auth"
)

func Authenticate(credentials *auth_model.CredentialModel) (string, error) {
	token, err := grpc_service.Authenticate(credentials)
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

func RefreshToken(token string) (string, error) {
	newToken, err := grpc_service.RefreshToken(token)
	if err != nil {
		return "", err
	}
	return newToken.Token, nil
}

func IsTokenExpired(token string) (bool, error) {
	check, err := grpc_service.IsTokenExpired(token)
	if err != nil {
		return true, err
	}
	return check.IsExpired, nil
}
