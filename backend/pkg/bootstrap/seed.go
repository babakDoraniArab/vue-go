package bootstrap

import (
	"github.com/babakDoraniArab/vue-go/internal/database/seeder"
	"github.com/babakDoraniArab/vue-go/pkg/database"
)

func Seed() {
	// //set the Configs
	// config.Set()
	//connect to the database
	database.Connect()

	//seed the database
	seeder.Seed()
}
