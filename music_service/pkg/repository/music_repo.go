package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"music_service/pkg/config/global_vars_config"
	"music_service/pkg/model"
	"music_service/pkg/proto/music_service_proto"
	"music_service/pkg/service/logger"
	"time"
)

func Save(authorId string, r *music_service_proto.AddReq) (*mongo.InsertOneResult, error) {
	releaseDate := time.Now()
	result, err := global_vars_config.DBCollection.InsertOne(
		global_vars_config.DBContext,
		bson.D{
			{Key: "name", Value: r.SongName},
			{Key: "genre", Value: r.Genre},
			{Key: "author", Value: authorId},
			{Key: "collaborators", Value: r.Collaborators},
			{Key: "release_date", Value: releaseDate},
			{Key: "played", Value: []string{}},
			{Key: "listened", Value: []string{}},
			{Key: "liked", Value: []string{}},
			{Key: "disliked", Value: []string{}},
			{Key: "trend_coefficient", Value: 0},
			{Key: "exclusive", Value: r.Exclusive},
		},
	)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(Save) Error: couldn't add song to database. Details: %v", err))
		return nil, err
	}
	return result, nil
}

func GetById(id string) (*model.Song, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	song, err := GetByObjectId(oid)
	if err != nil {
		return nil, err
	}
	return song, nil
}

func GetByObjectId(id primitive.ObjectID) (*model.Song, error) {
	filter := bson.D{{"_id", id}}
	cursor := global_vars_config.DBCollection.FindOne(global_vars_config.DBContext, filter)
	var song model.Song
	err := cursor.Decode(&song)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(GetById) Error: couldn't decode music. Details: %v", err))
		return nil, err
	}
	return &song, nil
}

func GetSongsSortedByCoefficient() ([]model.Song, error) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "trend_coefficient", Value: -1}})
	cursor, err := global_vars_config.DBCollection.Find(global_vars_config.DBContext, filter, opts)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get sorted songs. Details: %v", err))
		return nil, err
	}
	var songs []model.Song
	if err := cursor.All(global_vars_config.DBContext, &songs); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode sorted songs. Details: %v", err))
		return nil, err
	}
	return songs, nil
}

func GetByAuthor(authorId string) ([]*model.Song, error) {
	filter := bson.D{{"author", authorId}}
	var songs []*model.Song
	cursor, err := global_vars_config.DBCollection.Find(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(GetByAuthor)Error: couldn't get song. Details: %v", err))
		return nil, err
	}
	if err := cursor.All(global_vars_config.DBContext, &songs); err != nil {
		logger.ErrorLog(fmt.Sprintf("(GetByAuthor)Error: couldn't decode song. Details: %v", err))
		return nil, err
	}
	return songs, nil
}

func GetByAuthorFeaturing(authorId string) ([]*model.Song, error) {
	filter := bson.D{{"collaborators", bson.D{{"$all", bson.A{authorId}}}}}
	cursor, err := global_vars_config.DBCollection.Find(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(GetByAuthorFeaturing)Error: couldn't get song: invalid parameters. Details: %v", err))
		return nil, err
	}
	var songs []*model.Song
	if err := cursor.All(global_vars_config.DBContext, &songs); err != nil {
		logger.ErrorLog(fmt.Sprintf("(GetByAuthorFeaturing)Error: couldn't decode song. Details: %v", err))
		return nil, err
	}
	return songs, nil
}

func Update(filter bson.D, update bson.D) error {
	_, err := global_vars_config.DBCollection.UpdateOne(global_vars_config.DBContext, filter, update)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(Update): error couldn't. Details: %v", err))
		return err
	}
	return nil
}

func Delete(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := global_vars_config.DBCollection.DeleteOne(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(Delete)Error: couldn't delete user. Details: %v", err))
		return err
	}
	return nil
}

func AddUserToPlayed(userId string, song *model.Song) error {
	song.Played = append(song.Played, userId)
	updateData := model.MusicUpdateFieldModel{
		SongId: song.Id,
		Field:  "played",
		Value:  song.Played,
	}
	if err := updateSongField(updateData); err != nil {
		return err
	}
	return nil
}

func AddUserToListened(userId string, song *model.Song) error {
	song.Listened = append(song.Listened, userId)
	updateData := model.MusicUpdateFieldModel{
		SongId: song.Id,
		Field:  "listened",
		Value:  song.Listened,
	}
	if err := updateSongField(updateData); err != nil {
		return err
	}
	return nil
}

func Like(userId string, song *model.Song) error {
	song.Liked = append(song.Liked, userId)
	return updateSongLikedData(song)
}

func Dislike(userId string, song *model.Song) error {
	song.Disliked = append(song.Disliked, userId)
	return updateSongLikedData(song)
}

func updateSongLikedData(song *model.Song) error {
	likeSongData := model.MusicUpdateFieldModel{
		SongId: song.Id,
		Field:  "liked",
		Value:  song.Liked,
	}
	dislikeSongData := model.MusicUpdateFieldModel{
		SongId: song.Id,
		Field:  "disliked",
		Value:  song.Disliked,
	}
	if err := updateSongField(dislikeSongData); err != nil {
		return err
	}
	return updateSongField(likeSongData)
}

func updateSongField(updateData model.MusicUpdateFieldModel) error {
	oid, err := primitive.ObjectIDFromHex(updateData.SongId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert id %s to object id. Details: %v", updateData.SongId, err))
		return err
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: updateData.Field, Value: updateData.Value},
		}},
	}
	_, err = global_vars_config.DBCollection.UpdateOne(global_vars_config.DBContext, filter, update)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update record. Details: %v", err))
		return err
	}
	return nil
}

func GetLikedSongs(userId string) ([]*model.Song, error) {
	filter := bson.D{
		{Key: "liked", Value: bson.D{
			{Key: "$all", Value: bson.A{userId}},
		}},
	}
	return getSongs(filter)
}

func GetDislikedSongs(userId string) ([]*model.Song, error) {
	filter := bson.D{
		{Key: "disliked", Value: bson.D{
			{Key: "$all", Value: bson.A{userId}},
		}},
	}
	return getSongs(filter)
}

func getSongs(filter bson.D) ([]*model.Song, error) {
	cursor, err := global_vars_config.DBCollection.Find(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't find songs by given filter. Details: %v", err))
		return nil, err
	}
	var songs []*model.Song
	if err := cursor.All(global_vars_config.DBContext, &songs); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error couldn't decode songs. Details: %v", err))
		return nil, err
	}
	return songs, nil
}
