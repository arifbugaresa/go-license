package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-license-altera/api/v1"
	"go-license-altera/businesses/client"
	"go-license-altera/businesses/license"
	"go-license-altera/businesses/license_mapping"
	"go-license-altera/businesses/product"
	"go-license-altera/businesses/user_info"
	configuration "go-license-altera/config"
	"go-license-altera/modules/database"
	"strconv"
)

var (
	config       = configuration.GetConfig()
	dbConnection = database.NewDatabaseConnection(config)
	e            = echo.New()
)

func main() {
	defer database.CloseDatabaseConnection(dbConnection)

	migrateDatabase()

	api.Controller(e)

	runServer()

}

func migrateDatabase() {
	dbConnection.AutoMigrate(
		product.Product{},
		license.License{},
		client.Client{},
		user_info.UserInfo{},
		license_mapping.LicenseMapping{},
	)

	log.Info("Success migrate database, " + strconv.Itoa(int(dbConnection.RowsAffected)) + " row affected.")
}

func runServer() {
	address := fmt.Sprintf("localhost:%s", config.AppPort)
	err := e.Start(address)
	if err != nil {
		log.Info("shutting down the server")
	}
}
