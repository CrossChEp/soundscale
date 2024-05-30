package comment_model

import (
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
	"os"
)

type CommentAddModel struct {
	UserId     string
	EntityId   string `json:"entity_id"`
	EntityType string
	Content    string `json:"content"`
}

func (m *CommentAddModel) ToModel(ctx *gin.Context) error {
	dir, _ := os.Getwd()
	if err := ctx.Bind(m); err != nil {
		logger.ErrorWithDebugLog("Error: couldn't bind comment model", err, dir)
		return err
	}
	return nil
}
