package employeesvc

import (
	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/techiebutler/database/tables"
	"github.com/saxenashivang/techiebutler/dbops/gorm/employees"
	"github.com/saxenashivang/techiebutler/utils"
)

type employeeSvcImpl struct {
	employeesGorm employees.GormInterface
}

// interface.
type Interface interface {
	CreateEmployee(ctx *gin.Context, employee CreateEmployeeRequest) (utils.BaseResponse, tables.Employee, error)
	GetEmployeeById(ctx *gin.Context, id string) (utils.BaseResponse, tables.Employee, error)
	GetAllEmployees(ctx *gin.Context) (utils.BaseResponse, []tables.Employee, error)
	UpdateEmployee(ctx *gin.Context, id string, employee tables.Employee) (utils.BaseResponse, tables.Employee, error)
	DeleteEmployee(ctx *gin.Context, id string) (utils.BaseResponse, error)
}

/* -------------------------------------------------------------------------- */
/*                              employeeSVC HANDLER                               */
/* -------------------------------------------------------------------------- */
func Handler(employeesGorm employees.GormInterface) *employeeSvcImpl {
	return &employeeSvcImpl{
		employeesGorm: employeesGorm,
	}
}
