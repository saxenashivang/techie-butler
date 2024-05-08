package seeds

import (
	"strconv"

	"github.com/saxenashivang/techiebutler/database/tables"
	"gorm.io/gorm"
)

var Names = []string{
	"Shivang Saxena",
	"John Doe",
	"Jane Doe",
	"John Smith",
	"Jane Smith",
	"John Wick",
	"Jane Wick",
	"John Cena",
	"Jane Cena",
	"John Legend",
}

var Positions = []string{
	"Software Engineer",
	"Senior Software Engineer",
	"Software Developer",
	"Senior Software Developer",
	"Software Architect",
	"Senior Software Architect",
	"Software Tester",
	"Senior Software Tester",
	"Software Analyst",
	"Senior Software Analyst",
}

var Salaries = []float64{
	100000,
	200000,
	300000,
	400000,
	500000,
	600000,
	700000,
	800000,
	900000,
	1000000,
}

func EmployeeSeed(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		employee := tables.Employee{
			PID:      "emp_7cdbb49388904a2cab68d882323d6c4" + strconv.Itoa(i),
			Name:     Names[i],
			Position: Positions[i],
			Salary:   Salaries[i],
		}
		db.Create(&employee)
	}

	return nil
}
