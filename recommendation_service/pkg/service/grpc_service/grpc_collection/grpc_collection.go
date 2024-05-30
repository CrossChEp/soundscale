package grpc_collection

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"recommendation_service/pkg/config/service_address_config"
	"recommendation_service/pkg/proto/collection_service_proto"
	"recommendation_service/pkg/service/logger"
)

func GetCollection(userId string) (*collection_service_proto.GetFavouritesResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	getReq := &collection_service_proto.GetCollectionReq{UserId: userId}
	response, err := collectionService.GetCollection(context.TODO(), getReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user %s grpc_collection. Details: %d", userId, err))
		return nil, err
	}
	return response, nil
}

func GetCollectionsByCreatedGenres(genres []string) (*collection_service_proto.GetCollectionsResp, error) {
	conn, err := grpc.Dial(*service_address_config.CollectionServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't connect to grpc_collection service. Details: %v", err))
		return nil, err
	}
	collectionService := collection_service_proto.NewCollectionServiceClient(conn)
	getReq := &collection_service_proto.GetCollectionsByCreatedGenresReq{
		Genres: genres,
	}
	response, err := collectionService.GetCollectionByCreatedGenres(context.TODO(), getReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get collection by genres %v. Details: %d", genres, err))
		return nil, err
	}
	return response, nil
}
