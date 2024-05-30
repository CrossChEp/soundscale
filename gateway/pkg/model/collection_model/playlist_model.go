package collection_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type PlaylistModel struct {
	UserId       string
	PlaylistsIds []string `json:"playlists"`
}

func (model *PlaylistModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request collection add playlist struct. Details: %v", err))
		return err
	}
	return nil
}
