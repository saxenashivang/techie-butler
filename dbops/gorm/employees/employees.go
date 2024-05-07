package employees

import (
	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/techiebutler/constants"
	"github.com/saxenashivang/techiebutler/database/tables"
	"github.com/saxenashivang/techiebutler/utils"
	"gorm.io/gorm"
)

type GormInterface interface {
	CreateEmployee(ctx *gin.Context, employee tables.Employee) (tables.Employee, error)
	GetEmployeeById(ctx *gin.Context, id string) (tables.Employee, error)
	GetAllEmployees(ctx *gin.Context) ([]tables.Employee, error)
	UpdateEmployee(ctx *gin.Context, id string, employee tables.Employee) (tables.Employee, error)
	DeleteEmployee(ctx *gin.Context, id string) error
}

/* -------------------------------------------------------------------------- */
/*                                  HANDLEŖ                                  */
/* -------------------------------------------------------------------------- */
func Gorm(gormDB *gorm.DB) *employeeGormImpl {
	return &employeeGormImpl{
		DB: gormDB,
	}
}

type employeeGormImpl struct {
	DB *gorm.DB
}

func (e *employeeGormImpl) CreateEmployee(ctx *gin.Context, employee tables.Employee) (tables.Employee, error) {
	employee.PID = utils.UUIDWithPrefix(constants.EmployeeTablePrefix)
	err := e.DB.Create(&employee).Error
	return employee, err
}

func (e *employeeGormImpl) GetEmployeeById(ctx *gin.Context, id string) (tables.Employee, error) {
	var employee tables.Employee
	err := e.DB.Where("id = ?", id).First(&employee).Error
	return employee, err
}

func (e *employeeGormImpl) GetAllEmployees(ctx *gin.Context) ([]tables.Employee, error) {
	var employees []tables.Employee
	err := e.DB.Find(&employees).Error
	return employees, err
}

func (e *employeeGormImpl) UpdateEmployee(ctx *gin.Context, id string, employee tables.Employee) (tables.Employee, error) {
	err := e.DB.Model(&tables.Employee{}).Where("id = ?", id).Updates(employee).Error
	return employee, err
}

func (e *employeeGormImpl) DeleteEmployee(ctx *gin.Context, id string) error {
	err := e.DB.Where("id = ?", id).Delete(&tables.Employee{}).Error
	return err
}