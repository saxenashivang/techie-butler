package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/saxenashivang/techiebutler/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, *sql.DB) {
	dsn := "host=" + config.DB.Host +
		" user=" + config.DB.Username +
		" password=" + config.DB.Password +
		" dbname=" + config.DB.Database +
		" port=" + config.DB.Port +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Connection], Error in opening db")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("[Connection], Error in setting sqldb")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("database connected successfully")
	return db, sqlDB
}
