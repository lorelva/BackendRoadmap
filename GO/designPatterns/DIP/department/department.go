package department

import (
	"github.com/lorelva/designPatterns/DIP/department/employee"
)

type Department struct {
	Employees []employee.Employee
}

func (d *Department) AddEmployee(e employee.Employee) {
	d.Employees = append(d.Employees, e)
}

func (d *Department) GetEmployeeNames() (res []string) {
	for _, e := range d.Employees {
		res = append(res, e.GetName())
	}
	return
}

func (d *Department) GetEmployee(id int) employee.Employee {
	for _, e := range d.Employees {
		if e.GetID() == id {
			return e
		}
	}
	return nil
}
