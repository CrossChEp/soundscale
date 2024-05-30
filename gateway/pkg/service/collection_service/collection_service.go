package collection_service

import (
	"gateway/pkg/model/collection_model"
	"gateway/pkg/service/grpc_service/grpc_collection"
)

func AddSongs(userId string, songs []string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.AddSongs(userId, songs)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func RemoveSongs(userId string, songs []string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.RemoveSongs(userId, songs)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func AddPlaylists(userId string, playlists []string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.AddPlaylists(userId, playlists)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func AddGenres(genresModel *collection_model.GenresModel) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.AddGenres(genresModel)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func RemovePlaylists(userId string, playlists []string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.RemovePlaylists(userId, playlists)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func AddAlbums(userId string, albums []string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.AddAlbums(userId, albums)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func RemoveAlbums(userId string, albums []string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.RemoveAlbums(userId, albums)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func GetCollection(userId string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.GetCollection(userId)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func Follow(userId string, musicianId string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.Follow(userId, musicianId)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func Unfollow(userId string, musicianId string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.Unfollow(userId, musicianId)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func Subscribe(userId string, musicianId string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.Subscribe(userId, musicianId)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}

func Unsubscribe(userId string, musicianId string) (*collection_model.CollectionGetModel, error) {
	collectionResp, err := grpc_collection.Unsubscribe(userId, musicianId)
	if err != nil {
		return nil, err
	}
	collection := &collection_model.CollectionGetModel{}
	collection.ToModel(collectionResp)
	return collection, nil
}
