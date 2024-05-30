package photo_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type UploadSongCoverModel struct {
	UserId string
	SongId string `json:"song_id"`
	File   string `json:"file"`
}

func (model *UploadSongCoverModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request to model. Details: %v", err))
		return err
	}
	return nil
}
