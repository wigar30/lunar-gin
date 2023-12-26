package user

import (
	"lunar/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (uc UserController) GetProfile(c *gin.Context) {
	claim, _ := c.Get("claims")

	claims, ok := claim.(*model.JwtClaims)
	if !ok {
		model.OnError(c, &model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "error type assertion",
		})
		return
	}

	userID, err := strconv.ParseInt(claims.ID, 10, 64)
	if err, errC := err.(*model.ErrorResponse); errC {
		model.OnError(c, &model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	resp, err := uc.userUseCase.GetUserByID(userID)
	if err, errC := err.(*model.ErrorResponse); errC {
		model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
		return
	}

	model.OnSuccess(c, resp)
}
