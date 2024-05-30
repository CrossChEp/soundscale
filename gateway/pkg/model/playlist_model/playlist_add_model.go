package playlist_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type PlaylistAddModel struct {
	Name     string `json:"name"`
	AuthorId string
	Songs    []string `json:"songs"`
	Type     string   `json:"type"`
}

func (model *PlaylistAddModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't bind request to struct. Details: %v", err))
		return err
	}
	return nil
}
