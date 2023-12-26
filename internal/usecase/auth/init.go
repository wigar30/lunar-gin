package auth

import "lunar/internal/model"

type AuthUseCase struct {
	userRepo model.UserRepositoryInterface
	config   *model.EnvConfigs
}

func NewAuthUseCase(userRepo model.UserRepositoryInterface, config *model.EnvConfigs) model.AuthUseCaseInterface {
	return &AuthUseCase{
		userRepo: userRepo,
		config:   config,
	}
}
