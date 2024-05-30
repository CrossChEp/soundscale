package collection_model

import (
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
	"os"
)

type GenresModel struct {
	UserId string
	Genres []string `json:"genres"`
}

func (m *GenresModel) ToModel(ctx *gin.Context) error {
	dir, _ := os.Getwd()
	if err := ctx.Bind(m); err != nil {
		logger.ErrorWithDebugLog("Error: couldn't bind genres request to model", err, dir)
		return err
	}
	return nil
}
