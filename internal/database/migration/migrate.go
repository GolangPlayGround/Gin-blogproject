package migration

import (
	articleModels "Gin-blogproject/internal/modules/article/models"
	userModels "Gin-blogproject/internal/modules/user/models"
	"Gin-blogproject/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}
	fmt.Println("Migration done")
}
