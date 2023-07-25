package seeder

import (
	"fmt"
	"log"

	articleModel "github.com/babakDoraniArab/vue-go/internal/modules/article/models"
	userModel "github.com/babakDoraniArab/vue-go/internal/modules/user/models"
	"github.com/babakDoraniArab/vue-go/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	// Seeder logic
	db := database.Connection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while hashing password")
		panic(err)
	}
	user := userModel.User{Name: "Random Name ", Email: "randomEmail@gmail.com", Password: string(hashedPassword)}

	db.Create(&user) // pass pointer of data to Create

	log.Printf("User Created Successfully")

	// Seeding the 10 Articles
	for i := 0; i < 10; i++ {

		article := articleModel.Article{
			Title:   fmt.Sprintf("Article Title %d", i),
			Content: fmt.Sprintf("Article Content %d", i),
			UserID:  1,
		}
		db.Create(&article) // pass pointer of data to Create
		log.Printf("article %d Created Successfully", i)
	}
	fmt.Println("Seed Executed Successfully .... ")
}
