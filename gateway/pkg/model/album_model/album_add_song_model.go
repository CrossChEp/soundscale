package album_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type AddSongToAlbumModel struct {
	AuthorId   string
	PlaylistId string   `json:"album_id"`
	SongIds    []string `json:"songs"`
}

func (model *AddSongToAlbumModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request to struct. Details: %v", err))
		return err
	}
	return nil
}
