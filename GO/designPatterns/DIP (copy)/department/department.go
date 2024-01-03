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

func (d *Department) GetEmployee(e employee.Employee, id int, name string) {
	e.GetByID(d.DB, id)
	e.GetByName(d.DB, name)
}
