package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/techiebutler/utils"
)

// CreateEmployee creates a new employee.
func (c *employeeHandler) CreateEmployee(ctx *gin.Context) {
	//validate request
	req, err := validateCreateEmployeeRequest(ctx)
	if err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// service call
	baseRes, res, err := c.employeesvc.CreateEmployee(ctx, req)
	if err != nil {
		utils.InternalServer(ctx, err.Error())
		return
	}

	if baseRes.StatusCode != http.StatusCreated {
		utils.HandleServiceCodes(ctx, baseRes, err)
	}

	// transform response
	finalRes := createEmployeeTransformer(baseRes, res)

	// return response
	utils.ReturnJSONStruct(ctx, finalRes)
}

// GetEmployeeById fetches an employee by id.
func (c *employeeHandler) GetEmployeeById(ctx *gin.Context) {
	//validate request
	empId, err := validateGetEmployeeByIdRequest(ctx)
	if err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// service call
	baseRes, res, err := c.employeesvc.GetEmployeeById(ctx, empId)
	if err != nil {
		utils.InternalServer(ctx, err.Error())
		return
	}

	if baseRes.StatusCode != http.StatusOK {
		utils.HandleServiceCodes(ctx, baseRes, err)
	}

	// transform response
	finalRes := getEmployeeByIdTransformer(baseRes, res)

	// return response
	utils.ReturnJSONStruct(ctx, finalRes)
}

// UpdateEmployee updates an employee.
func (c *employeeHandler) UpdateEmployee(ctx *gin.Context) {
	//validate request
	req, err := validateUpdateEmployeeRequest(ctx)
	if err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// service call
	baseRes, res, err := c.employeesvc.UpdateEmployee(ctx, req)
	if err != nil {
		utils.InternalServer(ctx, err.Error())
		return
	}

	if baseRes.StatusCode != http.StatusOK {
		utils.HandleServiceCodes(ctx, baseRes, err)
	}

	// transform response
	finalRes := updateEmployeeTransformer(baseRes, res)

	// return response
	utils.ReturnJSONStruct(ctx, finalRes)
}

// DeleteEmployee deletes an employee.
func (c *employeeHandler) DeleteEmployee(ctx *gin.Context) {
	//validate request
	empId, err := validateGetEmployeeByIdRequest(ctx)
	if err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// service call
	baseRes, err := c.employeesvc.DeleteEmployee(ctx, empId)
	if err != nil {
		utils.InternalServer(ctx, err.Error())
		return
	}

	if baseRes.StatusCode != http.StatusOK {
		utils.HandleServiceCodes(ctx, baseRes, err)
	}

	// return response
	utils.ReturnJSONStruct(ctx, baseRes)
}

// GetEmployees fetches all employees.
func (c *employeeHandler) GetEmployees(ctx *gin.Context) {
	// validate request
	req, err := validateGetEmployeesRequest(ctx)
	if err != nil {
		utils.ValidationError(ctx, err)
		return
	}
	// service call
	baseRes, res, err := c.employeesvc.GetAllEmployees(ctx, req)
	if err != nil {
		utils.InternalServer(ctx, err.Error())
		return
	}

	if baseRes.StatusCode != http.StatusOK {
		utils.HandleServiceCodes(ctx, baseRes, err)
	}

	// transform response
	finalRes := getAllEmployeesTransformer(baseRes, res)

	// return response
	utils.ReturnJSONStruct(ctx, finalRes)
}
