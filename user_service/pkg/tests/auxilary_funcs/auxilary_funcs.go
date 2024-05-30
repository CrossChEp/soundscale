package tests

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"user_service/pkg/config/global_vars_config"
	"user_service/pkg/model/user_models"
	"user_service/pkg/proto/user_service_proto"
	"user_service/pkg/repository"
	"user_service/pkg/service/logger"
)

type AddTest struct {
	Arg, Arg2, Expected interface{}
}

var TestUser = &user_models.UserModel{
	Nickname:    "test",
	Email:       "test",
	PhoneNumber: "test",
	Password:    "test",
}

func GenerateSaveUserData() []AddTest {
	return []AddTest{
		{
			Arg: user_service_proto.AddRequest{
				Nickname:    TestUser.Nickname,
				Email:       TestUser.Email,
				PhoneNumber: TestUser.PhoneNumber,
				Password:    TestUser.Password,
			},
			Arg2:     nil,
			Expected: nil,
		},
	}
}

func CreateTestUser() error {
	user, err := repository.GetByNickname("test")
	if user != nil {
		TestUser.Id = user.Id
		TestUser.Nickname = user.Nickname
		TestUser.Email = user.Email
		TestUser.PhoneNumber = user.PhoneNumber
		return nil
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(TestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't hash user's password. Details: %v", err))
		return err
	}
	_, err = global_vars_config.DBCollection.InsertOne(
		global_vars_config.DBContext,
		bson.D{
			{Key: "nickname", Value: TestUser.Nickname},
			{Key: "email", Value: TestUser.Email},
			{Key: "phone_number", Value: TestUser.PhoneNumber},
			{Key: "password", Value: string(hashedPassword)},
		},
	)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't save user. Details: %v", err))
		return nil
	}
	return nil
}

func DeleteTestUser() error {
	objectId, err := primitive.ObjectIDFromHex(TestUser.Id)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't generate object id from user id. Details: %v", err))
		return err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	_, err = global_vars_config.DBCollection.DeleteOne(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete user. Details: %v", err))
		return err
	}
	return nil
}

func GetUserByObjectId(id mongo.InsertOneResult) (*user_models.UserModel, error) {
	oId, _ := id.InsertedID.(primitive.ObjectID)
	filter := bson.D{{Key: "_id", Value: oId}}
	cursor := global_vars_config.DBCollection.FindOne(global_vars_config.DBContext, filter)
	var user *user_models.UserModel
	if err := cursor.Decode(&user); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode user. Details: %v", err))
		return nil, err
	}
	return user, nil
}

func CheckUser(user *user_models.UserModel, expected *user_models.UserModel) bool {
	if user.Id != expected.Id {
		return false
	}
	if user.Email != expected.Email {
		return false
	}
	if user.Nickname != expected.Nickname {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(expected.Password))
	if err != nil {
		return false
	}
	if user.PhoneNumber != expected.PhoneNumber {
		return false
	}
	return true
}

func CheckGetResponse(user *user_models.UserModel, expected *user_service_proto.GetResponse) bool {
	if user.Id != expected.Id {
		return false
	}
	if user.Email != expected.Email {
		return false
	}
	if user.Nickname != expected.Nickname {
		return false
	}
	if user.PhoneNumber != expected.PhoneNumber {
		return false
	}
	return true
}

func CheckGetPrivateResponse(user *user_models.UserModel, expected *user_service_proto.GetPrivateResponse) bool {
	if user.Id != expected.Id {
		return false
	}
	if user.Email != expected.Email {
		return false
	}
	if user.Nickname != expected.Nickname {
		return false
	}
	if user.Password != expected.Password {
		return false
	}
	if user.PhoneNumber != expected.PhoneNumber {
		return false
	}
	return true
}
