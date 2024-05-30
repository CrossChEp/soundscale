package recommendation_model

import (
	"gateway/pkg/model/music_model"
	"gateway/pkg/model/post_model"
	"gateway/pkg/model/user_models"
	"gateway/pkg/proto/recommendation_service_proto"
	"gateway/pkg/service/auxilary"
	"gateway/pkg/service/collection_service"
	"gateway/pkg/service/music_service"
	"gateway/pkg/service/post_service"
	"gateway/pkg/service/user_service"
)

type RecommendationGetModel struct {
	Songs     music_model.SongsGetModel   `json:"songs"`
	Musicians []*user_models.UserGetModel `json:"musicians"`
	Posts     post_model.PostsGetModel    `json:"posts"`
}

func (m *RecommendationGetModel) ToModel(userId string, resp *recommendation_service_proto.GetRecommendationResp) {
	m.Songs = music_model.SongsGetModel{Songs: getSongs(userId, resp)}
	m.Musicians = getMusicians(resp)
	m.Posts = post_model.PostsGetModel{Posts: getPosts(resp)}
}

func getSongs(userId string, resp *recommendation_service_proto.GetRecommendationResp) []*music_model.SongGetModel {
	var songs []*music_model.SongGetModel
	userCollection, err := collection_service.GetCollection(userId)
	if err != nil {
		return nil
	}
	for _, songId := range resp.Songs {
		song, err := music_service.GetSongById(userId, songId)
		if song.Exclusive && song.AuthorId != userCollection.UserId &&
			!auxilary.IsElementInArr(song.AuthorId, userCollection.Subscribed) {
			continue
		}
		if err == nil {
			songs = append(songs, song)
		}
	}
	return songs
}

func getMusicians(resp *recommendation_service_proto.GetRecommendationResp) []*user_models.UserGetModel {
	var musicians []*user_models.UserGetModel
	for _, musicianId := range resp.Musicians {
		musician, err := user_service.GetById(musicianId)
		if err == nil {
			musicians = append(musicians, musician)
		}
	}
	return musicians
}

func getPosts(resp *recommendation_service_proto.GetRecommendationResp) []post_model.PostGetModel {
	var posts []post_model.PostGetModel
	for _, postId := range resp.Posts {
		post, err := post_service.GetPost(postId)
		if err == nil {
			posts = append(posts, *post)
		}
	}
	return posts
}
