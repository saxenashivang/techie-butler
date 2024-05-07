package database

import (
	"github.com/saxenashivang/techiebutler/database/tables"
	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(db *gorm.DB) []Migrate {
	var employees tables.Employee

	employeesM := Migrate{TableName: "employees",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&employees) }}

	return []Migrate{
		employeesM,
	}
}
