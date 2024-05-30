package photo_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type UploadPlaylistCoverModel struct {
	UserId string
	Id     string `json:"id"`
	File   string `json:"file"`
}

func (model *UploadPlaylistCoverModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return err
	}
	return nil
}
