package controller

import (
	"lunar/internal/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Ping(c *gin.Context) {
	model.OnSuccess(c, &model.PingResponse{
		Ping: "pong",
	})
}