package routing

import (
	"Gin-blogproject/internal/provider/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()

}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}
