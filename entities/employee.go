package entities

type Employee struct {
	ID       uint    `json:"id" gorm:"primary_key"`
	Name     string  `json:"name" gorm:"type:varchar(255)"`
	Position string  `json:"position" gorm:"type:varchar(255)"`
	Salary   float64 `json:"salary" gorm:"type:decimal(10,2)"`
}
