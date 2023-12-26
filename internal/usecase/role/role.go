package role

import (
	"lunar/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ru RoleUseCase) GetAll(c *gin.Context) (*model.RolesResponse, error) {
	roles, err := ru.roleRepo.GetAll()
	if err != nil {
		return nil, &model.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	var rolesResponses []*model.RoleResponse
	for _, role := range roles {
		rolesResponses = append(rolesResponses, &model.RoleResponse{
			// Copy values from role to RoleResponse
			ID:   role.ID,
			Name: role.Name,
			// Add other fields as needed
		})
	}

	return &model.RolesResponse{
		Items: rolesResponses,
	}, nil
}

func (ru RoleUseCase) GetByID(roleId int64) (*model.RoleResponse, error) {
	role, err := ru.roleRepo.GetByID(roleId)
	if errC, ok := err.(*model.ErrorResponse); ok {
		return nil, &model.ErrorResponse{
			Code: errC.Code,
			Message: errC.Error(),
		}
	}

	return &model.RoleResponse{
		ID: role.ID,
		Name: role.Name,
	}, nil
}