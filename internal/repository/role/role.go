package role

import (
	"errors"
	"lunar/internal/entity"
	"lunar/internal/model"
	"net/http"

	"gorm.io/gorm"
)

func (rr *RoleRepository) GetAll() ([]*entity.Role, error) {
	var roles []*entity.Role
	err := rr.db.Find(&roles).Error
	if err != nil {
		return nil, &model.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return roles, nil
}

func (rr *RoleRepository) GetByID(RoleId int64) (*entity.Role, error) {
	var role *entity.Role
	err := rr.db.First(&role, RoleId).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &model.ErrorResponse{
			Code: http.StatusNotFound,
			Message: err.Error(),
		}
	}
	if err != nil {
		return nil, &model.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return role, nil
}