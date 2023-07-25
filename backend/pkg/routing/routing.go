package routing

import (
	"github.com/babakDoraniArab/vue-go/internal/provider/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}

// GetRouter returns the route
func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	//let's get all the routes in the routes module and pass the router to it to register the routes
	// I passed the router to the Routes module and to do that I have a GetRouter function

	// we should not update the package whenever we update the logic so we are going to create a new package called routes in the internal/Provider folder and move the routes.go file to that folder and initialize it here and then register the routes
	// routes.Routes(GetRouter()) so I have to remove this line from here and move it to the internal/provider/routes/route.go file
	routes.RegisterRoutes(GetRouter())
}
