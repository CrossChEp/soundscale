package global_vars_config

import (
	"crypto/ecdsa"
	"github.com/gin-gonic/gin"
)

var (
	Router    *gin.Engine
	PublicKey *ecdsa.PublicKey
)
