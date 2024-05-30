package transport

import (
	"context"
	"errors"
	"fmt"
	"recommendation_service/pkg/proto/recommendation_service_proto"
	"recommendation_service/pkg/service/recommendation_service"
)

type RecommendationService struct {
	recommendation_service_proto.RecommendationServiceServer
}

func (s *RecommendationService) GetRecommendation(_ context.Context,
	req *recommendation_service_proto.GetRecommendationReq) (*recommendation_service_proto.GetRecommendationResp, error) {
	recommendation, err := recommendation_service.GetRecommendation(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't get recommendation for user %s. Deails: %v", req.UserId, err))
	}
	return recommendation.ToResp(), nil
}
