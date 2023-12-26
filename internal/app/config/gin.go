package config

import (
	"lunar/internal/model"

	"github.com/gin-gonic/gin"
)

func NewGinGonic(config *model.EnvConfigs) *gin.Engine {
	gin.SetMode(config.GinMode)
	app := gin.Default()

	return app
}