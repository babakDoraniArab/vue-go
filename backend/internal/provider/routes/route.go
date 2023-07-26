package routes

import (
	DashboardRoutes "github.com/babakDoraniArab/vue-go/internal/modules/dashboard/routes"
	UserRoutes "github.com/babakDoraniArab/vue-go/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	//let's get all the routes in the routes module and pass the router to it to register the routes . now every time we add a new module we have to add it here and we don't change the pkg/routing/routing.go file. we don't have to change the pkg/routing/routing.go file whenever we add a new route to the application

	// HomeRoutes.Routes(router)
	// ArticleRoutes.Routes(router)

	UserRoutes.Routes(router)
	DashboardRoutes.Routes(router)

}
