package router

import (
	"lunar/internal/presenter/http/controller"

	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine, ctrl *controller.Controller) {
	api := c.Group("/api")

	v1 := api.Group("/v1")
	v1.GET("/ping", ctrl.Ping)
}