package routes

import (
	userCtrl "github.com/babakDoraniArab/vue-go/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	apiV1 := router.Group("/api/v1")

	userController := userCtrl.New()
	apiV1.GET("/register", userController.Register)
	apiV1.POST("/register", userController.HandleRegister)

}
