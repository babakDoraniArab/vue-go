package service

import (
	"github.com/babakDoraniArab/vue-go/internal/modules/user/requests/auth"
	UserResponse "github.com/babakDoraniArab/vue-go/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Find(id int) (UserResponse.User, error)
	Create(request auth.RegisterRequest) (UserResponse.User, error)
}
