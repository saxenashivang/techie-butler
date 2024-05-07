package database

import (
	"gorm.io/gorm"
)

type Seed struct {
	TableName string
	Run       func(*gorm.DB) error
}

func Seeder(db *gorm.DB) []Seed {

	return []Seed{}
}
