package http_handlers

import (
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/transport/http_handlers/album_handler"
	"gateway/pkg/transport/http_handlers/auth_handler"
	"gateway/pkg/transport/http_handlers/collection_handler"
	"gateway/pkg/transport/http_handlers/comment_handler"
	"gateway/pkg/transport/http_handlers/music_handler"
	"gateway/pkg/transport/http_handlers/photo_handler"
	"gateway/pkg/transport/http_handlers/playlist_handler"
	"gateway/pkg/transport/http_handlers/post_handler"
	"gateway/pkg/transport/http_handlers/recommendation_handler"
	"gateway/pkg/transport/http_handlers/user_handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	global_vars_config.Router = gin.New()
	global_vars_config.Router.SetTrustedProxies([]string{"192.168.1.*"})
	user_handler.Init()
	auth_handler.Init()
	photo_handler.Init()
	music_handler.Init()
	playlist_handler.Init()
	album_handler.Init()
	collection_handler.Init()
	post_handler.Init()
	comment_handler.Init()
	recommendation_handler.Init()
}
