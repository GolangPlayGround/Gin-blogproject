package bootstrap

import (
	"Gin-blogproject/internal/database/migration"
	"Gin-blogproject/pkg/config"
	"Gin-blogproject/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
