package controllers

import (
	"log"
	"net/http"

	authRequest "github.com/babakDoraniArab/vue-go/internal/modules/user/requests/auth"
	UserService "github.com/babakDoraniArab/vue-go/internal/modules/user/services"
	"github.com/babakDoraniArab/vue-go/pkg/html"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.NewUserService(),
	}
}

func (controller *Controller) Register(c *gin.Context) {

	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})

}
func (controller *Controller) HandleRegister(c *gin.Context) {
	//validate user input
	var request authRequest.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	//create user

	user, err := controller.userService.Create(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not created",
			"error":   err.Error(),
		})
		return
	}

	// check if user created successfully without any error
	log.Printf("user created successfully: %s", user.Name)
	c.Redirect(http.StatusFound, "/")
	//redirect to login page

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "user created successfully",
	// 	"user":    user,
	// })
}
