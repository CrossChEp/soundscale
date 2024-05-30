package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"user_service/pkg/config/constants"
	"user_service/pkg/config/global_vars_config"
	"user_service/pkg/config/service_addresses_config"
	"user_service/pkg/model/user_models"
	"user_service/pkg/proto/user_service_proto"
	"user_service/pkg/service/logger"
)

func Save(request *user_service_proto.AddRequest) (*mongo.InsertOneResult, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't hash password. Details: %v", err))
		return nil, err
	}
	result, err := global_vars_config.DBCollection.InsertOne(
		global_vars_config.DBContext,
		bson.D{
			{Key: "nickname", Value: request.Nickname},
			{Key: "email", Value: request.Email},
			{Key: "phone_number", Value: request.PhoneNumber},
			{Key: "password", Value: string(hashedPassword)},
			{Key: "state", Value: constants.ListenerState},
		},
	)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't save user. Details: %v", err))
		return nil, err
	}
	return result, nil
}

func GetById(userId string) (*user_models.UserModel, error) {
	objectId, _ := primitive.ObjectIDFromHex(userId)
	return GetByObjectId(objectId)
}

func GetByObjectId(objectId primitive.ObjectID) (*user_models.UserModel, error) {
	collection := global_vars_config.DBClient.Database(*service_addresses_config.DBName).Collection("users")
	filter := bson.M{"_id": objectId}
	cursor := collection.FindOne(global_vars_config.DBContext, filter)
	var user user_models.UserModel
	err := cursor.Decode(&user)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couln't get user by id. Details: %v", err))
		return nil, err
	}
	return &user, nil
}

func GetByNickname(nickname string) (*user_models.UserModel, error) {
	filter := bson.M{"nickname": nickname}
	cursor := global_vars_config.DBCollection.FindOne(global_vars_config.DBContext, filter)
	var user user_models.UserModel
	err := cursor.Decode(&user)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user by nickname. Details: %v", err))
		return nil, err
	}
	return &user, nil
}

func Update(filter bson.D, updateData bson.D) error {
	_, err := global_vars_config.DBCollection.UpdateOne(global_vars_config.DBContext, filter, updateData)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update user. Details: %v", err))
		return err
	}
	return nil
}

func Delete(filter bson.D) error {
	_, err := global_vars_config.DBCollection.DeleteOne(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete user. Details: %v", err))
		return err
	}
	return nil
}
