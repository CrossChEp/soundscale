package photo_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type UploadPFPModel struct {
	UserId string
	Photo  string `json:"photo"`
}

func (model *UploadPFPModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request to model. Details: %v", err))
		return err
	}
	return nil
}
