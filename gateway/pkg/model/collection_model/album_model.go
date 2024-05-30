package collection_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
)

type AlbumModel struct {
	UserId    string
	AlbumsIds []string `json:"albums"`
}

func (model *AlbumModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(model); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert request collection add album struct. Details: %v", err))
		return err
	}
	return nil
}
