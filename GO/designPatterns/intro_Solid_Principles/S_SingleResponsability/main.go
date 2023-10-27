package main

import "fmt"

// S--> Single Responsibility Principle (SRP)

// EXAMPLE
type Employee struct {
	Name    string
	Salary  float64
	Address string
}

// STATEMENT: Each struct should have only one responsibility
//Split the responsibilities into two separate structs so:

// STRUCT 1
type EmployeeInfo struct {
	Name   string
	Salary float64
}

// STRUCT 2
type EmployeeAddress struct {
	Address string
}

// OBJECTIVE: Have separate functions that handle the different responsibilities of each struct:
// FOR STRUCT 1
func (e EmployeeInfo) GetName() string {
	return e.Name
}

func (e EmployeeInfo) GetSalary() float64 {
	return e.Salary
}

// FOR STRUCT 2
func (e EmployeeAddress) GetAddress() string {
	return e.Address
}

func main() {

	//Create instances of EmployeeInfo and EmployeeAddress(structs)
	n := EmployeeInfo{Name: "Lorena Valle"}
	s := EmployeeInfo{Salary: 2000.90}
	a := EmployeeAddress{Address: "Mariano Matamoros"}

	name := n.GetName()
	salary := s.GetSalary()
	address := a.GetAddress()

	fmt.Println(name, salary, address)

}
