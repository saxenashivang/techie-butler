package tables

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	PID      string  `gorm:"column:employee_pid;unique;not null;type:varchar(40)"`
	Name     string  `json:"name" gorm:"type:varchar(255)"`
	Position string  `json:"position" gorm:"type:varchar(255)"`
	Salary   float64 `json:"salary" gorm:"type:decimal(10,2)"`
}
