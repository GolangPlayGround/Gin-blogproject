package bootstrap

import (
	"Gin-blogproject/pkg/config"
	"Gin-blogproject/pkg/html"
	"Gin-blogproject/pkg/routing"
	"Gin-blogproject/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
