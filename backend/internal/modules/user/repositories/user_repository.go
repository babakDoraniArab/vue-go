package repositories

import (
	UserModel "github.com/babakDoraniArab/vue-go/internal/modules/user/models"
	"github.com/babakDoraniArab/vue-go/pkg/database"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	db := database.Connection()
	return &UserRepository{
		DB: db,
	}
}

func (userRepository UserRepository) Find(id int) UserModel.User {
	var users UserModel.User
	userRepository.DB.First(&users, id)
	return users
}
func (userRepository UserRepository) Create(user UserModel.User) UserModel.User {
	var newUser UserModel.User
	userRepository.DB.Create(&user).Scan(&newUser)
	return newUser
}
