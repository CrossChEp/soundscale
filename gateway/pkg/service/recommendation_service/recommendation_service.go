package recommendation_service

import (
	"gateway/pkg/model/recommendation_model"
	"gateway/pkg/service/grpc_service/grpc_recommendation"
)

func GetRecommendation(userId string) (*recommendation_model.RecommendationGetModel, error) {
	recRep, err := grpc_recommendation.GetRecommendation(userId)
	if err != nil {
		return nil, err
	}
	recommendation := &recommendation_model.RecommendationGetModel{}
	recommendation.ToModel(userId, recRep)
	return recommendation, nil
}
