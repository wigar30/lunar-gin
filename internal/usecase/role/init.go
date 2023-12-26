package role

import "lunar/internal/model"


type RoleUseCase struct {
	roleRepo model.RoleRepositoryInterface
}

func NewRoleUseCase(roleRepo model.RoleRepositoryInterface) model.RoleUseCaseInterface {
	return &RoleUseCase{
		roleRepo: roleRepo,
	}
}