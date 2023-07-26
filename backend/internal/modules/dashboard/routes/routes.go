package routes

import (
	dashboardCtrl "github.com/babakDoraniArab/vue-go/internal/modules/dashboard/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	apiV1 := router.Group("/api/v1/")

	dashboardController := dashboardCtrl.New()
	apiV1.GET("/", dashboardController.Dashboard)
	apiV1.GET("/test2", dashboardController.Test)
}
