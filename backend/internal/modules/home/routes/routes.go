package routes

import (
	homeCtrl "github.com/babakDoraniArab/vue-go/internal/modules/home/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)

}
