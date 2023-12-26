package user

import (
	"lunar/internal/app/driver"
	"lunar/internal/model"
)

type UserRepository struct {
	db *driver.Database
}

func NewUserRepository(db *driver.Database) model.UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}