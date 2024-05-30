package playlist_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type AddCurrentPlaylistModel struct {
	UserId string
	Songs  []string `json:"songs"`
}

func (model *AddCurrentPlaylistModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't bind request to struct. Details: %v", err))
		return err
	}
	return nil
}
