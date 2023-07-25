package repositories

import UserModel "github.com/babakDoraniArab/vue-go/internal/modules/user/models"

type UserRepositoryInterface interface {
	// Get all users
	Find(id int) UserModel.User
	Create(user UserModel.User) UserModel.User
}
