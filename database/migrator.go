package database

import (
	"github.com/saxenashivang/techiebutler/entities"
	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(db *gorm.DB) []Migrate {
	var employee entities.Employee

	employeeM := Migrate{TableName: "employees",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&employee) }}
	return []Migrate{
		employeeM,
	}
}
