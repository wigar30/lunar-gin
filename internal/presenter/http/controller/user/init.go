package user

import "lunar/internal/model"


type UserController struct {
	userUseCase model.UserUseCaseInterface
}

func NewUserController(userUseCase model.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}