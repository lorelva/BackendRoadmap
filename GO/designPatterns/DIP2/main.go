package main

import (
	"fmt"

	"github.com/lorelva/designPatterns/DIP/department"
	"github.com/lorelva/designPatterns/DIP/department/employee"
)

func main() {
	employees := []employee.Employee{}
	department := department.NewDepartment(employees)

	// abstracción de supervisor y worker
	//ver en los arvhivos de supervisor.go y worker.go
	supervisor := employee.NewSupervisor(1, "ch")
	worker := employee.NewWorker(2, "jfj")

	department.AddEmployee(supervisor.GetSupervisor())
	department.AddEmployee(worker.GetWorker())

	names := department.GetEmployeeNames()
	fmt.Printf("Employees: %s\n", names)

	id := department.GetEmployee(2)
	fmt.Println(id)
}
