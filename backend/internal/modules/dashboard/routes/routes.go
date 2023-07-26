package routes

import (
	"time"

	dashboardCtrl "github.com/babakDoraniArab/vue-go/internal/modules/dashboard/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	// apiV1 := router.Group("/api/v1/")

	dashboardController := dashboardCtrl.New()
	// apiV1.GET("/", dashboardController.Dashboard)
	router.GET("/api/v1/test2", dashboardController.Test)
}
