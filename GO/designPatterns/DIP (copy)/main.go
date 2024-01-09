package main

import (
	"fmt"

	"github.com/lorelva/designPatterns/DIP/database"
	"github.com/lorelva/designPatterns/DIP/department"
	"github.com/lorelva/designPatterns/DIP/department/employee/supervisor"
	"github.com/lorelva/designPatterns/DIP/department/employee/worker"
)

func main() {
	//la segunda base de datos para departamento
	departmentDB := database.ConnectToDepartment()

	d := department.Department{
		DB: departmentDB,
	}

	s := supervisor.Supervisor{
		Name: "Chris",
	}

	w := worker.Worker{
		Name: "Lore",
	}

	d.AddEmployee(&s)
	d.AddEmployee(&w)

	names := d.GetEmployeeNames(&s)
	fmt.Printf("Supervisors: %s\n", names)
	d.GetEmployee(&s, 0, "")

	names = d.GetEmployeeNames(&w)
	fmt.Printf("Workers: %s\n", names)
	d.GetEmployee(&w, 0, "")

	//concetar a la base de datos
	/*db := database.ConnectToTestTable()

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
	*/

}
