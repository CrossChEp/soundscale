package collection_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type SongsModel struct {
	UserId   string
	SongsIds []string `json:"songs"`
}

func (model *SongsModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request collection add song struct. Details: %v", err))
		return err
	}
	return nil
}
