package playlist_repo

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"playlist_serivce/pkg/config/constants"
	"playlist_serivce/pkg/config/global_vars_config"
	"playlist_serivce/pkg/models/playlist_models"
	"playlist_serivce/pkg/proto/playlist_service_proto"
	"playlist_serivce/pkg/service/base_functions"
	"playlist_serivce/pkg/service/logger"
	"time"
)

func Save(r *playlist_service_proto.AddPlaylistReq) (interface{}, error) {
	if !base_functions.IsPlaylistTypeExists(r.Type) {
		r.Type = constants.TypePlaylist
	}
	currentTime := time.Now().UTC().String()
	result, err := global_vars_config.DBCollection.InsertOne(
		global_vars_config.DBContext,
		bson.D{
			{Key: "name", Value: r.Name},
			{Key: "author", Value: r.AuthorId},
			{Key: "songs", Value: r.SongIds},
			{Key: "release_date", Value: currentTime},
			{Key: "last_update", Value: currentTime},
			{Key: "type", Value: r.Type},
			{Key: "liked", Value: []string{}},
			{Key: "disliked", Value: []string{}},
		},
	)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't save playlist_service. Details: %v", err))
		return nil, err
	}
	return result.InsertedID, nil
}

func AddToPlaylist(addData playlist_models.AddRemoveSongsModel) error {
	if addData.PlaylistType == constants.TypeAlbum {
		logger.ErrorLog(fmt.Sprintf("Error: can't add song to album"))
		return errors.New("error: can't add song to album")
	}
	playlist, err := Get(addData.PlaylistId, addData.PlaylistType)
	if err != nil {
		return err
	}
	filter := bson.D{
		{Key: "_id", Value: addData.PlaylistId},
		{Key: "type", Value: addData.PlaylistType},
	}
	playlist.Songs = append(playlist.Songs, addData.SongsIds...)
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "songs", Value: playlist.Songs},
		{Key: "last_update", Value: time.Now().UTC().String()},
	}}}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func DeleteFromPlaylist(deleteData playlist_models.AddRemoveSongsModel) error {
	if deleteData.PlaylistType == constants.TypeAlbum {
		logger.ErrorLog(fmt.Sprintf("Error: can't add song to album"))
		return errors.New("error: can't add song to album")
	}
	playlist, err := Get(deleteData.PlaylistId, deleteData.PlaylistType)
	if err != nil {
		return err
	}
	for _, song := range deleteData.SongsIds {
		playlist.Songs = base_functions.DeleteFromArr(playlist.Songs, song)
	}
	filter := bson.D{{Key: "_id", Value: deleteData.PlaylistId}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "songs", Value: playlist.Songs},
		{Key: "last_update", Value: time.Now().UTC().String()},
	}}}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func DeletePlaylist(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := global_vars_config.DBCollection.DeleteOne(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete playlist. Details: %v", err))
		return err
	}
	return nil
}

func LikePlaylist(likeModel *playlist_models.LikeDislikeModel) error {
	oid, err := primitive.ObjectIDFromHex(likeModel.Playlist.Id)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert id %s to object id. Details: %v", likeModel.Playlist.Id, err))
		return err
	}
	likeModel.Playlist.Liked = append(likeModel.Playlist.Liked, likeModel.UserId)
	filter := bson.D{{Key: "_id", Value: oid}}
	return updateLikes(filter, likeModel.Playlist)
}

func DislikedPlaylist(likeModel *playlist_models.LikeDislikeModel) error {
	oid, err := primitive.ObjectIDFromHex(likeModel.Playlist.Id)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert id %s to object id. Details: %v", likeModel.Playlist.Id, err))
		return err
	}
	likeModel.Playlist.Disliked = append(likeModel.Playlist.Liked, likeModel.UserId)
	filter := bson.D{{Key: "_id", Value: oid}}
	return updateLikes(filter, likeModel.Playlist)
}

func updateLikes(filter bson.D, playlist *playlist_models.PlaylistGetModel) error {
	updateData := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "liked", Value: playlist.Liked},
			{Key: "disliked", Value: playlist.Disliked},
		},
		},
	}
	return updateRecord(filter, updateData)
}

func updateRecord(filter bson.D, update bson.D) error {
	_, err := global_vars_config.DBCollection.UpdateOne(global_vars_config.DBContext, filter, update)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update record. Details: %v", err))
		return err
	}
	return nil
}

func Get(id primitive.ObjectID, playlistType string) (*playlist_models.PlaylistGetModel, error) {
	if !base_functions.IsPlaylistTypeExists(playlistType) {
		playlistType = constants.TypePlaylist
	}
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "type", Value: playlistType},
	}
	cursor := global_vars_config.DBCollection.FindOne(global_vars_config.DBContext, filter)
	var playlist playlist_models.PlaylistGetModel
	if err := cursor.Decode(&playlist); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't find playlist_service by id. Details: %v", err))
		return nil, err
	}
	return &playlist, nil
}

func GetByAuthor(authorId string, playlistType string) ([]playlist_models.PlaylistGetModel, error) {
	filter := bson.D{
		{Key: "author", Value: authorId},
	}
	if playlistType == constants.TypePlaylist || playlistType == constants.TypeAlbum {
		filter = append(filter, bson.E{Key: "type", Value: playlistType})
	}
	cursor, err := global_vars_config.DBCollection.Find(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get songs by author id. Details: %v", err))
		return nil, err
	}
	var playlists []playlist_models.PlaylistGetModel
	if err := cursor.All(global_vars_config.DBContext, &playlists); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode songs. Details: %v", err))
		return nil, err
	}
	return playlists, nil
}

func GetUserFavourite(userId string) (*playlist_models.PlaylistGetModel, error) {
	filter := bson.D{
		{Key: "author", Value: userId},
		{Key: "type", Value: constants.TypeFavourite},
	}
	favourite := global_vars_config.DBCollection.FindOne(global_vars_config.DBContext, filter)
	var playlist playlist_models.PlaylistGetModel
	if err := favourite.Decode(&playlist); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert playlist to struct. Details: %v", err))
		return nil, err
	}
	return &playlist, nil
}

func GetUserLiked(userId string, playlistType string) ([]playlist_models.PlaylistGetModel, error) {
	filter := bson.D{
		{Key: "type", Value: playlistType},
		{Key: "liked", Value: bson.D{
			{Key: "$all", Value: bson.A{userId}},
		}},
	}
	return getPlaylists(filter)
}

func GetUserDisliked(userId string, playlistType string) ([]playlist_models.PlaylistGetModel, error) {
	filter := bson.D{
		{Key: "type", Value: playlistType},
		{Key: "disliked", Value: bson.D{
			{Key: "$all", Value: bson.A{userId}},
		}},
	}
	return getPlaylists(filter)
}

func getPlaylists(filter bson.D) ([]playlist_models.PlaylistGetModel, error) {
	cursor, err := global_vars_config.DBCollection.Find(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get playlists. Details: %v", err))
		return nil, err
	}
	var playlists []playlist_models.PlaylistGetModel
	if err := cursor.All(global_vars_config.DBContext, &playlists); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode playlists. Details: %v", err))
		return nil, err
	}
	return playlists, nil
}
