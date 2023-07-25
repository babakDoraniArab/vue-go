package service

import (
	"errors"

	UserModel "github.com/babakDoraniArab/vue-go/internal/modules/user/models"
	UserRepository "github.com/babakDoraniArab/vue-go/internal/modules/user/repositories"
	"github.com/babakDoraniArab/vue-go/internal/modules/user/requests/auth"
	UserResponse "github.com/babakDoraniArab/vue-go/internal/modules/user/responses"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: UserRepository.NewUserRepository(),
	}
}

func (userService UserService) Find(id int) (UserResponse.User, error) {

	user := userService.userRepository.Find(id)
	var response UserResponse.User
	if user.ID == 0 {
		return response, errors.New("user not found")
	}
	return UserResponse.ToUser(user), nil
}
func (userService UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {

	hashedPAssword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse.User{}, err
	}

	user := UserModel.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPAssword),
	}

	newUser := userService.userRepository.Create(user)
	return UserResponse.ToUser(newUser), nil
}
