package router

import (
	"lunar/internal/presenter/http/controller"
	"lunar/internal/presenter/http/middleware"

	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine, ctrl *controller.Controller, m *middleware.Middleware) {
	api := c.Group("/api")

	v1 := api.Group("/v1")
	v1.GET("/ping", ctrl.Ping)

	auth := v1.Group("auth")
	auth.POST("/login", ctrl.Auth.Login)

	user := v1.Group("user", m.AuthMiddleware.ValidateToken())
	user.GET("/profile", ctrl.User.GetProfile)

	role := v1.Group("role", m.AuthMiddleware.ValidateToken())
	role.GET("/", ctrl.Role.GetAll)
	role.GET("/:id", ctrl.Role.GetByID)
}