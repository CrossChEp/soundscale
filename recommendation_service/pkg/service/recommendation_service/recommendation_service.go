package recommendation_service

import (
	"recommendation_service/pkg/model/recommendation_model"
	"recommendation_service/pkg/proto/collection_service_proto"
	"recommendation_service/pkg/proto/recommendation_service_proto"
	"recommendation_service/pkg/service/grpc_service/grpc_collection"
	"recommendation_service/pkg/service/grpc_service/grpc_music"
	"recommendation_service/pkg/service/grpc_service/grpc_post"
	"recommendation_service/pkg/service/grpc_service/grpc_user"
)

func GetRecommendation(r *recommendation_service_proto.GetRecommendationReq) (*recommendation_model.RecommendationGetModel, error) {
	user, err := grpc_user.GetUserById(r.UserId)
	if err != nil {
		return nil, err
	}
	collection, err := grpc_collection.GetCollection(user.Id)
	if err != nil {
		return nil, err
	}
	collectionsOfMusicians, err := grpc_collection.GetCollectionsByCreatedGenres(collection.Genres)
	if err != nil {
		return nil, err
	}
	return getData(collection, collectionsOfMusicians), nil
}

func getData(userCollection *collection_service_proto.GetFavouritesResp, collections *collection_service_proto.GetCollectionsResp) *recommendation_model.RecommendationGetModel {
	recommendation := &recommendation_model.RecommendationGetModel{}
	recommendation = getRecommendationData(recommendation, userCollection.Subscribed)
	recommendation = getRecommendationData(recommendation, userCollection.Followed)
	for _, collection := range collections.Collections {
		songs, err := getSongs(collection.UserId)
		if err != nil {
			continue
		}
		posts, err := getPosts(collection.UserId)
		if err != nil {
			continue
		}
		recommendation.Musicians = append(recommendation.Musicians, collection.UserId)
		recommendation.Songs = append(recommendation.Songs, songs...)
		recommendation.Posts = append(recommendation.Posts, posts...)
	}
	return recommendation
}

func getRecommendationData(recommendationModel *recommendation_model.RecommendationGetModel,
	arr []string) *recommendation_model.RecommendationGetModel {
	for _, musicianId := range arr {
		posts, err := getPosts(musicianId)
		if err != nil {
			continue
		}
		songs, err := getSongs(musicianId)
		if err != nil {
			continue
		}
		recommendationModel.Posts = append(recommendationModel.Posts, posts...)
		recommendationModel.Songs = append(recommendationModel.Songs, songs...)
	}
	return recommendationModel
}

func getSongs(singerId string) ([]string, error) {
	var songIds []string
	songs, err := grpc_music.GetBySinger(singerId)
	if err != nil {
		return nil, err
	}
	for _, song := range songs.Songs {
		songIds = append(songIds, song.Id)
	}
	return songIds, nil
}

func getPosts(userId string) ([]string, error) {
	var postIds []string
	posts, err := grpc_post.GetUserPosts(userId)
	if err != nil {
		return nil, err
	}
	for _, post := range posts.Posts {
		postIds = append(postIds, post.Id)
	}
	return postIds, nil
}
