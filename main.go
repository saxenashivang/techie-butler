package main

import (
	"github.com/saxenashivang/techiebutler/app"
	"github.com/saxenashivang/techiebutler/config"
	"github.com/saxenashivang/techiebutler/database"
)

func main() {
	config.LoadConfigs()

	initPostgres()
	app.RegisterRoutesAndMiddlewares()

}

// initPostgres method to connect to PSQL DBs
func initPostgres() {
	database.Connection()
}
