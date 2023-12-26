package auth

import (
	"lunar/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func (ac AuthController) Login(c *gin.Context) {
	var req model.LoginRequest
	// req := new(model.LoginRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		model.OnError(c, &model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err := Validator.Struct(req)
	if err != nil {
		model.OnError(c, &model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
		} 

	resp, err := ac.authUseCase.Login(c, &req)
	if err, errC := err.(*model.ErrorResponse); errC {
		model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
		return
	}

	model.OnSuccess(c, resp)
}
