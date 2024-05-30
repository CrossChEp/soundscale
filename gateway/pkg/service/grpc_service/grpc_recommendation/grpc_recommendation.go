package grpc_recommendation

import (
	"context"
	"fmt"
	"gateway/pkg/config/service_address_config"
	"gateway/pkg/proto/recommendation_service_proto"
	"gateway/pkg/service/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetRecommendation(userId string) (*recommendation_service_proto.GetRecommendationResp, error) {
	conn, err := grpc.Dial(*service_address_config.RecommendationServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to recommendation service. Details: %v", err))
		return nil, err
	}
	recommendationService := recommendation_service_proto.NewRecommendationServiceClient(conn)
	req := &recommendation_service_proto.GetRecommendationReq{UserId: userId}
	resp, err := recommendationService.GetRecommendation(context.TODO(), req)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get recommendations to user %s. Details: %v", userId, err))
		return nil, err
	}
	return resp, nil
}
