package bootstrap

import (
	"Gin-blogproject/pkg/config"
	"Gin-blogproject/pkg/html"
	"Gin-blogproject/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
