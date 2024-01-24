package bootstrap

import (
	"Gin-blogproject/pkg/config"
	"Gin-blogproject/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}
