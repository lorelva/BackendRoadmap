package department

import (
	"github.com/lorelva/designPatterns/DIP/department/employee"
)

type Department struct {
	Employees []employee.Employee
}

// constructor concepto heredado de POO: una funcion que devuelve el objeto que queremos con sus atributos
func NewDepartment(e []employee.Employee) *Department {
	return &Department{
		Employees: e,
	}

	// db -> agregar en db
}
