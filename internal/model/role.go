package model

import (
	"lunar/internal/entity"

	"github.com/gin-gonic/gin"
)

type RoleResponse struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
}

type RolesResponse struct {
	Items []*RoleResponse `json:"items"`
}

type RoleRepositoryInterface interface {
	// Create new role
	//  @param role *Role, role object
	GetAll() ([]*entity.Role, error)
	GetByID(int64) (*entity.Role, error)
}

type RoleUseCaseInterface interface {
	GetAll(c *gin.Context) (*RolesResponse, error)
	GetByID(int64) (*RoleResponse, error)
}
