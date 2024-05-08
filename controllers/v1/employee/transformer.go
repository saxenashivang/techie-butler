package employee

import (
	"github.com/saxenashivang/techiebutler/database/tables"
	dbops "github.com/saxenashivang/techiebutler/dbops/gorm"
	"github.com/saxenashivang/techiebutler/services/employeesvc"
	"github.com/saxenashivang/techiebutler/utils"
)

func createEmployeeTransformer(baseRes utils.BaseResponse, employee tables.Employee) utils.BaseResponse {

	var res utils.BaseResponse
	var dataRes employeesvc.Employee

	dataRes.PID = employee.PID
	dataRes.Name = employee.Name
	dataRes.Position = employee.Position
	dataRes.Salary = employee.Salary

	res.StatusCode = baseRes.StatusCode
	res.Message = baseRes.Message
	res.Success = baseRes.Success
	res.Data = dataRes

	return res

}

func getEmployeeByIdTransformer(baseRes utils.BaseResponse, employee tables.Employee) utils.BaseResponse {

	var res utils.BaseResponse
	var dataRes employeesvc.Employee

	dataRes.PID = employee.PID
	dataRes.Name = employee.Name
	dataRes.Position = employee.Position
	dataRes.Salary = employee.Salary

	res.StatusCode = baseRes.StatusCode
	res.Message = baseRes.Message
	res.Success = baseRes.Success
	res.Data = dataRes

	return res

}

func updateEmployeeTransformer(baseRes utils.BaseResponse, employee tables.Employee) utils.BaseResponse {

	var res utils.BaseResponse
	var dataRes employeesvc.Employee

	dataRes.PID = employee.PID
	dataRes.Name = employee.Name
	dataRes.Position = employee.Position
	dataRes.Salary = employee.Salary

	res.StatusCode = baseRes.StatusCode
	res.Message = baseRes.Message
	res.Success = baseRes.Success
	res.Data = dataRes

	return res

}

func getAllEmployeesTransformer(baseRes utils.BaseResponse, pagination dbops.Pagination) utils.BaseResponse {

	var res utils.BaseResponse
	var dataRes employeesvc.GetAllEmployeesResponse

	dataRes.Limit = pagination.Limit
	dataRes.Page = pagination.Page
	dataRes.Sort = pagination.Sort
	dataRes.TotalRows = pagination.TotalRows
	dataRes.TotalPages = pagination.TotalPages
	dataRes.Employees = make([]employeesvc.Employee, 0)

	for _, employee := range pagination.Rows.([]tables.Employee) {
		var emp employeesvc.Employee
		emp.PID = employee.PID
		emp.Name = employee.Name
		emp.Position = employee.Position
		emp.Salary = employee.Salary
		dataRes.Employees = append(dataRes.Employees, emp)
	}

	res.StatusCode = baseRes.StatusCode
	res.Message = baseRes.Message
	res.Success = baseRes.Success
	res.Data = dataRes

	return res

}
