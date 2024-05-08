package employee

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/techiebutler/constants"
	"github.com/saxenashivang/techiebutler/database"
	dbops "github.com/saxenashivang/techiebutler/dbops/gorm"
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

func validateGetEmployeesRequest(ctx *gin.Context) (dbops.Pagination, error) {
	var pagination dbops.Pagination
	var err error

	limit := ctx.DefaultQuery("limit", "10")
	page := ctx.DefaultQuery("page", "0")
	sort := ctx.DefaultQuery("sort", "created_at desc")

	limitI, err := strconv.Atoi(limit)
	if err != nil {
		return pagination, err
	}

	pageI, err := strconv.Atoi(page)
	if err != nil {
		return pagination, err
	}

	pagination.Sort = sort
	pagination.Limit = limitI
	pagination.Page = pageI

	return pagination, err
}
