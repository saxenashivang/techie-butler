package employee

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/techiebutler/constants"
	"github.com/saxenashivang/techiebutler/database"
	"github.com/saxenashivang/techiebutler/dbops/gorm/employees"
	"github.com/saxenashivang/techiebutler/services/employeesvc"
	"gorm.io/gorm"
)

func validateCreateEmployeeRequest(ctx *gin.Context) (employeesvc.CreateEmployeeRequest, error) {
	var reqBody employeesvc.CreateEmployeeRequest
	var err error

	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		return reqBody, err
	}

	if reqBody.Name == "" {
		return reqBody, errors.New("name is required")
	}

	if reqBody.Position == "" {
		return reqBody, errors.New("position is required")
	}

	if reqBody.Salary == 0 {
		return reqBody, errors.New("salary is required")
	}

	return reqBody, nil
}

func validateGetEmployeeByIdRequest(ctx *gin.Context) (string, error) {

	empId := ctx.Param("id")

	if empId == "" {
		return empId, errors.New("id is required")
	}

	// check emp id started with prefix emp
	if empId[:3] != constants.EmployeeTablePrefix {
		return empId, errors.New("invalid id")
	}
	return empId, nil
}

func validateUpdateEmployeeRequest(ctx *gin.Context) (employeesvc.Employee, error) {
	var reqBody employeesvc.Employee
	var err error

	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		return reqBody, err
	}

	if reqBody.PID == "" {
		return reqBody, errors.New("pid is required")
	} else {
		gormDB, _ := database.Connection()
		employeeGorm := employees.Gorm(gormDB)

		_, err = employeeGorm.GetEmployeeById(ctx, reqBody.PID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return reqBody, errors.New("employee not found")
			}
			return reqBody, err
		}

	}

	return reqBody, nil
}
