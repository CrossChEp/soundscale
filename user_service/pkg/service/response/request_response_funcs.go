package response

import (
	"user_service/pkg/model/user_models"
	"user_service/pkg/proto/user_service_proto"
)

func CreateGetResponse(user *user_models.UserModel) *user_service_proto.GetResponse {
	return &user_service_proto.GetResponse{
		Id:          user.Id,
		Nickname:    user.Nickname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		State:       user.State,
	}
}

func CreateGetPrivateResponse(user *user_models.UserModel) *user_service_proto.GetPrivateResponse {
	return &user_service_proto.GetPrivateResponse{
		Id:          user.Id,
		Nickname:    user.Nickname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		State:       user.State,
	}
}

func CreateErrorGetResponse(error string) *user_service_proto.GetResponse {
	return &user_service_proto.GetResponse{
		Error: error,
	}
}

func CreateErrorGetPrivateResponse(error string) *user_service_proto.GetPrivateResponse {
	return &user_service_proto.GetPrivateResponse{
		Error: error,
	}
}
