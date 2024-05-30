package grpc_service

import (
	"context"
	"fmt"
	"gateway/pkg/config/service_address_config"
	"gateway/pkg/model/auth_model"
	"gateway/pkg/proto/token_service_proto"

	"gateway/pkg/service/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Authenticate(credentials *auth_model.CredentialModel) (*token_service_proto.Token, error) {
	conn, err := grpc.Dial(*service_address_config.TokenServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog("Error: couldn't connect to user service")
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	tokenService := token_service_proto.NewTokenServiceClient(conn)
	req := &token_service_proto.Credentials{
		Login:    credentials.Nickname,
		Password: credentials.Password,
	}
	resp, err := tokenService.GetToken(context.Background(), req)
	if err != nil {
		logger.ErrorLog("Error: couldn't authenticate user")
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return resp, nil
}

func RefreshToken(token string) (*token_service_proto.Token, error) {
	conn, err := grpc.Dial(*service_address_config.TokenServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to token service. Details: %v", err))
		return nil, err
	}
	tokenService := token_service_proto.NewTokenServiceClient(conn)
	req := &token_service_proto.Token{Token: token}
	resp, err := tokenService.RefreshToken(context.Background(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't refresh token. Details: %v", err))
		return nil, err
	}
	return resp, nil
}

func IsTokenExpired(token string) (*token_service_proto.IsExpiredResponse, error) {
	conn, err := grpc.Dial(*service_address_config.TokenServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to token service"))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	tokenService := token_service_proto.NewTokenServiceClient(conn)
	req := &token_service_proto.Token{Token: token}
	resp, err := tokenService.IsTokenExpired(context.Background(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't check token"))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return resp, nil
}
