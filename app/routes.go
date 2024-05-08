package app

import (
	"github.com/gin-gonic/gin"
	"github.com/saxenashivang/techiebutler/controllers/v1/employee"
	"github.com/saxenashivang/techiebutler/database"
	"github.com/saxenashivang/techiebutler/dbops/gorm/employees"
	"github.com/saxenashivang/techiebutler/services/employeesvc"
)

func MapURL() {
	router := gin.Default()

	// database connection
	gormDB, _ := database.Connection()

	/* -------------------------------------------------------------------------- */
	/*                                GORM HANDLERS                               */
	/* -------------------------------------------------------------------------- */
	employeeGorm := employees.Gorm(gormDB)

	/* -------------------------------------------------------------------------- */
	/*                            SERVICE & CONTROLLER                            */
	/* -------------------------------------------------------------------------- */
	employeeSvc := employeesvc.Handler(employeeGorm)
	employeeHandler := employee.Handler(employeeSvc)

	/* -------------------------------------------------------------------------- */
	/*                                   ROUTES                                   */
	/* -------------------------------------------------------------------------- */

	v1 := router.Group("/v1")
	{
		employeeRoutes := v1.Group("/employee")
		employeeRoutes.POST("/create", employeeHandler.CreateEmployee)
		employeeRoutes.GET("/get/:id", employeeHandler.GetEmployeeById)
		employeeRoutes.GET("/getall", employeeHandler.GetEmployees)
		employeeRoutes.PUT("/update", employeeHandler.UpdateEmployee)
		employeeRoutes.DELETE("/delete/:id", employeeHandler.DeleteEmployee)
	}

	err := router.Run(":8080")
	if err != nil {
		panic(err.Error() + "MapURL router not able to run")
	}
}
