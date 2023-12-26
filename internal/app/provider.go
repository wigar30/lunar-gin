package app

import (
	"lunar/internal/app/config"
	"lunar/internal/app/driver"
	"lunar/internal/presenter/http/controller"

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
	)

	AllSet = wire.NewSet(
		AppSet,
		ControllerSet,
	)
)