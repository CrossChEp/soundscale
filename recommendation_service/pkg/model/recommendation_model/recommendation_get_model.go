package recommendation_model

import "recommendation_service/pkg/proto/recommendation_service_proto"

type RecommendationGetModel struct {
	Musicians []string
	Posts     []string
	Songs     []string
}

func (m *RecommendationGetModel) ToResp() *recommendation_service_proto.GetRecommendationResp {
	resp := &recommendation_service_proto.GetRecommendationResp{}
	resp.Musicians = m.Musicians
	resp.Posts = m.Posts
	resp.Songs = m.Songs
	return resp
}
