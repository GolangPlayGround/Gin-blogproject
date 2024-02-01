package bootstrap

import (
	"Gin-blogproject/internal/database/seeder"
	"Gin-blogproject/pkg/config"
	"Gin-blogproject/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()

}
