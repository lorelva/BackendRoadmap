package main

import (
	"fmt"

	"github.com/lorelva/designPatterns/DIP/database"
	"github.com/lorelva/designPatterns/DIP/database/queries"
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

	//concetar a la base de datos
	db := database.ConnectToTestTable()

	//Crear tabla , accediendo a queries.go
	//queries.CrearTabla(db)

	//Agregar datos a la tabla
	//queries.InsertarSucursal(db, "lol", 3)
	//queries.InsertarSucursal(db, "Revolución", 4)
	//queries.InsertarSucursal(db, "Lomas de Cortes", 5)

	//Actualizar datos a la tabla
	//queries.ActualizarSucursal(db, "Centro", 3)

	//Eliminar datos de la tabla
	//queries.EliminarSucursal(db, 5)

	queries.ObtenerSucursal2(db)

}
