package employeesvc

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/techiebutler/database/tables"
	"github.com/saxenashivang/techiebutler/utils"
)

// CreateEmployee creates a new employee.
func (e *employeeSvcImpl) CreateEmployee(ctx *gin.Context, employee CreateEmployeeRequest) (utils.BaseResponse, tables.Employee, error) {
	var baseRes utils.BaseResponse
	var err error

	// initialize the base response
	baseRes = utils.BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Something went wrong",
		Success:    false,
	}

	// create employee
	employeeM := tables.Employee{
		Name:     employee.Name,
		Position: employee.Position,
		Salary:   employee.Salary,
	}
	employeeM, err = e.employeesGorm.CreateEmployee(ctx, employeeM)
	if err != nil {
		log.Println("unable to create employee", err)
		baseRes.Message = "unable to create employee"
		return baseRes, employeeM, err
	}

	// set the base response
	baseRes.StatusCode = http.StatusCreated
	baseRes.Message = "Employee created successfully"
	baseRes.Success = true

	return baseRes, employeeM, nil
}

// GetEmployeeById fetches an employee by id.
func (e *employeeSvcImpl) GetEmployeeById(ctx *gin.Context, id string) (utils.BaseResponse, tables.Employee, error) {
	var baseRes utils.BaseResponse
	var err error
	var employee tables.Employee

	// initialize the base response
	baseRes = utils.BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Something went wrong",
		Success:    false,
	}

	// get employee by id
	employee, err = e.employeesGorm.GetEmployeeById(ctx, id)
	if err != nil {
		log.Println("unable to fetch employee", err)
		baseRes.Message = "unable to fetch employee"
		return baseRes, employee, err
	}

	// set the base response
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Employee fetched successfully"
	baseRes.Success = true

	return baseRes, employee, nil
}

// GetAllEmployees fetches all employees.
func (e *employeeSvcImpl) GetAllEmployees(ctx *gin.Context) (utils.BaseResponse, []tables.Employee, error) {
	var baseRes utils.BaseResponse
	var err error
	var employees []tables.Employee

	// initialize the base response
	baseRes = utils.BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Something went wrong",
		Success:    false,
	}

	// get all employees
	employees, err = e.employeesGorm.GetAllEmployees(ctx)
	if err != nil {
		log.Println("unable to fetch employees", err)
		baseRes.Message = "unable to fetch employees"
		return baseRes, employees, err
	}

	// set the base response
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Employees fetched successfully"
	baseRes.Success = true

	return baseRes, employees, nil
}

// UpdateEmployee updates an employee.
func (e *employeeSvcImpl) UpdateEmployee(ctx *gin.Context, id string, employee tables.Employee) (utils.BaseResponse, tables.Employee, error) {
	var baseRes utils.BaseResponse
	var err error

	// initialize the base response
	baseRes = utils.BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Something went wrong",
		Success:    false,
	}

	// update employee
	employee, err = e.employeesGorm.UpdateEmployee(ctx, id, employee)
	if err != nil {
		log.Println("unable to update employee", err)
		baseRes.Message = "unable to update employee"
		return baseRes, employee, err
	}

	// set the base response
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Employee updated successfully"
	baseRes.Success = true

	return baseRes, employee, nil
}

// DeleteEmployee deletes an employee.
func (e *employeeSvcImpl) DeleteEmployee(ctx *gin.Context, id string) (utils.BaseResponse, error) {
	var baseRes utils.BaseResponse
	var err error

	// initialize the base response
	baseRes = utils.BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Something went wrong",
		Success:    false,
	}

	// delete employee
	err = e.employeesGorm.DeleteEmployee(ctx, id)
	if err != nil {
		log.Println("unable to delete employee", err)
		baseRes.Message = "unable to delete employee"
		return baseRes, err
	}

	// set the base response
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Employee deleted successfully"
	baseRes.Success = true

	return baseRes, nil
}
