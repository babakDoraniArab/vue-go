package routes

import (
	dashboardCtrl "github.com/babakDoraniArab/vue-go/internal/modules/dashboard/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	dashboardController := dashboardCtrl.New()
	router.GET("/", dashboardController.Dashboard)
	router.GET("/test2", dashboardController.Test)
}
