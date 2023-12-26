package role

import (
	"lunar/internal/app/driver"
	"lunar/internal/model"
)

type RoleRepository struct {
	db *driver.Database
}

func NewRoleRepository(db *driver.Database) model.RoleRepositoryInterface {
	return &RoleRepository{
		db: db,
	}
}