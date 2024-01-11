package department

import (
	"database/sql"

	"github.com/lorelva/designPatterns/DIP/department/employee"
)

type Department struct {
	DB *sql.DB
}

func (d *Department) AddEmployee(e employee.Employee) {
	e.Add(d.DB)
}

func (d *Department) GetEmployeeNames(e employee.Employee) []string {
	names := e.GetAllNames(d.DB)
	return names
}

// en este no funciona porque se repitería que lo anterior
func (d *Department) GetEmployee(e employee.Employee, name string) {
	//e.GetByID(d.DB, id)
	e.GetByName(d.DB, name)
}
