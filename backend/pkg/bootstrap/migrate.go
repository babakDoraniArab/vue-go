package bootstrap

import (
	"github.com/babakDoraniArab/vue-go/internal/database/migration"
	"github.com/babakDoraniArab/vue-go/pkg/database"
)

func Migrate() {
	// //set the Configs
	// config.Set()
	//connect to the database
	database.Connect()
	//migrate the database
	migration.Migrate()

}
