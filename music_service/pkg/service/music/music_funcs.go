package music

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"music_service/pkg/config/constants"
	"music_service/pkg/model"
	"music_service/pkg/proto/music_service_proto"
	"music_service/pkg/proto/user_service_proto"
	"music_service/pkg/repository"
	"music_service/pkg/repository/redis_repo"
	"music_service/pkg/service/auxiliary"
	"music_service/pkg/service/checker"
	"music_service/pkg/service/grpc_funcs"
	"music_service/pkg/service/logger"
)

func Add(r *music_service_proto.AddReq) (*model.Song, error) {
	user, err := grpc_funcs.GetUserById(r.AuthorId)
	if err != nil {
		return nil, err
	}
	if isUserInFeaturing(user.Id, r.Collaborators) {
		logger.ErrorLog("Error: author can't be in featuring!")
		return nil, errors.New("author can't be in featuring")
	}
	result, err := repository.Save(user.Id, r)
	if err != nil {
		return nil, err
	}
	if err := updateUserState(user, r.Collaborators); err != nil {
		return nil, err
	}
	if err := addGenreToMusicianGenres(user, r.Genre); err != nil {
		return nil, err
	}
	song, err := repository.GetByObjectId(result.InsertedID.(primitive.ObjectID))
	if err != nil {
		return nil, err
	}
	return song, nil
}

func updateUserState(user *user_service_proto.GetResponse, collaborators []string) error {
	if err := updateState(user); err != nil {
		return err
	}
	for _, collaborator := range collaborators {
		collab, err := grpc_funcs.GetUserById(collaborator)
		if err != nil {
			continue
		}
		if err := updateState(collab); err != nil {
			return err
		}
	}
	return nil
}

func addGenreToMusicianGenres(user *user_service_proto.GetResponse, genre string) error {
	if _, err := grpc_funcs.AddCreatedGenres(user.Id, []string{genre}); err != nil {
		return err
	}
	return nil
}

func updateState(user *user_service_proto.GetResponse) error {
	if user.State != constants.AdminState && user.State != constants.MusicianState {
		_, err := grpc_funcs.UpdateUser(&model.UserUpdateModel{
			State:  constants.MusicianState,
			UserId: user.Id,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetBySinger(r *music_service_proto.GetBySingerReq) ([]*model.Song, error) {
	songs, err := repository.GetByAuthor(r.SingerId)
	if err != nil {
		return nil, err
	}
	if songs == nil {
		songs = []*model.Song{}
	}
	collabSongs, err := repository.GetByAuthorFeaturing(r.SingerId)
	if err != nil {
		return nil, err
	}
	if collabSongs != nil {
		songs = append(songs, collabSongs...)
	}
	return songs, nil
}

func Update(r *music_service_proto.UpdateReq) (*model.Song, error) {

	if !isUserExists(r.AuthorId) {
		logger.ErrorLog("Error: user id is invalid")
		return nil, errors.New("user id is invalid")
	}
	song, err := repository.GetById(r.SongId)
	if err != nil {
		return nil, err
	}
	if song.AuthorId != r.AuthorId {
		logger.ErrorLog("Error: user is not an author of this song")
		return nil, errors.New("user is not an author of this song")
	}
	if r.Collaborators != nil && isUserInFeaturing(r.AuthorId, r.Collaborators) {
		logger.ErrorLog("Error: author can't be in featuring")
		return nil, errors.New("author can't be in featuring")
	}
	if err = updateSongData(song.Id, r); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update song. Details: %v", err))
		return nil, err
	}
	song, err = repository.GetById(song.Id)
	if err != nil {
		return nil, err
	}
	return song, nil
}

func Delete(r *music_service_proto.DeleteReq) error {
	if !isUserExists(r.AuthorId) {
		logger.ErrorLog("Error: user with such id doesn't exists")
		return errors.New("user with such id doesn't exist")
	}
	song, err := repository.GetById(r.SongId)
	if err != nil {
		return err
	}
	if song.AuthorId != r.AuthorId {
		logger.ErrorLog("Error: user is not an author of this song!")
		return errors.New("user is not an author of this song")
	}
	oid, _ := primitive.ObjectIDFromHex(r.SongId)
	if err := repository.Delete(oid); err != nil {
		return err
	}
	return nil
}

func isUserExists(userId string) bool {
	userResp, err := grpc_funcs.GetUserById(userId)
	if err != nil {
		return false
	}
	if userResp.Error != "" {
		logger.ErrorLog("Error: couldn't get user with such claims")
		return false
	}
	return true
}

func isUserInFeaturing(uid string, authors []string) bool {
	for _, aid := range authors {
		if aid == uid {
			return true
		}
	}
	return false
}

func updateSongData(songId string, r *music_service_proto.UpdateReq) error {
	oid, err := primitive.ObjectIDFromHex(songId)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", oid}}
	update := bson.D{{"$set", parseSong(r)}}
	err = repository.Update(filter, update)
	if err != nil {
		return err
	}
	return nil
}

func parseSong(r *music_service_proto.UpdateReq) *bson.D {
	var updateData bson.D
	if r.SongName != "" {
		updateBson(&updateData, "name", r.SongName)
	}
	if r.Collaborators != nil {
		updateBson(&updateData, "collaborators", r.Collaborators)
	}
	return &updateData
}

func updateBson(data *bson.D, key string, value interface{}) {
	*data = append(*data, bson.E{Key: key, Value: value})
}

func AddUserToPlayed(r *music_service_proto.IncreasePlayedQuantityReq) (*model.Song, error) {
	song, err := repository.GetById(r.SongId)
	if err != nil {
		return nil, err
	}
	if song.AuthorId == r.UserId {
		return song, nil
	}
	if err := repository.AddUserToPlayed(r.UserId, song); err != nil {
		return nil, err
	}
	song, err = repository.GetById(song.Id)
	if err != nil {
		return nil, err
	}
	if err := updateTrendCoefficient(song); err != nil {
		return nil, err
	}
	return repository.GetById(song.Id)
}

func AddUserToListeners(r *music_service_proto.IncreaseListenedQuantityReq) (*model.Song, error) {
	song, err := repository.GetById(r.SongId)
	if err != nil {
		return nil, err
	}
	if isUserInListeners(r.UserId, song) {
		return song, nil
	}
	if err := repository.AddUserToListened(r.UserId, song); err != nil {
		return nil, err
	}
	song, err = repository.GetById(song.Id)
	if err != nil {
		return nil, err
	}
	if err := updateTrendCoefficient(song); err != nil {
		return nil, err
	}
	return repository.GetById(song.Id)
}

func isUserInListeners(userId string, song *model.Song) bool {
	for _, user := range song.Listened {
		if user == userId {
			return true
		}
	}
	return false
}

func AddUserToLiked(r *music_service_proto.LikeSongReq) (*model.Song, error) {
	song, err := repository.GetById(r.SongId)
	if err != nil {
		return nil, err
	}
	if checker.IsUserInArray(r.UserId, song.Liked) {
		logger.ErrorLog(fmt.Sprintf("Error: user already liked this song"))
		return song, nil
	}
	if checker.IsUserInArray(r.UserId, song.Disliked) {
		song.Disliked = auxiliary.DeleteUserFromArray(r.UserId, song.Disliked)
	}
	if err := repository.Like(r.UserId, song); err != nil {
		return nil, err
	}
	song, err = repository.GetById(r.SongId)
	if err != nil {
		return nil, err
	}
	if err := updateTrendCoefficient(song); err != nil {
		return nil, err
	}
	return repository.GetById(r.SongId)
}

func AddUserToDisliked(r *music_service_proto.DislikeSongReq) (*model.Song, error) {
	song, err := repository.GetById(r.SongId)
	if err != nil {
		return nil, err
	}
	if checker.IsUserInArray(r.UserId, song.Disliked) {
		logger.ErrorLog(fmt.Sprintf("Error: user already disliked this song"))
		return song, nil
	}
	if checker.IsUserInArray(r.UserId, song.Liked) {
		song.Liked = auxiliary.DeleteUserFromArray(r.UserId, song.Liked)
	}
	if err := repository.Dislike(r.UserId, song); err != nil {
		return nil, err
	}
	song, err = repository.GetById(r.SongId)
	if err != nil {
		return nil, err
	}
	if err := updateTrendCoefficient(song); err != nil {
		return nil, err
	}
	return repository.GetById(r.SongId)
}

func GetUserLiked(r *music_service_proto.GetLikedSongsReq) ([]*model.Song, error) {
	return repository.GetLikedSongs(r.UserId)
}

func GetUserDisliked(r *music_service_proto.GetDislikedSongsReq) ([]*model.Song, error) {
	return repository.GetDislikedSongs(r.UserId)
}

func GetTrends() ([]model.Song, error) {
	trends, err := redis_repo.GetTrends()
	if err != nil {
		if err := setTrends(); err != nil {
			return nil, err
		}
		trends, err = GetTrends()
		if err != nil {
			return nil, err
		}
	}
	return trends, nil
}

func setTrends() error {
	songs, err := repository.GetSongsSortedByCoefficient()
	if err != nil {
		return err
	}
	if err := redis_repo.SetTrends(songs); err != nil {
		return err
	}
	return nil
}

func updateTrendCoefficient(song *model.Song) error {
	var updateData bson.D
	updateBson(&updateData, "trend_coefficient", getTrendCoefficient(song))
	update := bson.D{
		{Key: "$set", Value: updateData},
	}
	oid, err := primitive.ObjectIDFromHex(song.Id)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get object id from id %s. Details: %v", song.Id, err))
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	if err := repository.Update(filter, update); err != nil {
		return err
	}
	return nil
}

func getTrendCoefficient(song *model.Song) float64 {
	return 1 / (1 + math.Pow(math.E, getTrendBasis(song)))
}

func getTrendBasis(song *model.Song) float64 {
	songPlayed := len(song.Played)
	if songPlayed == 0 {
		songPlayed = 1
	}
	return (float64(len(song.Liked)) + float64(len(song.Disliked)) +
		float64(len(song.Listened)) + math.Log(float64(songPlayed))) / 2
}
