package transport

import (
	"collection_service/pkg/proto/collection_service_proto"
	"collection_service/pkg/service/collection_service"
	"collection_service/pkg/service/response"
	"context"
	"errors"
	"fmt"
)

type CollectionService struct {
	collection_service_proto.CollectionServiceServer
}

func (s *CollectionService) InitCollection(_ context.Context,
	r *collection_service_proto.InitReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.Init(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't initialize collection. Details: %v", err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) AddPlaylists(_ context.Context,
	r *collection_service_proto.PlaylistsReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.AddPlaylists(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: couldn't add playlists to user %s collection. Details: %v", r.UserId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) AddAlbums(_ context.Context,
	r *collection_service_proto.AlbumsReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.AddAlbums(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't add albums to user %s collection. Details: %v", r.UserId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) AddSongs(_ context.Context,
	r *collection_service_proto.SongsReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.AddSongs(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't add songs to user %s collection. Details: %v", r.UserId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) AddGenres(_ context.Context,
	r *collection_service_proto.AddGenresReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.AddGenres(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't add genres to user %s collection. Details: %v", r.UserId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) AddCreatedGenres(_ context.Context,
	r *collection_service_proto.AddCreatedGenresReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.AddCreatedGenres(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't add created genres to user %s collection. Details: %v", r.UserId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) RemovePlaylists(_ context.Context,
	r *collection_service_proto.PlaylistsReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.RemovePlaylists(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't remove playlists of user %s collection. Details: %v", r.UserId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) RemoveAlbums(_ context.Context,
	r *collection_service_proto.AlbumsReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.RemoveAlbums(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't remove albums from user %s collection. Details: %v", r.UserId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) RemoveSongs(_ context.Context,
	r *collection_service_proto.SongsReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.RemoveSongs(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't remove songs from user %s collection. Details: %v", r.UserId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) GetCollection(_ context.Context,
	r *collection_service_proto.GetCollectionReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.GetCollection(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't get collection of user %s. Details: %v", r.UserId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) GetCollectionByCreatedGenres(_ context.Context,
	r *collection_service_proto.GetCollectionsByCreatedGenresReq) (*collection_service_proto.GetCollectionsResp, error) {
	collections, err := collection_service.GetCollectionsByCreatedGenre(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't get collection by created genres. Details: %v", err))
	}
	return response.ToCollectionsResponse(collections), nil
}

func (s *CollectionService) Follow(_ context.Context,
	r *collection_service_proto.FollowReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.Follow(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't follow by user %s on musician %s. Details: %v", r.UserId, r.MusicianId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) Unfollow(_ context.Context,
	r *collection_service_proto.UnfollowReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.Unfollow(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't unfollow by user %s on musician %s. Details: %v", r.UserId, r.MusicianId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) Subscribe(_ context.Context,
	r *collection_service_proto.SubscribeReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.Subscribe(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't subscribe by user %s on musician %s. Details: %v", r.UserId, r.MusicianId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}

func (s *CollectionService) Unsubscribe(_ context.Context,
	r *collection_service_proto.UnsubscribeReq) (*collection_service_proto.GetFavouritesResp, error) {
	collection, err := collection_service.Unsubscribe(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't unsubscribe by user %s on musician %s. Details: %v", r.UserId, r.MusicianId, err))
	}
	return response.ToFavouriteResponse(collection), nil
}
