package post_model

import (
	"fmt"
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
	"os"
)

type PostAddModel struct {
	AuthorId  string
	Content   string   `json:"content"`
	Songs     []string `json:"songs"`
	Playlists []string `json:"playlists"`
	Albums    []string `json:"albums"`
}

func (model *PostAddModel) ToModel(ctx *gin.Context) error {
	if err := ctx.Bind(&model); err != nil {
		logger.ErrorLog("Error: couldn't convert request to struct")
		dir, _ := os.Getwd()
		logger.DebugLog(fmt.Sprintf("PostAddModel %v: %v", dir, err))
	}
	return nil
}
