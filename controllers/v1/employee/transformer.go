package employee

import (
	"github.com/saxenashivang/techiebutler/database/tables"
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

func getAllEmployeesTransformer(baseRes utils.BaseResponse, employees []tables.Employee) utils.BaseResponse {

	var res utils.BaseResponse
	var dataRes []employeesvc.Employee

	for _, employee := range employees {
		var tempRes employeesvc.Employee
		tempRes.PID = employee.PID
		tempRes.Name = employee.Name
		tempRes.Position = employee.Position
		tempRes.Salary = employee.Salary
		dataRes = append(dataRes, tempRes)
	}

	res.StatusCode = baseRes.StatusCode
	res.Message = baseRes.Message
	res.Success = baseRes.Success
	res.Data = dataRes

	return res

}
