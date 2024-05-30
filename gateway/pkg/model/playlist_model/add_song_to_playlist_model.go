package playlist_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type AddSongsToPlaylistModel struct {
	AuthorId     string
	PlaylistId   string   `json:"playlist_id"`
	SongIds      []string `json:"songs"`
	PlaylistType string
}

func (model *AddSongsToPlaylistModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return err
	}
	return nil
}
