package controller

import (
	"lunar/internal/model"
	"lunar/internal/presenter/http/controller/auth"
	"lunar/internal/presenter/http/controller/role"
	"lunar/internal/presenter/http/controller/user"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Role *role.RoleController
	Auth *auth.AuthController
	User *user.UserController
}

func NewController(
	Role *role.RoleController,
	Auth *auth.AuthController,
	User *user.UserController,
) *Controller {
	return &Controller{
		Role: Role,
		Auth: Auth,
		User: User,
	}
}

func (ctrl *Controller) Ping(c *gin.Context) {
	model.OnSuccess(c, &model.PingResponse{
		Ping: "pong gin gonic",
	})
}