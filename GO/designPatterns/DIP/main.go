package main

import (
	"fmt"

	"github.com/lorelva/designPatterns/DIP/department"
	"github.com/lorelva/designPatterns/DIP/department/employee/supervisor"
	"github.com/lorelva/designPatterns/DIP/department/employee/worker"
)

func main() {

	d := department.Department{}

	s := supervisor.Supervisor{
		ID:   1,
		Name: "Christian",
	}

	w := worker.Worker{
		ID:   2,
		Name: "Lorena",
	}

	d.AddEmployee(&s)
	d.AddEmployee(&w)

	names := d.GetEmployeeNames()
	fmt.Printf("Employees: %s\n", names)

	id := d.GetEmployee(2)
	fmt.Println(id)

	/*
		//fmt.Println(d, s, w)
		d.GetEmployee(s.ID)
		d.GetEmployee(w.ID)
		d.GetEmployeeNames(s.GetName)
	*/

}
