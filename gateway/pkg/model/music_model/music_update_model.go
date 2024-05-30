package music_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type MusicUpdateModel struct {
	SongId string `json:"song_id"`
	MusicAddModel
}

func (model *MusicUpdateModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't bind request to model. Details: %v", err))
		return err
	}
	return nil
}
