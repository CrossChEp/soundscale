package user_models

import (
	"github.com/gin-gonic/gin"
)

type UserRegisterModel struct {
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (model *UserRegisterModel) ToModel(ctx *gin.Context) {
	model.Nickname = ctx.Request.PostFormValue("nickname")
	model.Email = ctx.Request.PostFormValue("email")
	model.PhoneNumber = ctx.Request.PostFormValue("phone_number")
	model.Password = ctx.Request.PostFormValue("password")
}
