package main

import (
	configuration "go-license-altera/config"
	"go-license-altera/modules/database"
)

var (
	config         = configuration.GetConfig()
	dbGormPostgres = database.NewDatabaseConnection(config)
)

func main() {
	defer database.CloseDatabaseConnection(dbGormPostgres)

}
