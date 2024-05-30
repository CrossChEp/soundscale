package user_service

import (
	"errors"
	"fmt"
	"gateway/pkg/model/user_models"
	"gateway/pkg/service/collection_service"
	"gateway/pkg/service/grpc_service/grpc_collection"
	user_grpc_service "gateway/pkg/service/grpc_service/user"
	"gateway/pkg/service/logger"
	"gateway/pkg/service/post_service"
	"net/mail"
	"os"
)

func GetById(userId string) (*user_models.UserGetModel, error) {
	userResp, err := user_grpc_service.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	user := &user_models.UserGetModel{}
	user.ConvertToModel(userResp)
	addUserRelationsToUserModel(user)
	return user, nil
}

func GetByNickname(nickname string) (*user_models.UserGetModel, error) {
	userResp, err := user_grpc_service.GetUserByNickname(nickname)
	if err != nil {
		return nil, err
	}
	user := &user_models.UserGetModel{}
	user.ConvertToModel(userResp)
	addUserRelationsToUserModel(user)
	return user, nil
}

func addUserRelationsToUserModel(user *user_models.UserGetModel) {
	addPostsToModel(user)
	addCollectionToModel(user)
}

func addPostsToModel(user *user_models.UserGetModel) {
	posts, err := post_service.GetUserPosts(user.Id)
	if err == nil {
		user.Posts = posts
	}
}

func addCollectionToModel(user *user_models.UserGetModel) {
	collection, err := collection_service.GetCollection(user.Id)
	if err == nil {
		user.Collection = collection
	}
}

func Register(userData *user_models.UserRegisterModel) (*user_models.UserGetPrivateModel, error) {
	if err := validateUserRegisterData(userData); err != nil {
		return nil, err
	}
	user, err := user_grpc_service.AddUser(userData)
	if err != nil {
		return nil, err
	}
	if _, err := grpc_collection.InitCollection(user.Id); err != nil {
		return nil, err
	}
	return &user_models.UserGetPrivateModel{
		UserGetModel: user_models.UserGetModel{
			Id:          user.Id,
			Nickname:    user.Nickname,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		},
		Password: user.Password,
	}, nil
}

func validateUserRegisterData(userData *user_models.UserRegisterModel) error {
	dir, _ := os.Getwd()
	if err := checkUserEmailAndNumber(userData); err != nil {
		return err
	}
	if userData.Email != "" {
		if _, err := mail.ParseAddress(userData.Email); err != nil {
			logger.ErrorWithDebugLog(fmt.Sprintf("Error: user email %s is not valid.", userData.Email), err, dir)
			return err
		}
	}
	if userData.PhoneNumber != "" {
		if !isPhoneNumberValid(userData.PhoneNumber) {
			err := errors.New(fmt.Sprintf("user phone number %s is not valid", userData.PhoneNumber))
			logger.ErrorWithDebugLog(fmt.Sprintf("Error: user phone number %s is not valid.", userData.PhoneNumber), err, dir)
			return err
		}
	}
	return nil
}

func checkUserEmailAndNumber(userData *user_models.UserRegisterModel) error {
	dir, _ := os.Getwd()
	if userData.PhoneNumber == "" && userData.Email == "" {
		err := errors.New(fmt.Sprintf("Error: user has no phone number or email. Email: %s. Phone number :%s", userData.Email, userData.PhoneNumber))
		logger.ErrorWithDebugLog(
			fmt.Sprintf("Error: user has no phone number or email. Email: %s. Phone number :%s", userData.Email, userData.PhoneNumber),
			err, dir)
		return errors.New("")
	}
	return nil
}

func isPhoneNumberValid(phoneNumber string) bool {
	if phoneNumber[:2] == "+7" && len(phoneNumber[2:]) == 10 {
		return true
	}
	return false
}

func Update(newUserData *user_models.UserUpdateModel) (*user_models.UserGetModel, error) {
	resp, err := user_grpc_service.UpdateUser(newUserData)
	if err != nil {
		return nil, err
	}
	user := user_models.UserGetModel{}
	user.ConvertToModel(resp)
	return &user, nil
}

func Delete(userId string) (*user_models.UserGetModel, error) {
	resp, err := user_grpc_service.DeleteUser(userId)
	if err != nil {
		return nil, err
	}
	userModel := &user_models.UserGetModel{}
	userModel.ConvertToModel(resp)
	return userModel, nil
}
