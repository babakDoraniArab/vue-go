package bootstrap

import (
	"github.com/babakDoraniArab/vue-go/pkg/database"
	"github.com/babakDoraniArab/vue-go/pkg/html"
	"github.com/babakDoraniArab/vue-go/pkg/routing"
	"github.com/babakDoraniArab/vue-go/pkg/static"
)

func Serve() {
	// //set the Configs
	// config.Set()
	//connect to the database
	database.Connect()

	routing.Init()

	// load HTML templates
	static.LoadStatic(routing.GetRouter())
	html.LoadHTML(routing.GetRouter())

	//let's get all the routes in the routes module and pass the router to it to register the routes
	routing.RegisterRoutes()

	routing.Serve()
}
