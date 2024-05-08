package employee

import "github.com/saxenashivang/techiebutler/services/employeesvc"

/* -------------------------------------------------------------------------- */
/*                                   HANDLER                                  */
/* -------------------------------------------------------------------------- */
type employeeHandler struct {
	employeesvc employeesvc.Interface
}

func Handler(employeesvc employeesvc.Interface) *employeeHandler {
	return &employeeHandler{
		employeesvc: employeesvc,
	}
}
