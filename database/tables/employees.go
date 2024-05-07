package tables

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	ID       int     `gorm:"column:user_id;primaryKey;autoIncrement"`
	PID      string  `gorm:"column:user_pid;unique;not null;type:varchar(40)"`
	Name     string  `json:"name" gorm:"type:varchar(255)"`
	Position string  `json:"position" gorm:"type:varchar(255)"`
	Salary   float64 `json:"salary" gorm:"type:decimal(10,2)"`
}
