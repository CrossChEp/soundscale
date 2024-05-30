package playlist_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type DeleteSongFromPlaylistModel struct {
	AddSongsToPlaylistModel
}

func (model *DeleteSongFromPlaylistModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return err
	}
	return nil
}
