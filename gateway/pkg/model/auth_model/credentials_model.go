package auth_model

import (
	"github.com/gin-gonic/gin"
)

type CredentialModel struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (model *CredentialModel) ToModel(ctx *gin.Context) {
	model.Nickname = ctx.Request.PostFormValue("nickname")
	model.Password = ctx.Request.PostFormValue("password")
}
