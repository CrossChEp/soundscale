package response

import (
	"collection_service/pkg/model/collection_model"
	"collection_service/pkg/proto/collection_service_proto"
)

func ToCollectionsResponse(collections []collection_model.CollectionGetModel) *collection_service_proto.GetCollectionsResp {
	var collectionResps []*collection_service_proto.GetFavouritesResp
	for _, collection := range collections {
		collectionResps = append(collectionResps, ToFavouriteResponse(&collection))
	}
	return &collection_service_proto.GetCollectionsResp{Collections: collectionResps}
}

func ToFavouriteResponse(collection *collection_model.CollectionGetModel) *collection_service_proto.GetFavouritesResp {
	return &collection_service_proto.GetFavouritesResp{
		Id:            collection.Id,
		UserId:        collection.UserId,
		Songs:         collection.Songs,
		Playlists:     collection.Playlists,
		Albums:        collection.Albums,
		Genres:        collection.Genres,
		CreatedGenres: collection.CreatedGenres,
		Subscribed:    collection.Subscribed,
		Followed:      collection.Followed,
	}
}
