package funcs

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"user_service/pkg/config/constants"
	"user_service/pkg/model/user_models"
	"user_service/pkg/proto/user_service_proto"
	"user_service/pkg/repository"
	"user_service/pkg/service/auxilary"
	"user_service/pkg/service/logger"
)

func AddUser(request *user_service_proto.AddRequest) (*user_models.UserModel, error) {
	user, _ := repository.GetByNickname(request.Nickname)
	if user != nil {
		logger.ErrorLog("Error: user with such nickname already exists")
		return nil, errors.New("user with such nickname already exists")
	}
	result, err := repository.Save(request)
	if err != nil {
		return nil, err
	}
	return getResultUser(result), nil
}

func getResultUser(result *mongo.InsertOneResult) *user_models.UserModel {
	oid, _ := result.InsertedID.(primitive.ObjectID)
	user, err := repository.GetByObjectId(oid)
	if err != nil {
		return nil
	}
	return user
}

func Update(request *user_service_proto.UpdateRequest) (*user_models.UserModel, error) {
	user, err := repository.GetByNickname(request.Nickname)
	if request.Nickname != "" && user.Nickname == request.Nickname {
		logger.ErrorLog(fmt.Sprintf("Error: can't update user as user with such nickname already exists."))
		return nil, errors.New("user with such nickname already exists")
	}
	objectId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get object id from user id. Details: %v", err))
		return nil, err
	}
	if request.UserState != "" && !auxilary.IsElInArr(request.UserState, constants.States) {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't change user state as state %s is not valid", request.UserState))
		return nil, errors.New(fmt.Sprintf("Error: couldn't change user state as state %s is not valid", request.UserState))
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	update := parseRequest(request)
	if err := repository.Update(filter, bson.D{{Key: "$set", Value: update}}); err != nil {
		return nil, err
	}
	user, err = repository.GetById(request.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func parseRequest(request *user_service_proto.UpdateRequest) bson.D {
	var update bson.D
	if request.Nickname != "" {
		updateBson(&update, "nickname", request.Nickname)
	}
	if request.PhoneNumber != "" {
		updateBson(&update, "phone_number", request.PhoneNumber)
	}
	if request.Email != "" {
		updateBson(&update, "email", request.Email)
	}
	if request.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		updateBson(&update, "password", string(hashedPassword))
	}
	if request.UserState != "" {
		updateBson(&update, "state", request.UserState)
	}
	return update
}

func updateBson(data *bson.D, key string, value interface{}) {
	*data = append(*data, bson.E{Key: key, Value: value})
}

func Delete(request *user_service_proto.DeleteRequest) (*user_models.UserModel, error) {
	user, err := repository.GetById(request.UserId)
	if err != nil {
		return nil, err
	}
	objectId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Couldn't convert user id to object id. Details: %v", err))
		return nil, err
	}

	if err := uDelete(objectId); err != nil {
		return nil, err
	}
	return user, nil
}

func uDelete(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	err := repository.Delete(filter)
	if err != nil {
		return err
	}
	return nil
}
