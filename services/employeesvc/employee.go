package employeesvc

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/techiebutler/database/tables"
	"github.com/saxenashivang/techiebutler/dbops/gorm"
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
func (e *employeeSvcImpl) GetAllEmployees(ctx *gin.Context, pagination gorm.Pagination) (utils.BaseResponse, gorm.Pagination, error) {
	var baseRes utils.BaseResponse
	var err error

	// initialize the base response
	baseRes = utils.BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Something went wrong",
		Success:    false,
	}

	// get all employees

	pagination, err = e.employeesGorm.GetAllEmployees(ctx, pagination)
	if err != nil {
		log.Println("unable to fetch employees", err)
		baseRes.Message = "unable to fetch employees"
		return baseRes, pagination, err
	}

	// set the base response
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Employees fetched successfully"
	baseRes.Success = true

	return baseRes, pagination, nil
}

// UpdateEmployee updates an employee.
func (e *employeeSvcImpl) UpdateEmployee(ctx *gin.Context, req Employee) (utils.BaseResponse, tables.Employee, error) {
	var baseRes utils.BaseResponse
	var employee tables.Employee
	var err error

	// initialize the base response
	baseRes = utils.BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Something went wrong",
		Success:    false,
	}

	if req.Name != "" {
		employee.Name = req.Name
	}
	if req.Position != "" {
		employee.Position = req.Position
	}
	if req.Salary != 0 {
		employee.Salary = req.Salary
	}
	employee.PID = req.PID
	// update employee
	employee, err = e.employeesGorm.UpdateEmployee(ctx, employee.PID, employee)
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
