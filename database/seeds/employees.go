package seeds

import (
	"github.com/saxenashivang/techiebutler/database/tables"
	"gorm.io/gorm"
)

func EmployeeSeed(db *gorm.DB) error {
	err := db.Create(&tables.Employee{
		Name:     "Shivang Saxena",
		PID:      "emp_uewuewiuewiuewiueiw",
		Position: "Software Engineer",
		Salary:   100000,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
