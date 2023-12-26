package role

import (
	"lunar/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (rc RoleController) GetAll(c *gin.Context) {
	// claimsCtx, ok := c.Locals("claims").(*model.JwtClaims)
	resp, err := rc.roleUseCase.GetAll(c)
	if err != nil {
		model.OnError(c, &model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	model.OnSuccess(c, resp)
}

func (rc RoleController) GetByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.ParseInt(paramsId, 10, 64)
	if err != nil {
		model.OnError(c, &model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	resp, err := rc.roleUseCase.GetByID(id)
	if err, errC := err.(*model.ErrorResponse); errC {
		model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
		return
	} else if !errC {
		model.OnError(c, &model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	model.OnSuccess(c, resp)
}
