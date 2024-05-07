package app

import (
	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/techiebutler/database"
)

func MapURL() {
	router := gin.Default()

	// database connection
	database.Connection()

	err := router.Run(":8080")
	if err != nil {
		panic(err.Error() + "MapURL router not able to run")
	}
}
