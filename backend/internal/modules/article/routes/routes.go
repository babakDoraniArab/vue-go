package routes

import (
	articleCtrl "github.com/babakDoraniArab/vue-go/internal/modules/article/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	articleController := articleCtrl.New()
	router.GET("/articles/:id", articleController.Show)

}
