package routes

import (
	userCtrl "github.com/babakDoraniArab/vue-go/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	userController := userCtrl.New()
	router.GET("/register", userController.Register)
	router.POST("/register", userController.HandleRegister)

}
