package collection_service

import (
	"collection_service/pkg/config/constants"
	"collection_service/pkg/model/collection_model"
	"collection_service/pkg/proto/collection_service_proto"
	"collection_service/pkg/repository/collection_repo"
	"collection_service/pkg/service/auxiliary"
	"collection_service/pkg/service/checkers"
	"collection_service/pkg/service/grpc_service/grpc_user"
	"collection_service/pkg/service/logger"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Init(r *collection_service_proto.InitReq) (*collection_model.CollectionGetModel, error) {
	if err := checkUser(r.UserId); err != nil {
		return nil, err
	}
	if checkers.IsUserCollectionExists(r.UserId) {
		logger.ErrorLog(fmt.Sprintf("Error: user with id %s already has collection", r.UserId))
		return nil, errors.New(fmt.Sprintf("Error: user with id %s already has collection", r.UserId))
	}
	result, err := collection_repo.InitCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	collection, err := collection_repo.GetByObjectId(oid)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func GetCollection(r *collection_service_proto.GetCollectionReq) (*collection_model.CollectionGetModel, error) {
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func AddPlaylists(r *collection_service_proto.PlaylistsReq) (*collection_model.CollectionGetModel, error) {
	if err := checkUser(r.UserId); err != nil {
		return nil, err
	}
	r.PlaylistIds = auxiliary.RemoveInvalidPlaylists(r.UserId, r.PlaylistIds)
	if err := collection_repo.AddPlaylists(r.UserId, r.PlaylistIds); err != nil {
		return nil, err
	}
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func AddGenres(r *collection_service_proto.AddGenresReq) (*collection_model.CollectionGetModel, error) {
	if err := checkUser(r.UserId); err != nil {
		return nil, err
	}
	if err := collection_repo.AddGenres(r.UserId, r.Genres); err != nil {
		return nil, err
	}
	return collection_repo.GetUserCollection(r.UserId)
}

func AddCreatedGenres(r *collection_service_proto.AddCreatedGenresReq) (*collection_model.CollectionGetModel, error) {
	if err := checkUser(r.UserId); err != nil {
		return nil, err
	}
	if err := collection_repo.AddCreatedGenres(r.UserId, r.Genres); err != nil {
		return nil, err
	}
	return collection_repo.GetUserCollection(r.UserId)
}

func AddAlbums(r *collection_service_proto.AlbumsReq) (*collection_model.CollectionGetModel, error) {
	if err := checkUser(r.UserId); err != nil {
		return nil, err
	}
	r.AlbumIds = auxiliary.RemoveInvalidAlbums(r.UserId, r.AlbumIds)
	if err := collection_repo.AddAlbums(r.UserId, r.AlbumIds); err != nil {
		return nil, err
	}
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func AddSongs(r *collection_service_proto.SongsReq) (*collection_model.CollectionGetModel, error) {
	if err := checkUser(r.UserId); err != nil {
		return nil, err
	}
	r.SongIds = auxiliary.RemoveInvalidSongs(r.SongIds)
	if err := collection_repo.AddSongs(r.UserId, r.SongIds); err != nil {
		return nil, err
	}
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func RemovePlaylists(r *collection_service_proto.PlaylistsReq) (*collection_model.CollectionGetModel, error) {
	if err := checkUser(r.UserId); err != nil {
		return nil, err
	}
	r.PlaylistIds = auxiliary.RemoveInvalidPlaylistsForRemoving(r.UserId, r.PlaylistIds)
	if err := collection_repo.RemoveFromPlaylists(r.UserId, r.PlaylistIds); err != nil {
		return nil, err
	}
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func RemoveAlbums(r *collection_service_proto.AlbumsReq) (*collection_model.CollectionGetModel, error) {
	if err := checkUser(r.UserId); err != nil {
		return nil, err
	}
	r.AlbumIds = auxiliary.RemoveInvalidAlbumsForeRemoving(r.UserId, r.AlbumIds)
	if err := collection_repo.RemoveFromAlbums(r.UserId, r.AlbumIds); err != nil {
		return nil, err
	}
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func RemoveSongs(r *collection_service_proto.SongsReq) (*collection_model.CollectionGetModel, error) {
	if err := checkUser(r.UserId); err != nil {
		return nil, err
	}
	r.SongIds = auxiliary.RemoveInvalidSongs(r.SongIds)
	if err := collection_repo.RemoveSongs(r.UserId, r.SongIds); err != nil {
		return nil, err
	}
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func Follow(r *collection_service_proto.FollowReq) (*collection_model.CollectionGetModel, error) {
	if err := checkFollowRequest(r.UserId, r.MusicianId); err != nil {
		return nil, err
	}
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	if checkers.IsElementInArr(r.MusicianId, collection.Followed) ||
		checkers.IsElementInArr(r.MusicianId, collection.Subscribed) {
		return collection, nil
	}
	if err := collection_repo.AddFollowing(r.UserId, r.MusicianId); err != nil {
		return nil, err
	}
	return collection_repo.GetUserCollection(r.UserId)
}

func Unfollow(r *collection_service_proto.UnfollowReq) (*collection_model.CollectionGetModel, error) {
	if err := checkFollowRequest(r.UserId, r.MusicianId); err != nil {
		return nil, err
	}
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	if !checkers.IsElementInArr(r.MusicianId, collection.Followed) {
		return collection, nil
	}
	if err := collection_repo.RemoveFollowing(r.UserId, r.MusicianId); err != nil {
		return nil, err
	}
	return collection_repo.GetUserCollection(r.UserId)
}

func Subscribe(r *collection_service_proto.SubscribeReq) (*collection_model.CollectionGetModel, error) {
	if err := checkFollowRequest(r.UserId, r.MusicianId); err != nil {
		return nil, err
	}
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	if checkers.IsElementInArr(r.MusicianId, collection.Subscribed) {
		return collection, nil
	}
	collection.Followed = auxiliary.RemoveFromArr(r.MusicianId, collection.Followed)
	collection.Subscribed = append(collection.Subscribed, r.MusicianId)
	if err := collection_repo.UpdateSubscribed(r.UserId, collection); err != nil {
		return nil, err
	}
	return collection_repo.GetUserCollection(r.UserId)
}

func Unsubscribe(r *collection_service_proto.UnsubscribeReq) (*collection_model.CollectionGetModel, error) {
	if err := checkFollowRequest(r.UserId, r.MusicianId); err != nil {
		return nil, err
	}
	collection, err := collection_repo.GetUserCollection(r.UserId)
	if err != nil {
		return nil, err
	}
	if !checkers.IsElementInArr(r.MusicianId, collection.Subscribed) {
		return collection, nil
	}
	collection.Subscribed = auxiliary.RemoveFromArr(r.MusicianId, collection.Subscribed)
	if err := collection_repo.UpdateSubscribed(r.UserId, collection); err != nil {
		return nil, err
	}
	return collection_repo.GetUserCollection(r.UserId)
}

func checkFollowRequest(userId string, musicianId string) error {
	if err := checkUser(userId); err != nil {
		return err
	}
	if userId == musicianId {
		logger.ErrorLog(fmt.Sprintf("Error: user id %s is equal to musician id %s", userId, musicianId))
		return errors.New(fmt.Sprintf("Error: user id %s is equal to musician id %s", userId, musicianId))
	}
	musician, err := grpc_user.GetUserById(musicianId)
	if err != nil {
		return err
	}
	if musician.State == constants.ListenerType {
		logger.ErrorLog(fmt.Sprintf("Error: user %s has no musician state. User state: %s", musician.Id, musician.State))
		return errors.New(fmt.Sprintf("Error: user %s has no musician state. User state: %s", musician.Id, musician.State))
	}
	return nil
}

func GetCollectionsByCreatedGenre(r *collection_service_proto.GetCollectionsByCreatedGenresReq) ([]collection_model.CollectionGetModel, error) {
	return collection_repo.GetCollectionsByCreatedGenres(r.Genres)
}

func checkUser(userId string) error {
	if !checkers.IsUserExists(userId) {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user with id %s.", userId))
		return errors.New(fmt.Sprintf("couldn't get user with id %s.", userId))
	}
	return nil
}
