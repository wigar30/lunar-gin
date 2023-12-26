package auth

import "lunar/internal/model"


type AuthController struct {
	authUseCase model.AuthUseCaseInterface
}

func NewAuthController(authUseCase model.AuthUseCaseInterface) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
	}
}