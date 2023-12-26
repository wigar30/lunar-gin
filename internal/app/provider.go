package app

import (
	"lunar/internal/app/config"
	"lunar/internal/app/driver"
	"lunar/internal/presenter/http/controller"
	roleCtrl "lunar/internal/presenter/http/controller/role"
	authCtrl "lunar/internal/presenter/http/controller/auth"
	userCtrl "lunar/internal/presenter/http/controller/user"
	"lunar/internal/presenter/http/middleware"
	roleRepo "lunar/internal/repository/role"
	userRepo "lunar/internal/repository/user"
	authUC "lunar/internal/usecase/auth"
	roleUC "lunar/internal/usecase/role"
	userUC "lunar/internal/usecase/user"

	"github.com/google/wire"
)

var (
	AppSet = wire.NewSet(
		config.NewViper,
		config.NewLogger,
		config.NewListenApp,
		config.NewGinGonic,
		driver.NewConnMySql,
	)

	ControllerSet = wire.NewSet(
		controller.NewController,
		roleCtrl.NewRoleController,
		authCtrl.NewAuthController,
		userCtrl.NewUserController,
	)

	UseCaseSet = wire.NewSet(
		roleUC.NewRoleUseCase,
		authUC.NewAuthUseCase,
		userUC.NewUserUseCase,
	)

	RepositorySet = wire.NewSet(
		roleRepo.NewRoleRepository,
		userRepo.NewUserRepository,
	)

	MiddlewareSet = wire.NewSet(
		middleware.NewMiddleware,
	)

	AllSet = wire.NewSet(
		AppSet,
		ControllerSet,
		UseCaseSet,
		RepositorySet,
		MiddlewareSet,
	)
)