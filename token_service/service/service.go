package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"token_service/config"
	"token_service/protos/token_service_proto"
	"token_service/protos/user_service_proto"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
)

type TokenService struct {
	token_service_proto.TokenServiceServer
}

func BuildAddress(service_name string) string {
	service, ok := config.CONFIG.Services[service_name]
	if !ok {
		panic("buildAdress: service name is not in config!")
	}
	return fmt.Sprintf("%s:%d", service.Address, service.Port)
}

func (s *TokenService) GetToken(ctx context.Context, creds *token_service_proto.Credentials) (*token_service_proto.Token, error) {
	if creds.Login == "" || creds.Password == "" {
		return nil, errors.New(fmt.Sprintf("fileds login and password must be not null. Login: %s. Password: %s", creds.Login, creds.Password))
	}
	caller, _ := peer.FromContext(ctx)
	config.Logger.Info(fmt.Sprintf("GetToken: Called by %s", caller.Addr.String()))
	conn, err := grpc.Dial(BuildAddress("user_service"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		config.Logger.Errorln("Connection error! Couldn't connect to UserService!")
	}
	defer conn.Close()
	userService := user_service_proto.NewUserServiceClient(conn)
	data := user_service_proto.GetByNicknameRequest{
		Nickname: creds.Login,
	}
	config.Logger.Info("Getting user...")
	user, err := userService.GetByNicknamePrivate(ctx, &data)
	if err != nil {
		config.Logger.Errorln(err)
		return nil, err
	}
	config.Logger.Info("User is gotten!")
	config.Logger.Info("Validating credentials...")
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))

	if err != nil {
		config.Logger.WithFields(logrus.Fields{
			"login":    creds.Login,
			"password": creds.Password,
		}).Warning("Failed to validate credentials!")
		return nil, err
	}
	config.Logger.Info("Credentials are valid")
	config.Logger.Info("Generating token...")
	token, err := generateToken(user)
	if err != nil {
		config.Logger.WithError(err).Error("Failed to generate token!")
		config.Logger.Debugf("uid:%d, username:%d", user.Id, user.Nickname)
		return nil, err
	}
	config.Logger.WithField("username", user.Nickname).Info("Token was generated!")
	return &token_service_proto.Token{
		Token: token,
	}, nil
}

func (s *TokenService) IsTokenExpired(ctx context.Context, msg *token_service_proto.Token) (*token_service_proto.IsExpiredResponse, error) {
	if msg.Token == "" {
		return nil, errors.New("token must be not null")
	}
	caller, _ := peer.FromContext(ctx)
	config.Logger.Infof("IsTokenExpired: Called by %s", caller.Addr)
	_, err := jwt.Parse(msg.Token, func(t *jwt.Token) (interface{}, error) { return config.PUBLIC_KEY, nil })
	if err != nil {
		config.Logger.WithField("error", err).Error("Token validation error!")
		return &token_service_proto.IsExpiredResponse{IsExpired: true}, nil
	}
	config.Logger.Info("Token was successfully validated")
	return &token_service_proto.IsExpiredResponse{IsExpired: false}, nil
}

func (s *TokenService) RefreshToken(ctx context.Context, msg *token_service_proto.Token) (*token_service_proto.Token, error) {
	if msg.Token == "" {
		return nil, errors.New("token must be not null")
	}
	caller, _ := peer.FromContext(ctx)
	config.Logger.Infof("RefreshToken: Called by %s", caller.Addr)
	claims := TokenClaims{}
	_, err := jwt.ParseWithClaims(msg.Token, &claims, func(t *jwt.Token) (interface{}, error) { return config.PUBLIC_KEY, nil })
	if err != nil && !strings.Contains(err.Error(), "token is expired") {
		config.Logger.WithError(err).Error("Failed to validate token!")
		return nil, err
	}
	claims.ExpiresAt = time.Now().Add(config.EXPIRATION_TIME).Unix()
	token := *jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token_str, err := token.SignedString(config.PRIVATE_KEY)
	if err != nil {
		config.Logger.WithError(err).Error("Failed to sign token!")
		return nil, err
	}
	return &token_service_proto.Token{Token: token_str}, nil
}
