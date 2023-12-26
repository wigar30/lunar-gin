package config

import (
	"fmt"
	"lunar/internal/model"
	"lunar/internal/presenter/http/controller"
	"lunar/internal/presenter/http/router"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type HTTPServiceInterface interface {
	ListenApp() error
}

type HttpService struct {
	app    *gin.Engine
	ctrl   *controller.Controller
	config *model.EnvConfigs
	log    *model.Logger
}

func NewListenApp(app *gin.Engine, ctrl *controller.Controller, config *model.EnvConfigs, log *model.Logger) HTTPServiceInterface {
	return &HttpService{
		app:    app,
		ctrl:   ctrl,
		config: config,
		log:    log,
	}
}

func (g *HttpService) ListenApp() error {
	g.app.Use(
		cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
		}),
		gin.Recovery(),
	)

	router.Router(g.app, g.ctrl)

	port := g.config.AppPort
	return g.app.Run(fmt.Sprintf(":%s", port))
}
