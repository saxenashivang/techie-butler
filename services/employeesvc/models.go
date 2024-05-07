package employeesvc

type CreateEmployeeRequest struct {
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}
