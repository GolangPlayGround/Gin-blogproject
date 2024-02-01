package seeder

import (
	articleModel "Gin-blogproject/internal/modules/article/models"
	userModel "Gin-blogproject/internal/modules/user/models"
	"Gin-blogproject/pkg/database"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Seed() {

	db := database.Connection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)

	if err != nil {
		log.Fatal("hash password error")
		return
	}

	user := userModel.User{Name: "saechimdaeki", Email: "saechim@daeki.com", Password: string(hashedPassword)}

	db.Create(&user)

	log.Printf("User created successfully with email %s \n", user.Email)

	for i := 1; i <= 10; i++ {
		article := articleModel.Article{Title: fmt.Sprintf("Random title: %d", i), Content: fmt.Sprintf("Random title: %d", i), UserID: 1}

		db.Create(&article)

		log.Printf("Article created successfully with title %s \n", article.Title)

	}
	log.Println("Seeder done")
}
