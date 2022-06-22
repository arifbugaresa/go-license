package main

import (
	"github.com/labstack/gommon/log"
	"go-license-altera/businesses/client"
	"go-license-altera/businesses/license"
	"go-license-altera/businesses/product"
	"go-license-altera/businesses/user_info"
	configuration "go-license-altera/config"
	"go-license-altera/modules/database"
	"strconv"
)

var (
	config         = configuration.GetConfig()
	dbGormPostgres = database.NewDatabaseConnection(config)
)

func main() {
	defer database.CloseDatabaseConnection(dbGormPostgres)

	migrateDatabase()

}

func migrateDatabase() {
	dbGormPostgres.AutoMigrate(
		product.Product{},
		license.License{},
		client.Client{},
		user_info.UserInfo{},
	)

	log.Info("Success migrate database, " + strconv.Itoa(int(dbGormPostgres.RowsAffected)) + " row affected.")
}