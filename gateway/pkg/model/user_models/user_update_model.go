package user_models

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type UserUpdateModel struct {
	UserId      string
	Nickname    string `json:"nickname"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (model *UserUpdateModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return err
	}
	return nil
}
