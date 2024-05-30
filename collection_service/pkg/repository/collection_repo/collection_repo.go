package collection_repo

import (
	"collection_service/pkg/config/global_vars_config"
	"collection_service/pkg/model/collection_model"
	"collection_service/pkg/service/logger"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitCollection(userId string) (*mongo.InsertOneResult, error) {
	newCollection := collection_model.CollectionAddModel{
		UserId:        userId,
		Songs:         []string{},
		Playlists:     []string{},
		Albums:        []string{},
		Genres:        []string{},
		CreatedGenres: []string{},
		Followed:      []string{},
		Subscribed:    []string{},
	}
	res, err := global_vars_config.DbCollection.InsertOne(global_vars_config.DbContext, newCollection)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't init collection for user with id %s. Details: %v", userId, err))
		return nil, err
	}
	return res, nil
}

func AddPlaylists(userId string, playlists []string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	collection.Playlists = append(collection.Playlists, playlists...)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "playlists", Value: collection.Playlists},
		}}}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func AddAlbums(userId string, albums []string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	collection.Playlists = append(collection.Albums, albums...)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "albums", Value: collection.Playlists},
		}}}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func AddSongs(userId string, songs []string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	collection.Songs = append(collection.Songs, songs...)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "songs", Value: collection.Songs},
		}},
	}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func AddGenres(userId string, genres []string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	genres = removeExistingElements(genres, collection.Genres)
	collection.Genres = append(collection.Genres, genres...)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "genres", Value: collection.Genres},
		}},
	}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func AddCreatedGenres(userId string, genres []string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	genres = removeExistingElements(genres, collection.CreatedGenres)
	collection.CreatedGenres = append(collection.CreatedGenres, genres...)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "created_genres", Value: collection.CreatedGenres},
		}},
	}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func RemoveFromPlaylists(userId string, playlists []string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	collection.Playlists = removeFromArr(playlists, collection.Playlists)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "playlists", Value: collection.Playlists},
		}},
	}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func RemoveFromAlbums(userId string, albums []string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	collection.Albums = removeFromArr(albums, collection.Albums)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "albums", Value: collection.Albums},
		}},
	}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func RemoveSongs(userId string, songs []string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	collection.Songs = removeFromArr(songs, collection.Songs)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "songs", Value: collection.Songs},
		}},
	}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func AddFollowing(userId string, musicianId string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	collection.Followed = append(collection.Followed, musicianId)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "followed", Value: collection.Followed},
		}},
	}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func RemoveFollowing(userId string, musicianId string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	collection.Followed = remove(musicianId, collection.Followed)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "followed", Value: collection.Followed},
		}},
	}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func UpdateSubscribed(userId string, collection *collection_model.CollectionGetModel) error {
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "subscribed", Value: collection.Subscribed},
			{Key: "followed", Value: collection.Followed},
		}},
	}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func RemoveSubscribed(userId string, musicianId string) error {
	collection, err := GetUserCollection(userId)
	if err != nil {
		return err
	}
	collection.Subscribed = remove(musicianId, collection.Subscribed)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "subscribed", Value: collection.Subscribed},
		}},
	}
	if err := updateRecord(filter, update); err != nil {
		return err
	}
	return nil
}

func updateRecord(filter bson.D, update bson.D) error {
	_, err := global_vars_config.DbCollection.UpdateOne(global_vars_config.DbContext, filter, update)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update record. Details: %v", err))
		return err
	}
	return nil
}

func removeExistingElements(arr []string, arr2 []string) []string {
	var newArr []string
	for _, el := range arr {
		if !isElementInArr(el, arr2) {
			newArr = append(newArr, el)
		}
	}
	return newArr
}

func removeFromArr(elements []string, arr []string) []string {
	var newArr []string
	for _, element := range elements {
		newArr = remove(element, arr)
	}
	return newArr
}

func remove(element string, arr []string) []string {
	var newArr []string
	flag := false
	for _, el := range arr {
		if el == element && !flag {
			flag = true
			continue
		}
		newArr = append(newArr, el)
	}
	return newArr
}
func isElementInArr(element string, arr []string) bool {
	for _, el := range arr {
		if el == element {
			return true
		}
	}
	return false
}

func Get(collId string) (*collection_model.CollectionGetModel, error) {
	oid, err := primitive.ObjectIDFromHex(collId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert id %s to object id. Details: %v", collId, err))
		return nil, err
	}
	collection, err := GetByObjectId(oid)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func GetByObjectId(oid primitive.ObjectID) (*collection_model.CollectionGetModel, error) {
	filter := bson.D{{"_id", oid}}
	var collection collection_model.CollectionGetModel
	cursor := global_vars_config.DbCollection.FindOne(global_vars_config.DbContext, filter)
	if err := cursor.Decode(&collection); err != nil {
		logger.ErrorLog(fmt.Sprintf("Could't convert record to struct. Details: %v", err))
		return nil, err
	}
	return &collection, nil
}

func GetUserCollection(userId string) (*collection_model.CollectionGetModel, error) {
	filter := bson.D{{Key: "user_id", Value: userId}}
	cursor := global_vars_config.DbCollection.FindOne(global_vars_config.DbContext, filter)
	var collection collection_model.CollectionGetModel
	if err := cursor.Decode(&collection); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode record to struct. Details: %v", err))
		return nil, err
	}
	return &collection, nil
}

func GetCollectionsByCreatedGenres(genres []string) ([]collection_model.CollectionGetModel, error) {
	filter := bson.D{
		{Key: "created_genres", Value: bson.D{
			{Key: "$in", Value: genres},
		}},
	}
	cursor, err := global_vars_config.DbCollection.Find(global_vars_config.DbContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get collections by genres. Details: %v", err))
		return nil, err
	}
	var collections []collection_model.CollectionGetModel
	if err := cursor.All(global_vars_config.DbContext, &collections); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: coudln't decode cursor to collection model. Details: %v", err))
		return nil, err
	}
	return collections, nil
}
