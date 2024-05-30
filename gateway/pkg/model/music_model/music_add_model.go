package music_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type MusicAddModel struct {
	AuthorId      string
	Collaborators []string `json:"collaborators"`
	SongName      string   `json:"song_name"`
	Genre         string   `json:"genre"`
	Exclusive     bool     `json:"exclusive"`
}

func (model *MusicAddModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't bind request to struct. Details: %v", err))
		return err
	}
	return nil
}
