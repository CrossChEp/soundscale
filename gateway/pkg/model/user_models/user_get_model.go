package user_models

import (
	"gateway/pkg/model/collection_model"
	"gateway/pkg/model/post_model"
	"gateway/pkg/proto/user_service_proto"
)

type UserGetModel struct {
	Id          string                               `json:"id"`
	Nickname    string                               `json:"nickname"`
	Email       string                               `json:"email"`
	PhoneNumber string                               `json:"phone_number"`
	Collection  *collection_model.CollectionGetModel `json:"collection"`
	Posts       *post_model.PostsGetModel            `json:"posts"`
}

type UserGetPrivateModel struct {
	UserGetModel
	Password string `json:"password"`
}

func (model *UserGetModel) ConvertToModel(response *user_service_proto.GetResponse) {
	model.Id = response.Id
	model.Nickname = response.Nickname
	model.Email = response.Email
	model.PhoneNumber = response.PhoneNumber
}
