package database

import (
	"github.com/saxenashivang/techiebutler/database/seeds"
	"gorm.io/gorm"
)

type Seed struct {
	TableName string
	Run       func(*gorm.DB) error
}

func Seeder(db *gorm.DB) []Seed {
	employees := Seed{TableName: "employees", Run: func(d *gorm.DB) error { return seeds.EmployeeSeed(db) }}

	return []Seed{
		employees,
	}
}
