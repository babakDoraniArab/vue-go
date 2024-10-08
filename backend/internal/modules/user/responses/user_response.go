package responses

import (
	"fmt"

	"github.com/babakDoraniArab/vue-go/internal/modules/user/models"
)

type User struct {
	ID    uint
	Name  string
	Email string
	Image string
}

type Users struct {
	Data []User
}

func ToUser(user models.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
	}
}
