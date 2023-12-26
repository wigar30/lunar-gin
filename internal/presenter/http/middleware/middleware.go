package middleware

import "lunar/internal/model"

type Middleware struct {
	AuthMiddleware *AuthMiddleware
}

func NewMiddleware(userRepo model.UserRepositoryInterface, config *model.EnvConfigs) *Middleware {
	return &Middleware{
		AuthMiddleware: NewAuthMiddleware(userRepo, config),
	}
}