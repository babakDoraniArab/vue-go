package migration

import (
	"fmt"
	"log"

	articleModel "github.com/babakDoraniArab/vue-go/internal/modules/article/models"
	userModel "github.com/babakDoraniArab/vue-go/internal/modules/user/models"
	"github.com/babakDoraniArab/vue-go/pkg/database"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModel.User{}, &articleModel.Article{})
	if err != nil {
		fmt.Println("Error in migration")
		log.Fatal(err)
	}
	fmt.Println("Migrate successfully")

}
