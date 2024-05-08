package employees

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/techiebutler/config"
	"github.com/saxenashivang/techiebutler/constants"
	"github.com/saxenashivang/techiebutler/database"
	"github.com/saxenashivang/techiebutler/database/tables"
	"github.com/saxenashivang/techiebutler/utils"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadConfigs()
}

func TestCreateEmployee(t *testing.T) {
	gormDB, _ := database.Connection()
	gormHandler := Gorm(gormDB)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   nil,
	}

	var employeeData tables.Employee

	employeeData.PID = utils.UUIDWithPrefix(constants.EmployeeTablePrefix)
	employeeData.Name = "Test Employee"
	employeeData.Position = "Software Engineer"
	employeeData.Salary = 100000

	res, err := gormHandler.CreateEmployee(c, employeeData)
	assert.Empty(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, res.Name, employeeData.Name)
	assert.Equal(t, res.Position, employeeData.Position)
	assert.Equal(t, res.Salary, employeeData.Salary)
	assert.NotEmpty(t, res.PID)

	t.Cleanup(func() {
		gormDB.Model(&tables.Employee{}).Where("employee_pid = ?", res.PID).Delete(&res)
	})
}

func TestGetEmployeeById(t *testing.T) {
	gormDB, _ := database.Connection()
	gormHandler := Gorm(gormDB)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   nil,
	}

	// from seed
	employeeID := "emp_7cdbb49388904a2cab68d882323d6c40"

	employee, err := gormHandler.GetEmployeeById(c, employeeID)
	assert.Empty(t, err)
	assert.NotEmpty(t, employee)
	assert.Equal(t, employee.PID, employeeID)

}

func TestUpdateEmployee(t *testing.T) {
	gormDB, _ := database.Connection()
	gormHandler := Gorm(gormDB)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   nil,
	}

	// from seed
	employeeID := "emp_7cdbb49388904a2cab68d882323d6c40"

	employee, err := gormHandler.GetEmployeeById(c, employeeID)
	assert.Empty(t, err)
	assert.NotEmpty(t, employee)
	assert.Equal(t, employee.PID, employeeID)

	employee.Name = "Updated Employee"
	employee.Position = "Senior Software Engineer"
	employee.Salary = 200000

	res, err := gormHandler.UpdateEmployee(c, employeeID, employee)
	assert.Empty(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, res.Name, employee.Name)
	assert.Equal(t, res.Position, employee.Position)
	assert.Equal(t, res.Salary, employee.Salary)
	assert.Equal(t, res.PID, employeeID)

}

func TestDeleteEmployee(t *testing.T) {
	gormDB, _ := database.Connection()
	gormHandler := Gorm(gormDB)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   nil,
	}

	var employeeData tables.Employee

	employeeData.PID = utils.UUIDWithPrefix(constants.EmployeeTablePrefix)
	employeeData.Name = "Test Employee"
	employeeData.Position = "Software Engineer"
	employeeData.Salary = 100000

	res, err := gormHandler.CreateEmployee(c, employeeData)
	assert.Empty(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, res.Name, employeeData.Name)
	assert.Equal(t, res.Position, employeeData.Position)
	assert.Equal(t, res.Salary, employeeData.Salary)
	assert.NotEmpty(t, res.PID)

	err = gormHandler.DeleteEmployee(c, res.PID)
	assert.Empty(t, err)

	t.Cleanup(func() {
		gormDB.Unscoped().Delete(&employeeData)
	})
}
