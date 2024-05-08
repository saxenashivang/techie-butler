package employeesvc

type CreateEmployeeRequest struct {
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

type Employee struct {
	PID      string  `json:"pid"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

type GetAllEmployeesResponse struct {
	Limit      int        `json:"limit"`
	Page       int        `json:"page"`
	Sort       string     `json:"sort"`
	TotalRows  int64      `json:"total_rows"`
	TotalPages int        `json:"total_pages"`
	Employees  []Employee `json:"employees"`
}
