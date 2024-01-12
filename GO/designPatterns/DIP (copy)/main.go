package main

import (
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
		Name: "Reyes",
	}

	//s.UpdateByID(departmentDB, 4)
	//s.DeleteByID(departmentDB, 6)

	w := worker.Worker{
		Name: "	Maya",
	}

	//w.UpdateByID(departmentDB, 4)
	//w.DeleteByID(departmentDB, 1)

	//d.AddEmployee(&s)
	//d.AddEmployee(&w)

	/*names := d.GetEmployeeNames(&s)
	fmt.Printf("Supervisors: %s\n", names)
	d.GetEmployee(&s, "Carlos")
	*/

	/*names = d.GetEmployeeNames(&w)
	fmt.Printf("Workers: %s\n", names)
	d.GetEmployee(&w, "C")
	*/
	d.GetEmployee(&w, 4, "Lore")
	d.GetEmployee(&s, 3, "holi")

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
