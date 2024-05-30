package comment_model

import (
	"gateway/pkg/service/logger"
	"github.com/gin-gonic/gin"
	"os"
)

type CommentUpdateModel struct {
	UserId     string
	CommentId  string
	NewContent string `json:"new_content"`
}

func (m *CommentUpdateModel) ToModel(ctx *gin.Context) error {
	dir, _ := os.Getwd()
	if err := ctx.Bind(m); err != nil {
		logger.ErrorWithDebugLog("Error: couldn't bind comment model", err, dir)
		return err
	}
	return nil
}
